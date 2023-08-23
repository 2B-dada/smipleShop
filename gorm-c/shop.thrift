namespace go api

enum ErrCode {
    Success     =   0,
    ServiceErr  =   101,
    ParamErr    =   102,
    NotFoundErr =   104,
    CreateErr   =   201,
    BuyErr      =   301
}

struct Item {
    1: string Name,
    2: i64 ID,
    3: i64 Price,
    4: i64 Number
}

struct CreateRequest {
    1: Item item
}

struct CreateResponse {
    1: i64 Code
}

struct ResqResponse {
    1: i64 Code
    2: list<Item> Items
}

struct ProductReqByPrice {
    1: i64 Price1,
    2: i64 Price2
}

struct ProductReqByName {
    1: string Name
}

struct BuyReq {
    1: i64 ID
}

struct BuyResp {
    1: i64 Code
}

service shop {
    CreateResponse createProtuct(1: CreateRequest req),

    ResqResponse requsetAll(),
    ResqResponse requsetByName(1: ProductReqByName req),
    ResqResponse requsetByPrice(1: ProductReqByPrice req),
    
    BuyResp buy(1: BuyReq req)
}