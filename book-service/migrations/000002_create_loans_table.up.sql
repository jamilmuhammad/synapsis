CREATE TABLE IF NOT EXISTS loans (
    id SERIAL PRIMARY KEY,
    loan_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP DEFAULT NULL,
    book_id SERIAL,
    user_id SERIAL,
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

CREATE TRIGGER update_loan_modtime
BEFORE UPDATE ON loans
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

-- CREATE TRIGGER update_loan_modtime
-- BEFORE UPDATE ON loans
-- FOR EACH ROW
-- EXECUTE FUNCTION update_modified_column();
-- CREATE OR REPLACE FUNCTION validate_loan_count()
-- RETURNS TRIGGER AS $$
-- DECLARE
--   book_count INTEGER;
--   loan_count INTEGER;
-- BEGIN
--   -- Get the count of copies for the book
--   SELECT copies INTO book_count
--   FROM books
--   WHERE id = NEW.book_id;

--   -- Get the count of loans for the book
--   SELECT COUNT(*) INTO loan_count
--   FROM loans
--   WHERE book_id = NEW.book_id;

--   -- Check if the number of loans exceeds the number of copies
--   IF loan_count >= book_count THEN
--     RAISE EXCEPTION 'Loan count exceeds the number of copies for the book';
--   END IF;

--   RETURN NEW;
-- END;
-- $$ LANGUAGE 'plpgsql';

-- CREATE TRIGGER validate_loan_count_trigger
-- BEFORE INSERT ON loans
-- FOR EACH ROW
-- EXECUTE FUNCTION validate_loan_count();
