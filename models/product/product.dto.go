package models

import "riskmanagement/lib"

type ProductRequest struct {
	ID            int64   `json:"id"`
	KodeProduct   string  `json:"kode_product"`
	Product       string  `json:"product"`
	ActivityID    *string `json:"activity_id"`
	SubActivityID *string `json:"sub_activity_id"`
	LiniBisnisLv1 *string `json:"lini_bisnis_lv1"`
	LiniBisnisLv2 *string `json:"lini_bisnis_lv2"`
	LiniBisnisLv3 *string `json:"lini_bisnis_lv3"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"update_at"`
	Segment       *string `json:"segment"`
}

type ProductResponse struct {
	ID            int64   `json:"id"`
	KodeProduct   string  `json:"kode_product"`
	Product       string  `json:"product"`
	ActivityID    *string `json:"activity_id"`
	SubActivityID *string `json:"sub_activity_id"`
	LiniBisnisLv1 *string `json:"lini_bisnis_lv1"`
	LiniBisnisLv2 *string `json:"lini_bisnis_lv2"`
	LiniBisnisLv3 *string `json:"lini_bisnis_lv3"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"update_at"`
	Segment       *string `json:"segment"`
}

type PageRequest struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type ProductResponseNull struct {
	ID            lib.NullInt64  `json:id`
	KodeProduct   lib.NullString `json:"kode_product"`
	Product       lib.NullString `json:"product"`
	ActivityID    lib.NullString `json:"activity_id"`
	SubActivityID lib.NullString `json:"sub_activity_id"`
	LiniBisnisLv1 lib.NullString `json:"lini_bisnis_lv1"`
	LiniBisnisLv2 lib.NullString `json:"lini_bisnis_lv2"`
	LiniBisnisLv3 lib.NullString `json:"lini_bisnis_lv3"`
	CreatedAt     lib.NullString `json:"created_at"`
	UpdatedAt     lib.NullString `json:"update_at"`
	Segment       lib.NullString `json:"segment"`
}

type GetProductByActivityRequest struct {
	ActivityID string `json:"activity_id"`
}

type GetProductBySegmentRequest struct {
	Segment string `json:"segment"`
}

type GetProductByActivityResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type GetProductByActivityResponseNull struct {
	ID   lib.NullInt64  `json:"id"`
	Code lib.NullString `json:"code"`
	Name lib.NullString `json:"name"`
}

type GetProductBySegmentResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type KodeProduct struct {
	KodeProduct string `json:kode_product`
}

type KeywordRequest struct {
	Order      string `json:"order"`
	Sort       string `json:"sort"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Keyword    string `json:"keyword"`
	Segment    string `json:"segment"`
	ActivityID string `json:"activity_id"`
}

func (p ProductRequest) ParseRequest() Product {
	return Product{
		ID:            p.ID,
		KodeProduct:   p.KodeProduct,
		Product:       p.Product,
		ActivityID:    p.ActivityID,
		SubActivityID: p.SubActivityID,
		LiniBisnisLv1: p.LiniBisnisLv1,
		LiniBisnisLv2: p.LiniBisnisLv2,
		LiniBisnisLv3: p.LiniBisnisLv3,
		Segment:       p.Segment,
	}
}

type ProductResponses struct {
	ID      int64  `json:"id"`
	Product string `json:"name"`
}

type ProductResponsesNull struct {
	ID      lib.NullInt64  `json:"id"`
	Product lib.NullString `json:"product"`
}

func (p ProductResponse) ParseRequest() Product {
	return Product{
		ID:            p.ID,
		KodeProduct:   p.KodeProduct,
		Product:       p.Product,
		ActivityID:    p.ActivityID,
		SubActivityID: p.SubActivityID,
		LiniBisnisLv1: p.LiniBisnisLv1,
		LiniBisnisLv2: p.LiniBisnisLv2,
		LiniBisnisLv3: p.LiniBisnisLv3,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		Segment:       p.Segment,
	}
}

func (pr ProductRequest) TableName() string {
	return "product"
}

func (pr ProductResponse) TableName() string {
	return "product"
}
