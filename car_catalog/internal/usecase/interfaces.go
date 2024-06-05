package usecase

import (
	"context"
	"github.com/T4jgat/cobalt/internal/entity"
)

type (
	Catalog interface {
		Catalogs(context.Context) ([]entity.Catalog, error)
	}

	CatalogRepo interface {
	}

	CatalogWebApi interface {
	}
)
