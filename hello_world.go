package hello

import (
	"fmt"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
