package db

func GetBuildings() ([]string, error) {
	db := GetDBInstance()
	rows, err := db.Query("SELECT * FROM buildings")
	if err != nil {
		// handle this error better than this
		return nil, err
	}
	var buildingsNamesArr []string
	defer rows.Close()
	for rows.Next() {
		var id int
		var description string
		var address string
		var country string
		var categoryId int
		var guestsNum int
		var roomsNum int
		var bathroomsNum int
		var priceDay int
		var avalableFrom string
		var avalableUntill string
		var userId int
		err = rows.Scan(&id, &description, &address, &country, &categoryId, &guestsNum, &roomsNum, &bathroomsNum, &priceDay, &avalableFrom, &avalableUntill, &userId)

		if err != nil {
			// handle this error
			return nil, err
		}

		buildingsNamesArr = append(buildingsNamesArr, description)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buildingsNamesArr, nil
}
