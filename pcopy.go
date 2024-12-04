package pcopy

import "reflect"

// CopyProperties copy properties from src to dst
func CopyProperties(src, dst interface{}) {
	srcValue := getRealValue(src)
	dstValue := getRealValue(dst)
	srcType := getRealType(src)
	dstType := getRealType(dst)

	for i := range dstType.NumField() {
		dstField := dstType.Field(i)
		if _, exists := srcType.FieldByName(dstField.Name); exists {
			// log.Printf("Dest: %v, Src: %v\n", dstField, srcField)
			dstFieldValue := dstValue.FieldByName(dstField.Name)
			srcFieldValue := srcValue.FieldByName(dstField.Name)

			if dstFieldValue.Kind() == reflect.Ptr && srcFieldValue.Kind() != reflect.Ptr {
				dstFieldValue.Elem().Set(srcFieldValue)
			} else if dstFieldValue.Kind() != reflect.Ptr && srcFieldValue.Kind() == reflect.Ptr {
				dstFieldValue.Set(srcFieldValue.Elem())
			} else {
				dstFieldValue.Set(srcFieldValue.Convert(dstFieldValue.Type()))
			}
		}
	}
}

func getRealType(object interface{}) reflect.Type {
	t := reflect.TypeOf(object)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func getRealValue(object interface{}) reflect.Value {
	v := reflect.ValueOf(object)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
