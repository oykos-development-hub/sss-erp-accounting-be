package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// MovementArticle struct
type MovementArticle struct {
	ID          int       `db:"id,omitempty"`
	Title       string    `db:"title"`
	MovementID  int       `db:"movement_id"`
	Description string    `db:"description"`
	Amount      int       `db:"amount"`
	CreatedAt   time.Time `db:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *MovementArticle) Table() string {
	return "movement_articles"
}

// GetAll gets all records from the database, using upper
func (t *MovementArticle) GetAll(page *int, size *int, condition *up.AndExpr) ([]*MovementArticle, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*MovementArticle
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
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
func (t *MovementArticle) Get(id int) (*MovementArticle, error) {
	var one MovementArticle
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *MovementArticle) Update(m MovementArticle) error {
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
func (t *MovementArticle) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *MovementArticle) Insert(m MovementArticle) (int, error) {
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
