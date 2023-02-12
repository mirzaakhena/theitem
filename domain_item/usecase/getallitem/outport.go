package getallitem

import "theitem/domain_item/model/repository"

type Outport interface {
	repository.FindAllItemRepo
}
