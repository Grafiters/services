package ukerkelolaan

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/ukerkelolaan"
	ukRepo "riskmanagement/repository/ukerkelolaan"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
)

var (
	UUID = uuid.NewString()
)

type UkerKelolaanDefinition interface {
	GetAllWithPaginate(request models.KeywordRequest) (responses []models.UkerKelolaanResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.UkerKelolaanResponseOne, status bool, err error)
	Store(request models.UkerKelolaanRequest) (responses bool, err error)
	Update(request *models.UkerKelolaanRequest) (responses bool, err error)
	Delete(request models.UkerKelolaanRequest) (response bool, err error)
	FilterUkerKelolaan(request models.KeywordRequest) (responses []models.UkerKelolaanResponse, pagination lib.Pagination, err error)
	GetListUkerKelolaan(request *models.PencarianUker) (response []models.UkerList, err error)
}

type UkerKelolaanService struct {
	db     lib.Database
	logger logger.Logger
	ukRepo ukRepo.UkerKelolaanDefinition
}

func NewUkerKelolaanService(
	db lib.Database,
	logger logger.Logger,
	ukRepo ukRepo.UkerKelolaanDefinition,
) UkerKelolaanDefinition {
	return UkerKelolaanService{
		db:     db,
		logger: logger,
		ukRepo: ukRepo,
	}
}

// Delete implements UkerKelolaanDefinition
func (uk UkerKelolaanService) Delete(request models.UkerKelolaanRequest) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := uk.db.DB.Begin()

	updateData := &models.UkerKelolaanReqDelete{
		ID:        request.ID,
		UpdatedAt: &timeNow,
		Status:    false,
	}

	include := []string{
		"id",
		"updated_at",
		"status",
	}

	_, err = uk.ukRepo.Delete(updateData, include, tx)

	if err != nil {
		tx.Rollback()
		uk.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()

	return true, err
}

