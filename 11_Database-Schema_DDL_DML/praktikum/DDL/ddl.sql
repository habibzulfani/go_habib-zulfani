CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE users (
    userID INT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255), 
    birthdate DATE,
    status INT,
    gender ENUM('MALE', 'FEMALE'),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE product (
    productID INT PRIMARY KEY,
    product_name VARCHAR(255),
    product_typeID INT,
    product_descriptionID INT UNIQUE,
    operatorID INT,
    created_at DATETIME,
    updated_at DATETIME,
);

    ALTER TABLE product
    ADD CONSTRAINT fk_product_type
    FOREIGN KEY (product_typeID)
    REFERENCES product_type(product_typeID),
    ADD CONSTRAINT fk_product_desc
    FOREIGN KEY (product_descriptionID)
    REFERENCES product_description(product_descriptionID),
    ADD CONSTRAINT fk_operator
    FOREIGN KEY (operatorID)
    REFERENCES operator(operatorID);

CREATE TABLE product_type (
    product_typeID INT PRIMARY KEY,
    type_name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE operator (
    operatorID INT PRIMARY KEY,
    operator_name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE product_description (
    product_descriptionID INT PRIMARY KEY,
    description LONGTEXT,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE payment_methods (
    payment_methodsID INT PRIMARY KEY,
    method_name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE transaction (
    transactionID INT PRIMARY KEY,
    payment_methodsID INT,
    userID INT,
    transaction_detailsID INT UNIQUE,
    productID INT,
    method_name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
);

CREATE TABLE transaction_details (
    transaction_detailsID INT PRIMARY KEY,
    total_amount INT,
    total_price INT,
    created_at DATETIME,
    updated_at DATETIME
);

    ALTER TABLE transaction
    ADD CONSTRAINT fk_payment_methods
    FOREIGN KEY (payment_methodsID)
    REFERENCES payment_methods(payment_methodsID),
    ADD CONSTRAINT fk_user
    FOREIGN KEY (userID)
    REFERENCES user(userID),
    ADD CONSTRAINT fk_transaction_details
    FOREIGN KEY (transaction_detailsID)
    REFERENCES transaction_details(transaction_detailsID),
    ADD CONSTRAINT fk_product
    FOREIGN KEY (productID)
    REFERENCES product(productID);

CREATE TABLE kurir (
    kurirID INT PRIMARY KEY,
    name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

ALTER TABLE kurir 
ADD COLUMN ongkos_dasar INT;

RENAME TABLE kurir TO shipping;

DROP TABLE shipping;

ALTER TABLE payment_methods
ADD payment_methods_descriptionID INT UNIQUE;

CREATE TABLE payment_methods_description (
    payment_methods_descriptionID INT PRIMARY KEY,
    description LONGTEXT,
    created_at DATETIME,
    updated_at DATETIME
);

    ALTER TABLE payment_methods
    ADD CONSTRAINT fk_payment_methods_desc
    FOREIGN KEY (payment_methods_descriptionID)
    REFERENCES payment_methods_description(payment_methods_descriptionID);

ALTER TABLE user
ADD addressID INT;

CREATE TABLE address (
    addressID INT PRIMARY KEY AUTO_INCREMENT,
    street_address VARCHAR(255),
    city VARCHAR(100),
    state_province VARCHAR(100),
    postal_code VARCHAR(20),
    country VARCHAR(100)
);

    ALTER TABLE user
    ADD CONSTRAINT fk_address
    FOREIGN KEY (addressID)
    REFERENCES address(addressID);

CREATE TABLE user_payment_method_detail (
    userID INT,
    payment_methodsID INT,
    PRIMARY KEY (userID, payment_methodsID),
    FOREIGN KEY (userID) REFERENCES user(userID),
    FOREIGN KEY (payment_methodsID) REFERENCES payment_methods(payment_methodsID)
);
