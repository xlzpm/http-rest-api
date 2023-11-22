package author

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xlzpm/internal/apperror"
	"github.com/xlzpm/internal/author/service"
	"github.com/xlzpm/internal/handlers"
	"github.com/xlzpm/pkg/api/filter"
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
	router.Handle("GET", authorsURL, filter.Middleware(sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC), 10))
}

func (h *handler) GetList(c *gin.Context) error {

	filterOptions := c.Request.Context().Value(filter.OptionsContextKey).(filter.Options)

	name := c.Request.URL.Query().Get("name")
	if name != "" {
		err := filterOptions.AddField("name", filter.OperatorLike, name, filter.DataTypeStr)
		if err != nil {
			return err
		}
	}

	age := c.Request.URL.Query().Get("age")
	if age != "" {
		operator := filter.OperatorEq
		value := age
		if !strings.Contains(age, ":") {
			split := strings.Split(age, ":")
			operator = split[0]
			value = split[1]
		}
		err := filterOptions.AddField("age", operator, value, filter.DataTypeInt)
		if err != nil {
			return err
		}
	}

	isAlive := c.Request.URL.Query().Get("is_alive")
	if isAlive != "" {
		_, err := strconv.ParseBool(isAlive)
		if err != nil {
			validationErr := apperror.BadRequest("filter params validation failed",
				"bool value wrong parameter")
			validationErr.WithParams(map[string]string{
				"is_alive": "this field should be boolean: true or false",
			})

			return validationErr
		}
		err = filterOptions.AddField("is_alive", filter.OperatorEq, isAlive, filter.DataTypeBool)
		if err != nil {
			return err
		}
	}

	createdAt := c.Request.URL.Query().Get("created_at")
	if createdAt != "" {
		var operator string
		if !strings.Contains(createdAt, ":") {
			operator = filter.OperatorBetween
		} else {
			operator = filter.OperatorEq
		}
		err := filterOptions.AddField("is_alive", operator, createdAt, filter.DataTypeDate)
		if err != nil {
			return err
		}
	}

	var sortOption sort.Options
	var filterOption filter.Options
	if options, ok := c.Request.Context().Value(sort.OptionsContextKey).(sort.Options); ok {
		sortOption = options
	}

	all, err := h.service.GetAll(c.Request.Context(), filterOption, sortOption)
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
