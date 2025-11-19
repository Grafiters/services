package models

type BriefingMateriRequest struct {
	ID                int64   `json:"id"`
	BriefingID        int64   `json:"briefing_id"`
	ActivityID        int64   `json:"activity_id"`
	SubActivityID     int64   `json:"sub_activity_id"`
	ProductID         int64   `json:"product_id"`
	TitleMateries     string  `json:"title_materies"`
	JudulMateri       string  `json:"judul_materi"`
	RiskIssueCode     string  `json:"risk_issue_code"`
	RekomendasiMateri string  `json:"rekomendasi_materi"`
	MateriTambahan    string  `json:"materi_tambahan"`
	UpdatedAt         *string `json:"updated_at"`
	CreatedAt         *string `json:"created_at"`
}

type BriefingMateriRequests []map[string]interface{}

type BriefingMateriResponse struct {
	ID                int64   `json:"id"`
	BriefingID        int64   `json:"briefing_id"`
	ActivityID        int64   `json:"activity_id"`
	SubActivityID     int64   `json:"sub_activity_id"`
	ProductID         int64   `json:"product_id"`
	TitleMateries     string  `json:"title_materies"`
	JudulMateri       string  `json:"judul_materi"`
	RiskIssueCode     string  `json:"risk_issue_code"`
	RekomendasiMateri string  `json:"rekomendasi_materi"`
	MateriTambahan    string  `json:"materi_tambahan"`
	UpdatedAt         *string `json:"updated_at"`
	CreatedAt         *string `json:"created_at"`
}

type BriefingMateriResponses struct {
	ID                int64  `json:"id"`
	BriefingID        int64  `json:"briefing_id"`
	ActivityID        int64  `json:"activity_id"`
	ActivityText      string `json:"activity_text"`
	SubActivityID     int64  `json:"sub_activity_id"`
	SubActivityText   string `json:"sub_activity_text"`
	ProductID         int64  `json:"product_id"`
	ProductText       string `json:"product_text"`
	TitleMateries     string `json:"title_materies"`
	JudulMateri       string `json:"judul_materi"`
	RiskIssueCode     string `json:"risk_issue_code"`
	RekomendasiMateri string `json:"rekomendasi_materi"`
	MateriTambahan    string `json:"materi_tambahan"`
}

// func (p BriefingMateriRequest) ParseRequest() BriefingMateri {
// 	return BriefingMateri{
// 		ID:                p.ID,
// 		BriefingID:        p.BriefingID,
// 		ActivityID:        p.ActivityID,
// 		SubActivityID:     p.SubActivityID,
// 		ProductID:         p.ProductID,
// 		JudulMateri:       p.JudulMateri,
// 		RekomendasiMateri: p.RekomendasiMateri,
// 		MateriTambahan:    p.MateriTambahan,
// 	}
// }

// func (p BriefingMateriResponse) ParseRequest() BriefingMateri {
// 	return BriefingMateri{
// 		ID:                p.ID,
// 		BriefingID:        p.BriefingID,
// 		ActivityID:        p.ActivityID,
// 		ProductID:         p.ProductID,
// 		SubActivityID:     p.SubActivityID,
// 		JudulMateri:       p.JudulMateri,
// 		RekomendasiMateri: p.RekomendasiMateri,
// 		MateriTambahan:    p.MateriTambahan,
// 		CreatedAt:         p.CreatedAt,
// 		UpdatedAt:         p.UpdatedAt,
// 	}
// }

func (bm BriefingMateriRequest) TableName() string {
	return "briefing_materis"
}

func (bm BriefingMateriResponse) TableName() string {
	return "briefing_materis"
}
