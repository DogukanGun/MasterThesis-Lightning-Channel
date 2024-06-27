package recorder

import (
	"MasterThesis/logger"
	"database/sql"
	"errors"
	"reflect"
)

func Get(db *sql.DB, tableName string, where string, item interface{}) error {
	query := "SELECT * FROM " + tableName
	if where != "" {
		query += " WHERE " + where
	}

	rows, err := db.Query(query)
	if err != nil {
		logger.LogE(err.Error())
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.LogE(err.Error())
		return err
	}

	// Check if item is a pointer to a slice
	sliceVal := reflect.ValueOf(item)
	if sliceVal.Kind() != reflect.Ptr || sliceVal.Elem().Kind() != reflect.Slice {
		err := errors.New("item must be a pointer to a slice")
		logger.LogE(err.Error())
		return err
	}

	sliceElemType := sliceVal.Elem().Type().Elem()
	results := reflect.MakeSlice(reflect.SliceOf(sliceElemType), 0, 0)

	for rows.Next() {
		elemPtr := reflect.New(sliceElemType)
		elem := elemPtr.Elem()

		values := make([]interface{}, len(columns))
		for i := range values {
			field := elem.Field(i)
			if field.CanAddr() {
				values[i] = field.Addr().Interface()
			} else {
				values[i] = new(interface{})
			}
		}

		if err := rows.Scan(values...); err != nil {
			logger.LogE(err.Error())
			continue
		}

		results = reflect.Append(results, elem)
	}
	if err = rows.Err(); err != nil {
		logger.LogE(err.Error())
		return err
	}

	sliceVal.Elem().Set(results)
	return nil
}
