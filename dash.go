package structDiff

import (
	"reflect"

	"github.com/fatih/structs"
)

func getType(i interface{}) (valueType string) {
	switch i.(type) {
	case int:
		valueType = "int"
		break
	case string:
		valueType = "string"
		break
	case bool:
		valueType = "bool"
		break
	}
	return
}

// convert struct to map
func structToMap(s interface{}) (data map[string]interface{}) {
	switch s.(type) {
	case map[string]interface{}:
		return s.(map[string]interface{})
	}
	return structs.Map(s)
}

// Difference get the diff between current and original，
// they should be the same struct type
func Difference(current, original interface{}) (diff map[string]interface{}) {
	map1 := structToMap(current)
	map2 := structToMap(original)
	diff = make(map[string]interface{})
	for name, v1 := range map1 {
		v2 := map2[name]
		if !reflect.DeepEqual(v1, v2) {
			diff[name] = v1
		}
	}
	return
}

// IncludesString check the value(string) is in collection
func IncludesString(collection []string, value string) (found bool) {
	for _, v := range collection {
		if !found && v == value {
			found = true
		}
	}
	return
}

// IncludesInt check the value(int) is in collection
func IncludesInt(collection []int, value int) (found bool) {
	for _, v := range collection {
		if !found && v == value {
			found = true
		}
	}
	return
}

// Pick get the data from struct
func Pick(current interface{}, fields []string) (data map[string]interface{}) {
	data = make(map[string]interface{})
	m := structToMap(current)
	for k, v := range m {
		if IncludesString(fields, k) {
			data[k] = v
		}
	}
	return
}

// IsString check the value type is string
func IsString(i interface{}) bool {
	return getType(i) == "string"
}

// IsInt check the value type is int
func IsInt(i interface{}) bool {
	return getType(i) == "int"
}

// IsBool check the value type is bool
func IsBool(i interface{}) bool {
	return getType(i) == "bool"
}

// GetString get the first string param from arguments
func GetString(args ...interface{}) (value string, found bool) {
	for _, v := range args {
		if !found && IsString(v) {
			value = v.(string)
			found = true
		}
	}
	return
}

// GetBool get the first bool param from arguments
func GetBool(args ...interface{}) (value bool, found bool) {
	for _, v := range args {
		if !found && IsBool(v) {
			value = v.(bool)
			found = true
		}
	}
	return
}

// GetInt get the first int param from arguments
func GetInt(args ...interface{}) (value int, found bool) {
	for _, v := range args {
		if !found && IsInt(v) {
			value = v.(int)
			found = true
		}
	}
	return
}
