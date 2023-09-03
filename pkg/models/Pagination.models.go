package models

type PaginationDataBooks struct {
	Books        *[]Book `json:"books"`
	NextPage     int     `json:"nextPage"`
	PreviousPage int     `json:"previousPage"`
	CurrentPage  int     `json:"currentPage"`
	TotalPages   int     `json:"totalPages"`
}

type PaginationDataTesis struct {
	Tesis        *[]Tesis `json:"tesis"`
	NextPage     int      `json:"nextPage"`
	PreviousPage int      `json:"previousPage"`
	CurrentPage  int      `json:"currentPage"`
	TotalPages   int      `json:"totalPages"`
}

type PaginationDataProjects struct {
	Projects     *[]Projects `json:"projects"`
	NextPage     int         `json:"nextPage"`
	PreviousPage int         `json:"previousPage"`
	CurrentPage  int         `json:"currentPage"`
	TotalPages   int         `json:"totalPages"`
}
