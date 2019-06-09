# Fullstack Challenge

## Applicant

James Crisman

## Technologies

- React
  - Bulma for CSS
  - Jest, React Testing Library
- Golang
- Postgresql
- Docker
- Postman/Newman

## [Notes about the design and more documentation](./DESIGN.md)

### Logging in

default seed admin credentials:
```
user: admin
password: asdfasdfasdf
```
default seed user credentials:
```
user: user
password: asdfasdfasdf
```

## How to run

## With Docker

With the current configuration it sets up and runs with postgres. It'll start with a fresh database on start up. It is running these in "DEV" mode.

```
# from the base directory
docker-compose up
```

And access via `localhost:3000` once the frontend is ready.

> You can run the API tests on the docker container as well, but you'll need to make port 8000 available for the backend

## Without docker

### Frontend

```
cd ./frontend
npm install
npm start
```

Run tests with
```
cd ./frontend
npm run test
```

### Postgres

You can run with or without postgres, but the default postgres configs expect a `paypay` named database. However you can define a different connection with the `POSTGRES_CONNECTION` environment variable.

For setting up default postgres db you may need to run

```
createdb paypay
```

[some documentation on setting up postgresql here](https://www.postgresql.org/docs/10/tutorial-createdb.html)

### Backend

The backend is built in Go. Ensure you're using a version of golang with modules (1.12+).

the `-resetPostgres -seedDatabase` flags aren't necessary if you've run them once for postgres. And neither are required if you don't want to use postgres.

> `-resetPostgres -seedDatabase` are only available with the `ENV=DEV` environment variable. This is to prevent you from accidentally deleting/inserting test data into a production environment. It's not open by default in the `docker-compose.yml`

```
# using postgres
cd ./backend
go build -o backend
./backend -usePostgres -resetPostgres -seedDatabase

# or

cd ./backend
go run ./backend.go -usePostgres -resetPostgres -seedDatabase
```

```
# in memory/map-based database
cd ./backend
go build -o backend
./backend

# or 

cd ./backend
go run ./backend.go
```

Run unit tests with
```
cd ./backend
go test ./...
```

### API Tests

I built out some api tests using [postman](https://www.getpostman.com/). They have a commandline tool to run API collections/tests with them. You can install the cli from npm.

*if you set the JWT_SECRET environment variable, you'll need to unset it for the script*

*The saved tokens for the API Requests will expire in two years*

```
npm install -g newman
```

Here's a script that will build the backend, reset the postgres database, and run the API test suite. It'll bring the server back into foreground after finishing the tests.

> You'll need environment variable `ENV=DEV` for the `-usePostgres` version of these scripts.

```
cd ./backend
go build ./backend.go && ./backend -usePostgres -resetPostgres -seedDatabase & sleep 1 && newman run ../paypay-api.postman_collection.json; fg
```

Here's the script for an in-memory database

```
cd ./backend
go build ./backend.go && ./backend -seedDatabase & sleep 1 && newman run ../paypay-api.postman_collection.json; fg
```