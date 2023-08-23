package main

import (
	"context"
	"log"
	api "main/kitex_gen/api"

	db "main/db"
)

// ShopImpl implements the last service interface defined in the IDL.
type ShopImpl struct{}

func (s *ShopImpl) RequsetAll(ctx context.Context) (resp *api.ResqResponse, err error) {
	ret, err := db.GetAllProduct(ctx)
	if err != nil {
		//具体错误类懒得写了
		resp.Code = 101
	}
	resp = &api.ResqResponse{Items: ret, Code: 0}
	return
}

func (s *ShopImpl) RequsetByName(ctx context.Context, req *api.ProductReqByName) (resp *api.ResqResponse, err error) {
	ret, err := db.GetProductByName(ctx, req.Name)
	if err != nil {
		resp.Code = 101
	}
	resp = &api.ResqResponse{Items: ret, Code: 0}
	return
}

func (s *ShopImpl) RequsetByPrice(ctx context.Context, req *api.ProductReqByPrice) (resp *api.ResqResponse, err error) {
	ret, err := db.GetProductByPrice(ctx, req.Price1, req.Price2)
	if err != nil {
		resp.Code = 101
	}
	resp = &api.ResqResponse{Items: ret, Code: 0}
	return
}

func (s *ShopImpl) Buy(ctx context.Context, req *api.BuyReq) (resp *api.BuyResp, err error) {
	er := db.BuyProduct(ctx, req.ID)
	if er != nil {
		resp.Code = 301
	}
	resp = &api.BuyResp{Code: 0}
	return
}

func (s *ShopImpl) CreateProtuct(ctx context.Context, req *api.CreateRequest) (resp *api.CreateResponse, err error) {
	to := &db.Product{}
	to.ID = int(req.Item.ID)
	to.Name = req.Item.Name
	to.Number = int(req.Item.Number)
	to.Price = int(req.Item.Price)
	er := db.CreateNewProtuct(ctx, to)
	log.Println("called")
	if er != nil {
		resp.Code = 201
	}
	resp = &api.CreateResponse{Code: 0}
	return
}
