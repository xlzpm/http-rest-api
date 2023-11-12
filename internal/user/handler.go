package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserById)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(c *gin.Context) {
	c.String(http.StatusOK, "this is list of users")
}

func (h *handler) CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "create user")
}

func (h *handler) GetUserById(c *gin.Context) {
	c.String(http.StatusOK, "get user by id")
}

func (h *handler) UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "fully update user")
}

func (h *handler) PartiallyUpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "partially update user")
}

func (h *handler) DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "delete user")
}
