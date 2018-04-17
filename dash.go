package structDiff

import (
	"reflect"

	"github.com/fatih/structs"
)

func getType(i interface{}) (valueType string) {
	switch i.(type) {
	case string:
		valueType = "string"
		break
	case bool:
		valueType = "bool"
		break
	case int:
		valueType = "int"
		break
	case int8:
		valueType = "int8"
		break
	case int16:
		valueType = "int16"
		break
	case int32:
		valueType = "int32"
		break
	case int64:
		valueType = "int64"
	case uint:
		valueType = "uint"
		break
	case uint8:
		valueType = "uint8"
		break
	case uint16:
		valueType = "uint16"
		break
	case uint32:
		valueType = "uint32"
		break
	case uint64:
		valueType = "uint64"
		break
	case uintptr:
		valueType = "uintptr"
		break
	case float32:
		valueType = "float32"
		break
	case float64:
		valueType = "float64"
		break
	case complex64:
		valueType = "complex64"
		break
	case complex128:
		valueType = "complex128"
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

// IsInt8 check the value type is int8
func IsInt8(i interface{}) bool {
	return getType(i) == "int8"
}

// IsInt16 check the value type is int16
func IsInt16(i interface{}) bool {
	return getType(i) == "int16"
}

// IsInt32 check the value type is int32
func IsInt32(i interface{}) bool {
	return getType(i) == "int32"
}

// IsInt64 check the value type is int64
func IsInt64(i interface{}) bool {
	return getType(i) == "int64"
}

// IsBool check the value type is bool
func IsBool(i interface{}) bool {
	return getType(i) == "bool"
}

// IsUint8 check the value type is uint8
func IsUint8(i interface{}) bool {
	return getType(i) == "uint8"
}

// IsUint16 check the value type is uint16
func IsUint16(i interface{}) bool {
	return getType(i) == "uint16"
}

// IsUint32 check the value type is uint32
func IsUint32(i interface{}) bool {
	return getType(i) == "uint32"
}

// IsUint64 check the value type is uint64
func IsUint64(i interface{}) bool {
	return getType(i) == "uint64"
}

// IsUintptr check the value type is uintptr
func IsUintptr(i interface{}) bool {
	return getType(i) == "uintptr"
}

// IsFloat32 check the value type is float32
func IsFloat32(i interface{}) bool {
	return getType(i) == "float32"
}

// IsFloat64 check the value type is float64
func IsFloat64(i interface{}) bool {
	return getType(i) == "float64"
}

// IsComplex64 check the value type is complex64
func IsComplex64(i interface{}) bool {
	return getType(i) == "complex64"
}

// IsComplex128 check the value type is complex128
func IsComplex128(i interface{}) bool {
	return getType(i) == "complex128"
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

// GetInt8 get the first int8 param from arguments
func GetInt8(args ...interface{}) (value int8, found bool) {
	for _, v := range args {
		if !found && IsInt8(v) {
			value = v.(int8)
			found = true
		}
	}
	return
}

// GetInt16 get the first int16 param from arguments
func GetInt16(args ...interface{}) (value int16, found bool) {
	for _, v := range args {
		if !found && IsInt16(v) {
			value = v.(int16)
			found = true
		}
	}
	return
}

// GetInt32 get the first int32 param from arguments
func GetInt32(args ...interface{}) (value int32, found bool) {
	for _, v := range args {
		if !found && IsInt32(v) {
			value = v.(int32)
			found = true
		}
	}
	return
}

// GetInt64 get the first int64 param from arguments
func GetInt64(args ...interface{}) (value int64, found bool) {
	for _, v := range args {
		if !found && IsInt64(v) {
			value = v.(int64)
			found = true
		}
	}
	return
}

// GetUint8 get the first uint8 param from arguments
func GetUint8(args ...interface{}) (value uint8, found bool) {
	for _, v := range args {
		if !found && IsUint8(v) {
			value = v.(uint8)
			found = true
		}
	}
	return
}

// GetUint16 get the first uint16 param from arguments
func GetUint16(args ...interface{}) (value uint16, found bool) {
	for _, v := range args {
		if !found && IsUint16(v) {
			value = v.(uint16)
			found = true
		}
	}
	return
}

// GetUint32 get the first uint32 param from arguments
func GetUint32(args ...interface{}) (value uint32, found bool) {
	for _, v := range args {
		if !found && IsUint32(v) {
			value = v.(uint32)
			found = true
		}
	}
	return
}

// GetUint64 get the first uint64 param from arguments
func GetUint64(args ...interface{}) (value uint64, found bool) {
	for _, v := range args {
		if !found && IsUint64(v) {
			value = v.(uint64)
			found = true
		}
	}
	return
}
