package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"order_service/models"
	"order_service/usecase"
)

type CreateOrderRequest struct {
	Order models.Order `json:"order"`
}

type CreateOrderResponse struct {
	OrderID uint `json:"order_id"`
}

func makeCreateOrderEndpoint(svc usecase.OrderUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrderRequest)
		orderID, err := svc.CreateOrder(ctx, req.Order)
		if err != nil {
			return nil, err
		}
		return CreateOrderResponse{OrderID: orderID}, nil
	}
}

type Endpoints struct {
	CreateOrderEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc usecase.OrderUseCase) Endpoints {
	return Endpoints{
		CreateOrderEndpoint: makeCreateOrderEndpoint(svc),
	}
}
