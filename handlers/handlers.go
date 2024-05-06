package handlers

import (
	"net/http"

	"github.com/ihummer/go-firebase/database"
	"github.com/ihummer/go-firebase/models"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

// Registrar manejadores de solicitudes
func RegisterHandlers(e *echo.Echo, db *firestore.Client) {
	// Crear producto
	e.POST("/products", createProduct(db))

	// Obtener todos los productos
	e.GET("/products", getAllProducts(db))
}

func createProduct(db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(models.Product)
		if err := c.Bind(p); err != nil {
			return err
		}

		if err := p.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		_, _, err := database.CreateProduct(db, p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, p)
	}
}

func getAllProducts(db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := database.GetAllProducts(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, products)
	}
}
