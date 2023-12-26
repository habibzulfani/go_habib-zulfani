CREATE DATABASE alta_online_shop1;

USE alta_online_shop1;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    status SMALLINT,
    birthdate DATE,
    gender CHAR(1),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(100),
    product_code VARCHAR(50),
    product_typeID INT,
    product_descriptionID INT UNIQUE,
    operatorID INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product_types (
    id INT PRIMARY KEY AUTO_INCREMENT,
    type_name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE operators (
    id INT PRIMARY KEY AUTO_INCREMENT,
    operator_name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product_descriptions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    description LONGTEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE payment_methods (
    id INT PRIMARY KEY AUTO_INCREMENT,
    method_name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    payment_methodsID INT,
    userID INT,
    status VARCHAR(10),
    total_qty INT,
    total_price NUMERIC(25,2),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE transaction_details (
    transactionID INT,
    productID INT,
    PRIMARY KEY(transactionID, productID),
    status VARCHAR(10),
    qty INT,
    price NUMERIC(25,2),
    created_at DATETIME,
    updated_at DATETIME
);

    ALTER TABLE products
    ADD CONSTRAINT fk_product_type
    FOREIGN KEY (product_typeID)
    REFERENCES product_types(id),
    ADD CONSTRAINT fk_product_desc
    FOREIGN KEY (product_descriptionID)
    REFERENCES product_descriptions(id),
    ADD CONSTRAINT fk_operator
    FOREIGN KEY (operatorID)
    REFERENCES operators(id);
   
    ALTER TABLE transactions
    ADD CONSTRAINT fk_payment_methods
    FOREIGN KEY (payment_methodsID)
    REFERENCES payment_methods(id),
    ADD CONSTRAINT fk_user
    FOREIGN KEY (userID)
    REFERENCES users(id);

    ALTER TABLE transaction_details
    ADD CONSTRAINT fk_transaction
    FOREIGN KEY (transactionID)
    REFERENCES transactions(id),
    ADD CONSTRAINT fk_product
    FOREIGN KEY (productID)
    REFERENCES products(id);

INSERT INTO operators (operator_name, created_at, updated_at)
VALUES 
  ('JohnDoe', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  ('AliceSmith', '2023-07-31 12:30:00', '2023-07-31 12:30:00'),
  ('AlissaSmith', '2023-07-31 12:30:00', '2023-07-31 12:40:00'),
  ('AlexSmith', '2023-07-31 12:30:00', '2023-07-31 12:50:00'),
  ('AdamSmith', '2023-07-31 12:30:00', '2023-07-31 12:55:00');

SELECT * FROM operators;CREATE DATABASE alta_online_shop1;

USE alta_online_shop1;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    status SMALLINT,
    birthdate DATE,
    gender CHAR(1),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(100),
    product_code VARCHAR(50),
    product_typeID INT,
    product_descriptionID INT UNIQUE,
    operatorID INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product_types (
    id INT PRIMARY KEY AUTO_INCREMENT,
    type_name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE operators (
    id INT PRIMARY KEY AUTO_INCREMENT,
    operator_name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product_descriptions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    description LONGTEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE payment_methods (
    id INT PRIMARY KEY AUTO_INCREMENT,
    method_name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    payment_methodsID INT,
    userID INT,
    status VARCHAR(10),
    total_qty INT,
    total_price NUMERIC(25,2),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE transaction_details (
    transactionID INT,
    productID INT,
    PRIMARY KEY(transactionID, productID),
    status VARCHAR(10),
    qty INT,
    price NUMERIC(25,2),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

    ALTER TABLE products
    ADD CONSTRAINT fk_product_type
    FOREIGN KEY (product_typeID)
    REFERENCES product_types(id),
    ADD CONSTRAINT fk_product_desc
    FOREIGN KEY (product_descriptionID)
    REFERENCES product_descriptions(id),
    ADD CONSTRAINT fk_operator
    FOREIGN KEY (operatorID)
    REFERENCES operators(id);
   
    ALTER TABLE transactions
    ADD CONSTRAINT fk_payment_methods
    FOREIGN KEY (payment_methodsID)
    REFERENCES payment_methods(id),
    ADD CONSTRAINT fk_user
    FOREIGN KEY (userID)
    REFERENCES users(id);

    ALTER TABLE transaction_details
    ADD CONSTRAINT fk_transaction
    FOREIGN KEY (transactionID)
    REFERENCES transactions(id),
    ADD CONSTRAINT fk_product
    FOREIGN KEY (productID)
    REFERENCES products(id);

INSERT INTO operators (operator_name, created_at, updated_at)
VALUES 
  ('JohnDoe', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  ('AliceSmith', '2023-07-31 12:30:00', '2023-07-31 12:30:00'),
  ('AlissaSmith', '2023-07-31 12:30:00', '2023-07-31 12:40:00'),
  ('AlexSmith', '2023-07-31 12:30:00', '2023-07-31 12:50:00'),
  ('AdamSmith', '2023-07-31 12:30:00', '2023-07-31 12:55:00');

SELECT * FROM operators;

INSERT INTO product_types  (type_name, created_at, updated_at)
VALUES 
  ('Toy', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  ('Beverage', '2023-07-31 12:30:00', '2023-07-31 12:30:00'),
  ('Fashion', '2023-07-31 12:30:00', '2023-07-31 12:40:00');

SELECT * FROM product_types;

INSERT INTO product_descriptions (description, created_at, updated_at)
VALUES 
  	('Description for Product 3', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  	('Description for Product 4', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
 	('Description for Product 5', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
	('Description for Product 6', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
	('Description for Product 7', '2023-07-31 12:00:00', '2023-07-31 12:00:00'),
	('Description for Product 8', '2023-07-31 12:00:00', '2023-07-31 12:00:00');
 
 SELECT * FROM product_descriptions;

INSERT INTO products  (product_name, product_code, product_typeID, product_descriptionID, operatorID, created_at, updated_at)
VALUES 
  ('Products1', 'code1', 1, 1, 3,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  ('Products2', 'code2', 1, 2, 3,'2023-07-31 12:00:00', '2023-07-31 12:00:00');

SELECT * FROM products;

INSERT INTO products  (product_name, product_code, product_typeID, product_descriptionID, operatorID, created_at, updated_at)
VALUES 
  	('Products3', 'code3', 2, 3, 1,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  	('Products4', 'code4', 2, 4, 1,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
 	('Products5', 'code5', 2, 5, 1,'2023-07-31 12:00:00', '2023-07-31 12:00:00');

SELECT * FROM products;

INSERT INTO products  (product_name, product_code, product_typeID, product_descriptionID, operatorID, created_at, updated_at)
VALUES 
  	('Products6', 'code6', 3, 6, 4,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  	('Products7', 'code7', 3, 7, 4,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
 	('Products8', 'code8', 3, 8, 4,'2023-07-31 12:00:00', '2023-07-31 12:00:00');

SELECT * FROM products;

DELETE FROM users;

ALTER TABLE products AUTO_INCREMENT = 1;

INSERT INTO payment_methods  (method_name, status, created_at, updated_at)
VALUES 
  	('Method1', 1,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
  	('Method2', 0,'2023-07-31 12:00:00', '2023-07-31 12:00:00'),
 	('Method3', 1,'2023-07-31 12:00:00', '2023-07-31 12:00:00');

SELECT * FROM payment_methods;

INSERT INTO users  (name, status, birthdate, gender, created_at, updated_at)
VALUES 
  	('Adam', 1, '2023-11-18', 'm', '2023-11-14 12:00:00', NOW()),
  	('Jane', 1, '2023-11-18', 'f', '2023-11-12 12:00:00', NOW()),
  	('John', 0, '2023-11-18', 'm', '2023-11-16 12:00:00', NOW()),
  	('Alex', 1, '2023-11-18', 'f', NOW(), NOW()),
  	('Jack', 1, '2023-11-18', 'm', NOW(), NOW());
  
SELECT * FROM users;

INSERT INTO transactions  (payment_methodsID, userID, status, total_qty, total_price, created_at, updated_at)
VALUES 
  	(1, 1, 'Pending', 3, 15000.75, '2023-08-01 10:00:00', '2023-08-01 10:00:00'),
    (2, 1, 'Completed', 5, 25000.50, '2023-08-02 12:30:00', '2023-08-02 12:30:00'),
    (3, 1, 'Failed', 1, 30000.00, '2023-08-03 15:45:00', '2023-08-03 15:45:00'),
    (2, 2, 'Completed', 30, 15000.75, '2023-08-01 10:00:00', '2023-08-01 10:00:00'),
    (1, 2, 'Completed', 10, 25000.50, '2023-08-02 12:30:00', '2023-08-02 12:30:00'),
    (3, 2, 'Failed', 15, 30000.00, '2023-08-03 15:45:00', '2023-08-03 15:45:00'),
    (3, 3, 'Pending', 3, 15000.75, '2023-08-01 10:00:00', '2023-08-01 10:00:00'),
    (2, 3, 'Completed', 5, 25000.50, '2023-08-02 12:30:00', '2023-08-02 12:30:00'),
    (1, 3, 'Pending', 1, 30000.00, '2023-08-03 15:45:00', '2023-08-03 15:45:00'),
    (1, 4, 'Completed', 30, 15000.75, '2023-08-01 10:00:00', '2023-08-01 10:00:00'),
    (3, 4, 'Completed', 50, 25000.50, '2023-08-02 12:30:00', '2023-08-02 12:30:00'),
    (2, 4, 'Completed', 2, 30000.00, '2023-08-03 15:45:00', '2023-08-03 15:45:00'),
    (1, 5, 'Pending', 3, 15000.75, '2023-08-01 10:00:00', '2023-08-01 10:00:00'),
    (2, 5, 'Completed', 5, 25000.50, '2023-08-02 12:30:00', '2023-08-02 12:30:00'),
    (3, 5, 'Failed', 1, 30000.00, '2023-08-03 15:45:00', '2023-08-03 15:45:00');
  
SELECT * FROM transactions;

INSERT INTO transaction_details  (transactionID, productID, status, qty, price, created_at, updated_at)
VALUES 
  	(1, 1, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(1, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (1, 3, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (2, 3, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(2, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (2, 9, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (3, 1, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(3, 3, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (3, 2, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (4, 1, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(4, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (4, 3, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (5, 3, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(5, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (5, 1, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (6, 1, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(6, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (6, 3, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (7, 4, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(7, 5, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (7, 10, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (8, 9, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(8, 11, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (8, 1, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (9, 2, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(9, 4, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (9, 3, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (10, 5, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(10, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (10, 9, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (11, 11, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(11, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (11, 3, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (12, 1, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(12, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (12, 5, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (13, 4, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(13, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (13, 1, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (14, 1, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(14, 4, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (14, 2, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00'),
    (15, 9, 'Shipped', 2, 5000.00, '2023-08-01 10:15:00', '2023-08-01 10:15:00'),
 	(15, 2, 'Processing', 1, 3000.00, '2023-08-01 10:30:00', '2023-08-01 10:30:00'),
    (15, 11, 'Delivered', 3, 12000.75, '2023-08-02 13:00:00', '2023-08-02 13:00:00');
  
SELECT * FROM transaction_details;

SELECT * FROM  users WHERE gender = 'm';

SELECT * FROM products WHERE id = 3;

SELECT * FROM users WHERE created_at BETWEEN CURDATE() - INTERVAL 7 DAY AND CURDATE() AND name LIKE '%e%';

SELECT COUNT(gender) FROM users WHERE gender = 'f'; 

SELECT * FROM users ORDER BY name ASC; 

SELECT * FROM products LIMIT 5;

UPDATE products SET product_name='product dummy' WHERE id =1;

SELECT * FROM products WHERE id =1;

UPDATE transaction_details  SET qty=3 WHERE productID =1;

SELECT * FROM transaction_details td WHERE productID =1;

ALTER TABLE transaction_details
DROP FOREIGN KEY fk_product;

DELETE FROM products WHERE id =1;

DELETE FROM products WHERE product_typeID =1;

SELECT * FROM products;

ALTER TABLE transaction_details
ADD CONSTRAINT fk_product
FOREIGN KEY (productID)
REFERENCES products(id);

SELECT * FROM transaction_details;

INSERT INTO products  (id, product_name, product_code, product_typeID, product_descriptionID, operatorID, created_at, updated_at)
VALUES 
  (2, 'Products2', 'code1', 1, 2, 3,NOW(), NOW());
  
 SELECT * FROM products;
 
SELECT * FROM transactions WHERE userID =1 UNION SELECT * FROM transactions WHERE userID =2;

SELECT SUM(total_price) FROM transactions WHERE userID =1;

SELECT
    product_types.type_name AS product_type,
    COUNT(transactions.id) AS total_transactions,
    SUM(transactions.total_price) AS total_transaction_amount
FROM
    transaction_details
JOIN products ON transaction_details.productID = products.id
JOIN product_types ON products.product_typeID = product_types.id
JOIN transactions ON transaction_details.transactionID = transactions.id
WHERE
    products.product_typeID = 2
GROUP BY
    product_types.type_name;
