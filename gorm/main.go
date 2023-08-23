package main

import (
	"log"
	"main/db"
	api "main/kitex_gen/api/shop"
)

func main() {
	db.Init()

	svr := api.NewServer(new(ShopImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
