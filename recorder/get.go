package recorder

import (
	"MasterThesis/logger"
	"database/sql"
	"reflect"
)

func Get(db *sql.DB, tableName string, where string, item interface{}) []interface{} {
	query := "SELECT * FROM " + tableName
	if where != "" {
		query += " WHERE " + where
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.LogE(err.Error())
	}
	defer rows.Close()

	var results []interface{}
	columns, err := rows.Columns()
	if err != nil {
		logger.LogE(err.Error())
		return nil
	}
	for rows.Next() {
		// Create a new instance of item to hold the row data
		newItem := reflect.New(reflect.TypeOf(item)).Interface()

		// Create a slice of interface{} to hold each field value
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}

		if err := rows.Scan(values...); err != nil {
			logger.LogE(err.Error())
			continue
		}
		results = append(results, newItem)
	}
	if err = rows.Err(); err != nil {
		logger.LogE(err)
	}
	return results

}
