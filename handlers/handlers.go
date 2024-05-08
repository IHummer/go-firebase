package handlers

import (
	"net/http"

	"github.com/ihummer/go-firebase/database"
	"github.com/ihummer/go-firebase/models"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	// Cast err to *echo.HTTPError
	httpError := err.(*echo.HTTPError)

	// Creating the custom response
	response := map[string]interface{}{
		"message": httpError.Message,
		"status":  httpError.Code,
	}

	c.JSON(httpError.Code, response)
}

func RegisterHandlers(e *echo.Echo, db *firestore.Client) {

	e.POST("/products", createProduct(db))

	e.GET("/products", getAllProducts(db))

}

func createProduct(db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(models.Product)
		if err := c.Bind(p); err != nil {
			return err
		}

		if err := p.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		f := new(database.FirestoreDatabase)

		_, err := f.CreateProduct(db, p)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, p)
	}
}

func getAllProducts(db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		f := new(database.FirestoreDatabase)

		products, err := f.GetAllProducts(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, products)
	}
}
