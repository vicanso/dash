package structDiff

import (
	"reflect"
)

// convert struct to map
func structToMap(s interface{}) (data map[string]interface{}) {
	switch s.(type) {
	case map[string]interface{}:
		return s.(map[string]interface{})
	}
	e := reflect.ValueOf(s).Elem()
	count := e.NumField()
	typeOfT := e.Type()
	data = make(map[string]interface{})
	for index := 0; index < count; index++ {
		name := typeOfT.Field(index).Name
		value := e.FieldByName(name).Interface()
		data[name] = value
	}
	return
}

// Diff get the diff between current and original，
// they should be the same struct type
func Diff(current, original interface{}, args ...interface{}) (update map[string]interface{}) {
	map1 := structToMap(current)
	map2 := structToMap(original)
	update = make(map[string]interface{})
	for name, v1 := range map1 {
		v2 := map2[name]
		if !reflect.DeepEqual(v1, v2) {
			update[name] = v1
		}
	}
	return
}
