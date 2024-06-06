package v1

import (
	"testing"

	"github.com/T4jgat/cobalt/pkg/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	catalogsController := &CatalogsController{}
	log := logger.New("info")
	router := NewRouter(catalogsController, *log)

	assert.IsType(t, &httprouter.Router{}, router)
}
