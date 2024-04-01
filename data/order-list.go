package data

import (
	"time"

	"github.com/lib/pq"
	up "github.com/upper/db/v4"
)

// Order list struct
type OrderList struct {
	ID                    int           `db:"id,omitempty"`
	DateOrder             time.Time     `db:"date_order"`
	TotalPrice            float32       `db:"total_price"`
	PublicProcurementID   *int          `db:"public_procurement_id"`
	SupplierID            *int          `db:"supplier_id"`
	IsUsed                bool          `db:"is_used"`
	Status                string        `db:"status"`
	DateSystem            *time.Time    `db:"date_system"`
	ProFormaInvoiceDate   time.Time     `db:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string        `db:"pro_forma_invoice_number"`
	InvoiceDate           *time.Time    `db:"invoice_date"`
	InvoiceNumber         *string       `db:"invoice_number"`
	OrganizationUnitID    int           `db:"organization_unit_id"`
	GroupOfArticlesID     *int          `db:"group_of_articles_id"`
	OfficeID              *int          `db:"office_id"`
	AccountID             *int          `db:"account_id"`
	RecipientUserID       *int          `db:"recipient_user_id"`
	Description           *string       `db:"description"`
	PassedToFinance       bool          `db:"passed_to_finance"`
	UsedInFinance         bool          `db:"used_in_finance"`
	IsProFormaInvoice     bool          `db:"is_pro_forma_invoice"`
	OrderFile             *int          `db:"order_file"`
	ReceiveFile           pq.Int64Array `db:"receive_file"`
	MovementFile          *int          `db:"movement_file"`
	CreatedAt             time.Time     `db:"created_at,omitempty"`
	UpdatedAt             time.Time     `db:"updated_at"`
}

// Table returns the table name
func (t *OrderList) Table() string {
	return "order_lists"
}

// GetAll gets all records from the database, using upper
func (t *OrderList) GetAll(page *int, size *int, conditions *up.AndExpr, orders []interface{}) ([]*OrderList, *uint64, error) {
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

	err = res.OrderBy(orders...).All(&all)
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

func (t *OrderList) SendToFinance(id int) error {
	query := `update order_lists set passed_to_finance = true where id = $1`

	rows, err := upper.SQL().Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
