package vo

import (
	"regexp"
	"theitem/domain_item/model/errorenum"
)

type StringURL string

const urlPatternRegex = `(http|https):\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?`

var re = regexp.MustCompile(urlPatternRegex)

// Validate MUST be a valid URL
func (r StringURL) Validate(urlContext string) error {

	if !re.Match([]byte(r)) {
		return errorenum.InvalidURL.Var(urlContext)
	}

	return nil
}

func (r StringURL) String() string {
	return string(r)
}
