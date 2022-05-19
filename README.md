# Golang Notes

### <u>Find out type of variable</u>

```golang
import (
  "fmt"
  "reflect"
)

i := 10
s := "Hello"

fmt.Println(reflect.TypeOf(i))
fmt.Println(reflect.TypeOf(s))
```
