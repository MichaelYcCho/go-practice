package member

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	var req CreateRequest // request와 일치하는 dto struct를 선언
	// 실제 넘겨받은 request를 struct에 담아줌
	if err := c.Bind(&req); err != nil {
		return ErrBindingError
	}

	res, err := controller.service.Create(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) Read(c echo.Context) error {
	id := c.Param("id")
	res, err := controller.service.GetByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (controller Controller) Update(c echo.Context) error {
	var req UpdateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	res, _ := controller.service.Update(req)
	return c.JSON(res.Code, res)
}

func (controller Controller) Delete(c echo.Context) error {
	id := c.Param("id")
	res, err := controller.service.Delete(id)
	if err != nil {
		return err
	}
	return c.JSON(res.Code, res)
}
