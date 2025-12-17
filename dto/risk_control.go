package dto

type PreviewFile[T any] struct {
	PerRow     T      `json:"row"`
	Validation string `json:"validation"`
}

type PreviewFileImport[T any] struct {
	Header T                `json:"header"`
	Body   []PreviewFile[T] `json:"body"`
}

type ListAttributeMap struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
