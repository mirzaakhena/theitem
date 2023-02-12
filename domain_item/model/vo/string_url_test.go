package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringURL001(t *testing.T) {
	assert.Nil(t, StringURL("http://a.b").Validate(""))

	assert.NotNil(t, StringURL("http://a").Validate(""))

	assert.NotNil(t, StringURL("x").Validate(""))
}
