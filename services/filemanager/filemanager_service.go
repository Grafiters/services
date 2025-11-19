package filemanager

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"riskmanagement/lib"
	models "riskmanagement/models/filemanager"
	"strings"

	"github.com/google/uuid"

	minio_github "github.com/minio/minio-go/v7"

	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"
)

type FileManagerDefinition interface {
	MakeUpload(request models.FileManagerRequest) (responses models.FileManagerResponse, err error)
	GetFile(request models.FileManagerRequest) (responses models.FileManagerResponseUrl, err error)
	RemoveObject(request models.FileManagerRequest) (response bool, err error)
	ReadFile(request models.FileManagerRequest) (string, error)
}

type FileManagerService struct {
	logger logger.Logger
	minio  minio.Minio
}

func NewFileManagerService(
	minio minio.Minio,
	logger logger.Logger,
) FileManagerDefinition {
	return FileManagerService{
		logger: logger,
		minio:  minio,
	}
}

// GetFile implements FileManagerDefinition
func (filemanager FileManagerService) GetFile(request models.FileManagerRequest) (responses models.FileManagerResponseUrl, err error) {
	bucket, err := lib.GetVarEnv("BUCKET_NAME")
	if err != nil {
		return responses, err
	}

	subdir := request.Subdir
	filename := request.Filename

	var minioPath string

	bucketExist := filemanager.minio.BucketExist(filemanager.minio.Client(), bucket)
	if bucketExist {
		preSign := filemanager.minio.SignUrl(filemanager.minio.Client(), bucket, subdir, filename)
		minioPath = fmt.Sprint(preSign)
		// fmt.Println(filename)
		// fmt.Println("presing url", preSign)
		// fmt.Println(minioPath)

		// resp, err := http.Get("http://172.18.53.99:9000/riskmanagement/riskindicator/2023/5/30/2023/KRID_Pinjaman%20Briguna%20Kawan%20Skim%20Baru%20yang%20masih%20memiliki%20fasilitas%20pinjaman%20pekerja%20lainnya%20%28Pitung%20dan%20atau%20Briguna%20Kawan%20Skim%20Lama%29.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=minio%2F20230605%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230605T020755Z&X-Amz-Expires=1000&X-Amz-SignedHeaders=host&X-Amz-Signature=fd63131761c1b6a843862ee4a1ea57daba452a2832f178bf6659fc503f7ddc35")
		resp, err := http.Get(minioPath)
		if err != nil {
			fmt.Println("Error retrieving image:", err)
			return responses, err
		}
		defer resp.Body.Close()

		// Read the image data
		imageData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading image data:", err)
			return responses, err
		}

		// Encode the image data to base64
		base64Str := base64.StdEncoding.EncodeToString(imageData)

		// Print the base64 string
		// fmt.Println("url=>", base64Str)
		// Decode the Base64 string to binary data
		fileData, err := base64.StdEncoding.DecodeString(base64Str)
		if err != nil {
			log.Fatal("Error decoding Base64:", err)
		}

		contentType := http.DetectContentType(fileData)

		// Create a download link with the data URI
		downloadLink := fmt.Sprintf(`data:`+contentType+`;base64,%s`, base64Str)

		// fmt.Println("download Link=>", downloadLink)

		responses := models.FileManagerResponseUrl{
			MinioPath:  downloadLink,
			PreSignUrl: preSign,
			// Base64:     base64Str,
		}

		return responses, err
	} else {
		fmt.Println("Not Exist")
		return responses, err
	}
}

