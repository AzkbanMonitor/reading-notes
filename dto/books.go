package dto

type Books struct {
	BookName   string `json:"bookName"`
	Writer     string `json:"writer"`
	TotalPages int    `json:"totalPages"`
	Status     bool   `json:"status"`
}
