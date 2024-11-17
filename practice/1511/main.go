package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("key")

func isTokened() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "пользователь не авторизован"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userId := claims["userId"].(string)
			role := claims["role"].(string)
			c.Set("userId", userId)
			c.Set("role", role)
		}
		c.Next()
	}
}

func isAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "не авторизован"})
			c.Abort()
			return
		}
		if roleStr, ok := role.(string); ok && roleStr == "ADMIN" {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "доступ запрещен"})
			c.Abort()
			return
		}
	}
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

type Basket struct {
	ID     string   `json:"id"`
	Prods  []string `json:"products"`
	UserId string   `json:"userId"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
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
	{ID: "1", Prods: []string{"1", "2", "3"}, UserId: "1"},
	{ID: "2", Prods: []string{"2", "5", "6"}, UserId: "2"},
	{ID: "3", Prods: []string{"4", "1", "5"}, UserId: "3"},
}

var user = []User{
	{ID: "1", Name: "Ivan", Login: "ivan@mail.ru", Password: "pass"},
	{ID: "2", Name: "Dmitry", Login: "dima@mail.ru", Password: "play"},
	{ID: "3", Name: "Silvestr", Login: "silya@mail.ru", Password: "camelCasePassword"},
	{ID: "4", Name: "Admin", Login: "admin@mail.ru", Password: "admin", Role: "ADMIN"},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductsById)
	router.DELETE("/products/delete/:id", isTokened(), isAdmin(), deleteProduct)
	router.POST("/products/create", isTokened(), isAdmin(), createProduct)
	router.PUT("/products/put/:id", isTokened(), isAdmin(), updateProduct)

	router.GET("/category", getCategories)
	router.GET("/category/:id", getCategoriesById)
	router.DELETE("/category/:id", isTokened(), isAdmin(), deleteCategory)
	router.POST("/category/create", isTokened(), isAdmin(), createCategory)
	router.PUT("/category/put/:id", isTokened(), isAdmin(), updateCategory)

	router.GET("/basket", isTokened(), getBasket)
	router.GET("/baskets", isTokened(), isAdmin(), getBaskets)

	router.POST("/login", login)
	router.POST("register", register)

	router.Run(":8000")

}

type loginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func login(c *gin.Context) {

	var userData loginForm
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	for _, el := range user {
		if el.Login == userData.Login && el.Password == userData.Password {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   el.ID,
				"name":     el.Name,
				"login":    el.Login,
				"password": el.Password,
				"role":     el.Role,
			})

			tokenStr, err := token.SignedString(jwtSecret)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"jwt error": err})
			}

			c.JSON(http.StatusOK, tokenStr)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func register(c *gin.Context) {
	var createUser User
	if err := c.BindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	user = append(user, createUser)
	c.JSON(http.StatusCreated, createUser)
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

// возвращает список всех корзин
func getBaskets(c *gin.Context) {
	c.JSON(http.StatusOK, baskets)
}

// type jsonBasket struct {
// 	ID           string   `json:"id"`
// 	ProductsName []string `productsName:"id"`
// }

func getBasket(c *gin.Context) {

	if uId, ok := c.Get("userId"); ok {
		for _, basket := range baskets {
			if basket.UserId == uId {
				c.JSON(http.StatusOK, basket)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "корзина не найдена"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "неверный jwt"})
	}
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
