// This Package handles all the queries for any model
// Default tableName is model name to lowercase and plural
// For type User, tableName is users
// These are generic functions and are only compatible with
// go 1.18 and above
package config

// Get the first record of the model
func GetOne[T any](model T) (T, error) {
	var models T
	err := DB.Where(model).First(&models).Error
	return models, err
}

// Create a new record for the model
func CreateOne[T any](class T) (interface{}, error) {
	err := DB.Create(&class).Error
	return class, err
}

// Create multiple records for the model
func CreateBulk[T any](data []T) ([]T, error) {
	err := DB.Create(&data).Error
	return data, err
}

// Update a record for the model
func UpdateOne[T any](class T) (interface{}, error) {
	err := DB.Save(&class).Error
	return class, err
}

// Delete a record for the model
func DeleteOne[T any](class T) (interface{}, error) {
	err := DB.Delete(&class).Error
	return class, err
}

// Get all the records of the model with pagination
func GetAll[T any](model T, page int, limit int, order_by string) (PaginatedResult[T], error) {
	offset := (page - 1) * limit
	var results []T
	// Get all the records of the model
	if order_by == "" {
		err := DB.Where(model).Offset(offset).Limit(limit).Find(&results).Error
		result := PaginatedResult[T]{
			Results: results,
			Page:    page,
			Limit:   limit,
		}
		return result, err
	}

	err := DB.Where(model).Offset(offset).Limit(limit).Order(order_by).Find(&results).Error
	result := PaginatedResult[T]{
		Results: results,
		Page:    page,
		Limit:   limit,
	}
	return result, err
}

func JoinGetAll[T any](tableName string, selection string, joins []string, page int, limit int, query string, model T) (PaginatedResult[T], error) {
	if tableName == "" || selection == "" || len(joins) == 0 || page < 1 || query == "" {
		panic("Missing Parameters")
	}
	switch len(joins) {
	default:
		return JoinTwoGetAll(tableName, selection, joins[0], page, limit, query, model)
	}
}

func JoinTwoGetAll[T any](tableName string, selection string, join1 string, page int, limit int, query string, model T) (PaginatedResult[T], error) {
	offset := (page - 1) * limit
	var results []T
	// Get all the records of the model
	err := DB.Table(tableName).Select(selection).Joins(join1).Where(query).Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return PaginatedResult[T]{}, err
	}
	result := PaginatedResult[T]{
		Results: results,
		Page:    page,
		Limit:   limit,
	}
	return result, err
}
