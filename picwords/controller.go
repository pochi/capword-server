package picwords

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/words", index)
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, words")
}
