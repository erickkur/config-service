# config-service
### How to run migration
1. Create new file `local-env.sh` from `example-env.sh`
2. Please install golang-migrate, curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz
3. Adjust all env in accordance with local development
4. Run `make migrate`

### How to run unit test
1. Run `make test`

### How to run this application
1. Run `make dep`
2. Run `make migrate`
3. Run `make run`