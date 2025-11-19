package jenistask

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/jenistask"

	"gitlab.com/golang-package-library/logger"
)

type JenisTaskDefinition interface {
	GetData(request *models.TaskRequest) (response []models.JenisTask, err error)
}

type JenisTaskRepository struct {
	db     lib.Database
	logger logger.Logger
}

func NewJenisTaskRepository(
	db lib.Database,
	logger logger.Logger,
) JenisTaskDefinition {
	return JenisTaskRepository{
		db:     db,
		logger: logger,
	}
}

func (r JenisTaskRepository) GetData(request *models.TaskRequest) (response []models.JenisTask, err error) {
	query := r.db.DB.Table("admin_setting task").
		Select(`
			task.id,
			task.task_type as jenis_task,
			task.kegiatan,
			task.period,
			task.range,
			task.upload
		`).
		Joins(`JOIN admin_setting_role as roles ON roles.id_setting = task.id`).
		Where(`task.Status = 'Aktif'`).
		Group("task.id")

	query = query.Where(`roles.tipe_uker LIKE ?`, fmt.Sprintf("%%%s%%", request.TipeUker))

	if request.TipeUker == "KP" {
		query = query.Where(`roles.kostl = ?`, request.Kostl)

		if request.Hilfm != "" {
			query = query.Where(`roles.hilfm LIKE ?`, fmt.Sprintf("%%%s%%", request.Hilfm))
		}
	}

	if request.TipeUker == "KW" {
		if request.Stell != "" {
			query = query.Where(`roles.stell LIKE ?`, fmt.Sprintf("%%%s%%", request.Stell))
		}

		if request.Jgpg != "" {
			query = query.Where(`roles.jg LIKE ?`, fmt.Sprintf("%%%s%%", request.Jgpg))
		}
	}

	err = query.Scan(&response).Error

	if err != nil {
		r.logger.Zap.Errorf("Error get data", err)
		return response, err
	}

	return response, nil
}
