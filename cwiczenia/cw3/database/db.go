package database


//func addToDB (name, description string, priority Priority) (*DBItem, error) {
//	var priorityInt uint = 0
//
//	if priority == MEDIUM {
//		priorityInt = 1
//	} else if priority == HIGH {
//		priorityInt = 2
//	}
//
//	item := makeDBItem(name, description, priorityInt)
//	return &item, driver.Insert(item)
//}

//func getAllFromDB() ([]*DBItem, error) {
//	var output []*DBItem
//	err := driver.Open(DBItem{}).Get().AsEntity(&output)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return output, nil
//}

func removeFromDB(name string) error {
	item := getItemOnlyWithID(name)

	err := driver.Open(DBItem{}).Delete(item)

	return err
}