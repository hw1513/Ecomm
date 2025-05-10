package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hw1513/ecomm/common/config"
	rest "github.com/hw1513/ecomm/common/genrest"
	"github.com/hw1513/ecomm/common/server"
)

var cfg *config.Config
var orders = make(map[string]rest.Order)

func init() {
	var err error
	if cfg, err = config.LoadConfig("../../internal/common/config/config.dev.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {

	server.RunRESTServer(&cfg.Order, func(router *gin.Engine) {
		rest.RegisterHandlersWithOptions(router, HTTPServer{}, rest.GinServerOptions{
			BaseURL: "",
		})
	})
}

type HTTPServer struct{}

func (s HTTPServer) ListOrders(ctx *gin.Context, params rest.ListOrdersParams) {
	var userOrders []rest.Order
	for _, order := range orders {
		if order.UserId == params.UserId {
			userOrders = append(userOrders, order)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": userOrders,
	})
}

func (s HTTPServer) CreateOrder(ctx *gin.Context) {
	var req rest.CreateOrderJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 计算订单总金额
	var totalAmount float64
	for _, item := range req.Items {
		totalAmount += item.Subtotal
	}

	now := time.Now()
	order := rest.Order{
		Id:          uuid.New().String(),
		UserId:      req.UserId,
		Items:       req.Items,
		TotalAmount: totalAmount,
		Status:      rest.ORDERSTATUSPENDING,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	orders[order.Id] = order
	ctx.JSON(http.StatusOK, order)
}

func (s HTTPServer) GetOrder(ctx *gin.Context, orderId string) {
	order, exists := orders[orderId]
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (s HTTPServer) UpdateOrderStatus(ctx *gin.Context, orderId string) {
	var req rest.UpdateOrderStatusJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, exists := orders[orderId]
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	now := time.Now()
	order.Status = req.Status
	order.UpdatedAt = &now
	orders[orderId] = order

	ctx.JSON(http.StatusOK, order)
}
