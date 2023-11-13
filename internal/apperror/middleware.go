package apperror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appHandler func(c *gin.Context) error

func Middleware(h appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		var appErr *AppError

		err := h(c)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					c.Writer.WriteHeader(http.StatusNotFound)
					c.Writer.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				c.Writer.WriteHeader(http.StatusBadRequest)
				c.Writer.Write(appErr.Marshal())
				return
			}

			c.Writer.WriteHeader(http.StatusTeapot)
			c.Writer.Write(systemError(err).Marshal())
		}

	}

}
