# evtool
[Golang](http://golang.org/) package for email validation and gmail email normalization.

 - Format (simple regexp, "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
 - Valid domain
 - Normalize Gmail aliases into real address. (Removes periods and anything starting from + sign forward)

[![GoDoc](https://godoc.org/github.com/danarenasb.evtool?status.png)](https://godoc.org/github.com/danarenasb.evtool)

## Usage

### 1. Format & Valid Domain
```go
func main() {
    valid := evtool.IsEmailValid("ç$€§/az@gmail.com")
    if !valid {
        // If not valid do something
    }
    // If valid move on and continue
}
```


### 2. Normalize Gmail aliases
```go
func main() {
    email, err := evtool.NormalizeGmail("e.m.a.i.l.+alsoemail@gmail.com")
    if err != nil {
        fmt.Println(err)
        // If not a gmail address will return error
    }
    // If no error will return the email address normalized from the alias
}
```

## License

Evtool is licensed under the [MIT License](./LICENSE).