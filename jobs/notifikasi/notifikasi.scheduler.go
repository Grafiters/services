package notifikasi

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NotifikasiReminder(dbdefault *sql.DB, totalWorker, dbMaxConns, dbMaxIdleConns int) {
	fmt.Println("Notifikasi Remainder=>", time.Now())
	sQuery := `select id,maker_id from tasklists where CURRENT_DATE  = end_date - INTERVAL 1 DAY and task_type = 5 and UPPER(approval_status) = 'DISETUJUI' and id not in (select task_id from tasklist_notifikasis)`

	rows, err := dbdefault.Query(sQuery)

	if err != sql.ErrNoRows {
		for rows.Next() {
			var (
				id       int64
				maker_id string
			)
			if err := rows.Scan(&id, &maker_id); err != nil {
				log.Fatal(err)
			}

			sQuery := `INSERT INTO tasklist_notifikasis (tanggal,keterangan,status,jenis,task_id) VALUES (CURRENT_DATE(),?,0,?,?)`

			stmt, err := dbdefault.Prepare(sQuery)
			if err != nil {
				fmt.Println(err)
			}
			defer stmt.Close()
			fmt.Println(id)
			if _, err := stmt.Exec("Anda Mempunyai Task Baru", "TAKSLIST", id); err != nil {
				log.Fatal(err)
			}
		}
	}
}
