package mapping

import (
	"errors"
	"reflect"
)

func Mapping(list interface{}, key string) (map[string]interface{}, error) {

	mapped := map[string]interface{}{}

	switch reflect.TypeOf(list).Kind() {

	case reflect.Slice:
	case reflect.Array:
		l := reflect.ValueOf(list)
		for i := 0; i < l.Len(); i++ {

			// if l.Index(i).FieldByName(key).IsValid() {
			// 	return nil, errors.New(list.(string) + " must has key: " + key)
			// }

			mapped[l.Index(i).FieldByName(key).String()] = l.Index(i).Interface()
		}

	default:
		return nil, errors.New(list.(string) + " must be array or slice")

	}

	return mapped, nil
}
