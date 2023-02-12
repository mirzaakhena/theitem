package vo

import "theitem/domain_item/model/errorenum"

type Rating int

// Validate
// rating MUST accept only integers, where rating is >= 0 and <= 5
func (r Rating) Validate() error {

	if r < 0 || r > 5 {
		return errorenum.InvalidRatingValue
	}

	return nil
}

func (r Rating) Int() int {
	return int(r)
}
