package model

import (
	"github.com/guregu/null"
)

type BlibliDownload struct {
	OrderNo       null.String `db:"order_no"`
	OrderStatus   null.String `db:"order_status"`
	AWBNo         null.String `db:"awb_no"`
	NameProduct   null.String `db:"name_product"`
	SKUCode       null.String `db:"sku_code"`
	ProductPrice  null.Int    `db:"product_price"`
	TotalQuantity null.Int    `db:"total_quantity"`
}

type BlibliDownloadList []*BlibliDownload

// Map List of Blibli Download into List of Orders Download
func BlibliNewOrdersDownloadList(blibliDownloadList BlibliDownloadList) OrdersDownloadList {
	return OrdersDownloadList{
		Blibli: blibliDownloadList,
	}
}

type BlibliBrand struct {
	Brand null.String `db:"brand"`
}

type BlibliBrandList []*BlibliBrand

func BlibliNewOrdersBrandList(blibliBrandList BlibliBrandList) OrdersBrandList {
	return OrdersBrandList{
		Blibli: blibliBrandList,
	}
}

type BlibliStore struct {
	Store null.String `db:"store"`
}

type BlibliStoreList []*BlibliStore

func BlibliNewOrdersStoreList(blibliStoreList BlibliStoreList) OrdersStoreList {
	return OrdersStoreList{
		Blibli: blibliStoreList,
	}
}

type BlibliPreview struct {
	ID              int         `db:"id"`
	Brand           null.String `db:"brand"`
	Store           null.String `db:"store"`
	OrderNo         null.String `db:"order_no"`
	OrderStatus     null.String `db:"order_status"`
	SKUCode         null.String `db:"sku_code"`
	TotalQuantity   null.Int    `db:"total_quantity"`
	ProductPrice    null.Int    `db:"product_price"`
	OrderDate       null.Time   `db:"order_date"`
	NameProduct     null.String `db:"name_product"`
	OrderItemNo     null.String `db:"order_item_no"`
	BuyerName       null.String `db:"buyer_name"`
	PacketNo        null.String `db:"packet_no"`
	AWBNo           null.String `db:"awb_no"`
	PickupPointCode null.String `db:"pickup_point_code"`
	CreatedAt       null.Time   `db:"created_at"`
	UpdatedAt       null.Time   `db:"updated_at"`
	CreatedBy       null.String `db:"created_by"`
	UpdatedBy       null.String `db:"updated_by"`
}

type BlibliPreviewList []*BlibliPreview

func BlibliNewOrdersPreviewList(blibliPreviewList BlibliPreviewList) OrdersPreviewList {
	return OrdersPreviewList{
		Blibli: blibliPreviewList,
	}
}
