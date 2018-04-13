# struct-diff

[![Build Status](https://travis-ci.org/vicanso/struct-diff.svg?branch=master)](https://travis-ci.org/vicanso/struct-diff)


get the difference between struct's instance

## API

```go
// Book book struct
type Book struct {
	Name  string
	Price float64
}
b1 := &Book{
  Name:  "book1",
  Price: 123.12,
}
b2 := &Book{
  Name:  "book2",
  Price: 123.15,
}
update := Diff(b1, b2)
// map[Name:book1 Price:123.12]
fmt.Println(update)

b3 := &Book{
  Name: "book3",
  Price: 123.12,
}
update = Diff(b1, b3)
// map[Name:book1]
fmt.Println(update)
```