package models

type SubActivityRequest struct {
	ID              int64   `json:"id"`
	ActivityID      int64   `json:"activity_id"`
	KodeSubActivity string  `json:"kode_sub_activity"`
	Name            string  `json:"name"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
}

type SubActivityResponse struct {
	ID              int64   `json:"id"`
	ActivityID      int64   `json:"activity_id"`
	KodeSubActivity string  `json:"kode_sub_activity"`
	Name            string  `json:"name"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
}

type SubActivityLastId struct {
	totalRows int64 `json:"total_rows"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type SubActivityResponses struct {
	ID              int64   `json:"id"`
	ActivityID      int64   `json:"activity_id"`
	ActivityName    string  `json:"activity_name"`
	KodeSubActivity string  `json:"kode_sub_activity"`
	NameSubActivity string  `json:"name_sub_activity"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
}

type KodeSubActivityRequest struct {
	KodeActivity string `json:"kode_activity"`
}

type KodeSubActivityResponse struct {
	KodeActivity    string `json:"kode_activity"`
	KodeSubActivity string `json:"kode_sub_activity"`
}

func (p SubActivityRequest) ParseRequest() SubActivity {
	return SubActivity{
		ID:              p.ID,
		KodeSubActivity: p.KodeSubActivity,
		ActivityID:      p.ActivityID,
		Name:            p.Name,
		// CreatedAt:  p.CreatedAt,
		// UpdatedAt:  p.UpdatedAt,
	}
}

func (p SubActivityResponse) ParseResponse() SubActivity {
	return SubActivity{
		ID:              p.ID,
		KodeSubActivity: p.KodeSubActivity,
		ActivityID:      p.ActivityID,
		Name:            p.Name,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func (ar SubActivityRequest) TableName() string {
	return "sub_activity"
}

func (ar SubActivityLastId) TableName() string {
	return "sub_activity"
}

func (ar SubActivityResponse) TableName() string {
	return "sub_activity"
}
