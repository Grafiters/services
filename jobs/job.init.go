package jobs

import (
	"database/sql"
	"fmt"
	"log"
	jobs "riskmanagement/lib/jobs"
	"time"

	"github.com/robfig/cron/v3"
)

type ParameterizeJobs struct {
	Time      string
	Method    string
	Action    string
	Flag      string
	Status    string
	UpdatedAt time.Time
}

func JobsInit(dbdefault *sql.DB) {

	/**
	 * * Concurrent Proccess for parameterize Jobs
	**/
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})

	go func() {
		jTime, _ := time.LoadLocation("Asia/Jakarta")
		c := cron.New(cron.WithLocation(jTime))
		job := jobs.NewCronJob(dbdefault)
		var idJob = make(map[string]cron.EntryID)

		for {
			select {
			case <-ticker.C:
				data, _ := parameterizeJobs(dbdefault)
				fmt.Println(data)
				for _, param := range data {
					switch param.Action {
					case "add":
						log.Printf("Add job %s \n", idJob[param.Method])
						fmt.Println("Add job ", idJob[param.Method])
						idJob[param.Method], _ = c.AddFunc(param.Time, job.ListJobs(param.Method))
						_ = parameterizeJobsFlag(dbdefault, "RUN", param.Method) // update flag
					case "remove":
						log.Printf("Remove job %s \n", idJob[param.Method])
						fmt.Println("Remove job ", idJob[param.Method])
						c.Remove(idJob[param.Method])
						_ = parameterizeJobsFlag(dbdefault, "NOT RUN", param.Method) // update flag
					default:
						continue
					}

				}

				c.Start()
				printCronEntries(c.Entries())
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func printCronEntries(cronEntries []cron.Entry) {
	log.Printf("Cron Info: %+v\n", cronEntries)
}

func parameterizeJobs(db *sql.DB) (pjs []ParameterizeJobs, err error) {
	query := "SELECT * FROM ref_cronjobs WHERE flag = 1"
	rows, err := db.Query(query)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
		// filename, function, line := helper.WhereAmI()
		// helper.CreateLogErrorToDB(db, filename, function, line, "Parameterize jobs gagal", fmt.Sprintf("%v", err))
		return pjs, err
	}

	pj := ParameterizeJobs{}
	for rows.Next() {
		_ = rows.Scan(
			&pj.Time,
			&pj.Method,
			&pj.Action,
			&pj.Flag,
			&pj.Status,
			&pj.UpdatedAt,
		)

		pjs = append(pjs, pj)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		// filename, function, line := helper.WhereAmI()
		// helper.CreateLogErrorToDB(db, filename, function, line, "Parameterize jobs gagal", fmt.Sprintf("%v", err))
		return pjs, err
	}

	return pjs, nil
}

func parameterizeJobsFlag(db *sql.DB, status, method string) (err error) {
	tx, _ := db.Begin()
	if err != nil {
		fmt.Println(err)
		// filename, function, line := helper.WhereAmI()
		// helper.CreateLogErrorToDB(db, filename, function, line, "Parameterize jobs flag gagal", fmt.Sprintf("%v", err))
		return err
	}

	if _, err = tx.Exec("UPDATE ref_cronjobs SET flag = 2, status = ?, updated_at = NOW() WHERE method = ?", status, method); err != nil {
		tx.Rollback()

		fmt.Println(err)
		// filename, function, line := helper.WhereAmI()
		// helper.CreateLogErrorToDB(db, filename, function, line, "Parameterize jobs flag gagal", fmt.Sprintf("%v", err))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ParameterizeJobsFlagRun(db *sql.DB) (err error) {
	tx, _ := db.Begin()
	if err != nil {
		fmt.Println(err)
		// filename, function, line := helper.WhereAmI()
		// helper.CreateLogErrorToDB(db, filename, function, line, "Parameterize jobs flag gagal", fmt.Sprintf("%v", err))
		return err
	}

	if _, err = tx.Exec("UPDATE ref_cronjobs SET flag=1"); err != nil {
		tx.Rollback()

		fmt.Println(err)
		// filename, function, line := helper.WhereAmI()
		// helper.CreateLogErrorToDB(db, filename, function, line, "Parameterize jobs flag gagal", fmt.Sprintf("%v", err))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
