package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"theitem/shared/model/apperror"
)

func TestCaseCategory001(t *testing.T) {

	assert.Nil(t, Category("photo").Validate())
	assert.Nil(t, Category("sketch").Validate())
	assert.Nil(t, Category("cartoon").Validate())
	assert.Nil(t, Category("animation").Validate())

	assert.Equal(t, "ER0003", Category("other").Validate().(apperror.ErrorType).Code())

}
