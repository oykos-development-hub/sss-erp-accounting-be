package handlers

import "net/http"

type Handlers struct {
	OrderListHandler               OrderListHandler
	OrderProcurementArticleHandler OrderProcurementArticleHandler
}

type OrderListHandler interface {
	CreateOrderList(w http.ResponseWriter, r *http.Request)
	UpdateOrderList(w http.ResponseWriter, r *http.Request)
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
