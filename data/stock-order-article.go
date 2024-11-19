package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// StockOrderArticle struct
type StockOrderArticle struct {
	ID        int       `db:"id,omitempty"`
	ArticleID int       `db:"article_id"`
	StockID   int       `db:"stock_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *StockOrderArticle) Table() string {
	return "stock_order_articles"
}

// GetAll gets all records from the database, using upper
func (t *StockOrderArticle) GetAll(condition *up.Cond) ([]*StockOrderArticle, error) {
	collection := Upper.Collection(t.Table())
	var all []*StockOrderArticle
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *StockOrderArticle) Get(id int) (*StockOrderArticle, error) {
	var one StockOrderArticle
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *StockOrderArticle) Update(m StockOrderArticle) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *StockOrderArticle) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *StockOrderArticle) Insert(m StockOrderArticle) (int, error) {
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
