package linkcage

import (
	"fmt"
	"riskmanagement/lib"
	requestFile "riskmanagement/models/files"
	models "riskmanagement/models/linkcage"
	fileRepo "riskmanagement/repository/files"
	repository "riskmanagement/repository/linkcage"

	fileModel "riskmanagement/models/filemanager"
	filemanager "riskmanagement/services/filemanager"

	"github.com/google/uuid"

	"gitlab.com/golang-package-library/minio"

	"os"
	"strings"
)

type LinkcageDefinition interface {
	GetAll(request models.LinkcageRequest) (responses []models.LinkcageResponse, pagination lib.Pagination, err error)
	Store(request models.LinkcageRequest) (status bool, err error)
	SetStatus(request *models.LinkcageRequest) (response bool, err error)
	GetActive() (responses []models.LinkcageResponse, err error)
	Delete(request *models.LinkcageRequest) (response bool, err error)

	GetOne(request models.LinkcageRequest) (response models.LinkcageResponse, status bool, err error)
	Update(request *models.LinkcageRequest) (status bool, err error)
}

type LinkcageService struct {
	db          lib.Database
	minio       minio.Minio
	repository  repository.LinkcageDefinition
	fileRepo    fileRepo.FilesDefinition
	linkImg     repository.LinkcageImageDefinition
	filemanager filemanager.FileManagerDefinition
}

func NewLinkcageService(
	db lib.Database,
	minio minio.Minio,
	repository repository.LinkcageDefinition,
	fileRepo fileRepo.FilesDefinition,
	linkImg repository.LinkcageImageDefinition,
	filemanager filemanager.FileManagerDefinition,
) LinkcageDefinition {
	return LinkcageService{
		db:          db,
		minio:       minio,
		repository:  repository,
		fileRepo:    fileRepo,
		linkImg:     linkImg,
		filemanager: filemanager,
	}
}

