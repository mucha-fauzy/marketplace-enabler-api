package dto

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/model"
)

type OrdersDownloadResponseList struct {
	Blibli BlibliDownloadResponseList `json:"blibli"`
	Shopee ShopeeDownloadResponseList `json:"shopee"`
}

func ConvertDownloadResponses(ordersDownloadList model.OrdersDownloadList) OrdersDownloadResponseList {
	return OrdersDownloadResponseList{
		Blibli: convertBlibliDownloadResponses(ordersDownloadList),
		Shopee: convertShopeeDownloadResponses(ordersDownloadList),
	}
}

type OrdersBrandResponseList struct {
	Blibli BlibliBrandResponseList `json:"blibli"`
	Shopee ShopeeBrandResponseList `json:"shopee"`
}

func ConvertBrandResponses(ordersBrandList model.OrdersBrandList) OrdersBrandResponseList {
	return OrdersBrandResponseList{
		Blibli: convertBlibliBrandResponses(ordersBrandList),
		Shopee: convertShopeeBrandResponses(ordersBrandList),
	}
}

type OrdersStoreResponseList struct {
	Blibli BlibliStoreResponseList `json:"blibli"`
}

func ConvertStoreResponses(ordersStoreList model.OrdersStoreList) OrdersStoreResponseList {
	return OrdersStoreResponseList{
		Blibli: convertBlibliStoreResponses(ordersStoreList),
	}
}

type OrdersPreviewResponseList struct {
	Blibli BlibliPreviewResponseList `json:"blibli"`
}

func ConvertPreviewResponses(ordersPreviewList model.OrdersPreviewList) OrdersPreviewResponseList {
	return OrdersPreviewResponseList{
		Blibli: convertBlibliPreviewResponses(ordersPreviewList),
	}
}
