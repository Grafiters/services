package taskassignment

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"riskmanagement/consumers/publisher"
	fileModels "riskmanagement/models/filemanager"
	publisherModels "riskmanagement/models/publisher"
	repository "riskmanagement/repository/taskassignment"
	"strings"

	filemanager "riskmanagement/services/filemanager"

	"github.com/streadway/amqp"
	"github.com/xuri/excelize/v2"
	"gitlab.com/golang-package-library/logger"
	"go.uber.org/zap"

	fileext "path/filepath"
)

type ReadFileMinioConsumer struct {
	logger      logger.Logger
	queueName   string
	repository  repository.TaskAssignmentsDefinition
	filemanager filemanager.FileManagerDefinition
	publisher   publisher.PublisherInterface
}

func NewReadFileMinioConsumer(
	logger logger.Logger,
	repository repository.TaskAssignmentsDefinition,
	filemanager filemanager.FileManagerDefinition,
	publisher publisher.PublisherInterface,
) ReadFileMinioConsumer {
	return ReadFileMinioConsumer{
		logger:      logger,
		queueName:   "taskassignment-read-file-minio",
		repository:  repository,
		filemanager: filemanager,
		publisher:   publisher,
	}
}
func (r ReadFileMinioConsumer) Setup(Channel *amqp.Channel) error {
	// Declare the queue
	_, err := Channel.QueueDeclare(
		r.queueName, // Name of the queue
		true,        // Durable
		true,        // Delete when unused
		false,       // Exclusive
		false,       // No-wait
		nil,         // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", r.queueName, err)
	}

	// Set up a consumer
	msgs, err := Channel.Consume(
		r.queueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return fmt.Errorf("failed to consume queue %s: %w", r.queueName, err)
	}
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			var payload publisherModels.PublishMessageDTO
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				r.logger.Zap.Error("[TaskAssignment - Read File Minio] Failed to unmarshal message: %v", err)
				break
			}
			r.logger.Zap.Info("[TaskAssignment - Read File Minio] Received a message", zap.Any("payload", payload.Pattern))

			if payload.Pattern == "insert-tasklist" {
				err := r.InsertTasklist(payload.Body)
				if err != nil {
					r.logger.Zap.Error("[TaskAssignment - Read File Minio] Failed to InsertTasklist: %v", err)
				}
			}

			r.logger.Zap.Info("[TaskAssignment - Read File Minio] Message %v processed successfully", payload.Pattern)
		}
		<-forever
	}()
	return nil
}

