package recorder

import (
	"strings"
)

func ItemToKeyValue(item map[string]interface{}) (fieldsAsStr string, placeholdersAsStr string, valuesArr []interface{}) {
	var fields []string
	var placeholders []string
	var values []interface{}

	for key, value := range item {
		fields = append(fields, key)
		placeholders = append(placeholders, "?")
		values = append(values, value)
	}

	fieldsAsStr = strings.Join(fields, ",")
	placeholdersAsStr = strings.Join(placeholders, ",")
	valuesArr = values
	return
}
