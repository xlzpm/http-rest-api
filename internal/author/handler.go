package author

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xlzpm/internal/apperror"
	"github.com/xlzpm/internal/author/service"
	"github.com/xlzpm/internal/handlers"
	"github.com/xlzpm/pkg/api/sort"
	"github.com/xlzpm/pkg/logging"
)

const (
	authorsURL = "/authors"
	authorURL  = "/authors/:uuid"
)

type handler struct {
	logger  *logging.Logger
	service *service.Service
}

func NewHandler(service *service.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.Handle("GET", authorsURL, sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC))
}

func (h *handler) GetList(c *gin.Context) error {

	var sortOptions sort.Options
	if options, ok := c.Request.Context().Value(sort.OptionsContextKey).(sort.Options); ok {
		sortOptions = options
	}

	all, err := h.service.GetAll(c.Request.Context(), sortOptions)
	if err != nil {
		c.Writer.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(allBytes)

	return nil
}
