package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"theitem/shared/model/apperror"
	"time"
)

func TestItem001(t *testing.T) {
	itemObj, err := NewItem(ItemCreateRequest{
		UUID:         "a",
		Now:          time.Time{},
		Name:         "12312312345",
		Rating:       0,
		Category:     "cartoon",
		ImageURL:     "http://a.b",
		Reputation:   100,
		Price:        500,
		Availability: 10,
	})

	assert.Nil(t, err)
	assert.Equal(t, "red", itemObj.ReputationBadge)

}

func TestItem002(t *testing.T) {
	_, err := NewItem(ItemCreateRequest{
		UUID:         "a",
		Now:          time.Time{},
		Name:         "123123123",
		Rating:       0,
		Category:     "cartoon",
		ImageURL:     "http://a.b",
		Reputation:   0,
		Price:        500,
		Availability: 10,
	})

	assert.Equal(t, "ER0006", err.(apperror.ErrorType).Code())

}

func TestItem003(t *testing.T) {
	_, err := NewItem(ItemCreateRequest{
		UUID:         "a",
		Now:          time.Time{},
		Name:         "12312312345",
		Rating:       0,
		Category:     "cartoon",
		ImageURL:     "http://a.b",
		Reputation:   0,
		Price:        -5,
		Availability: 10,
	})

	assert.Equal(t, "ER0011", err.(apperror.ErrorType).Code())

}

func TestItem004(t *testing.T) {
	itemObj, err := NewItem(ItemCreateRequest{
		UUID:         "a",
		Now:          time.Time{},
		Name:         "12312312345",
		Rating:       0,
		Category:     "cartoon",
		ImageURL:     "http://a.b",
		Reputation:   0,
		Price:        500,
		Availability: 5,
	})

	assert.Nil(t, err)
	assert.Equal(t, "ER0008", itemObj.Purchase(7).(apperror.ErrorType).Code())

}

func TestItem005(t *testing.T) {
	itemObj, err := NewItem(ItemCreateRequest{
		UUID:         "a",
		Now:          time.Time{},
		Name:         "12312312345",
		Rating:       0,
		Category:     "cartoon",
		ImageURL:     "http://a.b",
		Reputation:   0,
		Price:        500,
		Availability: 5,
	})

	assert.Nil(t, err)

	assert.Nil(t, itemObj.Purchase(3))
	assert.Equal(t, 2, itemObj.Availability)

}
