CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(300),
    email VARCHAR(320) NOT NULL,
    password VARCHAR(500),
    created_at TIMESTAMPTZ default current_timestamp,
    updated_at TIMESTAMPTZ default current_timestamp
);
CREATE TRIGGER update_user_updated_at 
    BEFORE UPDATE ON users FOR EACH ROW 
        EXECUTE PROCEDURE update_updated_at_column();
