package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/accounting-api/contextutil"
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

type ArticlesFilter struct {
	Year        string `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	OfficeID    int    `json:"office_id"`
}

// Table returns the table name
func (t *Movement) Table() string {
	return "movements"
}

// GetAll gets all records from the database, using Upper
func (t *Movement) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*Movement, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Movement
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

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *Movement) Get(id int) (*Movement, error) {
	var one Movement
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Movement) Update(ctx context.Context, m Movement) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Movement) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Movement) Insert(ctx context.Context, m Movement) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *Movement) GetAllForReport(StartDate *string, EndDate *string, Title *string, OfficeID *int, Exception *bool, OrganizationUnitID *int) ([]ArticlesFilter, error) {
	var all []ArticlesFilter

	selectS := `SELECT s.year, s.title, s.description, sum(a.amount) as amount 
	FROM movement_articles a, movements m, stocks s 
	WHERE a.movement_id = m.id AND s.id = a.stock_id `
	groupBy := ` GROUP BY s.year, s.title, s.description ORDER BY s.title asc`

	var filters []interface{}
	var filterArgs []string

	if StartDate != nil && *StartDate != "" && EndDate != nil && *EndDate != "" {
		start := *StartDate
		filters = append(filters, start)
		filterArgs = append(filterArgs, "a.created_at > $"+strconv.Itoa(len(filterArgs)+1))
	}

	if EndDate != nil && *EndDate != "" {
		end := *EndDate
		filters = append(filters, end)
		filterArgs = append(filterArgs, "a.created_at < $"+strconv.Itoa(len(filterArgs)+1))
	}

	if Title != nil && *Title != "" {
		lowercaseTitle := strings.ToLower(*Title)
		filters = append(filters, "%"+lowercaseTitle+"%")
		filterArgs = append(filterArgs, "LOWER(s.title) LIKE $"+strconv.Itoa(len(filterArgs)+1))
	}

	if OfficeID != nil && *OfficeID != 0 {
		filters = append(filters, OfficeID)
		filterArgs = append(filterArgs, "m.office_id = $"+strconv.Itoa(len(filterArgs)+1))
	}

	if Exception != nil {
		filters = append(filters, Exception)
		filterArgs = append(filterArgs, "s.exception = $"+strconv.Itoa(len(filterArgs)+1))
	}

	if OrganizationUnitID != nil && *OrganizationUnitID != 0 {
		filters = append(filters, OrganizationUnitID)
		filterArgs = append(filterArgs, "m.organization_unit_id = $"+strconv.Itoa(len(filterArgs)+1))
	}

	if len(filters) > 0 {
		selectS += " AND " + strings.Join(filterArgs, " AND ")
	}

	query := selectS + groupBy

	rows, err := Upper.SQL().Query(query, filters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticlesFilter
		err := rows.Scan(&article.Year, &article.Title, &article.Description, &article.Amount)
		if err != nil {
			return nil, err
		}

		all = append(all, article)
	}

	return all, err
}
