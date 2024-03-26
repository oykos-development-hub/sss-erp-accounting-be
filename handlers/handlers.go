package handlers

import "net/http"

type Handlers struct {
	OrderListHandler               OrderListHandler
	OrderProcurementArticleHandler OrderProcurementArticleHandler
	MovementHandler                MovementHandler
	StockHandler                   StockHandler
	MovementArticleHandler         MovementArticleHandler
}

type OrderListHandler interface {
	CreateOrderList(w http.ResponseWriter, r *http.Request)
	UpdateOrderList(w http.ResponseWriter, r *http.Request)
	SendToFinance(w http.ResponseWriter, r *http.Request)
	DeleteOrderList(w http.ResponseWriter, r *http.Request)
	GetOrderListById(w http.ResponseWriter, r *http.Request)
	GetOrderLists(w http.ResponseWriter, r *http.Request)
}

type OrderProcurementArticleHandler interface {
	CreateOrderProcurementArticle(w http.ResponseWriter, r *http.Request)
	UpdateOrderProcurementArticle(w http.ResponseWriter, r *http.Request)
	DeleteOrderProcurementArticle(w http.ResponseWriter, r *http.Request)
	GetOrderProcurementArticleById(w http.ResponseWriter, r *http.Request)
	GetOrderProcurementArticles(w http.ResponseWriter, r *http.Request)
}

type MovementHandler interface {
	CreateMovement(w http.ResponseWriter, r *http.Request)
	UpdateMovement(w http.ResponseWriter, r *http.Request)
	DeleteMovement(w http.ResponseWriter, r *http.Request)
	GetMovementById(w http.ResponseWriter, r *http.Request)
	GetMovementList(w http.ResponseWriter, r *http.Request)
	GetMovementReport(w http.ResponseWriter, r *http.Request)
}

type StockHandler interface {
	CreateStock(w http.ResponseWriter, r *http.Request)
	UpdateStock(w http.ResponseWriter, r *http.Request)
	DeleteStock(w http.ResponseWriter, r *http.Request)
	GetStockById(w http.ResponseWriter, r *http.Request)
	GetStockList(w http.ResponseWriter, r *http.Request)
}

type MovementArticleHandler interface {
	CreateMovementArticle(w http.ResponseWriter, r *http.Request)
	UpdateMovementArticle(w http.ResponseWriter, r *http.Request)
	DeleteMovementArticle(w http.ResponseWriter, r *http.Request)
	GetMovementArticleById(w http.ResponseWriter, r *http.Request)
	GetMovementArticleList(w http.ResponseWriter, r *http.Request)
}
