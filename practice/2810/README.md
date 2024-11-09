# 28.10 практика 8 - REST API на Go

Предмет: Технологии индустриального программирования

Выполнил студент: Куценко Игорь Олегович

Группа: ЭФМО-02-24

# Маршруты:

/products - get, получение всех товаров           
/products/:id - get, получение товара по id <br/>
/products/delete/:id" - delete, удаление товара по id <br/>
/products/create - post, создание нового товара <br/>
/products/put/:id - put, изменение товара по id <br/>

/category - get, все категории <br/>
/category/:id" - get, получение категории по id <br/>
/category/:id - delete, удаление категории по id  <br/>
/category/create - post, создание нового категории <br/>
/category/put/:id" - put, изменение категории по id <br/>

/baskets - get, все корзины <br/>
/baskets/:id - получение корзины по id <br/>

# Результат работы Postman:

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/products.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/getProductsId.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/productsCreate.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/productsPut.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/productsDelete.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/category.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/categoryById.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/categoryCreate.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/categoryPut.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/categoryDelete.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/basketsGet.png)

![Image alt](https://github.com/Tengu-64/TIP/blob/main/practice/2810/postman/basketbyId.png)

