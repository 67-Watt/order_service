package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"order_service/endpoint"
	"order_service/utils"
)

// NewHTTPHandler creates and configures an HTTP handler for the Order Service
func NewHTTPHandler(endpoints endpoint.Endpoints, router *mux.Router, logger log.Logger) *mux.Router {
	router.Methods("POST").Path("/orders").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateOrderHandler(w, r, endpoints, logger)
	})
	return router
}

// httpError is a helper function to handle errors consistently in HTTP responses.
func httpError(w http.ResponseWriter, logger log.Logger, statusCode int, message string, err error) {
	w.WriteHeader(statusCode)
	errorMessage := err.Error()

	// Log the error
	err = logger.Log("error", message, "details", errorMessage)
	if err != nil {
		return
	}

	// Respond with the standardized error response
	err = json.NewEncoder(w).Encode(utils.ErrorResponse(message, errorMessage))
	if err != nil {
		return
	}
}

// CreateOrderHandler handles the creation of a new order.
//
// @Summary Create a new order
// @Description Create a new order in the system
// @Tags Order
// @Accept json
// @Produce json
// @Param order body endpoint.CreateOrderRequest true "Order Request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /orders [post]
func CreateOrderHandler(w http.ResponseWriter, r *http.Request, endpoints endpoint.Endpoints, logger log.Logger) {
	var req endpoint.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, logger, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	response, err := endpoints.CreateOrderEndpoint(context.Background(), req)
	if err != nil {
		httpError(w, logger, http.StatusInternalServerError, "Failed to create order", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(utils.SuccessResponse("Order created successfully", response))
	if err != nil {
		return
	}
}
