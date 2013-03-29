# Goes

A proof of concept for a simple URL routing API with Go

Example code:

```go
package main

import (
    "fmt"
    "github.com/mattupstate/goes"
)

func home() string {
    return "Hello World"
}

func products(a struct{Category string}) string {
    return fmt.Sprintf("The product category is %s", a.Category)
}

func main() {
    app := goes.App{}
    app.Route("/", home)
    app.Route("/products/{Category}", products)
    app.Run("0.0.0.0", 8080)
}
```