package main

import (
	"gitlab.sudovi.me/erp/accounting-api/handlers"
	"gitlab.sudovi.me/erp/accounting-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	//api
	app.Routes.Route("/api", func(rt chi.Router) {

		// order lists
		rt.Post("/order-lists", handlers.OrderListHandler.CreateOrderList)
		rt.Get("/order-lists/{id}", handlers.OrderListHandler.GetOrderListById)
		rt.Get("/order-lists", handlers.OrderListHandler.GetOrderLists)
		rt.Put("/order-lists/{id}", handlers.OrderListHandler.UpdateOrderList)
		rt.Delete("/order-lists/{id}", handlers.OrderListHandler.DeleteOrderList)

		// order procurement article
		rt.Post("/order-procurement-articles", handlers.OrderProcurementArticleHandler.CreateOrderProcurementArticle)
		rt.Get("/order-procurement-articles/{id}", handlers.OrderProcurementArticleHandler.GetOrderProcurementArticleById)
		rt.Get("/order-procurement-articles", handlers.OrderProcurementArticleHandler.GetOrderProcurementArticles)
		rt.Put("/order-procurement-articles/{id}", handlers.OrderProcurementArticleHandler.UpdateOrderProcurementArticle)
		rt.Delete("/order-procurement-articles/{id}", handlers.OrderProcurementArticleHandler.DeleteOrderProcurementArticle)
	
		rt.Post("/movements", handlers.MovementHandler.CreateMovement)
rt.Get("/movements/{id}", handlers.MovementHandler.GetMovementById)
rt.Get("/movements", handlers.MovementHandler.GetMovementList)
rt.Put("/movements/{id}", handlers.MovementHandler.UpdateMovement)
rt.Delete("/movements/{id}", handlers.MovementHandler.DeleteMovement)
	
		rt.Post("/stocks", handlers.StockHandler.CreateStock)
rt.Get("/stocks/{id}", handlers.StockHandler.GetStockById)
rt.Get("/stocks", handlers.StockHandler.GetStockList)
rt.Put("/stocks/{id}", handlers.StockHandler.UpdateStock)
rt.Delete("/stocks/{id}", handlers.StockHandler.DeleteStock)
	
		rt.Post("/movement-articles", handlers.MovementArticleHandler.CreateMovementArticle)
rt.Get("/movement-articles/{id}", handlers.MovementArticleHandler.GetMovementArticleById)
rt.Get("/movement-articles", handlers.MovementArticleHandler.GetMovementArticleList)
rt.Put("/movement-articles/{id}", handlers.MovementArticleHandler.UpdateMovementArticle)
rt.Delete("/movement-articles/{id}", handlers.MovementArticleHandler.DeleteMovementArticle)
	})

	return app.Routes
}
