package data

import (
	"fmt"

	up "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"

	"database/sql"
	"os"
)

//nolint:all
var db *sql.DB

//nolint:all
var Upper up.Session

type Models struct {
	OrderList               OrderList
	OrderProcurementArticle OrderProcurementArticle
	Movement                Movement
	Stock                   Stock
	MovementArticle         MovementArticle
	Log                     Log
	ErrorLog ErrorLog
		StockOrderArticle StockOrderArticle
	}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		Upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		Upper, _ = postgresql.New(databasePool)
	default:
		// do nothing
	}

	return Models{
		OrderList:               OrderList{},
		OrderProcurementArticle: OrderProcurementArticle{},
		Movement:                Movement{},
		Stock:                   Stock{},
		MovementArticle:         MovementArticle{},
		Log:                     Log{},
		ErrorLog: ErrorLog{},
		StockOrderArticle: StockOrderArticle{},
	}
}

//nolint:all
func getInsertId(i up.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}

func paginateResult(res up.Result, page int, pageSize int) up.Result {
	// Calculate the offset based on the page number and page size
	offset := (page - 1) * pageSize

	// Apply pagination to the query
	res = res.Offset(offset).Limit(pageSize)

	return res
}
