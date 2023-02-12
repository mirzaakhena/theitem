package vo

import "theitem/domain_item/model/errorenum"

type Category string

const (
	CategoryPhoto     Category = "photo"
	CategorySketch    Category = "sketch"
	CategoryCartoon   Category = "cartoon"
	CategoryAnimation Category = "animation"
)

var cs = []Category{
	CategoryPhoto,
	CategorySketch,
	CategoryCartoon,
	CategoryAnimation,
}

// Validate
// category can be one of photo, sketch, cartoon, animation and it should be a string
func (r Category) Validate() error {

	for _, c := range cs {
		if c == r {
			return nil
		}
	}

	return errorenum.InvalidCategory.Var(cs)
}

func (r Category) String() string {
	return string(r)
}
