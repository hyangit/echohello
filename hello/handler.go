package hello

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"hyan.com/echohello/singleton"
)

// Hello handler
func Hello(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := singleton.GetValidator().ValidateField(id, "gte=10,lte=100"); err != nil {
		return err
	}

	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	u.ID = id
	return c.JSON(http.StatusOK, u)
}

// Index handler
func Index(c echo.Context) error {
	html := singleton.GetResource().GetTpl("index.html")
	return c.HTML(http.StatusOK, html)
}
