package utils

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"order_service/models"
	"order_service/usecase"
)

// LoggingMiddleware adds logging to each method in the OrderUseCase interface.
func LoggingMiddleware(logger log.Logger) func(usecase.OrderUseCase) usecase.OrderUseCase {
	return func(next usecase.OrderUseCase) usecase.OrderUseCase {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   usecase.OrderUseCase
	logger log.Logger
}

func (mw *loggingMiddleware) CreateOrder(ctx context.Context, order models.Order) (uint, error) {
	err := level.Info(mw.logger).Log("method", "CreateOrder", "order", order)
	if err != nil {
		return 0, err
	}
	orderID, err := mw.next.CreateOrder(ctx, order)
	if err != nil {
		err := level.Error(mw.logger).Log("method", "CreateOrder", "error", err)
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	err = level.Info(mw.logger).Log("method", "CreateOrder", "order_id", orderID)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (mw *loggingMiddleware) UpdateOrderStatus(ctx context.Context, orderID uint, status string) error {
	err := level.Info(mw.logger).Log("method", "UpdateOrderStatus", "order_id", orderID, "status", status)
	if err != nil {
		return err
	}
	err = mw.next.UpdateOrderStatus(ctx, orderID, status)
	if err != nil {
		err := level.Error(mw.logger).Log("method", "UpdateOrderStatus", "order_id", orderID, "error", err)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (mw *loggingMiddleware) GetOrderDetails(ctx context.Context, orderID uint) (models.Order, error) {
	err := level.Info(mw.logger).Log("method", "GetOrderDetails", "order_id", orderID)
	if err != nil {
		return models.Order{}, err
	}
	order, err := mw.next.GetOrderDetails(ctx, orderID)
	if err != nil {
		err := level.Error(mw.logger).Log("method", "GetOrderDetails", "order_id", orderID, "error", err)
		if err != nil {
			return models.Order{}, err
		}
		return models.Order{}, err
	}
	err = level.Info(mw.logger).Log("method", "GetOrderDetails", "order", order)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

// MetricsMiddleware adds metrics tracking to each method in the OrderUseCase interface.
func MetricsMiddleware(counter *prometheus.CounterVec) func(usecase.OrderUseCase) usecase.OrderUseCase {
	return func(next usecase.OrderUseCase) usecase.OrderUseCase {
		return &metricsMiddleware{
			next:    next,
			counter: counter,
		}
	}
}

type metricsMiddleware struct {
	next    usecase.OrderUseCase
	counter *prometheus.CounterVec
}

func (mw *metricsMiddleware) CreateOrder(ctx context.Context, order models.Order) (uint, error) {
	mw.counter.WithLabelValues("CreateOrder").Inc()
	return mw.next.CreateOrder(ctx, order)
}

func (mw *metricsMiddleware) UpdateOrderStatus(ctx context.Context, orderID uint, status string) error {
	mw.counter.WithLabelValues("UpdateOrderStatus").Inc()
	return mw.next.UpdateOrderStatus(ctx, orderID, status)
}

func (mw *metricsMiddleware) GetOrderDetails(ctx context.Context, orderID uint) (models.Order, error) {
	mw.counter.WithLabelValues("GetOrderDetails").Inc()
	return mw.next.GetOrderDetails(ctx, orderID)
}
