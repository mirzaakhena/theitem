package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"theitem/shared/model/apperror"
)

func TestReputation001(t *testing.T) {

	assert.Nil(t, Reputation(0).Validate())
	assert.Nil(t, Reputation(1000).Validate())
	assert.Equal(t, "ER0001", Reputation(-1).Validate().(apperror.ErrorType).Code())
	assert.Equal(t, "ER0001", Reputation(1001).Validate().(apperror.ErrorType).Code())

}

func TestReputation002(t *testing.T) {

	assert.Equal(t, "red", Reputation(200).Badge())
	assert.Equal(t, "yellow", Reputation(600).Badge())
	assert.Equal(t, "green", Reputation(800).Badge())

}