func (r ReadFileMinioConsumer) InsertTasklist(payload map[string]interface{}) (err error) {
	tasklist_id := int64(payload["tasklist_id"].(float64))
	minioPath := payload["minioPath"].(string)
	filename := payload["filename"].(string)

	tasklist, err := r.repository.GetOneById(tasklist_id)
	if err != nil {
		tasklist.StatusFile = "Gagal mendapatkan Tasklist : " + err.Error()
		r.repository.UpdateTasklist(tasklist)
		return fmt.Errorf("failed to Get Tasklist (%v)", err)

	}

	pathFiles, err := r.filemanager.ReadFile(fileModels.FileManagerRequest{
		ObjectName: minioPath,
		Filename:   filename,
	})

	if err != nil {
		tasklist.StatusFile = "Gagal memndapatkan File dari MinIO : " + err.Error()
		r.repository.UpdateTasklist(tasklist)
		os.Remove(pathFiles)
		return fmt.Errorf("failed to Get file from MinIO (%v)", err)
	}

	var rows [][]string
	var headers []string

	ext := strings.ToLower(fileext.Ext(filename))

	fmt.Println("File extension:", ext)
	fmt.Println("PATH FILES:", pathFiles)

	switch ext {
	case ".xlsx", ".xls":
		fmt.Println("Excel file detected")

		file, err := excelize.OpenFile(pathFiles)
		if err != nil {
			tasklist.StatusFile = "WOYLAH xlsx : " + err.Error()
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("failed to open File (%v)", err)
		}

		records, err := file.GetRows("Sheet1")
		if err != nil {
			tasklist.StatusFile = "Gagal Membaca Sheet : " + err.Error()
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("failed to read Sheet (%v)", err)
		}

		if len(records) < 1 {
			tasklist.StatusFile = "Gagal : sheet harus bernama Sheet1 (Default)"
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("sheet name must be Sheet1 (Default)")
		}

		rows = records
	case ".csv":
		fmt.Println("CSV file detected")

		files, err := os.Open(pathFiles)
		if err != nil {
			// tasklist.StatusFile = "Gagal Membuka File : " + err.Error()
			tasklist.StatusFile = "WOYLAH CSV: " + err.Error()
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("failed to open CSV file: %v", err)
		}

		defer files.Close()

		bufReader := bufio.NewReader(files)
		firstLine, _ := bufReader.ReadString('\n')

		// Reset reader back to start of file
		files.Seek(0, 0)

		reader := csv.NewReader(files)

		// Detect delimiter based on first line
		switch {
		case strings.Contains(firstLine, ";"):
			reader.Comma = ';'
			fmt.Println("Detected delimiter: semicolon ;")
		case strings.Contains(firstLine, "\t"):
			reader.Comma = '\t'
			fmt.Println("Detected delimiter: tab \\t")
		default:
			reader.Comma = ',' // fallback
			fmt.Println("Detected delimiter: comma ,")
		}

		records, err := reader.ReadAll()
		if err != nil {
			tasklist.StatusFile = "Gagal Membaca CSV : " + err.Error()
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("failed to read CSV: %v", err)
		}

		if len(records) < 1 {
			tasklist.StatusFile = "Gagal : CSV kosong"
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("CSV file is empty")
		}

		rows = records
	default:
		tasklist.StatusFile = "Format file tidak didukung"
		r.repository.UpdateTasklist(tasklist)
		os.Remove(pathFiles)
		return fmt.Errorf("unsupported file format: %s", ext)
	}

	// Open the Excel file
	re := regexp.MustCompile(`[^a-zA-Z0-9_ ' ]`)
	for i := range rows {
		for j := range rows[i] {
			// Sanitize each string using the regex
			rows[i][j] = re.ReplaceAllString(rows[i][j], "")
		}
	}

	headers = rows[0]
	rows = rows[1:]

	fmt.Println("Rows read successfully:", len(rows))
	fmt.Println("Headers: ", headers)
	// fmt.Println("Data Rows:", rows)

	// //Batch Insert Tasklist Uker

	chunkSize := 1000

	// Batch Insert Lampiran
	for i := 0; i < len(rows); i += chunkSize {
		fmt.Println("Read File -> Processing rows from", rows)

		end := i + chunkSize
		if end > len(rows) {
			end = len(rows)
		}

		dataTasklistUker := map[string]interface{}{
			"tasklist_id": tasklist.ID,
			"filepath":    pathFiles,
			"rows":        rows[i:end],
		}

		tasklistUkerPayload := publisherModels.PublishMessageDTO{
			QueueName: "tasklist-uker",
			Pattern:   "insert-tasklist-uker",
			Body:      dataTasklistUker,
		}

		fmt.Println("Publishing to tasklist-uker queue:", tasklistUkerPayload)

		err = r.publisher.PublishMessage(tasklistUkerPayload)
		if err != nil {
			tasklist.StatusFile = "Gagal mengirim data ke Antrian  : " + err.Error()
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("failed to publish message rabbitmq : %v (%v)", r.queueName, err)
		}

		dataLampiranRAP := map[string]interface{}{
			"tasklist_id": tasklist.ID,
			"headers":     headers,
			"rows":        rows[i:end],
			"filepath":    pathFiles,
			"minioPath":   minioPath,
		}

		lampiranRAPPayload := publisherModels.PublishMessageDTO{
			QueueName: "lampiran-rap",
			Pattern:   "insert-lampiran-rap",
			Body:      dataLampiranRAP,
		}

		err = r.publisher.PublishMessage(lampiranRAPPayload)
		if err != nil {
			tasklist.StatusFile = "Gagal input file ke DB : " + err.Error()
			r.repository.UpdateTasklist(tasklist)
			os.Remove(pathFiles)
			return fmt.Errorf("failed to publish message rabbitmq : insert-lampiran-rap (%v)", err)
		}

	}

	_, err = r.filemanager.RemoveObject(fileModels.FileManagerRequest{
		ObjectName: payload["minioPath"].(string),
	})
	if err != nil {
		r.logger.Zap.Error("[TaskAssignment - Read File Minio] Failed to remove file: %v", err)
	}

	os.Remove(pathFiles)
	tasklist.StatusFile = "Selesai"
	r.repository.UpdateTasklist(tasklist)

	return nil
}
