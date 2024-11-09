package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

type Basket struct {
	ID    string   `json:"id"`
	Prods []string `json:"products"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: "1", Name: "name 1", Image: "img1", Description: "decs 1", Price: "100", CategoryId: "1"},
	{ID: "2", Name: "name 2", Image: "img2", Description: "decs 2", Price: "500", CategoryId: "2"},
	{ID: "3", Name: "name 3", Image: "img3", Description: "decs 3", Price: "600", CategoryId: "3"},
	{ID: "4", Name: "name 4", Image: "img4", Description: "decs 4", Price: "200", CategoryId: "1"},
	{ID: "5", Name: "name 5", Image: "img5", Description: "decs 5", Price: "900", CategoryId: "3"},
	{ID: "6", Name: "name 6", Image: "img6", Description: "decs 6", Price: "1121", CategoryId: "2"},
}

var category = []Category{
	{ID: "1", Name: "cat 1"},
	{ID: "2", Name: "cat 2"},
	{ID: "3", Name: "cat 3"},
}

var baskets = []Basket{
	{ID: "1", Prods: []string{"1", "2", "3"}},
	{ID: "2", Prods: []string{"2", "5", "6"}},
	{ID: "3", Prods: []string{"4", "1", "5"}},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductsById)
	router.DELETE("/products/delete/:id", deleteProduct)
	router.POST("/products/create", createProduct)
	router.PUT("/products/put/:id", updateProduct)

	router.GET("/category", getCategories)
	router.GET("/category/:id", getCategoriesById)
	router.DELETE("/category/:id", deleteCategory)
	router.POST("/category/create", createCategory)
	router.PUT("/category/put/:id", updateCategory)

	router.GET("/baskets", getBaskets)
	router.GET("/baskets/:id", getBasketsById)

	router.Run(":8000")

}

func getCategories(c *gin.Context) {
	c.JSON(http.StatusOK, category)
}

func getCategoriesById(c *gin.Context) {
	id := c.Param("id")

	for _, el := range category {
		if el.ID == id {
			c.JSON(http.StatusOK, el)
		}
	}
}

func deleteCategory(c *gin.Context) {
	id := c.Param("id")
	for i, el := range category {
		if el.ID == id {
			category = append(category[:i], category[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "категория успешно удалена"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Категория не найдена"})
}

func createCategory(c *gin.Context) {
	var addCategory Category
	if err := c.BindJSON(&addCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	category = append(category, addCategory)
	c.JSON(http.StatusCreated, gin.H{"message": "категория " + addCategory.Name + " создана"})
}

func updateCategory(c *gin.Context) {
	id := c.Param("id")
	var changedCat Category
	if err := c.BindJSON(&changedCat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	for i, el := range category {
		if el.ID == id {
			category[i] = changedCat
			c.JSON(http.StatusOK, gin.H{"message": "категория изменена"})
		}
	}
}

// возвращает список корзин
func getBaskets(c *gin.Context) {
	c.JSON(http.StatusOK, baskets)
}

type jsonBasket struct {
	ID           string   `json:"id"`
	ProductsName []string `productsName:"id"`
}

// возвращает название товаров из корзины
func getBasketsById(c *gin.Context) {
	id := c.Param("id")
	arr := make([]string, 0)

	for _, el := range baskets {

		if el.ID == id {

			for _, pr := range products {
				for _, prb := range el.Prods {

					if pr.ID == prb {
						arr = append(arr, pr.Name)
					}
				}
			}
			var res = jsonBasket{ID: id, ProductsName: arr}
			c.JSON(http.StatusOK, res)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "корзина не найдена"})
}

// возвращает список всех товаров
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// возвращает информацию товара по id
func getProductsById(c *gin.Context) {
	id := c.Param("id")
	for _, el := range products {
		if el.ID == id {
			c.JSON(http.StatusOK, el)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Товар не найден"})
}

// создание товара
func createProduct(c *gin.Context) {
	var addProduct Product
	if err := c.BindJSON(&addProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	products = append(products, addProduct)
	c.JSON(http.StatusCreated, addProduct)
}

// обновление информации товара по id
func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updateProduct Product
	if err := c.BindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	for i, el := range products {
		if el.ID == id {
			products[i] = updateProduct
			c.JSON(http.StatusOK, gin.H{"message": "товар успешно обновлен"})
			return
		}
	}
}

// удаление товара по id
func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, el := range products {
		if el.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "товар удален"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "товара с таким id нет"})
}
