package models

import "time"

type Goods struct {
	Id              string    `db:"id"`
	Name            string    `db:"name"`
	Goods_brief     string    `db:"goods_brief"`
	Goods_desc      string    `db:"goods_desc"`
	Sort_order      int       `db:"sort_order"`
	Goods_number    int       `db:"goods_number"`
	Sell_volume     int       `db:"sell_volume"`
	Primary_pic_url string    `db:"primary_pic_url"`
	List_pic_url    string    `db:"list_pic_url"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
