# BANKING-APP

##### Prerequisite
- Go 
- Docker


```
make postgres // download and start Postgres Container
make createdb // Create DB
make dropdb // Drop DB
make migrateup // Migrate UP - Update DB with all tables & constraints 
make migratedown // Migrate DOWN - Deleted all tables  & constraints
make sqlc // Generate SQLC code for interacting with DB
make test // Run Test Cases
make server // Start the Server (PORT - 8080)
make mock // Mock DB & Test API

```