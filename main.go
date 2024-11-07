package main

import (
    "fmt"
    "net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var reward = "Reward myself with Dreuge worst"
var taskItems = []string{shortGolang, fullGolang, reward}


func main() {

  
  fmt.Println("####### Welcome to our TodoList App #######")
  
  http.HandleFunc("/", helloUser)
  http.HandleFunc("/show-tasks", showTasks)

  http.ListenAndServe(":8080", nil)



}

func helloUser(writer http.ResponseWriter,request *http.Request) {
  var greeting = "Hello user. Welcome to our TodoList App"
  fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter,request *http.Request) {
  for _, task := range taskItems { 
  fmt.Fprintln(writer, task)
  }
}
