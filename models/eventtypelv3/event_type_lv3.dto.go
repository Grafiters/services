package models

type EventTypeLv3Request struct {
	ID               int64   `json:"id"`
	IDEventTypeLv2   string  `json:"id_event_type_lv2"`
	KodeEventTypeLv3 string  `json:"kode_event_type_lv3"`
	EventTypeLv3     string  `json:"event_type_lv3"`
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

type EventTypeLv3Response struct {
	ID               int64   `json:"id"`
	IDEventTypeLv2   string  `json:"id_event_type_lv2"`
	KodeEventTypeLv3 string  `json:"kode_event_type_lv3"`
	EventTypeLv3     string  `json:"event_type_lv3"`
	Deskripsi        string  `json:"deskripsi"`
	Status           bool    `json:"status"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type IDEventTypeLv2 struct {
	IDEventTypeLv2 string `json:"id_event_type_lv2"`
}

type KodeEventType struct {
	KodeEventType string `json:"kode_event_type"`
}

func (p EventTypeLv3Request) ParseRequest() EventTypeLv3 {
	return EventTypeLv3{
		ID:               p.ID,
		IDEventTypeLv2:   p.IDEventTypeLv2,
		KodeEventTypeLv3: p.KodeEventTypeLv3,
		EventTypeLv3:     "",
		Deskripsi:        p.Deskripsi,
		Status:           p.Status,
	}
}

func (p EventTypeLv3Response) ParseResponse() EventTypeLv3 {
	return EventTypeLv3{
		ID:               p.ID,
		IDEventTypeLv2:   p.IDEventTypeLv2,
		KodeEventTypeLv3: p.KodeEventTypeLv3,
		EventTypeLv3:     p.EventTypeLv3,
		Deskripsi:        p.Deskripsi,
		Status:           p.Status,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func (ET3 EventTypeLv3Request) TableName() string {
	return "event_type_lv3"
}

func (ET3 EventTypeLv3Response) TableName() string {
	return "event_type_lv3"
}
