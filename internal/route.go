package internal

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

const (
	statusPath  = "/status"
	userPath    = "/users"
	getUserPath = "/users/:name"
)

//WireRoutes creates a new echo instance and wires routes to domain paths
func WireRoutes() *echo.Echo {

	e := echo.New()

	e.GET(statusPath, status)
	e.GET(getUserPath, getUsers)
	e.POST(userPath, postUsers)

	return e
}

func status(c echo.Context) error {
	return c.JSON(200, ServerStatus{"UP"})
}

func getUsers(c echo.Context) error {

	name := c.Param("name")

	log.Info("Getting user for name: ", name)

	u, err := FindUser(name)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, u)
}

func postUsers(c echo.Context) error {

	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	log.Info("Creating new user: ", u)

	if err := u.Save(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
