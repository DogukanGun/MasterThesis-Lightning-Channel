package recorder

import (
	"MasterThesis/logger"
	"database/sql"
)

type Chat struct {
	Message string `json:"message"`
}

func Save(db *sql.DB, tableName string, item map[string]interface{}) {
	fieldsAsStr, placeholderAsStr, values := ItemToKeyValue(item)
	query := "INSERT INTO " + tableName + " ( " + fieldsAsStr + "  ) VALUES ( " + placeholderAsStr + " )"
	insert, err := db.Prepare(query)
	if err != nil {
		logger.LogE(err)
		return
	}

	defer insert.Close() // Close the prepared statement after execution completes

	_, err = insert.Exec(values...)
	if err != nil {
		logger.LogE(err)
		return
	}

	logger.LogS("Insert successful")
}
