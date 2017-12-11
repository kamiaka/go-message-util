# go-message-util

go-message-util is a utility for predefining the message catalog.

This package uses `golang.org/x/text/message` package internally.

## Usage

### Defining messages

```go
var Japanese = []Expr{
  String("Hello, %s", "こんにちは、%s"),
  String("Good bye, %s", "さようなら、%s"),
}
```

### Set messages to catalog

```go
c := catalog.NewBuilder()
SetLanguage(c, language.Japanese, Japanese)
```

### Print translated message

```go
t := language.Japanese
p := message.NewPrinter(t, message.Catalog(c))

name := "世界"
p.Printf("Hello, %s", p.Sprintf(name)) // こんにちは、世界
```

## License

MIT