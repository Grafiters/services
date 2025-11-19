package lib

const (
	defaultLimit = 10
	maxLimit     = 100
	defaultOrder = ""
	defaultSort  = "DESC"
)

type Pagination struct {
	TotalData   int `json:"total_data"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	if order == "" {
		order = defaultOrder
	}

	if sort == "" {
		sort = defaultSort
	}

	offset := (page - 1) * limit

	return offset, page, limit, order, sort
}

func SetPaginationResponse(page, limit, total int, totalData int) Pagination {
	var lastPage int

	if total > 0 {
		lastPage = total / limit
		if total%limit != 0 {
			lastPage++
		}
	}

	return Pagination{
		TotalData:   totalData,
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}
}
