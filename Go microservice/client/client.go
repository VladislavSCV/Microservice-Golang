package main

import (
	"log"
	"fmt"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}


func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	client.Call("API.GetDB", "", &db)
	a := Item{"hello", "world"}
	b := Item{"hello1", "world1"}
	c := Item{"hello2", "world2"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)

	client.Call("API.GetDB", "", &db)

	fmt.Println("Db: ", db)

}