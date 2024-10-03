package controllers

import (
	"estate/models"
	"estate/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// type Estate struct {
// 	Description string  `json:"description" xml:"description" form:"description" query:"description"`
// 	Price       float32 `json:"price" xml:"price" form:"price" query:"price"`
// }

func CreateEstate(c echo.Context) error {
	b := new(models.Estate)
	db := storage.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	book := &models.Estate{
		// Name:        b.Name,
		Description: b.Description,
	}

	if err := db.Create(&book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateEstate(c echo.Context) error {
	id := c.Param("id")
	b := new(models.Estate)
	db := storage.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_book := new(models.Estate)

	if err := db.First(&existing_book, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	// existing_book.Name = b.Name
	existing_book.Description = b.Description
	if err := db.Save(&existing_book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_book,
	}

	return c.JSON(http.StatusOK, response)
}

func GetEstate(c echo.Context) error {
	id := c.Param("id")
	db := storage.DB()

	var books []*models.Estate

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

func GetEstates(c echo.Context) error {
	// id := c.Param("id")
	db := storage.DB()

	var estates []*models.Estate

	// if res := db.Find(&estates, id); res.Error != nil {
	// 	data := map[string]interface{}{
	// 		"message": res.Error.Error(),
	// 	}

	qp := c.QueryParam("page")
	page, err := strconv.Atoi(qp)
	if err != nil {
		page = 1
	}
	orderBy := c.QueryParam("order_by")
	orderByQuery := "price asc"
	if orderBy == "-price" {
		// db.Order("price desc")
		orderByQuery = "price desc"
	}
	if orderBy == "price" || orderBy == "" {
		// db.Order("price asc")
		orderByQuery = "price asc"
	}
	if orderBy == "created_at" {
		// db.Order("created_at asc")
		orderByQuery = "created_at asc"
	}
	if orderBy == "-created_at" {
		// db.Order("created_at desc")
		orderByQuery = "created_at desc"
	}
	// page, err := strconv.Atoi(ob)
	// if err != nil {
	// 	page = 1
	// }
	// panic(page)
	if res := db.Order(orderByQuery).Scopes(storage.Paginate(page, 9)).Find(&estates); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	// var meta map[string]interface{}
	// meta["last_page"] := 6
	meta := map[string]interface{}{
		"last_page": 6,
	}
	response := map[string]interface{}{
		"data": estates,
		"meta": meta,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteEstate(c echo.Context) error {
	id := c.Param("id")
	db := storage.DB()

	book := new(models.Estate)

	err := db.Delete(&book, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a book has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
