package dto

import "time"

type ListIndicatorHeader struct {
	ID          string    `json:"id"`
	IndicatorID int       `json:"indicator_id"`
	HeaderKey   string    `json:"header_key"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListIndicatorHeaderResponse struct {
	List       []ListIndicatorHeader `json:"list"`
	Pagination Pagination            `json:"pagination"`
}

type UkerDevisionFilter struct {
	IndicatorID int `json:"indicator_id"`
}

type UkerDevision struct {
	IndicatorID  int    `json:"indicator_id"`
	ID           string `json:"id"`
	Code         string `json:"code"`
	UkerDivision string `json:"uker_division"`
	Selected     bool   `json:"selected"`
}

type UpdateSelectedUker struct {
	IndicatorID int    `json:"indicator_id"`
	UkerKey     string `json:"uker_key"`
	Selected    bool   `json:"selected"`
}

type RequestSelectedUker struct {
	Data []UpdateSelectedUker `json:"data"`
}

type IndicatorHeader struct {
	IndicatorID int      `json:"indicator_id" validate:"required"`
	HeaderKey   []string `json:"header_key" validate:"required"`
}

type IndicatorHeaderRequest struct {
	Data []IndicatorHeader `json:"data"`
}
