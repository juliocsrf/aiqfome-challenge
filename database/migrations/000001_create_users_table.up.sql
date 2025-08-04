CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

-- Password: admin
INSERT INTO users (id, name, email, password) VALUES ('01986709-c873-7525-bd98-20457930777c', 'Admin', 'admin@admin.com', '$2a$10$fLtpywS.uDkctCvp2oRk7.bpbh.obycMk3EWJU6toqx4A64j1nj6q');