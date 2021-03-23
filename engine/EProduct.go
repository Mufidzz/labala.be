package engine

import "ofcode.dev/labala-backend/structs"

func GetSingleProduct(idb *InDB, searchKey string, param interface{}) (data structs.Product, err error) {
	fr := idb.DB.
		Where(searchKey+" = ?", param).
		Last(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetSingleProductWProductExchange(idb *InDB, searchKey string, param interface{}) (data structs.ProductWProductExchange, err error) {
	fr := idb.DB.
		Preload("ProductExchanges").
		Where(searchKey+" = ?", param).
		Last(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreateProduct(idb *InDB, Product structs.ProductWProductExchange) (data structs.ProductWProductExchange, err error) {
	fr := idb.DB.
		Preload("ProductExchanges").
		Create(&Product)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = Product
	return data, nil
}

func UpdateProduct(idb *InDB, newProduct structs.Product) (data structs.Product, err error) {
	fr := idb.DB.
		Save(&newProduct)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = newProduct
	return data, nil
}

func DeleteProduct(idb *InDB, id string) (data structs.Product, err error) {
	data, err = GetSingleProduct(idb, "id", id)
	if err != nil {
		return data, err
	}
	fr := idb.DB.
		Delete(&data)
	if fr.Error != nil {
		return data, fr.Error
	}
	return data, nil
}
