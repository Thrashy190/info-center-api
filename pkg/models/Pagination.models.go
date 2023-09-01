package models

type PaginationDataBooks struct {
	Books        *[]Book `json:"books"`
	NextPage     int     `json:"nextPage"`
	PreviousPage int     `json:"previousPage"`
	CurrentPage  int     `json:"currentPage"`
	TotalPages   int     `json:"totalPages"`
}
