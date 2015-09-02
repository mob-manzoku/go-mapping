package mapping

import "reflect"

func Mapping(list interface{}, key string) (map[string]interface{}, error) {

	mapped := map[string]interface{}{}

	switch reflect.TypeOf(list).Kind() {

	case reflect.Slice:
	case reflect.Array:
		l := reflect.ValueOf(list)
		for i := 0; i < l.Len(); i++ {
			mapped[l.Index(i).FieldByName(key).String()] = l.Index(i).Interface()

		}

	}

	return mapped, nil
}
