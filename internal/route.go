package internal

import (
	"github.com/labstack/echo"
)

const (
	statusPath = "/status"
)

//WireRoutes creates a new echo instance and wires routes to domain paths
func WireRoutes() *echo.Echo {

	e := echo.New()

	e.GET(statusPath, status)

	return e
}

func status(ctx echo.Context) error {
	return ctx.JSON(200, ServerStatus{"UP"})
}
