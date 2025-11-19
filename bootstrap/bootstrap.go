package bootstrap

import (
	"context"

	"database/sql"
	"fmt"
	"log"
	"riskmanagement/consumers"
	"riskmanagement/controllers"
	cronjob "riskmanagement/jobs"
	"riskmanagement/lib"
	env "riskmanagement/lib/env"
	"riskmanagement/middlewares"
	"riskmanagement/repository"
	"riskmanagement/routes"
	"riskmanagement/services"
	"time"

	minioEnv "gitlab.com/golang-package-library/env"
	logger "gitlab.com/golang-package-library/logger"
	storageMinio "gitlab.com/golang-package-library/minio"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module exported for initializing application
var Module = fx.Options(
	routes.Module,
	lib.Module,
	controllers.Module,
	services.Module,
	repository.Module,
	middlewares.Module,
	consumers.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env env.Env,
	logger logger.Logger,
	middlewares middlewares.Middlewares,
	database lib.Database,
	databases lib.Databases,
	minioEnv minioEnv.Env,
	minio storageMinio.Minio,
	consumer consumers.Consumers,
) {
	// if err := rabbitmq.ConnectMQ(); err != nil {
	// 	log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	// }

	// if err := rabbitmq.CloseMQ(); err != nil {
	// 	log.Fatalf("Failed to Close Connection to RabbitMQ: %v", err)
	// }

	conn, _ := database.DB.DB()
	connection := databases.DB
	status := minio.MinioClient.IsOnline()
	logger.Zap.Info("Minio Status : ", status)

	// Create a channel to signal job restart
	// restartJobsChan := make(chan struct{})

	// Create a goroutine for periodic checking and job restart
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute): // Adjust the time interval as needed
				idleConnections := conn.Stats().Idle
				if idleConnections >= 50 {
					logger.Zap.Info("Restarting Jobs due to idle connections")
					go cronjob.JobsInit(connection) // Restart the jobs
				}
			}
		}
	}()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("---------------------")
			logger.Zap.Info("------- CLEAN -------")
			logger.Zap.Info("---------------------")

			conn.SetMaxOpenConns(100)
			conn.SetMaxIdleConns(50)

			connection.SetMaxOpenConns(100)
			connection.SetMaxIdleConns(50)

			/**
			* * Concurrent Proccess for parameterize Jobs
			**/
			go cronjob.JobsInit(connection)
			// go cronjob.ParameterizeJobsFlagRun(connection)
			/**
			* * Concurrent Proccess for parameterize Jobs
			**/

			go func() {
				middlewares.Setup()
				routes.Setup()
				err := consumer.Connect()
				if err != nil {
					log.Panicf("Failed to connect to RabbitMQ: %v", err)
				}
				_ = handler.Gin.Run(env.ServerPort)
			}()

			ticker := time.NewTicker(24 * time.Hour)

			// Jalankan scheduler pada setiap interval ticker
			go func() {
				for {
					select {
					case <-ticker.C:
						schedulerTL(connection)
					}
				}
			}()

			// now := time.Now()
			// Calculate the duration until the next midnight.
			// nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Sub(now)

			// Create a timer that triggers at the next midnight.
			// timer := time.NewTimer(nextMidnight)

			// timer := time.NewTicker(2 * time.Minute) //if in local

			// // Run Scheduller Update status pgs
			// go func() {
			// 	for {
			// 		select {
			// 		case <-timer.C:
			// 			schedulerUpdatePGS(connection)
			// 		}
			// 	}
			// }()

			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			err := conn.Close()
			if err != nil {
				logger.Zap.Error("Error while closing database connection", zap.Error(err))
			}

			err = connection.Close()
			if err != nil {
				logger.Zap.Error("Error while closing database connection", zap.Error(err))
			}

			err = consumer.Close()
			if err != nil {
				logger.Zap.Error("Error while closing RabbitMQ connection", zap.Error(err))
			}
			return nil
		},
	})
}

func schedulerTL(dbdefault *sql.DB) {
	fmt.Println("Notifikasi Reminder=>", time.Now())

	sQuery := `SELECT vptl.pic_id 'pic_id', v.no_pelaporan 'no_pelaporan', v.BRDESC 'branch' FROM verifikasi v 
	JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = v.id 
	WHERE CURRENT_DATE = tanggal_target_selesai - INTERVAL 1 DAY AND NOT action = 'SELESAI'`

	rows, err := dbdefault.Query(sQuery)

	if err != sql.ErrNoRows {
		for rows.Next() {
			var (
				pic_id       string
				no_pelaporan string
				branch       string
			)

			if err := rows.Scan(&pic_id, &no_pelaporan, &branch); err != nil {
				log.Fatal(err)
			}

			sQuery := `INSERT INTO tasklist_notifikasis (tanggal,keterangan,status,jenis,task_id,receiver,uker) VALUES (CURRENT_DATE(),?,0,?,?,?,?)`

			stmt, err := dbdefault.Prepare(sQuery)
			if err != nil {
				fmt.Println(err)
			}
			defer stmt.Close()
			if _, err := stmt.Exec("Ada 1 verifikasi yang perlu ditindaklanjuti.", "Tindak Lanjut", no_pelaporan, pic_id, branch); err != nil {
				log.Fatal(err)
				dbdefault.Exec("INSERT INTO job_logs (job_id, process, status, status_description, created_at, updated_at) VALUES (1, 'Insert notifikasi tasklist', 'Fail', 'Gagal insert notifikasi tasklist', CURRENT_DATE(), CURRENT_DATE())")
			} else {
				dbdefault.Exec("INSERT INTO job_logs (job_id, process, status, status_description, created_at, updated_at) VALUES (1, 'Insert notifikasi tasklist', 'Success', 'Berhasil insert notifikasi tasklist', CURRENT_DATE(), CURRENT_DATE())")
			}
		}
	}
}

// func schedulerUpdatePGS(engineDB *sql.DB) {
// 	fmt.Println("Update Pgs =>", time.Now())

// 	query := `UPDATE pgs_user pu SET pu.status = ?, pu.action = ? WHERE CURDATE() > pu.periode_akhir`
// 	// fmt.Println(query)

// 	stmt, err := engineDB.Prepare(query)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer stmt.Close()

// 	if _, err := stmt.Exec("03a", "Non Aktive"); err != nil {
// 		fmt.Println(err)
// 		engineDB.Exec("INSERT INTO job_logs (job_id, process, status, status_description, created_at, updated_at) VALUES (4, 'Update Status PGS', 'Fail', 'Gagal Update status PGS', CURRENT_DATE(), CURRENT_DATE())")
// 	} else {
// 		fmt.Println("Berhasil")
// 		engineDB.Exec("INSERT INTO job_logs (job_id, process, status, status_description, created_at, updated_at) VALUES (4, 'Update Status PGS', 'Success', 'Berhasil Update status PGS', CURRENT_DATE(), CURRENT_DATE())")
// 	}
// }
