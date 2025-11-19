package jobs

import (
	"database/sql"
	dio "riskmanagement/jobs/dio"
	notif "riskmanagement/jobs/notifikasi"
)

const (
	dbMaxIdleConns = 4
	dbMaxConns     = 100
	totalWorker    = 100
)

type JobsNeeded struct {
	DBdefault *sql.DB
}

func NewCronJob(DBdefault *sql.DB) *JobsNeeded {
	return &JobsNeeded{
		DBdefault: DBdefault,
	}
}

func (j JobsNeeded) ListJobs(m string) func() {
	switch m {
	// case "updateDataMasterSDM":
	// return j.updateDataMasterSDM()
	// case "updateDataDashboard":
	// 	return j.updateDataDashboard()
	// case "dioRemainder":
	// 	return j.dioRemainderBisnisReview()
	case "tasklistReminder":
		return j.NotifikasiTaskList()
	case "scheduler1":
		return j.dioRemainderBisnisReview()
	default:
		return func() {}
	}
}

func (j JobsNeeded) dioRemainderBisnisReview() func() {
	return func() {
		go dio.DioRemainder(j.DBdefault, totalWorker, dbMaxConns, dbMaxIdleConns)
	}
}

func (j JobsNeeded) NotifikasiTaskList() func() {
	return func() {
		go notif.NotifikasiReminder(j.DBdefault, totalWorker, dbMaxConns, dbMaxIdleConns)
	}
}
