# Gorb
shitty basic rest api boostrap

## how to get

```sh
go get github.com/Kaurikallaste/gorb
```

You might need to run go mod tidy with GOPRIVATE env param

```sh
GOPRIVATE=github.com/Kaurikallaste/gorb go mod tidy
```


## basic example how to use, controllers and registerers should be in their own submodules
```go
package main

import (
	"fmt"
	"net/http"
	"github.com/Kaurikallaste/gorb"
)

func controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func registerer() {
	http.HandleFunc("/", controller)
}

func main() {
	gorb.Bootstrap([]gorb.Registerer{registerer})
}
```
