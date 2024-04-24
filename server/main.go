package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// func helloHandler(w http.ResponseWriter,r * http.Request){

// 	fmt.Println(r.URL.Path,r.Method)

// 	if r.URL.Path == "/gopi" {
//        fmt.Fprintf(w,"gopi coming")
// 	   return
// 	}

// 	fmt.Fprintf(w, "Hello, World!")

// }

type Product struct {
    ID          int     `json:"id"`
    ProductName string  `json:"productName"`
    Image       string  `json:"image"`
    From        string  `json:"from"`
    Nutrients   string  `json:"nutrients"`
    Quantity    string  `json:"quantity"`
    Price       string  `json:"price"`
    Organic     bool    `json:"organic"`
    Description string  `json:"description"`
}


func (p * Product) load() {
  data,err :=  os.ReadFile("./dev-data/data.json")

  fmt.Println(data)
  if err != nil{
	fmt.Println("something went wrong")
  }
}

func main(){
	
	
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/",fileServer)

	


	fmt.Println("server up on port 8080")

	if err:= http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)  
	}


}