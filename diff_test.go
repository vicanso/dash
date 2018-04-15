package structDiff

import (
	"reflect"
	"testing"
)

// Book book struct
type Book struct {
	Name  string
	Price float64
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
		update := Diff(b1, b2)
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
		e := reflect.ValueOf(b2).Elem()
		typeOf := e.Type()
		data := make(map[string]interface{})
		for index := 0; index < e.NumField(); index++ {
			name := typeOf.Field(index).Name
			value := e.Field(index).Interface()
			data[name] = value
		}
		update := Diff(b1, data)
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
		update := Diff(bs1, bs2)
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
		update := Diff(bs1, bs2)
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
		update := Diff(bs1, bs2)
		if update["NO"].(int) != bs1.NO {
			t.Fatalf("diff fail, the NO should be diff")
		}

		if len(update) != 2 {
			t.Fatalf("diff fail, diff fields should be 2")
		}
	})
}
