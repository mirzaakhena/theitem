package entity

import (
	"theitem/domain_item/model/errorenum"
	"theitem/domain_item/model/vo"
	"time"
)

type Item struct {
	ID              vo.ItemID     `json:"id" bson:"_id"`
	Created         time.Time     `json:"created"`
	Updated         time.Time     `json:"updated"`
	Name            string        `json:"name"`
	Rating          vo.Rating     `json:"rating" `
	Category        vo.Category   `json:"category"`
	ImageURL        vo.StringURL  `json:"image"`
	Reputation      vo.Reputation `json:"reputation"`
	ReputationBadge string        `json:"reputation_badge"`
	Price           int           `json:"price"`
	Availability    int           `json:"availability"`
}

type ItemCreateRequest struct {
	UUID         string        `json:"-"`
	Now          time.Time     `json:"-"`
	Name         string        `json:"name"`
	Rating       vo.Rating     `json:"rating"`
	Category     vo.Category   `json:"category"`
	ImageURL     vo.StringURL  `json:"image"`
	Reputation   vo.Reputation `json:"reputation"`
	Price        int           `json:"price"`
	Availability int           `json:"availability"`
}

const minItemNameChar = 10

func (r ItemCreateRequest) Validate() error {

	var err error

	if len(r.Name) <= minItemNameChar {
		return errorenum.NameLengthMustGreaterThan.Var(minItemNameChar)
	}

	if r.Price < 0 {
		return errorenum.PriceMustGreaterOrEqualZero
	}

	err = r.Rating.Validate()
	if err != nil {
		return err
	}

	err = r.ImageURL.Validate("image")
	if err != nil {
		return err
	}

	err = r.Category.Validate()
	if err != nil {
		return err
	}

	err = r.Reputation.Validate()
	if err != nil {
		return err
	}

	return nil
}

func NewItem(req ItemCreateRequest) (*Item, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	var obj Item
	obj.ID = vo.ItemID(req.UUID)
	obj.Created = req.Now
	obj.Updated = req.Now

	obj.Name = req.Name
	obj.Category = req.Category
	obj.Reputation = req.Reputation
	obj.ReputationBadge = obj.Reputation.Badge()
	obj.Price = req.Price
	obj.ImageURL = req.ImageURL
	obj.Availability = req.Availability
	obj.Rating = req.Rating

	return &obj, nil
}

type ItemUpdateRequest struct {
	Now      time.Time    `json:"-"`
	Name     string       `json:"name"`
	Category vo.Category  `json:"category"`
	ImageURL vo.StringURL `json:"image"`
	Price    int          `json:"price"`
}

func (r ItemUpdateRequest) Validate() error {

	var err error

	if len(r.Name) < minItemNameChar {
		return errorenum.NameLengthMustGreaterThan.Var(minItemNameChar)
	}

	err = r.ImageURL.Validate("image")
	if err != nil {
		return err
	}

	err = r.Category.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (r *Item) Update(req ItemUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now
	r.Category = req.Category
	r.ImageURL = req.ImageURL
	r.Price = req.Price
	r.Name = req.Name

	return nil
}

func (r *Item) Purchase(quantity int) error {

	if r.Availability < quantity {
		return errorenum.UnavailableItemStock.Var(quantity, r.Availability)
	}

	r.Availability = r.Availability - quantity

	return nil
}
