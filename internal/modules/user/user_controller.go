package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quangdtptit/go-cli/internal/modules/common"
)

var _ common.EchoController = (*Controller)(nil)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Register(ec *echo.Group) {
	ec.GET("/users", c.GetAll)
	ec.GET("/users/:userId", c.GetDetailByUserId)
	ec.POST("/users", c.Create)
	ec.PUT("/users/:userId", c.Update)
	ec.DELETE("/users/:userId", c.DeleteByUserId)
}

func (c *Controller) GetAll(ctx echo.Context) error {
	users, _ := c.service.GetAll()
	return ctx.JSON(http.StatusOK, users)
}

func (c *Controller) GetDetailByUserId(ctx echo.Context) error {
	return nil
}

func (c *Controller) Create(ctx echo.Context) error {
	return nil
}

func (c *Controller) Update(ctx echo.Context) error {
	return nil
}

func (c *Controller) DeleteByUserId(ctx echo.Context) error {
	return nil
}
