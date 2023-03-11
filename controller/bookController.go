package controller

import (
	"booking/config"
	"booking/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(model.Book)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingBook := new(model.Book)

	if err := db.First(&existingBook, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingBook.Descriptions = b.Descriptions
	existingBook.Users = b.Users
	if err := db.Save(&existingBook).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingBook,
	}

	return c.JSON(http.StatusOK, response)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var books []*model.Book

	if res := db.Find(&books, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": books[0],
	}

	return c.JSON(http.StatusOK, response)
}
