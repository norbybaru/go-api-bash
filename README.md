# Go Api - Playground

## Prerequisite
- Install docker on host machine

## Getting Started with Docker

```bash
# Create a .env file by copying example values
cp .env.example .env

# start Docker service (Postgres, Redis, Application with hot reload)
make dc-up
```

## Getting Started without Docker
```bash
# Create a .env file by copying example values. Make sure to replace Postgres and Redis with correct credentials
cp .env.example .env

# Install dependencies
go mod download

# Build binary and start application
make start
```


## Help command
```bash
make help
```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `POST /api/v1/login`: Authenticates a user and generates a JWT
* `POST /api/v1/register`: Register a user
* `GET /api/v1/dishes`: Returns a paginated list of the dishes
* `GET /api/v1/dishes/:id`: Returns the detailed information of a dish
* `POST /api/v1/dishes`: Creates a new dish
* `PUT /api/v1/dishes/:id`: Updates an existing dish
* `DELETE /api/v1/dishes/:id`: deletes a dish
* `POST /api/v1/ratings`: Add a rating on a dish

Try the URL `http://localhost:8080/` in a browser, and you should see something like `"OK"` displayed.

To read through API docs visit page `http://localhost:8080/docs/` in your browser.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following
more complex scenarios:

```shell
# Register new user via: POST /api/v1/register
curl --location '127.0.0.1:8080/api/v1/auth/register' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Norby",
    "password": "secret",
    "nickname": "Baru",
    "email": "norby@example.com"
}'
# Should return user information created in JSON format

# authenticate the user via: POST /api/v1/login
curl --location '127.0.0.1:8080/api/v1/auth/login' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "norby@example.com",
    "password": "secret"
}'
# should return a JWT token like: {"access_token":"...JWT token here..."}

# with the above ACCESS token, access the dish resources, such as: GET /api/v1/dishes
curl --location '127.0.0.1:8080/api/v1/dishes' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer ${ACCESS_TOKEN}'
# should return a list of dishes in the JSON format
```
