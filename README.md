### go_api
A RESTful API built with Go, Gin, MongoDB, and Docker

### Deployment Instructions
* **Option 1:** Run `docker-compose build` to build the images, then run `docker-compose up -d` to run the containers in the background
* **Option 2:** Run `docker-compose up --build` to force building the images each time, even if they already exist
* To stop the containers, run `docker-compose down`

### Usage
* To create a new book, submit a POST request to `localhost:8080/api/book/new` with the following body structure:
```
{
  "title": "[Book Title]",
  "author": "[Author]"
}
```
* To get all books, submit a GET request to `localhost:8080/api/book/all`
* To get a book by ID, submit a GET request to `localhost:8080/api/book/id/[ID of the book]`
* To delete a book by ID, submit a POST request to `localhost:8080/api/book/delete/[ID of the book]`