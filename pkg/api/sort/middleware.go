package sort

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type optionString string

const (
	ASC                            = "ASC"
	DESC                           = "DESC"
	OptionsContextKey optionString = "sort_options"
)

func Middleware(h gin.HandlerFunc, defaultSortField, defaultSortOrder string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sortBy := c.Request.URL.Query().Get("sort_by")
		sortOrder := c.Request.URL.Query().Get("sort_order")

		if sortBy == "" {
			sortBy = defaultSortField
		}
		if sortOrder == "" {
			sortOrder = defaultSortOrder
		} else {
			upperSortOrder := strings.ToUpper(sortOrder)
			if upperSortOrder != ASC && upperSortOrder != DESC {
				c.Writer.WriteHeader(http.StatusBadRequest)
				c.Writer.Write([]byte("bad sort order"))
				return
			}

		}

		options := Options{
			Field: sortBy,
			Order: sortOrder,
		}

		ctx := context.WithValue(c.Request.Context(), OptionsContextKey, options)
		c.Request = c.Request.WithContext(ctx)

		h(c)
	}
}

type Options struct {
	Field, Order string
}
