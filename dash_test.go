package dash

import (
	"math/cmplx"
	"strings"
	"testing"
	"unsafe"

	"github.com/fatih/structs"
)

// Book book struct
type Book struct {
	Name   string
	Price  float64
	Author string
}

// BookShelf book shelf struct
type BookShelf struct {
	NO    int
	Books []*Book
}

// GetName get the book name
func (b *Book) GetName() string {
	return b.Name
}

func TestDiff(t *testing.T) {
	t.Run("struct to map", func(t *testing.T) {
		b1 := &Book{
			Name:   "book1",
			Price:  123.12,
			Author: "abc",
		}
		m := ToMap(b1)
		keys := []string{}
		for k := range m {
			keys = append(keys, k)
		}
		if strings.Join(keys, ",") != "Name,Price,Author" {
			t.Fatalf("struct to map fail")
		}
	})
	t.Run("diff normal struct", func(t *testing.T) {
		b1 := &Book{
			Name:  "book1",
			Price: 123.12,
		}
		b2 := &Book{
			Name:  "book2",
			Price: 123.15,
		}
		update := Difference(b1, b2)
		if len(update) != 2 {
			t.Fatalf("diff fail, diff field should be 2")
		}
		if update["Name"].(string) != b1.Name {
			t.Fatalf("diff fail, the name should be diff")
		}
		if update["Price"].(float64) != b1.Price {
			t.Fatalf("diff fail, the price should be diff")
		}
	})

	t.Run("diff to map[string]interface{}", func(t *testing.T) {
		b1 := &Book{
			Name:  "book1",
			Price: 123.12,
		}
		b2 := &Book{
			Name:  "book2",
			Price: 123.15,
		}
		data := structs.Map(b2)
		data["new-filed"] = "new-value"
		update := Difference(b1, data)
		if len(update) != 2 {
			t.Fatalf("diff fail, diff field should be 2")
		}
		if update["Name"].(string) != b1.Name {
			t.Fatalf("diff fail, the name should be diff")
		}
		if update["Price"].(float64) != b1.Price {
			t.Fatalf("diff fail, the price should be diff")
		}
	})

	t.Run("diff nesting struct(same)", func(t *testing.T) {
		b1 := &Book{
			Name:  "book1",
			Price: 123.12,
		}
		b2 := &Book{
			Name:  "book2",
			Price: 123.15,
		}
		bs1 := &BookShelf{
			NO: 1,
			Books: []*Book{
				b1,
				b2,
			},
		}
		bs2 := &BookShelf{
			NO: 2,
			Books: []*Book{
				b1,
				b2,
			},
		}
		update := Difference(bs1, bs2)
		if update["NO"].(int) != bs1.NO {
			t.Fatalf("diff fail, the NO should be diff")
		}

		if len(update) != 1 {
			t.Fatalf("diff fail, diff fields should be 1")
		}
	})

	t.Run("diff nesting struct(diff stuct same data)", func(t *testing.T) {
		b1 := &Book{
			Name:  "book1",
			Price: 123.12,
		}
		b2 := &Book{
			Name:  "book1",
			Price: 123.12,
		}
		bs1 := &BookShelf{
			NO: 1,
			Books: []*Book{
				b1,
			},
		}
		bs2 := &BookShelf{
			NO: 2,
			Books: []*Book{
				b2,
			},
		}
		update := Difference(bs1, bs2)
		if update["NO"].(int) != bs1.NO {
			t.Fatalf("diff fail, the NO should be diff")
		}

		if len(update) != 1 {
			t.Fatalf("diff fail, diff fields should be 1")
		}
	})

	t.Run("diff nesting struct(diff stuct diff data)", func(t *testing.T) {
		b1 := &Book{
			Name:  "book1",
			Price: 123.12,
		}
		b2 := &Book{
			Name:  "book2",
			Price: 123.15,
		}
		bs1 := &BookShelf{
			NO: 1,
			Books: []*Book{
				b1,
			},
		}
		bs2 := &BookShelf{
			NO: 2,
			Books: []*Book{
				b2,
			},
		}
		update := Difference(bs1, bs2)
		if update["NO"].(int) != bs1.NO {
			t.Fatalf("diff fail, the NO should be diff")
		}

		if len(update) != 2 {
			t.Fatalf("diff fail, diff fields should be 2")
		}
	})
}

