package taskassignment

import (
	"encoding/json"
	"fmt"
	"os"
	publisherModels "riskmanagement/models/publisher"
	models "riskmanagement/models/taskassignment"

	repository "riskmanagement/repository/taskassignment"
	filemanager "riskmanagement/services/filemanager"

	"github.com/streadway/amqp"
	"gitlab.com/golang-package-library/logger"
	"go.uber.org/zap"
)

type InserTasklistUkerConsumer struct {
	logger      logger.Logger
	queueName   string
	repository  repository.TaskAssignmentsDefinition
	filemanager filemanager.FileManagerDefinition
}

func NewInserTasklistUkerConsumer(
	logger logger.Logger,
	repository repository.TaskAssignmentsDefinition,
	filemanager filemanager.FileManagerDefinition,
) InserTasklistUkerConsumer {
	return InserTasklistUkerConsumer{
		logger:      logger,
		queueName:   "tasklist-uker",
		repository:  repository,
		filemanager: filemanager,
	}
}
func (r InserTasklistUkerConsumer) Setup(Channel *amqp.Channel) error {
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
				r.logger.Zap.Error("[TaskAssignment - Tasklist Uker] Failed to unmarshal message: %v", err)
				break
			}
			r.logger.Zap.Info("[TaskAssignment - Tasklist Uker] Received a message", zap.Any("payload", payload.Pattern))

			if payload.Pattern == "insert-tasklist-uker" {
				rowsInterface := payload.Body["rows"].([]interface{})
				rows := make([][]string, len(rowsInterface))

				fmt.Println("Converting rows for Tasklist Uker, total rows:", len(rowsInterface))

				for i, rowInterface := range rowsInterface {
					row := rowInterface.([]interface{})
					rows[i] = make([]string, len(row))
					for j, cell := range row {
						rows[i][j] = cell.(string)
					}
				}

				fmt.Println("Sample row data:", rows)

				err := r.BatchStoreTasklistUker(rows, int64(payload.Body["tasklist_id"].(float64)), payload.Body["filepath"].(string))
				if err != nil {
					r.logger.Zap.Error("[TaskAssignment - Tasklist Uker] Failed to Insert Tasklist Uker: %v", err)
					os.Remove(payload.Body["filepath"].(string))
				}
				os.Remove(payload.Body["filepath"].(string))
			}

			r.logger.Zap.Info("[TaskAssignment - Tasklist Uker] Message %v processed successfully", payload.Pattern)
		}
		<-forever
	}()
	return nil
}

func (r InserTasklistUkerConsumer) BatchStoreTasklistUker(rows [][]string, tasklist_id int64, filepath string) error {
	tasklist, err := r.repository.GetOneById(tasklist_id)

	fmt.Println("Starting BatchStoreTasklistUker for Tasklist:", tasklist)

	if err != nil {
		r.logger.Zap.Error("[TaskAssignment - Tasklist Uker] Failed to get Tasklist: %v", err)
		tasklist.StatusFile = "Gagal mendapatkan Tasklist : " + err.Error()
		r.repository.UpdateTasklist(tasklist)
		return fmt.Errorf("failed to Get Tasklist (%v)", err)
	}

	duplicateCounts := make(map[string]int)
	distinctData := make([]models.TasklistUker, 0, len(rows))

	fmt.Println("Rows to process for Tasklist Uker:", len(rows))
	// fmt.Println("Duplicated Count", duplicateCounts)
	// fmt.Println("Duplicate Data", distinctData)

	for _, row := range rows {

		fmt.Println("consumer uker => Processing row:", row)
		key := fmt.Sprintf("%s|%s|%s|%s|%s|%s", row[0], row[1], row[2], row[3], row[4], row[5])

		fmt.Println("Key:", key)
		duplicateCounts[key]++

		if duplicateCounts[key] == 1 {
			distinctData = append(distinctData, models.TasklistUker{
				TasklistId: tasklist_id,
				REGION:     row[0], RGDESC: row[1], MAINBR: row[2],
				MBDESC: row[3], BRANCH: row[4], BRDESC: row[5],
				JumlahNominatif: duplicateCounts[key],
			})
		} else {
			for i := range distinctData {
				if distinctData[i].REGION == row[0] && distinctData[i].RGDESC == row[1] &&
					distinctData[i].MAINBR == row[2] && distinctData[i].MBDESC == row[3] &&
					distinctData[i].BRANCH == row[4] && distinctData[i].BRDESC == row[5] {
					distinctData[i].JumlahNominatif = duplicateCounts[key]
					break
				}
			}
		}
	}

	err = r.repository.BatchStoreTasklistUker(distinctData)
	if err != nil {
		r.logger.Zap.Error("[TaskAssignment - Tasklist Uker] Failed to Batch Store Tasklist Uker: %v", err)
		tasklist.StatusFile = "Gagal input ke Table Tasklist User  : " + err.Error()
		r.repository.UpdateTasklist(tasklist)
		return fmt.Errorf("failed to insert to DB Tasklist User (%v)", err)
	}
	return nil
}
