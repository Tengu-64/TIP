# 11.11 практика 10 - БД

Предмет: Технологии индустриального программирования

Выполнил студент: Куценко Игорь Олегович

Группа: ЭФМО-02-24
```
CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(50) NOT NULL UNIQUE,
       email VARCHAR(100) NOT NULL UNIQUE,
       password VARCHAR(255) NOT NULL,
   ); 

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    userId INT REFERENCES users(id) ON DELETE CASCADE);

CREATE TABLE sellers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE products (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       description TEXT,
       price DECIMAL(10, 2) NOT NULL,
       stock INT NOT NULL,
sellerId INT REFERENCES sellers(id) ON DELETE CASCADE);


CREATE TABLE cartItems (
    id SERIAL PRIMARY KEY,
    cartId INT REFERENCES carts(id) ON DELETE CASCADE,
    productId INT REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL);


CREATE TABLE reviews (
        id SERIAL PRIMARY KEY,
        productId INT REFERENCES products(id) ON DELETE CASCADE,
        userId INT REFERENCES users(id) ON DELETE CASCADE,
        rating INT CHECK (rating >= 1 AND rating <= 5),
        comment TEXT,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );


CREATE TABLE orders (
       id SERIAL PRIMARY KEY,
       userid INT REFERENCES users(user_id) ON DELETE CASCADE,
       orderDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       status VARCHAR(20) NOT NULL
   );

CREATE TABLE orderItems (
       id SERIAL PRIMARY KEY,
       orderIOd INT REFERENCES orders(id) ON DELETE CASCADE,
       productId INT REFERENCES products(id) ON DELETE CASCADE,
       quantity INT NOT NULL,
       priceAtPurchase DECIMAL(10, 2) NOT NULL,
       createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

CREATE TABLE orderCheck (
   id SERIAL PRIMARY KEY,
    orderId INT REFERENCES orders(id) ON DELETE CASCADE,
    userId INT REFERENCES users(id) ON DELETE CASCADE,
    totalAmount DECIMAL(10, 2) NOT NULL,
    issuedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL
);


CREATE TABLE categories (
       id SERIAL PRIMARY KEY,
       name VARCHAR(50) NOT NULL UNIQUE
   );

CREATE TABLE productCategories (
 id SERIAL PRIMARY KEY,
       productId INT REFERENCES products(id) ON DELETE CASCADE,
       categoryId INT REFERENCES categories(categoryId) ON DELETE CASCADE,
       PRIMARY KEY (product_id, category_id)
   );
```