func (link LinkcageService) GetAll(request models.LinkcageRequest) (responses []models.LinkcageResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, 10, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	links, totalRows, totalData, err := link.repository.GetAll(&request)
	if err != nil {
		return responses, pagination, err
	}

	for _, link := range links {
		responses = append(responses, models.LinkcageResponse{
			ID:        link.ID,
			Name:      link.Name,
			URL:       link.URL,
			Status:    link.Status,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

func (link LinkcageService) Store(request models.LinkcageRequest) (status bool, err error) {
	tx := link.db.DB.Begin()
	timeNow := lib.GetTimeNow("timestime")

	reqType := &models.Linkcage{
		Name:      request.Name,
		URL:       request.URL,
		Status:    request.Status,
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	}

	dataLink, err := link.repository.Store(reqType, tx)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	//Begin Upload Banner
	bucket := os.Getenv("BUCKET_NAME")

	if request.Files[0].Filename != "" {
		for _, value := range request.Files {
			UUID := uuid.New()
			var destinationPath string
			bucketExist := link.minio.BucketExist(link.minio.Client(), bucket)

			pathSplit := strings.Split(value.Path, "/")
			sourcePath := fmt.Sprint(value.Path)
			destinationPath = pathSplit[1] + "/" +
				lib.GetTimeNow("year") + "/" +
				lib.GetTimeNow("month") + "/" +
				lib.GetTimeNow("day") + "/" +
				UUID.String() + "/" +
				pathSplit[2] + "/" + value.Filename

			if bucketExist {
				fmt.Println("Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				uploaded := link.minio.CopyObject(link.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			} else {
				fmt.Println("Not Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				link.minio.MakeBucket(link.minio.Client(), bucket, "")
				uploaded := link.minio.CopyObject(link.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			}

			files, err := link.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				return false, err
			}

			_, err = link.linkImg.Store(&models.LinkcageImage{
				LinkcageID: dataLink.ID,
				FileID:     files.ID,
				CreatedAt:  &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				return false, err
			}
		}
	}
	//End Upload Banner

	tx.Commit()
	return true, err
}

func (link LinkcageService) SetStatus(request *models.LinkcageRequest) (response bool, err error) {
	tx := link.db.DB.Begin()
	timeNow := lib.GetTimeNow("timestime")

	newStatus := &models.LinkcageRequest{
		ID:        request.ID,
		Status:    request.Status,
		UpdatedAt: &timeNow,
	}

	_, err = link.repository.SetStatus(newStatus, tx)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, err
}

func (link LinkcageService) GetActive() (responses []models.LinkcageResponse, err error) {
	links, err := link.repository.GetActive()
	if err != nil {
		return responses, err
	}

	for _, linkData := range links {
		var mini_link fileModel.FileManagerResponseUrl

		if linkData.Filename != "" {
			mini_link, err = link.filemanager.GetFile(fileModel.FileManagerRequest{
				Subdir:   linkData.Path,
				Filename: linkData.Filename,
			})

			if err != nil {
			}
		}

		responses = append(responses, models.LinkcageResponse{
			ID:     linkData.ID,
			Name:   linkData.Name,
			URL:    linkData.URL,
			Banner: mini_link.MinioPath,
			Status: linkData.Status,
		})
	}

	return responses, err
}

func (link LinkcageService) Delete(request *models.LinkcageRequest) (response bool, err error) {
	tx := link.db.DB.Begin()
	timeNow := lib.GetTimeNow("timestime")

	deleteLink := &models.LinkcageRequest{
		ID:        request.ID,
		UpdatedAt: &timeNow,
	}

	_, err = link.repository.Delete(deleteLink, tx)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, err
}

func (link LinkcageService) GetOne(request models.LinkcageRequest) (response models.LinkcageResponse, status bool, err error) {
	dataLink, err := link.repository.GetOne(&request)
	if err != nil {
		return response, false, err
	}

	if dataLink.ID != 0 {
		imgReq := &models.LinkcageRequest{
			ID: request.ID,
		}

		img, err := link.linkImg.GetLinkImage(imgReq)

		if err != nil {
			return response, false, err
		}

		response = models.LinkcageResponse{
			ID:     dataLink.ID,
			Name:   dataLink.Name,
			URL:    dataLink.URL,
			Files:  img,
			Status: dataLink.Status,
		}

		return response, true, err
	}

	return response, false, err
}

func (link LinkcageService) Update(request *models.LinkcageRequest) (status bool, err error) {
	tx := link.db.DB.Begin()
	timeNow := lib.GetTimeNow("timestime")

	updateLink := &models.Linkcage{
		ID:        request.ID,
		Name:      request.Name,
		URL:       request.URL,
		Status:    request.Status,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}

	dataLink, err := link.repository.Update(updateLink, tx)

	//Begin Update Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if request.Files[0].Filename != "" {
		err = link.linkImg.DeleteFilesByID(request.ID, tx)

		if err != nil {
			tx.Rollback()
			return false, err
		}

		for _, value := range request.Files {
			var destinationPath string
			bucketExist := link.minio.BucketExist(link.minio.Client(), bucket)

			pathSplit := strings.Split(value.Path, "/")
			sourcePath := fmt.Sprint(value.Path)
			destinationPath = pathSplit[1] + "/" +
				lib.GetTimeNow("year") + "/" +
				lib.GetTimeNow("month") + "/" +
				lib.GetTimeNow("day") + "/" +
				pathSplit[2] + "/" + value.Filename

			if bucketExist {
				fmt.Println("Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				uploaded := link.minio.CopyObject(link.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			} else {
				fmt.Println("Not Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)

				link.minio.MakeBucket(link.minio.Client(), bucket, "")
				uploaded := link.minio.CopyObject(link.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			}

			files, err := link.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				return false, err
			}

			_, err = link.linkImg.Store(&models.LinkcageImage{
				LinkcageID: dataLink.ID,
				FileID:     files.ID,
				CreatedAt:  &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				return false, err
			}
		}
	}

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, err
}
