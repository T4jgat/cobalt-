package repo

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/T4jgat/cobalt/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCatalogRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(db)
	car := &entity.Catalog{
		ID:    1,
		Model: "Model A",
		Brand: "Brand A",
		Color: "Red",
		Price: 10000,
	}

	mock.ExpectExec("INSERT INTO catalog").
		WithArgs(car.ID, car.Model, car.Brand, car.Color, car.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(car)
	assert.NoError(t, err)

	// Ensure the expected query was executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCatalogRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(db)

	// Test case 1: No filters
	rows := sqlmock.NewRows([]string{"id", "model", "brand", "color", "price"}).
		AddRow(1, "Model A", "Brand A", "Red", 10000).
		AddRow(2, "Model B", "Brand B", "Blue", 20000)

	mock.ExpectQuery(`SELECT id, model, brand, color, price FROM catalog WHERE 1=1 ORDER BY id`).
		WillReturnRows(rows)

	cars, err := repo.GetAll(nil, "id")
	assert.NoError(t, err)
	assert.Len(t, cars, 2)
	assert.Equal(t, cars[0].ID, 1)
	assert.Equal(t, cars[0].Model, "Model A")
	assert.Equal(t, cars[0].Brand, "Brand A")
	assert.Equal(t, cars[0].Color, "Red")
	assert.Equal(t, cars[0].Price, 10000)

	// Test case 2: With filters
	rows = sqlmock.NewRows([]string{"id", "model", "brand", "color", "price"}).
		AddRow(1, "Model A", "Brand A", "Red", 10000)

	mock.ExpectQuery(`SELECT id, model, brand, color, price FROM catalog WHERE 1=1 AND color = \$1 ORDER BY id`).
		WithArgs("Red").
		WillReturnRows(rows)

	cars, err = repo.GetAll(map[string]string{"color": "Red"}, "id")
	assert.NoError(t, err)
	assert.Len(t, cars, 1)
	assert.Equal(t, cars[0].ID, 1)
	assert.Equal(t, cars[0].Model, "Model A")
	assert.Equal(t, cars[0].Brand, "Brand A")
	assert.Equal(t, cars[0].Color, "Red")
	assert.Equal(t, cars[0].Price, 10000)

	// Ensure the expected query was executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestCatalogRepo_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(db)

	// Test case 1: Car found
	rows := sqlmock.NewRows([]string{"id", "model", "brand", "color", "price"}).
		AddRow(1, "Model A", "Brand A", "Red", 10000)

	mock.ExpectQuery(`SELECT id, model, brand, color, price FROM catalog WHERE id = \$1`). // Use regular expression
												WithArgs(1).
												WillReturnRows(rows)

	car, err := repo.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, car.ID, 1)
	assert.Equal(t, car.Model, "Model A")
	assert.Equal(t, car.Brand, "Brand A")
	assert.Equal(t, car.Color, "Red")
	assert.Equal(t, car.Price, 10000)

	// Test case 2: Car not found
	mock.ExpectQuery(`SELECT id, model, brand, color, price FROM catalog WHERE id = \$1`). // Use regular expression
												WithArgs(2).
												WillReturnError(sql.ErrNoRows)

	car, err = repo.Get(2)
	assert.NoError(t, err)
	assert.Nil(t, car)

	// Ensure the expected query was executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
