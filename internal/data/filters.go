package data

import "github.com/Infamous003/greenlight/internal/validator"

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
