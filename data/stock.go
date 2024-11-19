package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/accounting-api/pkg/errors"
)

// Stock struct
type Stock struct {
	ID                 int       `db:"id,omitempty"`
	Year               string    `db:"year"`
	Title              string    `db:"title"`
	NetPrice           float32   `db:"net_price"`
	VatPercentage      int       `db:"vat_percentage"`
	Description        string    `db:"description"`
	OrganizationUnitID int       `db:"organization_unit_id"`
	Amount             int       `db:"amount"`
	Exception          bool      `db:"exception"`
	CreatedAt          time.Time `db:"created_at,omitempty"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Stock) Table() string {
	return "stocks"
}

// GetAll gets all records from the database, using Upper
func (t *Stock) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*Stock, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Stock
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order")
	}

	return all, &total, err
}

func (t *Stock) GetAllForReport(date time.Time, organizationUnitID *int) ([]*Stock, error) {

	query := `SELECT 
			    s.id AS article_id,
			    s.title,
			    s.description,
			    s.year,
			    s.amount 
			        - COALESCE(SUM(ma.amount), 0) 
			        + COALESCE(SUM(opa.amount), 0) AS stock_amount
				FROM 
				    stocks s
				LEFT JOIN 
				    movement_articles ma ON s.id = ma.stock_id AND ma.created_at > $1
				LEFT JOIN 
				    order_procurement_articles opa ON s.id = opa.stock_id 
				LEFT JOIN 
					order_lists ol ON opa.order_list_id = ol.id AND ol.receipt_date <= $1
				WHERE 
				    AND ($2 = 0 OR s.organization_unit_id = $2)
				GROUP BY 
				    s.id, s.title, s.description, s.year;`

	if organizationUnitID == nil {
		zero := 0
		organizationUnitID = &zero
	}

	date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

	rows, err := Upper.SQL().Query(query, date, *organizationUnitID)

	if err != nil {
		return nil, newErrors.Wrap(err, "upper query")
	}

	defer rows.Close()

	var articles []*Stock

	for rows.Next() {
		var article Stock
		err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Year, &article.Amount)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper scan")
		}

		articles = append(articles, &article)
	}

	return articles, nil

}

// Get gets one record from the database, by id, using Upper
func (t *Stock) Get(id int) (*Stock, error) {
	var one Stock
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Stock) Update(m Stock) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Stock) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Stock) Insert(m Stock) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}
