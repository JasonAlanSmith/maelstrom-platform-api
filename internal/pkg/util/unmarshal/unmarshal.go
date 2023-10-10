package unmarshal

import (
	"encoding/json"
	"log/slog"
	"reflect"
	"strings"
)

func Unmarshal(data []byte, dts interface{}) (error, []string) {
	var val_err []string
	err := json.Unmarshal(data, dts)
	if err != nil {
		slog.Error(err.Error())
		return err, val_err
	}

	fields := reflect.ValueOf(dts).Elem()
	for i := 0; i < fields.NumField(); i++ {
		maelstromTags := fields.Type().Field(i).Tag.Get("maelstrom")
		if strings.Contains(maelstromTags, "required") && fields.Field(i).IsZero() {
			val_err = append(val_err, "Required field is missing: "+fields.Type().Field(i).Name)
		}
	}
	return nil, val_err
}
