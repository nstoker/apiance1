CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_users_email_key
    ON users (lower(email));