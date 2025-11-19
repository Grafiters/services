package models

type MstKriteria struct {
	ID           int64
	KodeCriteria string
	Criteria     string
	Restruck     int8
	Status       int8
	ActiveDate   *string
	InactiveDate *string
	CreatedAt    *string
	CreatedBy    string
	CreatedDesc  string
	EnabledDate  *string
	EnabledBy    string
	EnabledDesc  string
	DisabledDate *string
	DisabledBy   string
	DisabledDesc string
}

type MstKriteriaHistory struct {
	ID           int64
	IdCriteria   int64
	KodeCriteria string
	Criteria     string
	Restruck     int8
	Status       int8
	ActiveDate   *string
	CreatedAt    *string
	CreatedBy    *string
	CreatedDesc  *string
}
