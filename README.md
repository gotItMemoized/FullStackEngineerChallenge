# Fullstack Challenge

## Applicant

James Crisman

## Technologies

- React (leveraging create-react-app)
  - Bulma for (S)CSS
- Golang
- Postgresql
- Docker

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

With the current configuration it sets up and runs with postgres. It'll start with a fresh database on start up. It is running these in "DEV" mode. So you may see dev-only console warnings from things like react-router when you click on the link of a page you're already on. Those warnings aren't there in a production build.

```
# from the base directory
docker-compose up
```

## Without docker

### Frontend

```
cd ./frontend
npm install
npm start
```

### Backend

You can run with or without postgres, but the default configs expect a `paypay` named database. However you can define a different connection with the `POSTGRES_CONNECTION` environment variable.

For setting up postgres db you may need to run

```
createdb paypay
```

[if you have issues with setting up the postgres db](https://www.postgresql.org/docs/10/tutorial-createdb.html)

the `-resetPostgres -seedDatabase` flags aren't necessary if you've run them once for postgres. And neither are required if you don't want to use postgres.

```
# using postgres
cd ./backend
go build -o backend
./backend -usePostgres -resetPostgres -seedDatabase
```

```
# in memory/mapbased database
cd ./backend
go build -o backend
./backend
```

### API Tests

I built out some api tests using [postman](https://www.getpostman.com/). They have a commandline tool to run API collections/tests with them. You can install the tool from npm.

*if you set the JWT_SECRET environment variable, you'll need to unset it for the script*

*The saved tokens for the API Requests will expire in two years*

```
npm install -g newman
```

Here's a script that will build the backend, reset the database, and run the API test suite. It'll bring the server back into foreground after finishing the tests.

```
cd ./backend
go build ./backend.go && ./backend -usePostgres -resetPostgres -seedDatabase & sleep 1 && newman run ../paypay-api.postman_collection.json; fg
```

Here's the script for an in-memory database

```
cd ./backend
go build ./backend.go && ./backend -seedDatabase & sleep 1 && newman run ../paypay-api.postman_collection.json; fg
```