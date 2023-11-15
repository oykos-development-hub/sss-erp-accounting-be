package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Movement struct
type Movement struct {
	ID                 int       `db:"id,omitempty"`
	DateOrder          time.Time `db:"date_order"`
	OrganizationUnitID int       `db:"organization_unit_id"`
	OfficeID           int       `db:"office_id"`
	RecipientUserID    int       `db:"recipient_user_id"`
	Description        string    `db:"description"`
	FileID             int       `db:"file_id"`
	CreatedAt          time.Time `db:"created_at,omitempty"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Movement) Table() string {
	return "movements"
}

// GetAll gets all records from the database, using upper
func (t *Movement) GetAll(page *int, size *int, condition *up.AndExpr) ([]*Movement, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Movement
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
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

// Get gets one record from the database, by id, using upper
func (t *Movement) Get(id int) (*Movement, error) {
	var one Movement
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Movement) Update(m Movement) error {
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
func (t *Movement) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Movement) Insert(m Movement) (int, error) {
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
