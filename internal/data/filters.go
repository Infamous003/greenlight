package data

import (
	"strings"

	"github.com/Infamous003/greenlight/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string // contains all the available sort options, like year, title, genres, etc
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitzero"`
	PageSize     int `json:"page_size,omitzero"`
	FirstPage    int `json:"first_page,omitzero"`
	LastPage     int `json:"last_page,omitzero"`
	TotalRecords int `json:"total_records,omitzero"`
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than 0")
	v.Check(f.Page <= 10000000, "page", "must be lesser than 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than 0")
	v.Check(f.PageSize <= 100, "page", "must be lesser than 100")

	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}

// loops through SortSafeList and checks if the Sort is present in it, if it is
// then it returns the sort value without any prefixes
func (f *Filters) sortColumn() string {
	for _, sortValue := range f.SortSafeList {
		if f.Sort == sortValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort paramter: " + f.Sort)
}

func (f *Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f *Filters) limit() int {
	return f.PageSize
}

func (f *Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

func caclulateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}
