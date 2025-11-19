package models

type EventTypeLv1Request struct {
	ID            int64   `json:"id"`
	KodeEventType string  `json:"kode_event_type"`
	EventType     string  `json:"event_type"`
	Deskripsi     string  `json:"deskripsi"`
	Status        bool    `json:"status"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type EventTypeLv1Response struct {
	ID            int64   `json:"id"`
	KodeEventType string  `json:"kode_event_type"`
	EventType     string  `json:"event_type"`
	Deskripsi     string  `json:"deskripsi"`
	Status        bool    `json:"status"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type KodeEventType struct {
	KodeEventType string `json:"kode_event_type"`
}

func (p EventTypeLv1Request) ParseRequest() EventTypeLv1 {
	return EventTypeLv1{
		ID:            p.ID,
		KodeEventType: p.KodeEventType,
		EventType:     p.EventType,
		Deskripsi:     p.Deskripsi,
		Status:        p.Status,
	}
}

func (p EventTypeLv1Response) ParseResponse() EventTypeLv1 {
	return EventTypeLv1{
		ID:            p.ID,
		KodeEventType: p.KodeEventType,
		EventType:     p.EventType,
		Deskripsi:     p.Deskripsi,
		Status:        p.Status,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (ET1 EventTypeLv1Request) TableName() string {
	return "event_type_lv1"
}

func (ET1 EventTypeLv1Response) TableName() string {
	return "event_type_lv1"
}
