# BANKING-APP

### Prerequisite
- Go Lang
- Docker
- Migrate
```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo apt-get update
sudo apt-get install migrate
```
- Sqlc
```
sudo snap install sqlc
```

### Setup infrastructure

Start postgres container:
```
make postgres
```

Create simple_bank database:
```
make createdb           
```

Run db migration up all versions:
```
make migrateup
```

Run db migration down all versions:
```
make migratedown
```

### How to generate code

Generate SQL CRUD with sqlc:
```
make sqlc
```

Generate DB mock with gomock:
```
make mock
```

### How to run

Run server:
```
make server
```

Run test:
```
make test
```