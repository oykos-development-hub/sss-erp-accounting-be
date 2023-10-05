package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Order list struct
type OrderList struct {
	ID                  int        `db:"id,omitempty"`
	DateOrder           time.Time  `db:"date_order"`
	TotalPrice          float32    `db:"total_price"`
	PublicProcurementID int        `db:"public_procurement_id"`
	SupplierID          *int       `db:"supplier_id"`
	Status              string     `db:"status"`
	DateSystem          *time.Time `db:"date_system"`
	InvoiceDate         *time.Time `db:"invoice_date"`
	InvoiceNumber       *int       `db:"invoice_number"`
	OrganizationUnitID  int        `db:"organization_unit_id"`
	OfficeID            *int       `db:"office_id"`
	RecipientUserID     int        `db:"recipient_user_id"`
	Description         *string    `db:"description"`
	CreatedAt           time.Time  `db:"created_at,omitempty"`
	UpdatedAt           time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *OrderList) Table() string {
	return "order_lists"
}

// GetAll gets all records from the database, using upper
func (t *OrderList) GetAll(page *int, size *int, conditions *up.AndExpr) ([]*OrderList, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*OrderList
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

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *OrderList) Get(id int) (*OrderList, error) {
	var one OrderList
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *OrderList) Update(m OrderList) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *OrderList) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *OrderList) Insert(m OrderList) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
