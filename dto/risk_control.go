package dto

type PreviewRiskControl struct {
	PerRow     [10]string `json:"row"`
	Validation string     `json:"validation"`
}

type PreviewFileRiskControl struct {
	Header [10]string           `json:"header"`
	Body   []PreviewRiskControl `json:"body"`
}
