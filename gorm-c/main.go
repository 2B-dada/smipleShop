package main

import (
	"context"
	"log"
	"main/kitex_gen/api"
	"main/kitex_gen/api/shop"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := shop.NewClient("shop", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	client.Buy(context.Background(), &api.BuyReq{ID: 2})
}
