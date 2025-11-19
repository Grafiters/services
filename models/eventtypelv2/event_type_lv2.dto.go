package models

type EventTypeLv2Request struct {
	ID               int64   `json:"id"`
	IDEventTypeLv1   string  `json:"id_event_type_lv1"`
	KodeEventTypeLv2 string  `json:"kode_event_type_lv2"`
	EventTypeLv2     string  `json:"event_type_lv2"`
	Deskripsi        string  `json:"deskripsi"`
	Status           bool    `json:"status"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type EventTypeLv2Response struct {
	ID               int64   `json:"id"`
	IDEventTypeLv1   string  `json:"id_event_type_lv1"`
	KodeEventTypeLv2 string  `json:"kode_event_type_lv2"`
	EventTypeLv2     string  `json:"event_type_lv2"`
	Deskripsi        string  `json:"deskripsi"`
	Status           bool    `json:"status"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type KodeEventType struct {
	KodeEventType string `json:"kode_event_type"`
}

type IDEventTypeLv1 struct {
	IDEventTypeLv1 string `json:"id_event_type_lv1"`
}

func (p EventTypeLv2Request) ParseRequest() EventTypeLv2 {
	return EventTypeLv2{
		ID:               p.ID,
		IDEventTypeLv1:   p.IDEventTypeLv1,
		KodeEventTypeLv2: p.KodeEventTypeLv2,
		EventTypeLv2:     "",
		Deskripsi:        p.Deskripsi,
		Status:           p.Status,
	}
}

func (p EventTypeLv2Response) ParseResponse() EventTypeLv2 {
	return EventTypeLv2{
		ID:               p.ID,
		IDEventTypeLv1:   p.IDEventTypeLv1,
		KodeEventTypeLv2: p.KodeEventTypeLv2,
		EventTypeLv2:     p.EventTypeLv2,
		Deskripsi:        p.Deskripsi,
		Status:           p.Status,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func (ET2 EventTypeLv2Request) TableName() string {
	return "event_type_lv2"
}

func (ET2 EventTypeLv2Response) TableName() string {
	return "event_type_lv2"
}
