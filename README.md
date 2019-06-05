# Fullstack Challenge

## Applicant

James Crisman

## Technologies

- React (leveraging create-react-app)
  - Bulma for (S)CSS
- Golang
- Postgresql

## How to run

### Frontend

```
cd ./frontend
npm install
npm start
```

### Backend

Need postgresql running for the database. Currently expects default for the connection.

For setting up postgres db you'll need to run

```
createdb paypay
```

[if you have issues with setting up the postgres db](https://www.postgresql.org/docs/10/tutorial-createdb.html)

the `-resetDatabase -seedDatabase` flags aren't necessary if you've run them once.

```
cd ./backend
go build -o backend
./backend -resetDatabase -seedDatabase
```

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