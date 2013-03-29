// This code was taken from the beedb project, a simple ORM for Go
// beedb is licensed under BSD and can be found here:
// https://github.com/astaxie/beedb

package goes

import (
	"errors"
	"reflect"
	"strconv"
)

func mapToStruct(dStruct interface{}, mapping map[string]string) error {
	pStruct := reflect.Indirect(reflect.ValueOf(dStruct))
	
	if pStruct.Kind() != reflect.Struct {
		return errors.New("expected a pointer to a struct")
	}
	
	for key, data := range mapping {
		structField := pStruct.FieldByName(key)
		
		if !structField.CanSet() {
			continue
		}

		var v interface{}

		switch structField.Type().Kind() {
		case reflect.Slice:
			v = data
		case reflect.String:
			v = string(data)
		case reflect.Bool:
			v = string(data) == "1"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			x, err := strconv.Atoi(string(data))
			if err != nil {
				return errors.New("arg " + key + " as int: " + err.Error())
			}
			v = x
		case reflect.Int64:
			x, err := strconv.ParseInt(string(data), 10, 64)
			if err != nil {
				return errors.New("arg " + key + " as int: " + err.Error())
			}
			v = x
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(string(data), 64)
			if err != nil {
				return errors.New("arg " + key + " as float64: " + err.Error())
			}
			v = x
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(string(data), 10, 64)
			if err != nil {
				return errors.New("arg " + key + " as int: " + err.Error())
			}
			v = x
		default:
			return errors.New("unsupported type in Scan: " + reflect.TypeOf(v).String())
		}

		structField.Set(reflect.ValueOf(v))
	}
	return nil
}
