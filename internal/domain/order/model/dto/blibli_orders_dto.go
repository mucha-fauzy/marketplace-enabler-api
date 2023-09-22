package dto

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/guregu/null"
)

type BlibliDownloadResponse struct {
	OrderNo       null.String `json:"orderNo"`
	OrderStatus   null.String `json:"orderStatus"`
	AWBNo         null.String `json:"awbNo"`
	NameProduct   null.String `json:"nameProduct"`
	SKUCode       null.String `json:"skuCode"`
	ProductPrice  null.Int    `json:"productPrice"`
	TotalQuantity null.Int    `json:"totalQuantity"`
}

// Map dto Response of Blibli Download based on model Blibli Download
func mapBlibliDownloadResponses(blibliDownload model.BlibliDownload) BlibliDownloadResponse {
	return BlibliDownloadResponse{
		OrderNo:       blibliDownload.OrderNo,
		OrderStatus:   blibliDownload.OrderStatus,
		AWBNo:         blibliDownload.AWBNo,
		NameProduct:   blibliDownload.NameProduct,
		SKUCode:       blibliDownload.SKUCode,
		ProductPrice:  blibliDownload.ProductPrice,
		TotalQuantity: blibliDownload.TotalQuantity,
	}
}

// Response List of Blibli Orders Download
type BlibliDownloadResponseList []*BlibliDownloadResponse

// convert model List of Orders Download into dto Response List of Orders for specific marketplace (in this case Blibli)
func convertBlibliDownloadResponses(ordersDownloadList model.OrdersDownloadList) BlibliDownloadResponseList {
	var blibliDownloadResponseList BlibliDownloadResponseList = BlibliDownloadResponseList{}

	if len(ordersDownloadList.Blibli) > 0 {
		for _, blibliDownload := range ordersDownloadList.Blibli {
			blibliDownloadResponse := mapBlibliDownloadResponses(*blibliDownload)
			blibliDownloadResponseList = append(blibliDownloadResponseList, &blibliDownloadResponse)
		}
	}

	return blibliDownloadResponseList
}

type BlibliBrandResponse struct {
	Brand null.String `json:"brand"`
}

// Blibli dto mapper
func mapBlibliBrandResponses(blibliBrand model.BlibliBrand) BlibliBrandResponse {
	return BlibliBrandResponse{
		Brand: blibliBrand.Brand,
	}
}

type BlibliBrandResponseList []*BlibliBrandResponse

// Blibli dto converter
func convertBlibliBrandResponses(ordersBrandList model.OrdersBrandList) BlibliBrandResponseList {
	var blibliBrandResponseList BlibliBrandResponseList = BlibliBrandResponseList{}

	if len(ordersBrandList.Blibli) > 0 {
		for _, blibliBrand := range ordersBrandList.Blibli {
			blibliBrandResponse := mapBlibliBrandResponses(*blibliBrand)
			blibliBrandResponseList = append(blibliBrandResponseList, &blibliBrandResponse)
		}
	}

	return blibliBrandResponseList
}

type BlibliStoreResponse struct {
	Store null.String `json:"store"`
}

func mapBlibliStoreResponses(blibliStore model.BlibliStore) BlibliStoreResponse {
	return BlibliStoreResponse{
		Store: blibliStore.Store,
	}
}

type BlibliStoreResponseList []*BlibliStoreResponse

func convertBlibliStoreResponses(ordersStoreList model.OrdersStoreList) BlibliStoreResponseList {
	var blibliStoreResponseList BlibliStoreResponseList = BlibliStoreResponseList{}

	if len(ordersStoreList.Blibli) > 0 {
		for _, blibliStore := range ordersStoreList.Blibli {
			blibliStoreResponse := mapBlibliStoreResponses(*blibliStore)
			blibliStoreResponseList = append(blibliStoreResponseList, &blibliStoreResponse)
		}
	}
	return blibliStoreResponseList
}

type BlibliPreviewResponse struct {
	ID              int         `json:"id"`
	Brand           null.String `json:"brand"`
	Store           null.String `json:"store"`
	OrderNo         null.String `json:"orderNo"`
	OrderStatus     null.String `json:"orderStatus"`
	SKUCode         null.String `json:"skuCode"`
	TotalQuantity   null.Int    `json:"totalQuantity"`
	ProductPrice    null.Int    `json:"productPrice"`
	OrderDate       null.Time   `json:"orderDate"`
	NameProduct     null.String `json:"nameProduct"`
	OrderItemNo     null.String `json:"orderItemNo"`
	BuyerName       null.String `json:"buyerName"`
	PacketNo        null.String `json:"packetNo"`
	AWBNo           null.String `json:"awbNo"`
	PickupPointCode null.String `json:"pickupPointCode"`
	CreatedAt       null.Time   `json:"createdAt"`
	UpdatedAt       null.Time   `json:"updatedAt"`
	CreatedBy       null.String `json:"createdBy"`
	UpdatedBy       null.String `json:"updatedBy"`
}

func mapBlibliPreviewResponses(blibliPreview model.BlibliPreview) BlibliPreviewResponse {
	return BlibliPreviewResponse{
		ID:              blibliPreview.ID,
		Brand:           blibliPreview.Brand,
		Store:           blibliPreview.Store,
		OrderNo:         blibliPreview.OrderNo,
		OrderStatus:     blibliPreview.OrderStatus,
		SKUCode:         blibliPreview.SKUCode,
		TotalQuantity:   blibliPreview.TotalQuantity,
		ProductPrice:    blibliPreview.ProductPrice,
		OrderDate:       blibliPreview.OrderDate,
		NameProduct:     blibliPreview.NameProduct,
		OrderItemNo:     blibliPreview.OrderItemNo,
		BuyerName:       blibliPreview.BuyerName,
		PacketNo:        blibliPreview.PacketNo,
		AWBNo:           blibliPreview.AWBNo,
		PickupPointCode: blibliPreview.PickupPointCode,
		CreatedAt:       blibliPreview.CreatedAt,
		UpdatedAt:       blibliPreview.UpdatedAt,
		CreatedBy:       blibliPreview.CreatedBy,
		UpdatedBy:       blibliPreview.UpdatedBy,
	}
}

type BlibliPreviewResponseList []*BlibliPreviewResponse

func convertBlibliPreviewResponses(ordersPreviewList model.OrdersPreviewList) BlibliPreviewResponseList {
	var blibliPreviewResponseList BlibliPreviewResponseList = BlibliPreviewResponseList{}

	if len(ordersPreviewList.Blibli) > 0 {
		for _, blibliPreview := range ordersPreviewList.Blibli {
			blibliPreviewResponse := mapBlibliPreviewResponses(*blibliPreview)
			blibliPreviewResponseList = append(blibliPreviewResponseList, &blibliPreviewResponse)
		}
	}
	return blibliPreviewResponseList
}
