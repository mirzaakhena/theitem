package vo

type ItemID string

//func NewItemID(randomStringID string, now time.Time) (ItemID, error) {
//	var obj = ItemID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))
//
//	// you may change it as necessary ...
//
//	return obj, nil
//}

func (r ItemID) String() string {
	return string(r)
}
