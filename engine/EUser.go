package engine

import (
	"golang.org/x/crypto/bcrypt"
	"ofcode.dev/labala-backend/structs"
)

func GetSingleUser(idb *InDB, searchKey string, param interface{}) (data structs.User, err error) {
	fr := idb.DB.
		Where(searchKey+" = ?", param).
		Last(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreateUser(idb *InDB, user structs.User) (data structs.User, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return data, err
	}

	user.Password = string(hash)

	fr := idb.DB.
		Create(&user)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = user
	return data, nil
}
