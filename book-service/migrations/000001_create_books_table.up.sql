CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    copies INT DEFAULT 1,
    date_of_publication TIMESTAMP DEFAULT NULL,
    author_id SERIAL,
    category_id SERIAL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_book_modtime
BEFORE UPDATE ON books
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();
