package repo

import (
	"database/sql"
	"fmt"
	"github.com/T4jgat/cobalt/internal/entity"
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

func (r *CatalogRepo) GetAll(filters map[string]string, sort string) ([]*entity.Catalog, error) {
	query := "SELECT id, model, brand, color, price FROM catalog WHERE 1=1"
	args := []interface{}{}
	argID := 1

	if color, ok := filters["color"]; ok {
		query += fmt.Sprintf(" AND color = $%d", argID)
		args = append(args, color)
		argID++
	}

	if price, ok := filters["price"]; ok {
		query += fmt.Sprintf(" AND price = $%d", argID)
		args = append(args, price)
		argID++
	}

	if sort != "" {
		query += " ORDER BY " + sort
	}

	rows, err := r.db.Query(query, args...)
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

func (r *CatalogRepo) GetByModel(model string) ([]*entity.Catalog, error) {
	rows, err := r.db.Query("SELECT id, model, brand, color, price FROM catalog WHERE model = $1", model)
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
