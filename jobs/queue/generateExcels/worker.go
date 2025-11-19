package generateExcels

import (
	"fmt"
	"riskmanagement/lib"
	"sync"

	minio "gitlab.com/golang-package-library/minio"
)

func Worker(minio *minio.Minio, db *lib.Database, queueJob chan Job, wg *sync.WaitGroup) {
	for job := range queueJob {
		if job.RetryCode == "00" {
			// Insert Request
			insertId, errInsertId := SaveRequestInfo(db, job)
			if errInsertId != nil {
				// queueJob <- job
				fmt.Println("errInsertId", errInsertId)
			}

			fmt.Println(insertId)

			fmt.Println("Masuk Worker", job.GenerateInfo)

			// Call the Generator function
			JsonData, GeneratorStatus, err := Generator(db, job)

			if err != nil {
				fmt.Println("Generator failed with error:", err)

				status, err := UpdateStatusDownload(db, job, insertId, "Error Get Data")
				if err != nil && !status {
					fmt.Println("Error Update Status", err)
				}
			}

			if GeneratorStatus {
				job.JSONData = JsonData

				isSuccess, errisSuccess := convertJSONToExcel(job)
				if errisSuccess != nil {
					// queueJob <- job
					fmt.Println("errisSuccess", errisSuccess)
					status, err := UpdateStatusDownload(db, job, insertId, "Error Generate File")
					if err != nil && !status {
						fmt.Println("Error Update Status", err)
					}
				}

				fmt.Println("isSuccess", isSuccess)

				if isSuccess {
					// job.ExcelPath = ""
					minioLink, fileName, statusUpload, errMinioLink := uploadToMinIO(minio, job)
					fmt.Println("minioLink", minioLink)

					if !statusUpload {
						// queueJob <- job
						fmt.Println("errMinioLink", errMinioLink)
						UpdateDataFailUpload(db, job, insertId, fileName, "Error Upload Data")
					} else {
						isSuccessUpdate, errIsSuccessUpdate := saveMinIOLinkToDB(db, job, minioLink, fileName, insertId)
						if errIsSuccessUpdate != nil {
							// queueJob <- job
							fmt.Println("errIsSuccessUpdate", errIsSuccessUpdate)
							// status, err := UpdateStatusDownload(db, job, insertId, "Gagal Convert")
						}
						if err != nil {
							fmt.Println("Error Update Status", err)
						}
						fmt.Println("isSuccessUpdate", isSuccessUpdate)
						status, err := UpdateStatusDownload(db, job, insertId, "Done")

						//
						if err != nil && !status {
							fmt.Println("Error Update Status", err)
						}

						// After FileReady download delete temporary file on server
						deleteStatus, errStatus := DeleteTempFile(job)
						if !deleteStatus {
							fmt.Println("Error delete file =>", errStatus)
						}

						fmt.Println("Delete file =>", deleteStatus)
					}
				}
			}
		} else if job.RetryCode == "03" {
			fmt.Println("Retry Upload")
			fmt.Println("Job =>", job)
			minioLink, fileName, statusUpload, errMinioLink := uploadToMinIO(minio, job)
			fmt.Println("minioLink", minioLink)

			if !statusUpload {
				// queueJob <- job
				fmt.Println("errMinioLink", errMinioLink)
				UpdateDataFailUpload(db, job, job.ID, fileName, "Error Upload Data")
			} else {
				fmt.Println("Berhasil")
				isSuccessUpdate, errIsSuccessUpdate := saveMinIOLinkToDB(db, job, minioLink, fileName, job.ID)
				if errIsSuccessUpdate != nil {
					// queueJob <- job
					fmt.Println("errIsSuccessUpdate", errIsSuccessUpdate)
					// status, err := UpdateStatusDownload(db, job, insertId, "Gagal Convert")
				}

				fmt.Println("isSuccessUpdate", isSuccessUpdate)
				status, err := UpdateStatusDownload(db, job, job.ID, "Done")

				if err != nil && !status {
					fmt.Println("Error Update Status", err)
				}
				// After FileReady download delete temporary file on server
				deleteStatus, errStatus := DeleteTempFile(job)
				if !deleteStatus {
					fmt.Println("Error delete file =>", errStatus)
				}

				fmt.Println("Delete file =>", deleteStatus)
			}
		} else if job.RetryCode == "01" {
			fmt.Println("Retry Generate")

			fmt.Println("Masuk Worker", job.GenerateInfo)

			status, errInsertId := UpdateRequestInfo(db, job)
			if errInsertId != nil {
				// queueJob <- job
				fmt.Println("errInsertId", errInsertId)
			}

			fmt.Println("Update Request =>", status)

			// Call the Generator function
			JsonData, GeneratorStatus, err := Generator(db, job)
			if err != nil {
				fmt.Println("Generator failed with error:", err)

				status, err := UpdateStatusDownload(db, job, job.ID, "Error Get Data")

				if err != nil && !status {
					fmt.Println("Error Update Status", err)
				}

			}

			if GeneratorStatus {
				// fmt.Println(JsonData)
				job.JSONData = JsonData

				isSuccess, errisSuccess := convertJSONToExcel(job)
				if errisSuccess != nil {
					// queueJob <- job
					fmt.Println("errisSuccess", errisSuccess)
					status, err := UpdateStatusDownload(db, job, job.ID, "Error Generate File")
					if err != nil && !status {
						fmt.Println("Error Update Status", err)
					}
				}

				fmt.Println("isSuccess", isSuccess)

				if isSuccess {
					// job.ExcelPath = ""
					minioLink, fileName, statusUpload, errMinioLink := uploadToMinIO(minio, job)
					fmt.Println("minioLink", minioLink)

					if !statusUpload {
						// queueJob <- job
						fmt.Println("errMinioLink", errMinioLink)
						UpdateDataFailUpload(db, job, job.ID, fileName, "Error Upload Data")
					} else {
						// fmt.Println("Berhasil")
						isSuccessUpdate, errIsSuccessUpdate := saveMinIOLinkToDB(db, job, minioLink, fileName, job.ID)
						if errIsSuccessUpdate != nil {
							// queueJob <- job
							fmt.Println("errIsSuccessUpdate", errIsSuccessUpdate)
							// status, err := UpdateStatusDownload(db, job, insertId, "Gagal Convert")
						}
						if err != nil {
							fmt.Println("Error Update Status", err)
						}
						fmt.Println("isSuccessUpdate", isSuccessUpdate)
						status, err := UpdateStatusDownload(db, job, job.ID, "Done")

						if err != nil && !status {
							fmt.Println("Error Update Status", err)
						}
						// After FileReady download delete temporary file on server
						deleteStatus, errStatus := DeleteTempFile(job)
						if !deleteStatus {
							fmt.Println("Error delete file =>", errStatus)
						}

						fmt.Println("Delete file =>", deleteStatus)
					}
				}
			}
		}

		if _, ok := <-queueJob; !ok {
			break
		}
	}

	wg.Done()
}
