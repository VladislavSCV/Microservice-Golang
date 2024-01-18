package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}


type API int

var database []Item

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}


func (a *API) GetItem(title string) (Item, error) {
	var getItem Item

	for _, item := range database{
		if item.Title == title {
			getItem = item
		}
	}
	return getItem, nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, val := range database{
		if val.Title == edit.Title {
			database[idx] = edit
			changed = database[idx]
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database{
		if val.Body == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}


func main() {
	var api = new(API)

	err := rpc.Register(api)
	if err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on: %s", "8000")
	
	err = http.Serve(listen, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(database)

}