# Apiance1

Loosly following the [crud restful api with go gorm jwt postgres and testing](https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121) tutorial.

## Developer Notes

Goconvey command line `go convey --host 0.0.0.0`.

Updating dependancies `go get -u ./...`

## Database

The database used is postgres, the connection for that is expected to be in an environment variable. For local testing, the `.env` will contain the developers credentials. An example version is `.env.example`.

Migrations will be automatically applied on startup of the application.

To undo all migrations, from the command use the following snippit, replacing the `{{password}}` with the password to your local database. If you need to do this in production, panic.

```bash
export POSTGRESQL_URL=postgres://postgres:{{password}}@localhost:5432/go_api?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path api/migrate/files down

```

### Generate a new migration

Migrations are created using the `migrate` tool. You generate the migration files, then need to edit the files to add the actual migration.

```bash
migrate create -ext sql -dir api/migrate/file _what-the-migration-does_
# ls db/migrations
# 20191119133115_{{what_the_migration_does}}.down.sql
# 20191119133115_{{what_the_migration_does}}.up.sql

```

Edit the files (this example adds basic user details):

```bash
# db/migrations/20191119110032_create_table.down.sql
DROP TABLE IF EXISTS users;
```

```bash
# db/migrations/20191119110032_create_table.up.sql
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(300),
    email VARCHAR(300) UNIQUE NOT NULL,
    password VARCHAR(500),
    updated_at BIGINT
);
```

#### Migrating index

There is an issue with making index changes in production: postgres will lock the tables while updating indexes.

To minimise the effect that this will have use a [concurrent](https://www.postgresql.org/docs/9.5/sql-createindex.html) index:

```sql
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_users_email_key
    ON users
    (email ASC);

```
