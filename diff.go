package structDiff

import (
	"reflect"
)

// Diff get the diff between current and original，
// they should be the same struct type
func Diff(current, original interface{}, args ...interface{}) (update map[string]interface{}) {
	e1 := reflect.ValueOf(current).Elem()
	e2 := reflect.ValueOf(original).Elem()
	count := e1.NumField()
	typeOfT := e1.Type()
	update = make(map[string]interface{})
	for index := 0; index < count; index++ {
		name := typeOfT.Field(index).Name
		v1 := e1.FieldByName(name).Interface()
		v2 := e2.FieldByName(name).Interface()
		if !reflect.DeepEqual(v1, v2) {
			update[name] = v1
		}
	}
	return
}
