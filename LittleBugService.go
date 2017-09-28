package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type HTMLObj struct{
  PageTitle string
  Word string
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path[1:] == "" {
    fmt.Fprintf(w, "You need fix a little bugs")
  } else {
    output := HTMLObj{
      PageTitle: "You are fixing the bugs",
      Word : r.URL.Path[1:],
    }

    tmpl := template.Must(template.ParseFiles("template/layout.html"))
    tmpl.Execute(w, output)
  }
}

func main(){
    http.HandleFunc("/", handler)
    fmt.Println("Let's Start")
    http.ListenAndServe(":8080", nil)
}