package data

import (
	"strconv"
	"strings"
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

// GetAll gets all records from the database, using upper
func (t *Movement) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*Movement, *uint64, error) {
	collection := upper.Collection(t.Table())
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

func (t *Movement) GetAllForReport(StartDate *string, EndDate *string, Title *string, OfficeID *int, Exception *bool, OrganizationUnitID *int) ([]ArticlesFilter, error) {
	var all []ArticlesFilter

	selectS := `SELECT a.year, a.title, a.description, sum(a.amount) as amount`
	from := ` FROM movement_articles a, movements m`
	where := ` WHERE a.movement_id = m.id`
	groupBy := ` GROUP BY a.year, a.title, a.description`
	orderBy := ` ORDER BY a.title asc`

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
		filterArgs = append(filterArgs, "LOWER(a.title) LIKE $"+strconv.Itoa(len(filterArgs)+1))
	}

	if OfficeID != nil && *OfficeID != 0 {
		filters = append(filters, OfficeID)
		filterArgs = append(filterArgs, "m.office_id = $"+strconv.Itoa(len(filterArgs)+1))
	}

	if Exception != nil {
		filters = append(filters, Exception)
		filterArgs = append(filterArgs, "a.exception = $"+strconv.Itoa(len(filterArgs)+1))
	}

	if OrganizationUnitID != nil {
		selectS += ", m.office_id"
		groupBy += ", m.office_id"
		filters = append(filters, OrganizationUnitID)
		filterArgs = append(filterArgs, "m.organization_unit_id = $"+strconv.Itoa(len(filterArgs)+1))
	}

	if len(filters) > 0 {
		where += " AND " + strings.Join(filterArgs, " AND ")
	}

	query := selectS + from + where + groupBy + orderBy

	rows, err := upper.SQL().Query(query, filters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticlesFilter
		if OrganizationUnitID != nil {
			err := rows.Scan(&article.Year, &article.Title, &article.Description, &article.Amount, &article.OfficeID)
			if err != nil {
				return nil, err
			}
		} else {
			err := rows.Scan(&article.Year, &article.Title, &article.Description, &article.Amount)
			if err != nil {
				return nil, err
			}
		}
		all = append(all, article)
	}

	return all, err
}
