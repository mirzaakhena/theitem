package getoneitem

import "theitem/domain_item/model/repository"

type Outport interface {
	repository.FindOneItemRepo
}
