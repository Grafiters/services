package materi

import (
	"fmt"
	"os"
	"reflect"
	"riskmanagement/lib"
	models "riskmanagement/models/materi"
	repository "riskmanagement/repository/materi"
	"strings"

	requestFile "riskmanagement/models/files"
	fileRepo "riskmanagement/repository/files"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"
)

var (
	UUID = uuid.NewString()
)

type MateriDefinition interface {
	// GetAll() (responses []models.MateriFilesResponse, err error)
	GetAll() (responses []models.MateriAllResponse, err error)
	Store(request *models.MateriRequest) (status bool, err error)
	Delete(request models.MateriFilesRequest) (status bool, err error)
	GetMateriByActivityAndProduct(request models.GetMateriByActivityAndProductRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error)
	GetVerifikasiMateri(request models.GetMateriVerifikasiRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error)
}

type MateriService struct {
	db         lib.Database
	minio      minio.Minio
	logger     logger.Logger
	repository repository.MateriDefinition
	filesRepo  fileRepo.FilesDefinition
}

func NewMateriService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	repository repository.MateriDefinition,
	filesrepo fileRepo.FilesDefinition,
) MateriDefinition {
	return MateriService{
		db:         db,
		minio:      minio,
		logger:     logger,
		repository: repository,
		filesRepo:  filesrepo,
	}
}

// Delete implements MateriDefinition
func (materi MateriService) Delete(request models.MateriFilesRequest) (status bool, err error) {
	// tx := materi.db.DB.Begin()
	// status, err = materi.repository.Delete(request)
	// if err != nil {
	// 	tx.Rollback()
	// 	materi.logger.Zap.Error(err)
	// 	return false, err
	// }

	// bucket := os.Getenv("BUCKET_NAME")
	// ok := materi.minio.RemoveObject(minio.Minio.Client(), bucket, request)
	// err = materi.filesRepo.Delete(request.FilesID, tx)
	// if err != nil {
	// 	tx.Rollback()
	// 	materi.logger.Zap.Error(err)
	// 	return false, err
	// }
	// tx.Commit()
	// return true, err

	panic("undefined")
}

// GetAll implements MateriDefinition
func (materi MateriService) GetAll() (responses []models.MateriAllResponse, err error) {
	// return materi.repository.GetAllMateriFiles(1)
	return materi.repository.GetAll()
}

// Store implements MateriDefinition
func (materi MateriService) Store(request *models.MateriRequest) (status bool, err error) {
	timenow := lib.GetTimeNow("timestime")
	tx := materi.db.DB.Begin()
	bucket := os.Getenv("BUCKET_NAME")

	reqMateri := &models.Materi{
		Name:      request.Name,
		CreatedAt: &timenow,
	}

	dataMateri, err := materi.repository.Store(reqMateri, tx)

	if err != nil {
		tx.Rollback()
		materi.logger.Zap.Error(err)
		return false, err
	}

	fmt.Println("dataMateri", dataMateri.ID)

	for _, value := range request.Files {
		var destinationPath string
		bucketExist := materi.minio.BucketExist(materi.minio.Client(), bucket)

		pathSplit := strings.Split(value.Path, "/")
		sourcePath := fmt.Sprint(value.Path)
		destinationPath = pathSplit[1] + "/" +
			lib.GetTimeNow("year") + "/" +
			lib.GetTimeNow("month") + "/" +
			lib.GetTimeNow("day") + "/" +
			pathSplit[2] + "/" + value.Filename

		fmt.Println("bucketExist=>", reflect.TypeOf(bucketExist))

		/*destinationPath = "materi/" +
			lib.GetTimeNow("year") + "/" +
			lib.GetTimeNow("month") + "/" +
			lib.GetTimeNow("day")

		destinationPath = destinationPath + "/" + value.Filename*/

		if bucketExist {
			fmt.Println("Exist")
			fmt.Println(bucket)
			fmt.Println(sourcePath)
			fmt.Println(destinationPath)
			uploaded := materi.minio.CopyObject(materi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
			// uploaded := materi.minio.PutObject(materi.minio.MinioClient, bucket, destinationPath, sourcePath)

			fmt.Println(uploaded)
		} else {
			fmt.Println("Not Exist")
			fmt.Println(bucket)
			fmt.Println(sourcePath)
			fmt.Println(destinationPath)
			materi.minio.MakeBucket(materi.minio.Client(), bucket, "")
			uploaded := materi.minio.CopyObject(materi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
			// uploaded := materi.minio.PutObject(materi.minio.MinioClient, bucket, destinationPath, sourcePath)
			fmt.Println(uploaded)
		}

		files, err := materi.filesRepo.Store(&requestFile.Files{
			Filename:  value.Filename,
			Path:      destinationPath,
			Extension: value.Extension,
			Size:      value.Size,
			CreatedAt: &timenow,
		}, tx)

		if err != nil {
			tx.Rollback()
			materi.logger.Zap.Error(err)
			return false, err
		}

		_, err = materi.repository.StoreMateriFiles(&models.MateriRequest{
			MateriID: dataMateri.ID, //ID belum kebaca
			FilesID:  files.ID,
		}, tx)

		if err != nil {
			tx.Rollback()
			materi.logger.Zap.Error(err)
			return false, err
		}

	}

	tx.Commit()
	return true, err
}

func (materi MateriService) GetMateriByActivityAndProduct(request models.GetMateriByActivityAndProductRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error) {
	dataMateri, err := materi.repository.GetMateriByActivityAndProduct(&request)

	responses = dataMateri

	return responses, err
}

func (materi MateriService) GetVerifikasiMateri(request models.GetMateriVerifikasiRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error) {
	dataMateri, err := materi.repository.GetVerifikasiMateri(&request)

	responses = dataMateri

	return responses, err
}
