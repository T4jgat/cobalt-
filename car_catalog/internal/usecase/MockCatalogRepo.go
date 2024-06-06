package usecase

//
//import (
//	"github.com/T4jgat/cobalt/internal/entity"
//)
//
//// MockCatalogRepo представляет собой макет для тестирования репозитория каталогов.
//type MockCatalogRepo struct {
//	CreateFunc     func(catalog *entity.Catalog) error
//	GetAllFunc     func(filters map[string]string, sort string) ([]*entity.Catalog, error)
//	GetFunc        func(id int) (*entity.Catalog, error)
//	UpdateFunc     func(catalog *entity.Catalog) error
//	DeleteFunc     func(id int) error
//	GetByModelFunc func(model string) ([]*entity.Catalog, error)
//}
//
//// Create реализует метод Create интерфейса CatalogRepoInterface для MockCatalogRepo.
//func (m *MockCatalogRepo) Create(catalog *entity.Catalog) error {
//	if m.CreateFunc != nil {
//		return m.CreateFunc(catalog)
//	}
//	return nil
//}
//
//// GetAll реализует метод GetAll интерфейса CatalogRepoInterface для MockCatalogRepo.
//func (m *MockCatalogRepo) GetAll(filters map[string]string, sort string) ([]*entity.Catalog, error) {
//	if m.GetAllFunc != nil {
//		return m.GetAllFunc(filters, sort)
//	}
//	return nil, nil
//}
//
//// Get реализует метод Get интерфейса CatalogRepoInterface для MockCatalogRepo.
//func (m *MockCatalogRepo) Get(id int) (*entity.Catalog, error) {
//	if m.GetFunc != nil {
//		return m.GetFunc(id)
//	}
//	return nil, nil
//}
//
//// Update реализует метод Update интерфейса CatalogRepoInterface для MockCatalogRepo.
//func (m *MockCatalogRepo) Update(catalog *entity.Catalog) error {
//	if m.UpdateFunc != nil {
//		return m.UpdateFunc(catalog)
//	}
//	return nil
//}
//
//// Delete реализует метод Delete интерфейса CatalogRepoInterface для MockCatalogRepo.
//func (m *MockCatalogRepo) Delete(id int) error {
//	if m.DeleteFunc != nil {
//		return m.DeleteFunc(id)
//	}
//	return nil
//}
//
//// GetByModel реализует метод GetByModel интерф