// MakeUpload implements FileManagerDefinition
func (filemanager FileManagerService) MakeUpload(request models.FileManagerRequest) (responses models.FileManagerResponse, err error) {
	var allowedExtensions = map[string]bool{
		"csv":  true,
		"pdf":  true,
		"xlsx": true,
		"xls":  true,
		"jpg":  true,
		"jpeg": true,
		"jfif": true,
		"doc":  true,
		"docx": true,

		// Tambahan
		"png": true,
	}

	ext := filepath.Ext(request.File.Filename)
	extension := strings.ToLower(strings.TrimPrefix(ext, "."))

	ext2 := strings.Contains(request.File.Filename, ".")
	parts := strings.Split(request.File.Filename, ".")

	// println("extension ==>", extension)
	println("extension2 ==>", ext2)

	println("len =>", len(parts))

	// if ext2 && len(parts) > 2 {
	// 	return responses, errors.New("Double extension not allowed. Please upload a file with a single extension.")
	// } else {
	if !allowedExtensions[extension] {
		return responses, errors.New("Invalid file type. Only PDF, Excel, and JPG files are allowed.")
	} else {
		var minioPath string
		bucketName, err := lib.GetVarEnv("BUCKET_NAME")
		if err != nil {
			return responses, err
		}

		src, err := request.File.Open()
		if err != nil {
			filemanager.logger.Zap.Info(err)
		}
		defer src.Close()

		dir, err := os.Getwd()

		if err != nil {
			filemanager.logger.Zap.Error(err)
		}

		// Path direktori bucket
		bucketDir := filepath.Join(dir, "storage", "uploads", "temp")

		// Buat folder jika belum ada
		if _, err := os.Stat(bucketDir); os.IsNotExist(err) {
			err = os.MkdirAll(bucketDir, os.ModePerm)
			if err != nil {
				filemanager.logger.Zap.Error(err)
				return responses, err
			}
		}

		// Simpan file ke direktori lokal sementara
		fileLocation := filepath.Join(bucketDir, request.File.Filename)

		// fmt.Println(fileLocation)
		dst, err := os.Create(fileLocation)
		if err != nil {
			filemanager.logger.Zap.Error(err)
			return responses, err
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			dst.Close()
			filemanager.logger.Zap.Error("Error Copy file :", err)
		}

		// pathFile := "storage/uploads/" + bucketName + "/" + request.Filename

		bucketExist := filemanager.minio.BucketExist(filemanager.minio.Client(), bucketName)

		extension := ""

		uuid := uuid.New()
		minioPath = "tmp/" + request.Subdir + "/" + lib.GetTimeNow("year") + "/" + lib.GetTimeNow("month") + "/" + lib.GetTimeNow("day") + "/" + uuid.String() + "/" + request.File.Filename

		if bucketExist {
			dataFile, err := os.Open(fileLocation)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Data File", dataFile)

			contentType, err := GetFileContentType(dataFile)
			if err != nil {
				filemanager.logger.Zap.Error("Error Get Content", err)
			}

			_, err = filemanager.minio.UploadObject(filemanager.minio.Client(), bucketName, minioPath, fileLocation, contentType)
			if err != nil {
				filemanager.logger.Zap.Error("Error Upload Minio", err)
			}

			extension = contentType

			dataFile.Close()

			if err = os.Remove(fileLocation); err != nil {
				filemanager.logger.Zap.Error("Error removing temp file", err)
			}
		} else {
			filemanager.minio.MakeBucket(filemanager.minio.Client(), bucketName, "")

			dataFile, err := os.Open(fileLocation)
			if err != nil {
				filemanager.logger.Zap.Error("Error create bucket", err)
			}

			contentType, err := GetFileContentType(dataFile)
			if err != nil {
				filemanager.logger.Zap.Error("Error Get Content", err)
			}

			_, err = filemanager.minio.UploadObject(filemanager.minio.Client(), bucketName, minioPath, fileLocation, contentType)
			if err != nil {
				filemanager.logger.Zap.Error("Error upload minio", err)
			}

			extension = contentType

			dataFile.Close()

			if err = os.Remove(fileLocation); err != nil {
				filemanager.logger.Zap.Error("Error removing temp file", err)
			}
		}

		fileResponse := models.FileManagerResponse{
			Filename:  request.File.Filename,
			Path:      minioPath,
			Extension: extension,
			Size:      fmt.Sprint(request.File.Size),
		}

		return fileResponse, err
	}
	// }

}

// RemoveObject implements FileManagerDefinition
func (filemanager FileManagerService) RemoveObject(request models.FileManagerRequest) (response bool, err error) {
	bucket, err := lib.GetVarEnv("BUCKET_NAME")
	if err != nil {
		return false, err
	}
	objectName := request.ObjectName

	bucketExist := filemanager.minio.BucketExist(filemanager.minio.Client(), bucket)
	if bucketExist {
		remove := filemanager.minio.RemoveObject(filemanager.minio.Client(), bucket, objectName)

		if remove {
			return true, err
		} else {
			return false, err
		}
	} else {
		return false, err
	}
}

func GetFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func (filemanager FileManagerService) ReadFile(request models.FileManagerRequest) (string, error) {
	bucketName, err := lib.GetVarEnv("BUCKET_NAME")
	if err != nil {
		return "", err
	}

	// Check if the bucket exists
	exists := filemanager.minio.BucketExist(filemanager.minio.Client(), bucketName)
	if !exists {
		return "", fmt.Errorf("bucket %s does not exist", bucketName)
	}
	// Download file from MinIO
	filepath := "/storage/uploads/riskmanagement/" + request.Filename
	err = filemanager.minio.MinioClient.FGetObject(context.Background(), bucketName, request.ObjectName, filepath, minio_github.GetObjectOptions{})
	// success := filemanager.minio.GetObject(filemanager.minio.Client(), bucketName, request.ObjectName, "")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("filename : ", request.Filename)
	fmt.Println("objectname : ", request.ObjectName)
	return filepath, nil
}
