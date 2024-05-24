-- добавляем какие-то начальные записи
INSERT INTO warehouses (id, name, is_available) VALUES
(1, 'Warehouse 1', TRUE),
(2, 'Warehouse 2', TRUE),
(3, 'Warehouse 3', FALSE);

INSERT INTO products (name, size, code, quantity, warehouse_id) VALUES
('Product 1', 'Small', 'P001', 100, 1),
('Product 2', 'Medium', 'P002', 200, 1),
('Product 3', 'Large', 'P003', 150, 2),
('Product 4', 'Small', 'P004', 250, 2),
('Product 5', 'Medium', 'P005', 300, 3);

INSERT INTO reserved (name, size, code, quantity, warehouse_id) VALUES
('Product 1', 'Small', 'P001', 5, 1),
('Product 3', 'Large', 'P003', 1, 2),
('Product 5', 'Medium', 'P005', 3, 3);