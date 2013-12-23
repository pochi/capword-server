package capword

import (
  "appengine"
  "appengine/datastore"
  "encoding/json"
  "net/http"
)

type Word struct {
  Id   int64
  Name string
}

type Words struct {
  Elements []Word
}

func init() {
  http.HandleFunc("/words/index", index)
  http.HandleFunc("/words/create", create)
}

func index(response http.ResponseWriter, request *http.Request) {
  context := appengine.NewContext(request)
  key := datastore.NewKey(context, "Words", "default_words", 0, nil)
  query := datastore.NewQuery("Words").Ancestor(key).Order("Name").Limit(10)
  words := make([]Word, 0, 10)
  keys, _ := query.GetAll(context, &words)
  for i := 0; i < len(words); i++ {
    words[i].Id = keys[i].IntID()
  }

  response_words := Words{Elements: words}
  response.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(response)
  encoder.Encode(response_words)
}

func create(response http.ResponseWriter, request *http.Request) {
  context := appengine.NewContext(request)
  context.Infof("-----------------")
  context.Infof("Request url: %v", request.FormValue("Name"))
  word := Word{Name: request.FormValue("Name")}
  base_key := datastore.NewKey(context, "Words", "default_words", 0, nil)
  key := datastore.NewIncompleteKey(context, "Words", base_key)
  datastore_word, error := datastore.Put(context, key, &word)
  if error != nil {
    http.Error(response, error.Error(), http.StatusInternalServerError)
    return
  }
  response.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(response)
  encoder.Encode(datastore_word)
}
