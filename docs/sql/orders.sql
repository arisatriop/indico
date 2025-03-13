CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number varchar(50),
    order_type varchar(50),
    product_id INT not null,
    product_name varchar(50) not null,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50), 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(50), 
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50), 
    
    CONSTRAINT fk_orders_product FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE INDEX idx_orders_order_number ON orders(order_number);