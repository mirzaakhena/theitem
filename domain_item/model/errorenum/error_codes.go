package errorenum

import "theitem/shared/model/apperror"

const (
	UnknownError                apperror.ErrorType = "ER0000 unknown error"
	OutOfRangeReputation        apperror.ErrorType = "ER0001 out of range reputation. must between 0 to 1000"
	InvalidURL                  apperror.ErrorType = "ER0002 invalid url for '%s'"
	InvalidCategory             apperror.ErrorType = "ER0003 invalid category. must be one of %v"
	InvalidRatingValue          apperror.ErrorType = "ER0004 invalid rating value. must be integer between 0..5"
	ForbiddenWord               apperror.ErrorType = "ER0005 word '%s' is not allowed"
	NameLengthMustGreaterThan   apperror.ErrorType = "ER0006 name length must greater than %d"
	UnavailableItem             apperror.ErrorType = "ER0007 unavailable item with id '%s'"
	UnavailableItemStock        apperror.ErrorType = "ER0008 unavailable item stock. requested %d but availability is %d"
	ItemNameAlreadyExist        apperror.ErrorType = "ER0009 item with name '%s' already exist"
	InvalidReputationBadge      apperror.ErrorType = "ER0010 invalid reputation badge"
	PriceMustGreaterOrEqualZero apperror.ErrorType = "ER0011 price must greater or equal zero"
)
