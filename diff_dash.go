package structDiff

import (
	"fmt"
	"testing"

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
		fmt.Println(data["Name"])
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

	t.Run("is int", func(t *testing.T) {
		if !IsInt(1) {
			t.Fatalf("is int check fail")
		}
		if IsInt('a') {
			t.Fatalf("is int check fail")
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
}
