CREATE TABLE books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR(75) NOT NULL ,
    publisher VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);