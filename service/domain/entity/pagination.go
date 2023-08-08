package entity

type Paginate struct {
	List      interface{} `json:"list"`
	Limit     int64       `json:"limit"`
	Page      int64       `json:"page"`
	TotalData int64       `json:"total_data"`
	TotalPage float64     `json:"total_page"`
}
