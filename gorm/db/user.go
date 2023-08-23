package db

import (
	"context"
	"errors"
	"log"
	"main/kitex_gen/api"

	"gorm.io/gorm"
)

func GetProductByName(ctx context.Context, name string) (items []*api.Item, err error) {
	var retItems []Product
	ret := db.Where("product_name = ?", name).Find(&retItems)
	if ret.Error != nil {
		err = ret.Error
	}

	items = make([]*api.Item, len(retItems))
	for i, j := range retItems {
		items[i] = &api.Item{}
		items[i].ID = int64(j.ID)
		items[i].Name = j.Name
		items[i].Price = int64(j.Price)
		items[i].Name = j.Name
	}

	return
}

func GetProductByPrice(ctx context.Context, p1 int64, p2 int64) (items []*api.Item, err error) {
	var retItems []Product
	ret := db.Where("product_price between ? and ?", p1, p2).Find(&retItems)
	if ret.Error != nil {
		err = ret.Error
	}

	items = make([]*api.Item, len(retItems))
	for i, j := range retItems {
		items[i] = &api.Item{}
		items[i].ID = int64(j.ID)
		items[i].Name = j.Name
		items[i].Price = int64(j.Price)
		items[i].Name = j.Name
	}

	return
}

func GetAllProduct(ctx context.Context) (items []*api.Item, err error) {
	var retItems []Product
	ret := db.Find(&retItems)
	if ret.Error != nil {
		err = ret.Error
	}

	items = make([]*api.Item, len(retItems))
	for i, j := range retItems {
		items[i] = &api.Item{}
		items[i].ID = int64(j.ID)
		items[i].Name = j.Name
		items[i].Price = int64(j.Price)
		items[i].Name = j.Name
	}

	return
}

func CreateNewProtuct(ctx context.Context, p *Product) (err error) {
	ret := db.Create(&p)
	if ret.Error != nil {
		return ret.Error
	}
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	if !tx.Statement.Changed("product_number") {
		return errors.New("not changed. there are no such product or it already been seld out")
	}
	return
}

func BuyProduct(ctx context.Context, ID int64) (err error) {
	log.Println(ID)
	db.Model(&Product{}).Where("ID = ? and product_number > ?", ID, 0).Update("product_number", gorm.Expr("product_number - ?", 1))
	return
}
