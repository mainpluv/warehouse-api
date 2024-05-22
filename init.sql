-- init.sql

CREATE TABLE IF NOT EXISTS warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    availability BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(50) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    quantity INT NOT NULL,
    warehouse_id INT NOT NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);