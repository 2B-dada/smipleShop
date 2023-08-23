// Code generated by Kitex v0.7.0. DO NOT EDIT.

package shop

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "main/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return shopServiceInfo
}

var shopServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "shop"
	handlerType := (*api.Shop)(nil)
	methods := map[string]kitex.MethodInfo{
		"createProtuct":  kitex.NewMethodInfo(createProtuctHandler, newShopCreateProtuctArgs, newShopCreateProtuctResult, false),
		"requsetAll":     kitex.NewMethodInfo(requsetAllHandler, newShopRequsetAllArgs, newShopRequsetAllResult, false),
		"requsetByName":  kitex.NewMethodInfo(requsetByNameHandler, newShopRequsetByNameArgs, newShopRequsetByNameResult, false),
		"requsetByPrice": kitex.NewMethodInfo(requsetByPriceHandler, newShopRequsetByPriceArgs, newShopRequsetByPriceResult, false),
		"buy":            kitex.NewMethodInfo(buyHandler, newShopBuyArgs, newShopBuyResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "api",
		"ServiceFilePath": "shop.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}

func createProtuctHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShopCreateProtuctArgs)
	realResult := result.(*api.ShopCreateProtuctResult)
	success, err := handler.(api.Shop).CreateProtuct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopCreateProtuctArgs() interface{} {
	return api.NewShopCreateProtuctArgs()
}

func newShopCreateProtuctResult() interface{} {
	return api.NewShopCreateProtuctResult()
}

func requsetAllHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*api.ShopRequsetAllResult)
	success, err := handler.(api.Shop).RequsetAll(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopRequsetAllArgs() interface{} {
	return api.NewShopRequsetAllArgs()
}

func newShopRequsetAllResult() interface{} {
	return api.NewShopRequsetAllResult()
}

func requsetByNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShopRequsetByNameArgs)
	realResult := result.(*api.ShopRequsetByNameResult)
	success, err := handler.(api.Shop).RequsetByName(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopRequsetByNameArgs() interface{} {
	return api.NewShopRequsetByNameArgs()
}

func newShopRequsetByNameResult() interface{} {
	return api.NewShopRequsetByNameResult()
}

func requsetByPriceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShopRequsetByPriceArgs)
	realResult := result.(*api.ShopRequsetByPriceResult)
	success, err := handler.(api.Shop).RequsetByPrice(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopRequsetByPriceArgs() interface{} {
	return api.NewShopRequsetByPriceArgs()
}

func newShopRequsetByPriceResult() interface{} {
	return api.NewShopRequsetByPriceResult()
}

func buyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShopBuyArgs)
	realResult := result.(*api.ShopBuyResult)
	success, err := handler.(api.Shop).Buy(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopBuyArgs() interface{} {
	return api.NewShopBuyArgs()
}

func newShopBuyResult() interface{} {
	return api.NewShopBuyResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateProtuct(ctx context.Context, req *api.CreateRequest) (r *api.CreateResponse, err error) {
	var _args api.ShopCreateProtuctArgs
	_args.Req = req
	var _result api.ShopCreateProtuctResult
	if err = p.c.Call(ctx, "createProtuct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RequsetAll(ctx context.Context) (r *api.ResqResponse, err error) {
	var _args api.ShopRequsetAllArgs
	var _result api.ShopRequsetAllResult
	if err = p.c.Call(ctx, "requsetAll", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RequsetByName(ctx context.Context, req *api.ProductReqByName) (r *api.ResqResponse, err error) {
	var _args api.ShopRequsetByNameArgs
	_args.Req = req
	var _result api.ShopRequsetByNameResult
	if err = p.c.Call(ctx, "requsetByName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RequsetByPrice(ctx context.Context, req *api.ProductReqByPrice) (r *api.ResqResponse, err error) {
	var _args api.ShopRequsetByPriceArgs
	_args.Req = req
	var _result api.ShopRequsetByPriceResult
	if err = p.c.Call(ctx, "requsetByPrice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Buy(ctx context.Context, req *api.BuyReq) (r *api.BuyResp, err error) {
	var _args api.ShopBuyArgs
	_args.Req = req
	var _result api.ShopBuyResult
	if err = p.c.Call(ctx, "buy", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
