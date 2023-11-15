package main

import (
	"log"
	"os"

	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/handlers"
	"gitlab.sudovi.me/erp/accounting-api/middleware"
	"gitlab.sudovi.me/erp/accounting-api/services"

	"github.com/oykos-development-hub/celeritas"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "gitlab.sudovi.me/erp/accounting-api"

	models := data.New(cel.DB.Pool)

	OrderListService := services.NewOrderListServiceImpl(cel, models.OrderList)
	OrderListHandler := handlers.NewOrderListHandler(cel, OrderListService)

	OrderProcurementArticleService := services.NewOrderProcurementArticleServiceImpl(cel, models.OrderProcurementArticle)
	OrderProcurementArticleHandler := handlers.NewOrderProcurementArticleHandler(cel, OrderProcurementArticleService)

		
	MovementService := services.NewMovementServiceImpl(cel, models.Movement)
	MovementHandler := handlers.NewMovementHandler(cel, MovementService)

		
	StockService := services.NewStockServiceImpl(cel, models.Stock)
	StockHandler := handlers.NewStockHandler(cel, StockService)

		
	MovementArticleService := services.NewMovementArticleServiceImpl(cel, models.MovementArticle)
	MovementArticleHandler := handlers.NewMovementArticleHandler(cel, MovementArticleService)

	myHandlers := &handlers.Handlers{
		OrderListHandler:               OrderListHandler,
		OrderProcurementArticleHandler: OrderProcurementArticleHandler,
		MovementHandler: MovementHandler,
		StockHandler: StockHandler,
		MovementArticleHandler: MovementArticleHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
