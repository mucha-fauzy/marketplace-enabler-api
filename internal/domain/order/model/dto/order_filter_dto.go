package dto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
)

type DownloadOrdersByMarketFilterRequest struct {
	Marketplace string
	Brand       string
	Store       string
	DateFromStr string
	DateToStr   string
	DateFrom    time.Time
	DateTo      time.Time
}

func (f *DownloadOrdersByMarketFilterRequest) Validate() error {
	var err error
	validator := shared.GetValidator()

	//Parse and validate dates in DownloadOrdersByMarketFilterRequest struct, "" will be defaulted into 0001-01-01 00:00:00 +0000 UTC
	if f.DateFromStr != "" && f.DateToStr != "" {
		f.DateFrom, err = time.Parse(shared.DefaultDateFormat, f.DateFromStr)
		if err != nil {
			return failure.BadRequestFromString(fmt.Sprintf("Invalid dateFrom parameter %s", err))
		}

		f.DateTo, err = time.Parse(shared.DefaultDateFormat, f.DateToStr)
		if err != nil {
			return failure.BadRequestFromString(fmt.Sprintf("Invalid dateTo parameter %s", err))
		}

		if f.DateTo.Before(f.DateFrom) {
			return failure.BadRequestFromString("DateTo must be after DateFrom")
		}
	}
	return validator.Struct(f)
}

// Map field of DownloadOrdersByMarketFilterRequest based on the input (for init filter request variable)
func NewDownloadOrdersByMarketFilterRequests(marketplace, brand, store, dateFromStr, dateToStr string) DownloadOrdersByMarketFilterRequest {
	return DownloadOrdersByMarketFilterRequest{
		Marketplace: marketplace,
		Brand:       brand,
		Store:       store,
		DateFromStr: dateFromStr,
		DateToStr:   dateToStr,
	}
}

// Convert dto of Filter Request into model of Filter to interact with repo/db
func ConvertDownloadOrdersByMarketFilterRequests(downloadOrdersByMarketFilterRequest DownloadOrdersByMarketFilterRequest) model.DownloadOrdersByMarketFilter {
	return model.DownloadOrdersByMarketFilter{
		Brand:    downloadOrdersByMarketFilterRequest.Brand,
		Store:    downloadOrdersByMarketFilterRequest.Store,
		DateFrom: downloadOrdersByMarketFilterRequest.DateFrom,
		DateTo:   downloadOrdersByMarketFilterRequest.DateTo,
	}
}

type GetBrandsByMarketFilterRequest struct {
	Marketplace string
}

func (f *GetBrandsByMarketFilterRequest) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(f)
}

func NewGetBrandsByMarketFilterRequests(marketplace string) GetBrandsByMarketFilterRequest {
	return GetBrandsByMarketFilterRequest{
		Marketplace: marketplace,
	}
}

type GetStoresByMarketFilterRequest struct {
	Marketplace string
}

func (f *GetStoresByMarketFilterRequest) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(f)
}

func NewGetStoresByMarketFilterRequests(marketplace string) GetStoresByMarketFilterRequest {
	return GetStoresByMarketFilterRequest{
		Marketplace: marketplace,
	}
}

type PreviewOrdersByMarketFilterRequest struct {
	Marketplace string
	PageStr     string
	SizeStr     string
	Page        int
	Size        int
}

func (f *PreviewOrdersByMarketFilterRequest) Validate() error {
	var err error
	validator := shared.GetValidator()

	if f.PageStr != "" && f.SizeStr != "" {
		f.Page, err = strconv.Atoi(f.PageStr)
		if err != nil {
			return failure.BadRequestFromString(fmt.Sprintf("Invalid page parameter %s", err))
		}

		if f.Page < shared.PageLowerLimit {
			f.Page = shared.DefaultPage
		}

		f.Size, err = strconv.Atoi(f.SizeStr)
		if err != nil {
			return failure.BadRequestFromString(fmt.Sprintf("Invalid size parameter %s", err))
		}

		if f.Size < shared.SizeLowerLimit || f.Size > shared.SizeUpperLimit {
			f.Size = shared.DefaultSize
		}

	} else {
		f.Page = shared.DefaultPage
		f.Size = shared.DefaultSize
	}

	return validator.Struct(f)
}

func NewPreviewOrdersByMarketFilterRequests(marketplace, pageStr, sizeStr string) PreviewOrdersByMarketFilterRequest {
	return PreviewOrdersByMarketFilterRequest{
		Marketplace: marketplace,
		PageStr:     pageStr,
		SizeStr:     sizeStr,
	}
}

func ConvertPreviewOrdersByMarketFilterRequests(previewOrdersByMarketFilterRequest PreviewOrdersByMarketFilterRequest) model.PreviewOrdersByMarketFilter {
	return model.PreviewOrdersByMarketFilter{
		Page: previewOrdersByMarketFilterRequest.Page,
		Size: previewOrdersByMarketFilterRequest.Size,
	}
}
