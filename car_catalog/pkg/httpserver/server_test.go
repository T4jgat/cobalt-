package httpserver

import (
	"testing"

	v1 "github.com/T4jgat/cobalt/internal/controller/http/v1"
	"github.com/T4jgat/cobalt/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	catalogsController := &v1.CatalogsController{}
	log := logger.New("info")
	router := New(catalogsController, *log)

	assert.IsType(t, &mux.Router{}, router)
}
