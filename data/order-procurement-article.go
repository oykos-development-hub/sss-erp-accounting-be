package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Order procurement article struct
type OrderProcurementArticle struct {
	ID            int       `db:"id,omitempty"`
	OrderID       int       `db:"order_id"`
	ArticleID     *int      `db:"article_id"`
	Title         *string   `db:"title"`
	Description   *string   `db:"description"`
	Year          string    `db:"year"`
	NetPrice      float32   `db:"net_price"`
	VatPercentage int       `db:"vat_percentage"`
	Amount        int       `db:"amount"`
	CreatedAt     time.Time `db:"created_at,omitempty"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *OrderProcurementArticle) Table() string {
	return "order_procurement_articles"
}

// GetAll gets all records from the database, using Upper
func (t *OrderProcurementArticle) GetAll(page *int, size *int, conditions *up.AndExpr) ([]*OrderProcurementArticle, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*OrderProcurementArticle
	var res up.Result

	if conditions != nil {
		res = collection.Find(conditions)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, err
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *OrderProcurementArticle) Get(id int) (*OrderProcurementArticle, error) {
	var one OrderProcurementArticle
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *OrderProcurementArticle) Update(m OrderProcurementArticle) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *OrderProcurementArticle) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *OrderProcurementArticle) Insert(m OrderProcurementArticle) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
