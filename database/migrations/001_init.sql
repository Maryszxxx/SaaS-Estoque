CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       role VARCHAR(50) NOT NULL CHECK (role IN ('ADMIN', 'EMPLOYEE')),
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL
);

CREATE TABLE products (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT NOT NULL,
                          price DOUBLE PRECISION NOT NULL,
                          created_at TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP NOT NULL,
                          category_id BIGINT NOT NULL,
                          quantity BIGINT NOT NULL DEFAULT 0
);