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
