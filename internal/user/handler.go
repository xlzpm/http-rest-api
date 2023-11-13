package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xlzpm/internal/apperror"
	"github.com/xlzpm/internal/handlers"
	"github.com/xlzpm/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.Handle("GET", usersURL, apperror.Middleware(h.GetList))
	router.Handle("POST", usersURL, apperror.Middleware(h.CreateUser))
	router.Handle("GET", userURL, apperror.Middleware(h.GetUserById))
	router.Handle("PUT", userURL, apperror.Middleware(h.UpdateUser))
	router.Handle("PATCH", userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.Handle("DELETE", userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(c *gin.Context) error {

	return apperror.ErrNotFound
}

func (h *handler) CreateUser(c *gin.Context) error {

	return fmt.Errorf("this is API ERROR")
}

func (h *handler) GetUserById(c *gin.Context) error {

	return apperror.NewAppError(nil, "test", "test", "t13")
}

func (h *handler) UpdateUser(c *gin.Context) error {
	c.String(http.StatusOK, "fully update user")

	return nil
}

func (h *handler) PartiallyUpdateUser(c *gin.Context) error {
	c.String(http.StatusOK, "partially update user")

	return nil
}

func (h *handler) DeleteUser(c *gin.Context) error {
	c.String(http.StatusOK, "delete user")

	return nil
}
