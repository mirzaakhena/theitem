package vo

import "theitem/domain_item/model/errorenum"

type Reputation int

// Validate
// The reputation MUST be an integer >= 0 and <= 1000
func (r Reputation) Validate() error {

	if r < 0 || r > 1000 {
		return errorenum.OutOfRangeReputation
	}

	return nil
}

func (r Reputation) Int() int {
	return int(r)
}

// Badge
//
//	If reputation is <= 500 the value is red
//	If reputation is <= 799 the value is yellow
//	Otherwise the value is green
func (r Reputation) Badge() string {

	if r <= 500 {
		return "red"
	}

	if r <= 799 {
		return "yellow"
	}

	return "green"
}
