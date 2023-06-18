# Bank-Transaction

An application to learn golang and focusing in databases

### Running the application

To run the application first mak sure to have the `make` cli tool installed.

Then just enter the following command:

```bash
make execute
```

### Migrations

To generate new migrations just hit the code below:
  
```shell
make create_migration file_name=example_schema
```

A new file should be created under `app/internal/db/migrations` with the name `000002_schema_migration_up.sql` and `000002_schema_migration_down.sql`.
After editing these two sql files just hit the code below to apply.

```shell
make migrate_up 
make migrate_down
```

### QUERIES

To generate new code for the queries located at `app/internal/db/queries`, first make sure that you have the sqlc tool installed.

```shell
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

After that just hit the code below:
  
```shell
make query
```

Files under `app/internal/db/sqlc` should be generated (DO NOT EDIT), they are like CRUD automatically generated from .sql files.
