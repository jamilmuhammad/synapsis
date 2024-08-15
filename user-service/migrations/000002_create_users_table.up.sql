CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) CHECK (role IN ('Member', 'Librarian', 'Admin')),
    status VARCHAR(50) CHECK (role IN ('Pending', 'Verified', 'Rejected')),
    deleted_at TIMESTAMPTZ DEFAULT NULL AT TIME ZONE 'UTC',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP AT TIME ZONE 'UTC',
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
);

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC';
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_modtime
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();
