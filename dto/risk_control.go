package dto

type PreviewFile[T any] struct {
	PerRow     T      `json:"row"`
	Validation string `json:"validation"`
}

type PreviewFileImport[T any] struct {
	Header T                `json:"header"`
	Body   []PreviewFile[T] `json:"body"`
}
