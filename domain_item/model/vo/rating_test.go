package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"theitem/shared/model/apperror"
)

func TestCaseRating001(t *testing.T) {

	for i := 0; i <= 5; i++ {
		assert.Nil(t, Rating(i).Validate())
	}

	assert.Equal(t, "ER0004", Rating(-1).Validate().(apperror.ErrorType).Code())
	assert.Equal(t, "ER0004", Rating(6).Validate().(apperror.ErrorType).Code())

}
