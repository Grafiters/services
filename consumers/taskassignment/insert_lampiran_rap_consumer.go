package taskassignment

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	publisherModels "riskmanagement/models/publisher"
	models "riskmanagement/models/taskassignment"
	repository "riskmanagement/repository/taskassignment"

	filemanager "riskmanagement/services/filemanager"

	"github.com/streadway/amqp"
	"gitlab.com/golang-package-library/logger"
	"go.uber.org/zap"
)

type InsertLampiranRAPConsumer struct {
	logger      logger.Logger
	queueName   string
	repository  repository.TaskAssignmentsDefinition
	filemanager filemanager.FileManagerDefinition
}

func NewInsertLampiranRAPConsumer(
	logger logger.Logger,
	repository repository.TaskAssignmentsDefinition,
	filemanager filemanager.FileManagerDefinition,
) InsertLampiranRAPConsumer {
	return InsertLampiranRAPConsumer{
		logger:      logger,
		queueName:   "lampiran-rap",
		repository:  repository,
		filemanager: filemanager,
	}
}
func (r InsertLampiranRAPConsumer) Setup(Channel *amqp.Channel) error {
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
				r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Failed to unmarshal message: %v", err)
				continue
			}
			r.logger.Zap.Info("[TaskAssignment - Lampiran RAP]] Received a message", zap.Any("payload", payload.Pattern))

			if payload.Pattern == "insert-lampiran-rap" {
				headersInterface, ok := payload.Body["headers"]
				if !ok {
					r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid headers format")
					continue
				}

				headersSlice, ok := headersInterface.([]interface{})
				if !ok {
					r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid headers type")
					continue
				}

				headers := make([]string, len(headersSlice))
				for i, header := range headersSlice {
					headerStr, ok := header.(string)
					if !ok {
						r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid header type")
						continue
					}
					headers[i] = headerStr
				}

				rowsInterface, ok := payload.Body["rows"]
				if !ok {
					r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid rows format")
					continue
				}

				rowsSlice, ok := rowsInterface.([]interface{})
				if !ok {
					r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid rows type")
					continue
				}

				rows := make([][]string, len(rowsSlice))
				for i, row := range rowsSlice {
					rowSlice, ok := row.([]interface{})
					if !ok {
						r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid row type")
						continue
					}

					rows[i] = make([]string, len(rowSlice))
					for j, cell := range rowSlice {
						cellStr, ok := cell.(string)
						if !ok {
							r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Invalid cell type")
							continue
						}
						rows[i][j] = cellStr
					}
				}

				err := r.InserLampiranRAP(int64(payload.Body["tasklist_id"].(float64)), headers, rows, payload.Body["filepath"].(string))
				if err != nil {
					r.logger.Zap.Error("[TaskAssignment - Lampiran RAP] Failed to Insert Tasklist Uker: %v", err)
				}
				// err := r.InserLampiranRAP(int64(payload.Body["tasklist_id"].(float64)), payload.Body["headers"].([]interface{}), payload.Body["rows"].([]interface{}), payload.Body["filepath"].(string))
				// if err != nil {
				// 	r.logger.Zap.Error("[TaskAssignment - Lampiran RAP]] Failed to Insert Tasklist Uker: %v", err)
				// }
			}

			r.logger.Zap.Info("[TaskAssignment - Lampiran RAP] Message %v processed successfully", payload.Pattern)

		}
		<-forever
	}()
	return nil
}

func (r InsertLampiranRAPConsumer) InserLampiranRAP(tasklist_id int64, headers []string, rows [][]string, filepath string) error {
	tasklist, err := r.repository.GetOneById(int64(tasklist_id))
	if err != nil {
		tasklist.StatusFile = "Gagal mendapatkan Tasklsit : " + err.Error()
		r.repository.UpdateTasklist(tasklist)
	}
	tableName := "lampiran_rap_" + strconv.FormatInt(tasklist.RiskIssueID, 10) + "_" + strconv.FormatInt(tasklist.RiskIndicatorID, 10)

	isExist, err := r.repository.ValidatFileToDB(tableName, headers)

	fmt.Println("Table Exist", isExist)

	if err != nil {
		tasklist.StatusFile = "Gagal validasi file : " + err.Error()
		r.repository.UpdateTasklist(tasklist)
		os.Remove(filepath)

	} else {
		if isExist {
			err := r.repository.BatchStoreLampiranRAP(tableName, headers, rows, tasklist.ID)
			if err != nil {
				tasklist.StatusFile = "Gagal input file ke DB : " + err.Error()
				r.repository.UpdateTasklist(tasklist)
				os.Remove(filepath)
			}
		} else if !isExist {
			err := r.repository.CreateTableLampiranRAP(tableName, headers)
			if err != nil {
				tasklist.StatusFile = "Gagal Membuat Table Lampiran RAP : " + err.Error()
				r.repository.UpdateTasklist(tasklist)
				os.Remove(filepath)
			}

			err = r.repository.InsertLampiranIndicator(models.LampiranIndicatorRequest{
				RiskIssueId:       tasklist.RiskIssueID,
				RiskIndicatorId:   tasklist.RiskIndicatorID,
				NamaTable:         tableName,
				JumlahKolom:       int64(len(headers)),
				RiskIndicatorDesc: tasklist.RiskIndicator,
			})
			if err != nil {
				tasklist.StatusFile = "Gagal input Lampiran Indikator : " + err.Error()
				r.repository.UpdateTasklist(tasklist)
				os.Remove(filepath)
			}

			err = r.repository.BatchStoreLampiranRAP(tableName, headers, rows, tasklist.ID)
			if err != nil {
				tasklist.StatusFile = "Gagal input file ke DB : " + err.Error()
				r.repository.UpdateTasklist(tasklist)
				os.Remove(filepath)
			}
		}
	}
	return nil
}
