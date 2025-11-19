package taskassignment

import (
	"encoding/json"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/taskassignment"
	repository "riskmanagement/repository/taskassignment"

	"github.com/streadway/amqp"
	"gitlab.com/golang-package-library/logger"
	"go.uber.org/zap"
)

type ApprovalTaskAssignmentConsumer struct {
	logger     logger.Logger
	queueName  string
	repository repository.TaskAssignmentsDefinition
}

func NewApprovalTaskAssignmentConsumer(
	logger logger.Logger,
	repository repository.TaskAssignmentsDefinition,
) ApprovalTaskAssignmentConsumer {
	return ApprovalTaskAssignmentConsumer{
		logger:     logger,
		queueName:  "taskassignment-approval",
		repository: repository,
	}
}

type ApprovalMessage struct {
	BrcList  []models.GetBRCResponse `json:"brcList"`  // Matches the "brcList" key in the JSON
	Tasklist models.Task             `json:"tasklist"` // Matches the "tasklist" key in the JSON
	// Task     string                  `json:"task"`     // Matches the "task" key in the JSON
}

func (r ApprovalTaskAssignmentConsumer) Setup(Channel *amqp.Channel) error {
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
			var raw map[string]interface{}
			if err := json.Unmarshal(msg.Body, &raw); err != nil {
				r.logger.Zap.Error("[TaskAssignment - Approval] Failed to unmarshal message", zap.ByteString("body", msg.Body), zap.Error(err))
				continue
			}

			var payload ApprovalMessage
			bodyRaw, ok := raw["body"]
			if !ok {
				r.logger.Zap.Warn("[TaskAssignment - Approval] 'body' field not found in message")
				continue
			}
			bodyBytes, err := json.Marshal(bodyRaw)
			if err != nil {
				r.logger.Zap.Error("[TaskAssignment - Approval] Failed to marshal 'body' field", zap.Any("body", bodyRaw), zap.Error(err))
				continue
			}

			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				r.logger.Zap.Error("[TaskAssignment - Approval] Failed to unmarshal 'body' field", zap.ByteString("body", bodyBytes), zap.Error(err))
				continue
			}

			r.logger.Zap.Info("[TaskAssignment - Approval] Received a message", zap.Any("payload", payload))

			err = r.InsertTaskToday(payload.BrcList, payload.Tasklist)
			if err != nil {
				r.logger.Zap.Error("[TaskAssignment - Approval] Failed to insert task: %v", err)
				continue
			}

			r.logger.Zap.Info("[TaskAssignment - Approval] Message processed successfully", zap.Any("payload", payload))
		}
		<-forever
	}()
	return nil
}

func (r ApprovalTaskAssignmentConsumer) InsertTaskToday(brcList []models.GetBRCResponse, tasklist models.Task) (err error) {
	timeNow := lib.GetTimeNow("timestime")
	tasklistsToday := []models.TasklistsToday{}
	notif := []models.TasklistNotif{}

	for _, brc := range brcList {
		sample := tasklist.Sample
		if tasklist.Sample == 0 {
			sample = brc.JumlahNominatif
		}

		tasklistsToday = append(tasklistsToday, models.TasklistsToday{
			TasklistID:      tasklist.ID,
			ActivityID:      tasklist.ActivityID,
			ProductID:       tasklist.ProductID,
			Product:         tasklist.ProductName,
			RiskIssueID:     tasklist.RiskIssueID,
			RiskIssue:       tasklist.RiskIssue,
			RiskIndicatorID: tasklist.RiskIndicatorID,
			RiskIndicator:   tasklist.RiskIndicator,
			StartDate:       tasklist.StartDate,
			EndDate:         tasklist.EndDate,
			Status:          tasklist.Status,
			TaskType:        tasklist.TaskType,
			TaskTypeName:    tasklist.TaskTypeName,
			Kegiatan:        tasklist.Kegiatan,
			Period:          tasklist.Period,
			Sample:          sample,
			Progres:         0,
			Persentase:      0,
			MakerID:         tasklist.MakerID,
			PERNR:           brc.PN,
			Assigned:        brc.SNAME,
			REGION:          brc.REGION,
			RGDESC:          brc.RGDESC,
			MAINBR:          brc.MAINBR,
			MBDESC:          brc.MBDESC,
			BRANCH:          brc.BRANCH,
			BRDESC:          brc.BRDESC,
			CreatedAt:       &timeNow,
			UpdatedAt:       &timeNow,
		})

		notif = append(notif, models.TasklistNotif{
			TaskID:     tasklist.ID,
			Tanggal:    &timeNow,
			Keterangan: "Ada 1 tasklist baru yang harus dikerjakan",
			Status:     0,
			Jenis:      "Tasklist",
			Receiver:   brc.PN,
			Uker:       brc.BRDESC,
		})
	}

	err = r.repository.BatchStoreTasklistToday(tasklistsToday)
	if err != nil {
		return err
	}

	err = r.repository.BatchStoreNotifTask(notif)
	if err != nil {
		return err
	}

	return nil
}
