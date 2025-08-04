CREATE TABLE favorites (
    customer_id UUID NOT NULL,
    product_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (customer_id, product_id),
    
    CONSTRAINT fk_favorites_customer 
        FOREIGN KEY (customer_id) 
        REFERENCES customers(id) 
        ON DELETE CASCADE
);

CREATE INDEX idx_favorites_product_id ON favorites(product_id);