-- Insert initial data into warehouses table
INSERT INTO warehouses (name, is_available) VALUES
('Warehouse 1', TRUE),
('Warehouse 2', TRUE),
('Warehouse 3', FALSE);

-- Insert initial data into products table
INSERT INTO products (name, size, code, quantity, warehouse_id) VALUES
('Product 1', 'Small', 'P001', 100, 1),
('Product 2', 'Medium', 'P002', 200, 1),
('Product 3', 'Large', 'P003', 150, 2),
('Product 4', 'Small', 'P004', 250, 2),
('Product 5', 'Medium', 'P005', 300, 3);
