package repo

import (
	"database/sql"
	"fmt"
	"github.com/T4jgat/cobalt/internal/entity"
	"strings"
)

type CatalogRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *CatalogRepo {
	return &CatalogRepo{db: db}
}

func (r *CatalogRepo) Create(car *entity.Catalog) error {

	_, err := r.db.Exec("INSERT INTO catalog (id, model, brand, color, price) VALUES ($1, $2, $3, $4, $5)",
		car.ID, car.Model, car.Brand, car.Color, car.Price)
	return err
}

func (r *CatalogRepo) GetAll(filters map[string]string) ([]*entity.Catalog, error) {
	baseQuery := "SELECT id, model, brand, color, price FROM catalog"
	var conditions []string
	var args []interface{}

	for key, value := range filters {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", key, len(args)+1))
		args = append(args, value)
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []*entity.Catalog{}
	for rows.Next() {
		car := &entity.Catalog{}
		err := rows.Scan(&car.ID, &car.Model, &car.Brand, &car.Color, &car.Price)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (r *CatalogRepo) Get(id int) (*entity.Catalog, error) {
	row := r.db.QueryRow("SELECT id, model, brand, color, price FROM catalog WHERE id = $1", id)

	car := &entity.Catalog{}
	err := row.Scan(&car.ID, &car.Model, &car.Brand, &car.Color, &car.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No row found
		}
		return nil, err
	}
	return car, nil
}

func (r *CatalogRepo) Update(car *entity.Catalog) error {
	_, err := r.db.Exec("UPDATE catalog SET model = $1, brand = $2, color = $3, price = $4 WHERE id = $5",
		car.Model, car.Brand, car.Color, car.Price, car.ID)
	return err
}

func (r *CatalogRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM catalog WHERE id = $1", id)
	return err
}
