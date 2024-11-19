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

	ErrorLogService := services.NewErrorLogServiceImpl(cel, models.ErrorLog)
	ErrorLogHandler := handlers.NewErrorLogHandler(cel, ErrorLogService)

	OrderListService := services.NewOrderListServiceImpl(cel, models.OrderList)
	OrderListHandler := handlers.NewOrderListHandler(cel, OrderListService, ErrorLogService)

	OrderProcurementArticleService := services.NewOrderProcurementArticleServiceImpl(cel, models.OrderProcurementArticle)
	OrderProcurementArticleHandler := handlers.NewOrderProcurementArticleHandler(cel, OrderProcurementArticleService, ErrorLogService)

	MovementService := services.NewMovementServiceImpl(cel, models.Movement)
	MovementHandler := handlers.NewMovementHandler(cel, MovementService, ErrorLogService)

	StockService := services.NewStockServiceImpl(cel, models.Stock)
	StockHandler := handlers.NewStockHandler(cel, StockService, ErrorLogService)

	MovementArticleService := services.NewMovementArticleServiceImpl(cel, models.MovementArticle)
	MovementArticleHandler := handlers.NewMovementArticleHandler(cel, MovementArticleService, ErrorLogService)

	LogService := services.NewLogServiceImpl(cel, models.Log)
	LogHandler := handlers.NewLogHandler(cel, LogService, ErrorLogService)

	StockOrderArticleService := services.NewStockOrderArticleServiceImpl(cel, models.StockOrderArticle)
	StockOrderArticleHandler := handlers.NewStockOrderArticleHandler(cel, StockOrderArticleService)

	myHandlers := &handlers.Handlers{
		OrderListHandler:               OrderListHandler,
		OrderProcurementArticleHandler: OrderProcurementArticleHandler,
		MovementHandler:                MovementHandler,
		StockHandler:                   StockHandler,
		MovementArticleHandler:         MovementArticleHandler,
		LogHandler:                     LogHandler,
		ErrorLogHandler:                ErrorLogHandler,
		StockOrderArticleHandler:       StockOrderArticleHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
