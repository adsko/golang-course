package database

func AddItem (name, description string, priority Priority) (*DBItem, error) {
	item, err := addToDB(name, description, priority)

	if err != nil {
		return item, err
	}


	return item, nil
}

func GetAllItems () ([]*DBItem, error) {
	items, err := getAllFromDB()

	if err != nil {
		return nil, err
	}

	return items, nil
}

func Delete (name string) error {
	return removeFromDB(name)
}