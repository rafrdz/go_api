### go_api
Building a RESTful API with Go, Gin, MongoDB, and Docker

### Deployment Instructions
* Run `docker-compose build` to build the images
* Run `docker-compose up -d` to run the containers in the background

### Usage
* To create a new book, submit a POST request to `localhost:8080/api/book/new` with the following body structure:
```
{
  "title": "[Book Title]",
  "author": "[Author]"
}
```
* To get all books, submit a GET request to `localhost:8080/api/book/all`