CREATE TABLE IF NOT EXISTS warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(50),
    code VARCHAR(100) UNIQUE NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    warehouse_id INT REFERENCES warehouses(id)
);

CREATE TABLE IF NOT EXISTS reserved (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(50),
    code VARCHAR(100) UNIQUE NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    warehouse_id INT
);

