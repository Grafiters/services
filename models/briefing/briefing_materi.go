package models

type BriefingMateri struct {
	ID                int64
	BriefingID        int64
	ActivityID        int64
	SubActivityID     int64
	ProductID         int64
	TitleMateries     string
	JudulMateri       string
	RiskIssueCode     string
	RekomendasiMateri string
	MateriTambahan    string
	UpdatedAt         *string
	CreatedAt         *string
}

type BriefingMateriUpdate struct {
	ID                int64
	BriefingID        int64
	ActivityID        int64
	SubActivityID     int64
	ProductID         int64
	TitleMateries     string
	JudulMateri       string
	RiskIssueCode     string
	RekomendasiMateri string
	MateriTambahan    string
	UpdatedAt         *string
}

func (bm BriefingMateriUpdate) TableName() string {
	return "briefing_materis"
}
