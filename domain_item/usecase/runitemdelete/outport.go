package runitemdelete

import "theitem/domain_item/model/repository"

type Outport interface {
	repository.DeleteOneItemRepo
}
