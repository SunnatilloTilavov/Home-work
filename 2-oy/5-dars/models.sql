CREATE TABLE IF NOT EXISTS Category(
    id uuid PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS Product(
    id uuid PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    price INTEGER UNIQUE,
    category_id uuid References Category(id),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);

Insert into product(id,name,category_id,price) 
values('fa82a9f1-0826-4c75-b5df-a7b6176286c7','qwasd','87c01bbe-ecfa-4dfc-92fa-e890e6f3d1e1',123123),


