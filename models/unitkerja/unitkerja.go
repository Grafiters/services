package models

type UnitKerja struct {
	ID         int64
	KodeUker   int64
	NamaUker   string
	KodeCabang int64
	NamaCabang string
	KanwilID   int64
	KodeKanwil string
	Kanwil     string
	Status     int64
	CreatedAt  *string
	UpdatedAt  *string
}
