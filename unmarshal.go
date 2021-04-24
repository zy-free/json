package json

import (
	"encoding/json"
	"reflect"
	"strconv"
)

const tagName = "default"

func setDefaultValue(val interface{}) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(val).Elem()
	v := reflect.ValueOf(val).Elem()
	for i := 0; i < t.NumField(); i++ {
		if _, exists := t.Field(i).Tag.Lookup(tagName); !exists {
			continue
		}

		if v.Field(i).IsZero() {
			tagVal := t.Field(i).Tag.Get(tagName)
			setValue(v.Field(i), tagVal)
		}
	}
}

func setValue(v reflect.Value, defVal string) {
	switch v.Kind() {
	case reflect.String:
		v.SetString(defVal)
		break
	case reflect.Bool:
		tmp, err := strconv.ParseBool(defVal)
		if err != nil {
			panic(err.Error())
		}
		v.SetBool(tmp)
		break
	case reflect.Int64, reflect.Int32, reflect.Int, reflect.Int8:
		tmp, err := strconv.ParseInt(defVal, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.SetInt(tmp)
		break
	case reflect.Uint64, reflect.Uint32, reflect.Uint, reflect.Uint8:
		tmp, err := strconv.ParseUint(defVal, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.SetUint(tmp)
		break
	default:
		panic("unsupported type :" + v.Type().Kind().String())
	}
}

// Unmarshal with JSON
func Unmarshal(data []byte, val interface{}) error {
	err := json.Unmarshal(data, &val)
	return err
}

// UnmarshalDefault set default value by tag if value is empty
func UnmarshalDefault(data []byte, val interface{}) error {
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	setDefaultValue(val)
	return nil
}

