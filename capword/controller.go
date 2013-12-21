package capword

import (
  "encoding/json"
  "net/http"
)

type Word struct {
  Name string
}

func init() {
  http.HandleFunc("/words", index)
}

func index(response http.ResponseWriter, request *http.Request) {
  response.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(response)
  word := Word{"happy"}
  encoder.Encode(word)
}
