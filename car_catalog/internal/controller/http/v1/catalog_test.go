package v1

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/T4jgat/cobalt/internal/entity"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//// MockCatalogRepo mocks the CatalogRep interface.
//type MockCatalogRepo struct {
//	mock.Mock
//}
//
//// Create implements the CatalogRep Create method.
//func (m *MockCatalogRepo) Create(car *entity.Catalog) error {
//	args := m.Called(car)
//	return args.Error(0)
//}
//
//// GetAll implements the CatalogRep GetAll method.
//func (m *MockCatalogRepo) GetAll(filters map[string]string, sort string) ([]*entity.Catalog, error) {
//	args := m.Called(filters, sort)
//	return args.Get(0).([]*entity.Catalog), args.Error(1)
//}
//
//// Get implements the CatalogRep Get method.
//func (m *MockCatalogRepo) Get(id int) (*entity.Catalog, error) {
//	args := m.Called(id)
//	return args.Get(0).(*entity.Catalog), args.Error(1)
//}
//
//// Update implements the CatalogRep Update method.
//func (m *MockCatalogRepo) Update(car *entity.Catalog) error {
//	args := m.Called(car)
//	return args.Error(0)
//}
//
//// Delete implements the CatalogRep Delete method.
//func (m *MockCatalogRepo) Delete(id int) error {
//	args := m.Called(id)
//	return args.Error(0)
//}
//
//func TestCatalogsController_CreateCar(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Create", mock.Anything).Return(nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		reqBody := []byte(`{"model": "Model A", "brand": "Brand A", "color": "Red", "price": 10000}`)
//		req, err := http.NewRequest(http.MethodPost, "/catalogs", bytes.NewReader(reqBody))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.CreateCar(rr, req)
//
//		assert.Equal(t, http.StatusCreated, rr.Code)
//
//		var responseMap map[string]interface{}
//		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
//		require.NoError(t, err)
//
//		assert.Equal(t, "Catalog entry successfully created", responseMap["message"])
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("ErrorDecoding", func(t *testing.T) {
//		controller := New(&MockCatalogRepo{}) // Передаем mockRepo как интерфейс
//		reqBody := []byte(`{ "model": "Model A", "brand": "Brand A", "color": "Red"`)
//		req, err := http.NewRequest(http.MethodPost, "/catalogs", bytes.NewReader(reqBody))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.CreateCar(rr, req)
//
//		assert.Equal(t, http.StatusBadRequest, rr.Code)
//	})
//
//	t.Run("ErrorCreating", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Create", mock.Anything).Return(fmt.Errorf("database error"))
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		reqBody := []byte(`{"model": "Model A", "brand": "Brand A", "color": "Red", "price": 10000}`)
//		req, err := http.NewRequest(http.MethodPost, "/catalogs", bytes.NewReader(reqBody))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.CreateCar(rr, req)
//
//		assert.Equal(t, http.StatusInternalServerError, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//}
//
//func TestCatalogsController_GetAllCars(t *testing.T) {
//	t.Run("SuccessAscending", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("GetAll", mock.Anything, "price").Return([]*entity.Catalog{
//			{ID: 1, Model: "Model A", Brand: "Brand A", Color: "Red", Price: 10000},
//			{ID: 2, Model: "Model B", Brand: "Brand B", Color: "Blue", Price: 20000},
//		}, nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodGet, "/catalogs?sort=price&order=asc", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetAllCars(rr, req)
//
//		assert.Equal(t, http.StatusOK, rr.Code)
//
//		var responseMap map[string]interface{}
//		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
//		require.NoError(t, err)
//
//		cars := responseMap["cars"].([]interface{})
//		assert.Len(t, cars, 2)
//		assert.Equal(t, 10000.0, cars[0].(map[string]interface{})["price"])
//		assert.Equal(t, 20000.0, cars[1].(map[string]interface{})["price"])
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("SuccessDescending", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("GetAll", mock.Anything, "price").Return([]*entity.Catalog{
//			{ID: 2, Model: "Model B", Brand: "Brand B", Color: "Blue", Price: 20000},
//			{ID: 1, Model: "Model A", Brand: "Brand A", Color: "Red", Price: 10000},
//		}, nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodGet, "/catalogs?sort=price&order=desc", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetAllCars(rr, req)
//
//		assert.Equal(t, http.StatusOK, rr.Code)
//
//		var responseMap map[string]interface{}
//		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
//		require.NoError(t, err)
//
//		cars := responseMap["cars"].([]interface{})
//		assert.Len(t, cars, 2)
//		assert.Equal(t, 20000.0, cars[0].(map[string]interface{})["price"])
//		assert.Equal(t, 10000.0, cars[1].(map[string]interface{})["price"])
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("ErrorGettingCars", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("GetAll", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("database error"))
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodGet, "/catalogs", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetAllCars(rr, req)
//
//		assert.Equal(t, http.StatusInternalServerError, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//}
//
//func TestCatalogsController_GetCar(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Get", 1).Return(&entity.Catalog{ID: 1, Model: "Model A", Brand: "Brand A", Color: "Red", Price: 10000}, nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodGet, "/catalogs/1", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetCar(rr, req)
//
//		assert.Equal(t, http.StatusOK, rr.Code)
//
//		var responseMap map[string]interface{}
//		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
//		require.NoError(t, err)
//
//		car := responseMap["car"].(map[string]interface{})
//		assert.Equal(t, 1, int(car["id"].(float64)))
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("InvalidID", func(t *testing.T) {
//		controller := New(&MockCatalogRepo{}) // Передаем mockRepo как интерфейс
//		req, err := http.NewRequest(http.MethodGet, "/catalogs/abc", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetCar(rr, req)
//
//		assert.Equal(t, http.StatusBadRequest, rr.Code)
//	})
//
//	t.Run("CarNotFound", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Get", 1).Return(nil, nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodGet, "/catalog
//		req, err := http.NewRequest(http.MethodGet, "/catalogs/1", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetCar(rr, req)
//
//		assert.Equal(t, http.StatusNotFound, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("ErrorGettingCar", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Get", 1).Return(nil, fmt.Errorf("database error"))
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodGet, "/catalogs/1", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.GetCar(rr, req)
//
//		assert.Equal(t, http.StatusInternalServerError, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//}
//
//func TestCatalogsController_UpdateCar(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Update", mock.Anything).Return(nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		reqBody := []byte(`{"id": 1, "model": "Model A Updated", "brand": "Brand A", "color": "Red", "price": 12000}`)
//		req, err := http.NewRequest(http.MethodPut, "/catalogs/1", bytes.NewReader(reqBody))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.UpdateCar(rr, req)
//
//		assert.Equal(t, http.StatusOK, rr.Code)
//
//		var responseMap map[string]interface{}
//		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
//		require.NoError(t, err)
//
//		assert.Equal(t, "Catalog entry successfully updated", responseMap["message"])
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("ErrorDecoding", func(t *testing.T) {
//		controller := New(&MockCatalogRepo{}) // Передаем mockRepo как интерфейс
//		reqBody := []byte(`{ "id": 1, "model": "Model A Updated", "brand": "Brand A", "color": "Red"`)
//		req, err := http.NewRequest(http.MethodPut, "/catalogs/1", bytes.NewReader(reqBody))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.UpdateCar(rr, req)
//
//		assert.Equal(t, http.StatusBadRequest, rr.Code)
//	})
//
//	t.Run("ErrorUpdating", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Update", mock.Anything).Return(fmt.Errorf("database error"))
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		reqBody := []byte(`{"id": 1, "model": "Model A Updated", "brand": "Brand A", "color": "Red", "price": 12000}`)
//		req, err := http.NewRequest(http.MethodPut, "/catalogs/1", bytes.NewReader(reqBody))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.UpdateCar(rr, req)
//
//		assert.Equal(t, http.StatusInternalServerError, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//}
//
//func TestCatalogsController_DeleteCar(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Delete", 1).Return(nil)
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodDelete, "/catalogs/1", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.DeleteCar(rr, req)
//
//		assert.Equal(t, http.StatusOK, rr.Code)
//
//		var responseMap map[string]interface{}
//		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
//		require.NoError(t, err)
//
//		assert.Equal(t, "Catalog entry successfully deleted", responseMap["message"])
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("InvalidID", func(t *testing.T) {
//		controller := New(&MockCatalogRepo{}) // Передаем mockRepo как интерфейс
//		req, err := http.NewRequest(http.MethodDelete, "/catalogs/abc", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.DeleteCar(rr, req)
//
//		assert.Equal(t, http.StatusBadRequest, rr.Code)
//	})
//
//	t.Run("CarNotFound", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Delete", 1).Return(fmt.Errorf("not found"))
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodDelete, "/catalogs/1", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.DeleteCar(rr, req)
//
//		assert.Equal(t, http.StatusNotFound, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//
//	t.Run("ErrorDeleting", func(t *testing.T) {
//		mockRepo := &MockCatalogRepo{}
//		mockRepo.On("Delete", 1).Return(fmt.Errorf("database error"))
//		controller := New(mockRepo) // Передаем mockRepo как интерфейс
//
//		req, err := http.NewRequest(http.MethodDelete, "/catalogs/1", nil)
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		controller.DeleteCar(rr, req)
//
//		assert.Equal(t, http.StatusInternalServerError, rr.Code)
//		mockRepo.AssertExpectations(t)
//	})
//}
