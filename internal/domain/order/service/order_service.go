package service

import (
	"context"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/evermos/boilerplate-go/internal/domain/order/model/dto"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/rs/zerolog/log"
)

type OrderExtensionService interface {
	DownloadOrdersByMarket(downloadCtx context.Context, downloadFilter dto.DownloadOrdersByMarketFilterRequest) (dto.OrdersDownloadResponseList, error)
	GetBrandsByMarket(brandCtx context.Context, brandFilter dto.GetBrandsByMarketFilterRequest) (dto.OrdersBrandResponseList, error)
	GetStoresByMarket(storeCtx context.Context, storeFilter dto.GetStoresByMarketFilterRequest) (dto.OrdersStoreResponseList, error)
	PreviewOrdersByMarket(previewCtx context.Context, previewFilter dto.PreviewOrdersByMarketFilterRequest) (dto.OrdersPreviewResponseList, error)
}

func (s *OrderServiceImpl) DownloadOrdersByMarket(downloadCtx context.Context, downloadFilter dto.DownloadOrdersByMarketFilterRequest) (dto.OrdersDownloadResponseList, error) {
	var downloadResults model.OrdersDownloadList
	var err error

	downloadQueryFilter := dto.ConvertDownloadOrdersByMarketFilterRequests(downloadFilter)
	switch downloadFilter.Marketplace {
	case shared.SHOPEE:
		downloadResults, err = s.OrderRepository.DownloadShopeeOrders(downloadCtx, downloadQueryFilter)
	case shared.BLIBLI:
		downloadResults, err = s.OrderRepository.DownloadBlibliOrders(downloadCtx, downloadQueryFilter)
	default:
		return dto.OrdersDownloadResponseList{}, nil
	}

	if err != nil {
		if failure.GetCode(err) == http.StatusInternalServerError {
			log.Error().
				Err(err).
				Msg("[DownloadOrdersByMarket] Internal server error occurred")
			err = failure.InternalError(shared.InternalErrorSystem)
			return dto.OrdersDownloadResponseList{}, err
		}
		log.Warn().
			Msg("[DownloadOrdersByMarket] Error occurred")
		return dto.OrdersDownloadResponseList{}, err
	}

	return dto.ConvertDownloadResponses(downloadResults), nil
}

func (s *OrderServiceImpl) GetBrandsByMarket(brandCtx context.Context, brandFilter dto.GetBrandsByMarketFilterRequest) (dto.OrdersBrandResponseList, error) {
	var brandResults model.OrdersBrandList
	var err error

	switch brandFilter.Marketplace {
	case shared.SHOPEE:
		brandResults, err = s.OrderRepository.GetShopeeBrands(brandCtx)
	case shared.BLIBLI:
		brandResults, err = s.OrderRepository.GetBlibliBrands(brandCtx)
	default:
		return dto.OrdersBrandResponseList{}, nil
	}

	if err != nil {
		if failure.GetCode(err) == http.StatusInternalServerError {
			log.Error().
				Err(err).
				Msg("[GetBrandsByMarket] Internal server error occurred")
			err = failure.InternalError(shared.InternalErrorSystem)
			return dto.OrdersBrandResponseList{}, err
		}
		log.Warn().
			Msg("[GetBrandsByMarket] Error occurred")
		return dto.OrdersBrandResponseList{}, err
	}

	return dto.ConvertBrandResponses(brandResults), nil
}

func (s *OrderServiceImpl) GetStoresByMarket(storeCtx context.Context, storeFilter dto.GetStoresByMarketFilterRequest) (dto.OrdersStoreResponseList, error) {
	var storeResults model.OrdersStoreList
	var err error

	switch storeFilter.Marketplace {
	case shared.BLIBLI:
		storeResults, err = s.OrderRepository.GetBlibliStores(storeCtx)
	default:
		return dto.OrdersStoreResponseList{}, nil
	}

	if err != nil {
		if failure.GetCode(err) == http.StatusInternalServerError {
			log.Error().
				Err(err).
				Msg("[GetStoresByMarket] Internal server error occurred")
			err = failure.InternalError(shared.InternalErrorSystem)
			return dto.OrdersStoreResponseList{}, err
		}
		log.Warn().
			Msg("[GetStoresByMarket] Error occurred")
		return dto.OrdersStoreResponseList{}, err
	}
	return dto.ConvertStoreResponses(storeResults), nil
}

func (s *OrderServiceImpl) PreviewOrdersByMarket(previewCtx context.Context, previewFilter dto.PreviewOrdersByMarketFilterRequest) (dto.OrdersPreviewResponseList, error) {
	var previewResults model.OrdersPreviewList
	var err error

	previewQueryFilter := dto.ConvertPreviewOrdersByMarketFilterRequests(previewFilter)
	switch previewFilter.Marketplace {
	case shared.BLIBLI:
		previewResults, err = s.OrderRepository.PreviewBlibliOrders(previewCtx, previewQueryFilter)
	default:
		return dto.OrdersPreviewResponseList{}, nil
	}

	if err != nil {
		if failure.GetCode(err) == http.StatusInternalServerError {
			log.Error().
				Err(err).
				Msg("[PreviewOrdersByMarket] Internal server error occurred")
			err = failure.InternalError(shared.InternalErrorSystem)
			return dto.OrdersPreviewResponseList{}, err
		}
		log.Warn().
			Msg("[PreviewOrdersByMarket] Error occurred")
		return dto.OrdersPreviewResponseList{}, err
	}

	return dto.ConvertPreviewResponses(previewResults), nil
}
