-- +migrate Up
-- +migrate StatementBegin
DROP TABLE IF EXISTS customer;
DROP TABLE IF EXISTS restaurant;
DROP TABLE IF EXISTS food;
DROP TABLE IF EXISTS orders;

CREATE TABLE food (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE customer (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    order_id INT,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE restaurant (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    order_id INT,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE orders (
  id BIGSERIAL PRIMARY KEY,
  cust_id INT,
  resto_id INT,
  status VARCHAR(255),
  food JSON NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_cust_id 
  FOREIGN KEY (cust_id) REFERENCES customer(id),
  
  CONSTRAINT fk_resto_id 
  FOREIGN KEY (resto_id) REFERENCES restaurant(id)
);

-- +migrate StatementEnd