package generateExcels

import (
	"os"
	"riskmanagement/lib"
	"sync"

	minio "gitlab.com/golang-package-library/minio"
)

// func Start(minio *minio.Minio, db *lib.Database, jsonData []map[string]interface{}, columnNames []string, headers map[string]string, generateInfo GenerateInfo, fileName string) {
func Start(minio *minio.Minio, db *lib.Database, columnNames []string, headers map[string]string, generateInfo GenerateInfo, fileName string, RetryCode string) {

	maxWorkers := 1
	queueJob := make(chan Job, 1) // the number is accoarding to how many action should run

	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1) // the number is accoarding to how many action should run
		go Worker(minio, db, queueJob, &wg)
	}

	// // Add jobs to the queue
	job := Job{
		ID:       generateInfo.ID,
		ReportId: generateInfo.ReportId,
		// JSONData:     jsonData,
		ColumnNames:  columnNames,
		Headers:      headers,
		ExcelPath:    os.Getenv("GeneratePath") + fileName + `.xlsx`, //dev staging
		FileName:     fileName,
		GenerateInfo: generateInfo,
		RetryCode:    RetryCode,
	}
	queueJob <- job

	// wg.Wait()
	close(queueJob)
}
