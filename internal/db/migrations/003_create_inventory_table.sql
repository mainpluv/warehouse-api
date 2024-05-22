CREATE TABLE inventory (
    warehouse_id INT REFERENCES warehouses(id),
    item_id INT REFERENCES items(id),
    quantity INT NOT NULL,
    PRIMARY KEY (warehouse_id, item_id)
);