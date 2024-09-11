package schemas

import "reflect"

type baseModel struct{}

func (b *baseModel) ToMap(model interface{}) map[string]any {
	s := reflect.ValueOf(model)
	t := reflect.TypeOf(model)
	resMap := make(map[string]any, s.NumField())

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		var value any
		if field.Kind() == reflect.Ptr {
			value = field.Elem().Interface()
		} else {
			value = field.Interface()
		}

		dbTag := fieldType.Tag.Get("db")
		if dbTag != "" {
			resMap[dbTag] = value
		} else {
			resMap[fieldType.Name] = value
		}
	}

	return resMap
}
