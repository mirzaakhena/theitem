package runitemcreate

import (
	"theitem/domain_item/model/repository"
)

type Outport interface {
	repository.SaveItemRepo
	repository.ExistInForbiddenNameListRepo
	repository.FindOneItemByNameRepo
}
