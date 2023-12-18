package paginator

import (
	"fmt"
	"strings"
)

var (
	// DefaultPageSize specifies the default per page size
	DefaultPageSize = 15
	// MaxPageSize specifies the maximum per page size
	MaxPageSize = 100
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PerPageVar specifies the query parameter name for page size limit
	PerPageVar = "limit"
)

type PaginatorResult struct {
	Paginator *Paginator
	Records   interface{}
}

type Paginator struct {
	NextPage     int `json:"next_page,omitempty" example:"4"`
	PreviousPage int `json:"prev_page,omitempty" example:"2"`
	PerPage      int `json:"per_page,omitempty" example:"15"`
	CurrentPage  int `json:"current_page,omitempty" example:"3"`
	TotalPages   int `json:"to,omitempty" example:"20"`
	TotalRecords int `json:"total,omitempty" example:"50"`
}

// Object struct for pagination links
type PageLink struct {
	First    string `json:"first,omitempty" example:"http://127.0.0.1:8080/api/v1/dishes?page=3"`
	Next     string `json:"next,omitempty" example:"http://127.0.0.1:8080/api/v1/dishes?page=4"`
	Last     string `json:"last,omitempty" example:"http://127.0.0.1:8080/api/v1/dishes?page=5"`
	Previous string `json:"prev,omitempty" example:"http://127.0.0.1:8080/api/v1/dishes?page=2"`
}

func (p *Paginator) Offset() int {
	return (p.CurrentPage - 1) * p.PerPage
}

func (p *Paginator) Limit() int {
	return p.PerPage
}

func NewPaginator(currentPage int, perPage int, totalRecords int) *Paginator {
	var paginator Paginator

	if perPage <= 0 {
		perPage = DefaultPageSize
	}

	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}

	if currentPage < 1 {
		currentPage = 1
	}

	pageCount := (totalRecords / perPage)

	paginator.TotalRecords = totalRecords
	paginator.TotalPages = pageCount
	paginator.CurrentPage = currentPage
	paginator.PerPage = perPage
	paginator.NextPage = currentPage + 1

	// Calculate total pages to show all data
	remainder := (totalRecords % perPage)
	if remainder != 0 {
		paginator.TotalPages = pageCount + 1
	}

	if currentPage > 0 && currentPage <= paginator.TotalPages {
		paginator.PreviousPage = currentPage - 1
	}

	if currentPage == paginator.TotalPages {
		paginator.NextPage = 0
	}

	if currentPage >= paginator.TotalPages {
		paginator.NextPage = 0
	}

	return &paginator
}

// Build pagination links: first, prev, next, and last links corresponding to the pagination.
func (p *Paginator) BuildLinks(baseURL string, defaultPerPage int) *PageLink {
	var pageLink PageLink

	if strings.Contains(baseURL, "?") {
		baseURL += "&"
	} else {
		baseURL += "?"
	}

	if p.CurrentPage >= 1 {
		pageLink.First = fmt.Sprintf("%v%v=%v", baseURL, PageVar, 1)
		pageLink.Last = fmt.Sprintf("%v%v=%v", baseURL, PageVar, p.TotalPages)
	}

	if p.NextPage > 0 {
		pageLink.Next = fmt.Sprintf("%v%v=%v", baseURL, PageVar, p.NextPage)
	}

	if p.PreviousPage > 0 {
		pageLink.Previous = fmt.Sprintf("%v%v=%v", baseURL, PageVar, p.PreviousPage)
	}

	return &pageLink
}
