# Dancing Pony

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

### Objective

Your assignment is to implement a restaurant REST API using Go and any framework.

### Brief

Frogo Baggins, a hobbit from the Shire, has a great idea. He wants to build a restaurant that serves traditional dishes from the world of Middle Earth. The restaurant will be called "The Dancing Pony" and will have a cozy atmosphere.

Frogo has hired you to build the website for his restaurant. As payment, he has offered you either a chest of gold or a ring. Choose wisely.

### Tasks

-   Implement assignment using:
    -   Language: **Go**
    -   Framework: **any framework**
-   Implement a REST API returning JSON
-   Implement a custom user model with a "nickname" field
-   Implement a dish model. Each dish should have a name, description, image, and price.
    -   Choose the appropriate data type for each field
    -   Add validation to the dish model fields to ensure that the name and description fields are unique
-   Provide an endpoint to authenticate with the API using username, password and return a JWT
-   Provide REST resources for the authenticated user for the Dish resource
    -   Implement the following CRUD (Create, Read, Update, Delete) operations for this resource:
        -   **Create**: Allow authenticated users to create new dishes
        -   **Read**: Allow authenticated users to view details of specific dishes, as well as a list of dishes
            -   Make the List resource searchable with query parameters
            -   Implement pagination for the /dishes resource. Allow users to set a limit and offset for the number of dishes returned in the response
        -   **Update**: Allow authenticated users to update dishes
        -   **Delete**: Allow authenticated users to delete dishes
    -   Implement an endpoint to rate a dish (POST)
        -   Store the rating on the Dish model
    -   Implement rate limiting for the /dishes resource to prevent abuse

### Evaluation Criteria

-   **Go** best practices
-   Make sure that users can only rate dishes once
-   Bonus: Make sure the user _Sméagol_ is unable to rate any dish at "The Dancing Pony"
-   Bonus: Write an API test for the rating endpoint

### CodeSubmit

Please organize, design, test, and document your code as if it were going into production - then push your changes to the master branch. After you have pushed your code, you may submit the assignment on the assignment page.

Best of luck, and happy coding!

The TFG Labs Team
