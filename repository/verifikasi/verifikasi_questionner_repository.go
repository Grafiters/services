package verifikasi

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiQuestionnerDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiQuestionnerRepository
	GetOneByVerifikasi(id int64) (response []models.VerifikasiQuestionnerResponse, err error)
	Store(request *models.VerifikasiQuestionner, tx *gorm.DB) (response *models.VerifikasiQuestionner, err error)
	UpdateStatusVerifikasi(request *models.UpdateStatusVerifikasi, tx *gorm.DB) (response bool, err error)
	AcceptValidasi(request *models.AcceptValidasiRequest, tx *gorm.DB) (response bool, err error)
	RejectValidasi(request *models.RejectValidasiRequest, tx *gorm.DB) (response bool, err error)
}

type VerifikasiQuestionnerRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewVerifikatiQuestionnerRepository(
	db lib.Database,
	logger logger.Logger,
) VerifikasiQuestionnerDefinition {
	return VerifikasiQuestionnerRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetOneByVerifikasi implements VerifikasiQuestionnerDefinition
func (q VerifikasiQuestionnerRepository) GetOneByVerifikasi(id int64) (response []models.VerifikasiQuestionnerResponse, err error) {
	db := q.db.DB.Table("verifikasi_questionner").
		Select(`id, 
				verifikasi_id, 
				questionner, 
				data_sumber, 
				checker, 
				signer, 
				approval_ord,
				jenis_fraud,
				status_validasi_rmc,
				status_validasi_signer,
				status_validasi_ord
			`)

	db = db.Where(`verifikasi_id = ?`, id)

	err = db.Find(&response).Error

	fmt.Println("data on repo =====>", response)

	return response, err
}

// Store implements VerifikasiQuestionnerDefinition
func (VerifikasiQuestionnerRepository) Store(request *models.VerifikasiQuestionner, tx *gorm.DB) (response *models.VerifikasiQuestionner, err error) {
	return request, tx.Table("verifikasi_questionner").Save(&request).Error
}

// WithTrx implements VerifikasiQuestionnerDefinition
func (q VerifikasiQuestionnerRepository) WithTrx(trxHandle *gorm.DB) VerifikasiQuestionnerRepository {
	if trxHandle == nil {
		q.logger.Zap.Error("transaction Database not found in gin context.")
		return q
	}

	q.db.DB = trxHandle
	return q
}

// AcceptValidasi implements VerifikasiQuestionnerDefinition
func (q VerifikasiQuestionnerRepository) AcceptValidasi(request *models.AcceptValidasiRequest, tx *gorm.DB) (response bool, err error) {
	// err = tx.Table("verifikasi_questionner").Save(&request).Error
	var dataValidasi interface{}

	switch request.ValidationBy {
	case "RMC":
		dataValidasi = models.AcceptValidasiRequest{
			StatusValidasiRmc:            request.StatusValidasiRmc,
			TindakLanjutIndikasiFraudRmc: request.TindakLanjutIndikasiFraudRmc,
			TindakLanjutRmc:              request.TindakLanjutRmc,
			CatatanRmc:                   request.CatatanRmc,
		}
	case "RRM":
		dataValidasi = models.AcceptValidasiRequest{
			StatusValidasiSigner:            request.StatusValidasiSigner,
			TindakLanjutIndikasiFraudSigner: request.TindakLanjutIndikasiFraudSigner,
			TindakLanjutSigner:              request.TindakLanjutSigner,
			CatatanSigner:                   request.CatatanSigner,
		}
	default:
		dataValidasi = models.AcceptValidasiRequest{
			StatusValidasiOrd:        request.StatusValidasiOrd,
			ValidasiIndikasiFraudOrd: request.ValidasiIndikasiFraudOrd,
			TindakLanjutOrd:          request.TindakLanjutOrd,
			CatatanOrd:               request.CatatanOrd,
		}
	}

	fmt.Println("dataValidasi => ", dataValidasi)
	err = tx.Table("verifikasi_questionner").Where("id = ?", request.ID).Updates(dataValidasi).Error

	if err != nil {
		q.logger.Zap.Error("transaction Database not found in gin context.")
		return false, err
	}

	return true, nil
}

// UpdateStatusVerifikasi implements VerifikasiQuestionnerDefinition
func (q VerifikasiQuestionnerRepository) UpdateStatusVerifikasi(request *models.UpdateStatusVerifikasi, tx *gorm.DB) (response bool, err error) {
	// var statusUpdate interface{}

	// switch request.IsTask {
	// case "Validasi":
	// 	fmt.Println("Validasi")
	// 	statusUpdate = models.UpdateStatusVerifikasi{
	// 		Status: request.Status,
	// 		Action: request.Action,
	// 	}
	// default:
	// 	fmt.Println("Reject")
	// 	statusUpdate = models.UpdateStatusVerifikasi{
	// 		IndikasiFraud: false,
	// 		Status:        request.Status,
	// 		Action:        request.Action,
	// 	}
	// }

	// fmt.Println("updateStatus =>", statusUpdate)

	// err = tx.Table("verifikasi").Where("id = ?", request.ID).Updates(statusUpdate).Error

	// if err != nil {
	// 	q.logger.Zap.Error("transaction Database not found in gin context.")
	// 	return false, err
	// }

	// return true, nil
	updateFields := make(map[string]interface{})

	switch request.IsTask {
	case "Validasi":
		fmt.Println("Validasi")
		updateFields["status_indikasi_fraud"] = request.StatusIndikasiFraud
		updateFields["action_indikasi_fraud"] = request.ActionIndikasiFraud
	default:
		fmt.Println("Reject")
		updateFields["indikasi_fraud"] = false
		updateFields["status_indikasi_fraud"] = request.StatusIndikasiFraud
		updateFields["action_indikasi_fraud"] = request.ActionIndikasiFraud
	}

	fmt.Println("updateFields =>", updateFields)

	err = tx.Table("verifikasi").Where("id = ?", request.ID).Updates(updateFields).Error

	if err != nil {
		q.logger.Zap.Error("Error updating record:", err)
		return false, err
	}

	return true, nil
}

// RejectValidasi implements VerifikasiQuestionnerDefinition
func (q VerifikasiQuestionnerRepository) RejectValidasi(request *models.RejectValidasiRequest, tx *gorm.DB) (response bool, err error) {
	var dataReject interface{}

	switch request.RejectBy {
	case "RMC":
		dataReject = models.RejectValidasiRequest{
			StatusValidasiRmc: request.StatusValidasiRmc,
			CatatanRmc:        request.CatatanRmc,
		}
	case "RRM":
		dataReject = models.RejectValidasiRequest{
			StatusValidasiSigner: request.StatusValidasiSigner,
			CatatanSigner:        request.CatatanSigner,
		}
	default:
		dataReject = models.RejectValidasiRequest{
			StatusValidasiOrd: request.StatusValidasiOrd,
			CatatanOrd:        request.CatatanOrd,
		}
	}

	fmt.Println("dataReject => ", dataReject)

	err = tx.Table("verifikasi_questionner").Where("id = ?", request.ID).Updates(dataReject).Error

	if err != nil {
		q.logger.Zap.Error("transaction Database not found in gin context.")
		return false, err
	}

	return true, nil
}
