package models

type Product struct {
	ID            int64
	KodeProduct   string
	Product       string
	ActivityID    *string
	SubActivityID *string
	LiniBisnisLv1 *string
	LiniBisnisLv2 *string
	LiniBisnisLv3 *string
	CreatedAt     *string
	UpdatedAt     *string
	Segment       *string
}