// GetAllWithPaginate implements UkerKelolaanDefinition
func (uk UkerKelolaanService) GetAllWithPaginate(request models.KeywordRequest) (responses []models.UkerKelolaanResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataUker, totalRows, totalData, err := uk.ukRepo.GetAllWithPaginate(&request)
	if err != nil {
		uk.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		uk.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataUker {
		responses = append(responses, models.UkerKelolaanResponse{
			ID:        response.ID,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
			ExpiredAt: response.ExpiredAt,
			IsTemp:    response.IsTemp,
			Pn:        response.Pn,
			SNAME:     response.SNAME,
			REGION:    response.REGION,
			RGDESC:    response.RGDESC,
			MAINBR:    response.MAINBR,
			MBDESC:    response.MBDESC,
			BRANCH:    response.BRANCH,
			BRDESC:    response.BRDESC,
			Status:    response.Status,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// Store implements UkerKelolaanDefinition
func (uk UkerKelolaanService) Store(request models.UkerKelolaanRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	cekExist, err := uk.ukRepo.CekBRCKelolaan(request.Pn)

	if cekExist < 1 {
		fmt.Println("massokk")
		fmt.Println("jumlah ===>", cekExist)
		tx := uk.db.DB.Begin()
		status, err := uk.ukRepo.StoreMstUker(&models.SaveMstRequest{
			Pn:    request.Pn,
			Sname: request.SNAME,
			Aktif: request.Status,
		}, tx)

		if !status {
			tx.Rollback()
			uk.logger.Zap.Error(err)
			return false, err
		}

		if len(request.ListUker) != 0 {
			for _, value := range request.ListUker {
				_, err = uk.ukRepo.Store(&models.UkerKelolaan{
					CreatedAt: &timeNow,
					IsTemp:    false,
					Pn:        request.Pn,
					SNAME:     request.SNAME,
					REGION:    value.REGION,
					RGDESC:    value.RGDESC,
					MAINBR:    value.MAINBR,
					MBDESC:    value.MBDESC,
					BRANCH:    value.BRANCH,
					BRDESC:    value.BRDESC,
					Status:    true,
				}, tx)

				if err != nil {
					tx.Rollback()
					uk.logger.Zap.Error(err)
					return false, err
				}
			}
		} else {
			if err != nil {
				tx.Rollback()
				uk.logger.Zap.Error(err)
				return false, err
			}
		}

		tx.Commit()
		return true, err
	}

	return false, err
}

// Update implements UkerKelolaanDefinition
func (uk UkerKelolaanService) Update(request *models.UkerKelolaanRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := uk.db.DB.Begin()

	status, err := uk.ukRepo.StoreMstUker(&models.SaveMstRequest{
		Id:    request.ID,
		Pn:    request.Pn,
		Sname: request.SNAME,
		Aktif: request.Status,
	}, tx)

	if !status {
		tx.Rollback()
		uk.logger.Zap.Error(err)
		return false, err
	}

	if len(request.ListUker) != 0 {
		for _, value := range request.ListUker {
			updateData := &models.UkerKelolaan{
				ID:        value.ID,
				CreatedAt: request.CreatedAt,
				UpdatedAt: &timeNow,
				// ExpiredAt: "",
				IsTemp: false,
				Pn:     request.Pn,
				SNAME:  request.SNAME,
				REGION: value.REGION,
				RGDESC: value.RGDESC,
				MAINBR: value.MAINBR,
				MBDESC: value.MBDESC,
				BRANCH: value.BRANCH,
				BRDESC: value.BRDESC,
				Status: request.Status,
			}

			include := []string{
				"id",
				"created_at",
				"updated_at",
				"is_tempt",
				"pn",
				"SNAME",
				"REGION",
				"RGDESC",
				"MAINBR",
				"MBDESC",
				"BRANCH",
				"BRDESC",
				"status",
			}

			_, err = uk.ukRepo.Update(updateData, include, tx)

			if err != nil {
				tx.Rollback()
				uk.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			uk.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

// FilterUkerKelolaan implements UkerKelolaanDefinition
func (uk UkerKelolaanService) FilterUkerKelolaan(request models.KeywordRequest) (responses []models.UkerKelolaanResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataUker, totalRows, totalData, err := uk.ukRepo.FilterUkerKelolaan(&request)
	if err != nil {
		uk.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		uk.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataUker {
		responses = append(responses, models.UkerKelolaanResponse{
			ID:        response.ID,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
			ExpiredAt: response.ExpiredAt,
			IsTemp:    response.IsTemp,
			Pn:        response.Pn,
			SNAME:     response.SNAME,
			REGION:    response.REGION,
			RGDESC:    response.RGDESC,
			MAINBR:    response.MAINBR,
			MBDESC:    response.MBDESC,
			BRANCH:    response.BRANCH,
			BRDESC:    response.BRDESC,
			Status:    response.Status,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (uk UkerKelolaanService) GetOne(id int64) (responses models.UkerKelolaanResponseOne, status bool, err error) {

	fmt.Println("pn ===> ", id)
	// dataKelolaan, err := uk.ukRepo.GetOne(pn)
	dataKelolaan, err := uk.ukRepo.GetDetailData(id)

	fmt.Println(dataKelolaan)

	if dataKelolaan.ID != 0 {
		listUker, err := uk.ukRepo.GetListUker(dataKelolaan.Pn)

		responses = models.UkerKelolaanResponseOne{
			ID:        dataKelolaan.ID,
			CreatedAt: dataKelolaan.CreatedAt,
			UpdatedAt: dataKelolaan.UpdatedAt,
			ExpiredAt: dataKelolaan.ExpiredAt,
			IsTemp:    dataKelolaan.IsTemp,
			Pn:        dataKelolaan.Pn,
			SNAME:     dataKelolaan.SNAME,
			ListUker:  listUker,
			Status:    dataKelolaan.Status,
		}

		return responses, true, err
	}

	return responses, false, err

}

// GetListUkerKelolaan implements UkerKelolaanDefinition.
func (uk UkerKelolaanService) GetListUkerKelolaan(request *models.PencarianUker) (response []models.UkerList, err error) {
	listData, err := uk.ukRepo.GetListUkerKelolaan(request)

	// fmt.Println(listData)
	if err != nil {
		uk.logger.Zap.Error(err)
		return nil, err
	}

	for _, value := range listData {
		response = append(response, models.UkerList{
			REGION: value.REGION,
			RGDESC: value.RGDESC,
			MAINBR: value.MAINBR,
			MBDESC: value.MBDESC,
			BRANCH: value.BRANCH,
			BRDESC: value.BRDESC,
		})
	}

	return response, err
}
