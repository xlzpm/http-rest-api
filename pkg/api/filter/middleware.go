package filter

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type optionString string

const (
	OptionsContextKey optionString = "filter_options"
)

func Middleware(h gin.HandlerFunc, defaultLimit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitFromQuery := c.Request.URL.Query().Get("limit")

		limit := defaultLimit
		var limitParseErr error
		if limitFromQuery != "" {
			if limit, limitParseErr = strconv.Atoi(limitFromQuery); limitParseErr != nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				c.Writer.Write([]byte("bad limit"))
				return
			}
		}

		optionsI := NewOptions(limit)

		ctx := context.WithValue(c.Request.Context(), OptionsContextKey, optionsI)
		c.Request = c.Request.WithContext(ctx)

		h(c)
	}
}
