# Auth with Clean Architecture
This mini project describes about authentication, but I will implement clean architecture as well. Here are the guidelines you to follow:
* [Setup](#scenario)
* [How to Run](#how-to-run)
* [Postman Collection](#postman-collection)
* [Tech Spec](#tech-spec)

## Setup
1. Install `Go` and `MySQL`.
    * [`Go`](https://go.dev/doc/install)
    * [`MySQL`](https://dev.mysql.com/doc/)
1. Cloning this project out with the following command:
    * Using SSH URL
    ```bash
    $ git clone git@github.com:pejuangreceh/mini-project-go-2
    ```
    * Using HTTPS URL
    ```bash
    $ git clone https://github.com/pejuangreceh/mini-project-go-2.git
    ```
1. Do the command to install the dependencies:
    ```bash
    $ go mod download
    ```
1. Set up the database configuration at `initDB()`. Here is the default example:
    ```text
func initDB() (*gorm.DB, error) {
	//default local
	dsn := "root@tcp(localhost:3306)/db_project?parseTime=true"
	//production
	//dsn := "root:@tcp(db:3306)/app?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
    ```


## How to Run
Use this command to run API app from root directory:
```bash
go run main.go
```

## Postman Collection
You can get postman collection for Admin [here](https://api.postman.com/collections/17613454-1b8485a7-4d94-4dea-bf3d-e0bc56204147?access_key=PMAT-01H4Z0CWGYYQKRDY6XN9HAYF3K).
You can get postman collection for Customer [here](https://api.postman.com/collections/17613454-b815da22-794f-4135-8f0d-57dafdcf52b4?access_key=PMAT-01H4Z0EC633YZA1DTN6D08T2F2).


## Tech Spec
### Feature Description
there are three endpoint that could use for appointment process, which are:
1. Login and show profile
1. Create, read, update and delete User
1. Create, read, update and delete Customer

### Architecture and Design
This service using onion architecture, there are 5 layers from inner to outer which are entity, repository, use case, controller, and request handler. the usage and responsibility of each layer are follow:
1. Entity: this layer contains the domain model or entities of the system. These are the core objects that represent the business concepts and rules.
1. Repository: This layer provides an interface for the application to access and manipulate the entities. It encapsulates the data access logic and provides a way to abstract the database implementation details.
1. Use case : This layer contains the business logic or use cases of the system. It defines the operations that can be performed on the entities and orchestrates the interactions between the entities and the repository layer.
1. Controller: This layer handles the HTTP requests and responses. It maps the incoming requests to the appropriate use case and returns the response to the client.
1. Request handler: This layer is responsible for handling the incoming HTTP requests and passing them on to the controller layer.

### Data Flow Diagram
There are 5 table that dedicated on this services which are users, customers, roles, permissions and role_permission. For complete explanation of database structure you can follow this ddf.

![data flow diagram](./docs/)