func TestIncludes(t *testing.T) {
	t.Run("IncludesInt", func(t *testing.T) {
		found := IncludesInt([]int{1, 2, 3}, 1)
		if !found {
			t.Fatalf("includes int check found fail")
		}

		found = IncludesInt([]int{1, 2, 3}, 4)
		if found {
			t.Fatalf("includes int check not found fail")
		}

		found = IncludesUint8([]uint8{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes uint8 check found fail")
		}

		found = IncludesUint8([]uint8{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes uint8 check found fail")
		}

		found = IncludesUint16([]uint16{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes uint16 check found fail")
		}

		found = IncludesUint16([]uint16{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes uint16 check found fail")
		}

		found = IncludesUint32([]uint32{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes uint32 check found fail")
		}

		found = IncludesUint32([]uint32{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes uint32 check found fail")
		}

		found = IncludesUint64([]uint64{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes uint64 check found fail")
		}

		found = IncludesUint64([]uint64{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes uint64 check found fail")
		}

		found = IncludesInt8([]int8{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes int8 check found fail")
		}

		found = IncludesInt8([]int8{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes int8 check found fail")
		}

		found = IncludesInt16([]int16{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes int16 check found fail")
		}

		found = IncludesInt16([]int16{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes int16 check found fail")
		}

		found = IncludesInt32([]int32{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes int32 check found fail")
		}

		found = IncludesInt32([]int32{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes int32 check found fail")
		}

		found = IncludesInt64([]int64{0, 1, 2}, 1)
		if !found {
			t.Fatalf("includes int64 check found fail")
		}

		found = IncludesInt64([]int64{0, 1, 2}, 3)
		if found {
			t.Fatalf("includes int64 check found fail")
		}
	})

	t.Run("IncludesString", func(t *testing.T) {
		found := IncludesString([]string{"a", "b", "c"}, "b")
		if !found {
			t.Fatalf("includes string check found fail")
		}

		found = IncludesString([]string{"a", "b", "c"}, "d")
		if found {
			t.Fatalf("includes string check not found fail")
		}

	})
}

func TestPick(t *testing.T) {
	t.Run("Pick", func(t *testing.T) {
		b := &Book{
			Name:   "book1",
			Price:  1234.12,
			Author: "robot",
		}
		fields := []string{"Name", "Price"}
		data := Pick(b, fields)
		if len(data) != len(fields) {
			t.Fatalf("pick data fail")
		}
		if data["Name"] != b.Name {
			t.Fatalf("pick data fail, the name is not equal")
		}
		if data["Price"] != b.Price {
			t.Fatalf("pick data fail, the price is not equal")
		}
	})
}

func TestIs(t *testing.T) {
	t.Run("is string", func(t *testing.T) {
		if !IsString("abc") {
			t.Fatalf("is string check fail")
		}
		if IsString(1) {
			t.Fatalf("is string check fail")
		}
	})

	t.Run("is bool", func(t *testing.T) {
		if !IsBool(true) {
			t.Fatalf("is bool check fail")
		}

		if IsBool(0) {
			t.Fatalf("is bool check fail")
		}
	})

	t.Run("is int", func(t *testing.T) {
		if !IsInt(1) {
			t.Fatalf("is int check fail")
		}
		if IsInt('a') {
			t.Fatalf("is int check fail")
		}
	})

	t.Run("is int8", func(t *testing.T) {
		var v int8 = 10
		if !IsInt8(v) {
			t.Fatalf("is int8 check fail")
		}

		if IsInt8(10) {
			t.Fatalf("is int8 check fail")
		}
	})

	t.Run("is int16", func(t *testing.T) {
		var v int16 = 10
		if !IsInt16(v) {
			t.Fatalf("is int16 check fail")
		}

		if IsInt16(10) {
			t.Fatalf("is int16 check fail")
		}
	})

	t.Run("is int32", func(t *testing.T) {
		var v int32 = 10
		if !IsInt32(v) {
			t.Fatalf("is int32 check fail")
		}

		if !IsInt32('a') {
			t.Fatalf("is int32 check fail")
		}

		if IsInt32(0) {
			t.Fatalf("is int32 check fail")
		}
	})

	t.Run("is int64", func(t *testing.T) {
		var v int64 = 10
		if !IsInt64(v) {
			t.Fatalf("is int64 check fail")
		}

		if IsInt64(10) {
			t.Fatalf("is int64 check fail")
		}
	})

	t.Run("is uint", func(t *testing.T) {
		var v uint = 10
		if !IsUint(v) {
			t.Fatalf("is uint check fail")
		}

		if IsUint(10) {
			t.Fatalf("is uint check fail")
		}
	})

	t.Run("is uint8", func(t *testing.T) {
		if !IsUint8([]byte("a")[0]) {
			t.Fatalf("is uint8 check fail")
		}

		if IsUint8(0) {
			t.Fatalf("is uint8 check fail")
		}
	})

	t.Run("is uint16", func(t *testing.T) {
		var v uint16 = 10
		if !IsUint16(v) {
			t.Fatalf("is uint16 check fail")
		}
		if IsUint16(10) {
			t.Fatalf("is uint16 check fail")
		}
	})

	t.Run("is uint32", func(t *testing.T) {
		var v uint32 = 10
		if !IsUint32(v) {
			t.Fatalf("is uint32 check fail")
		}

		if IsUint32(10) {
			t.Fatalf("is uint32 check fail")
		}
	})

	t.Run("is uint64", func(t *testing.T) {
		var v uint64 = 10
		if !IsUint64(v) {
			t.Fatalf("is uint64 check fail")
		}

		if IsUint64(10) {
			t.Fatalf("is uint64 check fail")
		}
	})

	t.Run("is uintptr", func(t *testing.T) {
		i := 10
		ptr := uintptr(unsafe.Pointer(&i))
		if !IsUintptr(ptr) {
			t.Fatalf("is uintptr check fail")
		}

		if IsUintptr(&i) {
			t.Fatalf("is uintptr check fail")
		}
	})

	t.Run("is float32", func(t *testing.T) {
		var v float32 = 10.1
		if !IsFloat32(v) {
			t.Fatalf("is float32 check fail")
		}

		if IsFloat32(10.1) {
			t.Fatalf("is float32 check fail")
		}
	})

	t.Run("is float64", func(t *testing.T) {
		if !IsFloat64(10.1) {
			t.Fatalf("is float64 check fail")
		}

		if IsFloat64(10) {
			t.Fatalf("is float64 check fail")
		}
	})

	t.Run("is complex128", func(t *testing.T) {
		v := cmplx.Sqrt(-5 + 12i)
		if !IsComplex128(v) {
			t.Fatalf("is complex128 check fail")
		}
	})
}

func TestGetParam(t *testing.T) {
	t.Run("get string param", func(t *testing.T) {
		s, found := GetString(1, true, "name")
		if !found || s != "name" {
			t.Fatalf("get string fail")
		}
	})

	t.Run("get bool param", func(t *testing.T) {
		b, found := GetBool(1, true, "name")
		if !found || !b {
			t.Fatalf("get bool fail")
		}
	})

	t.Run("get int param", func(t *testing.T) {
		i, found := GetInt(1, true, "name")
		if !found || i != 1 {
			t.Fatalf("get int fail")
		}
	})

	t.Run("get int8 param", func(t *testing.T) {
		i, found := GetInt8(int8(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get int8 fail")
		}
	})

	t.Run("get int16 param", func(t *testing.T) {
		i, found := GetInt16(int16(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get int16 fail")
		}
	})

	t.Run("get int32 param", func(t *testing.T) {
		i, found := GetInt32(int32(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get int32 fail")
		}
	})

	t.Run("get int64 param", func(t *testing.T) {
		i, found := GetInt64(int64(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get int64 fail")
		}
	})

	t.Run("get uint8 param", func(t *testing.T) {
		i, found := GetUint8(uint8(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get uint8 fail")
		}
	})

	t.Run("get uint16 param", func(t *testing.T) {
		i, found := GetUint16(uint16(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get uint16 fail")
		}
	})

	t.Run("get uint32 param", func(t *testing.T) {
		i, found := GetUint32(uint32(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get uint32 fail")
		}
	})

	t.Run("get uint64 param", func(t *testing.T) {
		i, found := GetUint64(uint64(1), true, "name")
		if !found || i != 1 {
			t.Fatalf("get uint64 fail")
		}
	})
}

func TestFill(t *testing.T) {
	t.Run("fill with no such field", func(t *testing.T) {
		b := &Book{}
		m := make(map[string]interface{})
		m["Name"] = "book"
		m["ISBN"] = "123123123"
		err := Fill(b, m)
		if err == nil || err.Error() != "No such field: ISBN in dest" {
			t.Fatalf("fill should be return no such field error")
		}
	})
	t.Run("fill success", func(t *testing.T) {
		b := &Book{}
		m := make(map[string]interface{})
		m["Name"] = "book"
		m["Price"] = 100.12
		err := Fill(b, m)
		if err != nil {
			t.Fatalf("fill fail, %v", err)
		}
		if b.Name != m["Name"] || b.Price != m["Price"] {
			t.Fatalf("fill fail")
		}
	})

	t.Run("fill by first letter upper", func(t *testing.T) {
		b := &Book{}
		m := make(map[string]interface{})
		m["name"] = "book"
		m["price"] = 100.12
		err := Fill(b, m, true)
		if err != nil {
			t.Fatalf("fill fail, %v", err)
		}
		if b.Name != m["name"] || b.Price != m["price"] {
			t.Fatalf("fill fail")
		}
	})
}

func TestUniq(t *testing.T) {
	t.Run("uniq int", func(t *testing.T) {
		data := UniqInt([]int{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq int fail, result is:%v", data)
		}
	})

	t.Run("uniq int8", func(t *testing.T) {
		data := UinqInt8([]int8{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq int8 fail, result is:%v", data)
		}
	})

	t.Run("uniq int16", func(t *testing.T) {
		data := UinqInt16([]int16{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq int16 fail, result is:%v", data)
		}
	})

	t.Run("uniq int32", func(t *testing.T) {
		data := UinqInt32([]int32{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq int32 fail, result is:%v", data)
		}
	})

	t.Run("uniq int64", func(t *testing.T) {
		data := UinqInt64([]int64{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq int64 fail, result is:%v", data)
		}
	})

	t.Run("uniq uint8", func(t *testing.T) {
		data := UinqUint8([]uint8{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq uint8 fail, result is:%v", data)
		}
	})

	t.Run("uniq uint16", func(t *testing.T) {
		data := UinqUint16([]uint16{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq uint16 fail, result is:%v", data)
		}
	})

	t.Run("uniq uint32", func(t *testing.T) {
		data := UinqUint32([]uint32{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq uint32 fail, result is:%v", data)
		}
	})

	t.Run("uniq uint64", func(t *testing.T) {
		data := UinqUint64([]uint64{1, 4, 1, 3, 4, 5})
		if len(data) != 4 {
			t.Fatalf("uniq uint64 fail, result is:%v", data)
		}
	})

	t.Run("uniq string", func(t *testing.T) {
		data := UinqString([]string{
			"a",
			"b",
			"c",
			"d",
			"a",
			"c",
		})
		if len(data) != 4 {
			t.Fatalf("uniq string fail")
		}
	})
}
