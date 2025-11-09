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
