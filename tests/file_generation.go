package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

type Page struct {
  Title string
  Num   int
  Body  []byte
}

func (p *Page) save() error {
  filename := p.Title + fmt.Sprintf("%d.txt", p.Num)
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func removeFile(title string, i int) {
  filename := "./" + title + fmt.Sprintf("%d.txt", i)
  os.Remove(filename)
  if err := os.Remove(filename); err != nil {
    return
  }
}

func loadPage(title string, i int) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  p := &Page{Title: title, Num: i, Body: []byte(body)}
  err = p.save()
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Num: i, Body: []byte(body)}, nil
}

var done = make(chan bool)

func main() {

  args := os.Args

  if args[1] == "write" {
    fmt.Printf("--------------------------")
    for i := 0; i < 100000; i++ {
      go loadPage("test", i)
    }
    <-done
  }

  if args[1] == "delete" {
    fmt.Printf("--------------------------")
    for i := 0; i < 100000; i++ {
      go removeFile("test", i)
    }
    <-done
  }

}
