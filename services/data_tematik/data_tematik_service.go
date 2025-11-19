package datatematik

import (
	"fmt"
	models "riskmanagement/models/data_tematik"
	repo "riskmanagement/repository/data_tematik"
)

type DataTematikServiceDefinition interface {
	GetSampleDataTematik(request models.DataTematikRequest) (response models.DataTematikResponse, totalData int64, err error)
	UpdateStatusDataSample(request models.UpdaterData) (response bool, err error)
}

type DataTematikService struct {
	repo repo.DataTematikDefinition
}

func NewDataTematikService(
	repo repo.DataTematikDefinition,
) DataTematikServiceDefinition {
	return DataTematikService{
		repo: repo,
	}
}

// GetSampleDataTematik implements DataTematikServiceDefinition.
func (dt DataTematikService) GetSampleDataTematik(request models.DataTematikRequest) (response models.DataTematikResponse, totalData int64, err error) {
	dataResponse, totalData, err := dt.repo.GetSampleDataTematik(&request)

	return dataResponse, totalData, err
}

// UpdateStatusDataSample implements DataTematikServiceDefinition.
func (dt DataTematikService) UpdateStatusDataSample(request models.UpdaterData) (response bool, err error) {
	if len(request.RequestUpdate) < 1 {
		return false, fmt.Errorf("Data Sample Kosong")
	}

	fmt.Println("Updateer DATA TEMATIK =>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", request)

	for _, data := range request.RequestUpdate {
		_, err := dt.repo.UpdateStatusDataSample(&models.RequestUpdate{
			NamaTable:    data.NamaTable,
			Id:           data.Id,
			Status:       data.Status,
			NoVerifikasi: data.NoVerifikasi,
		})

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
