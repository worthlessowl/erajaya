# Erajaya technical test

## Running the Application

### Requirement
You need to setup docker to run this webserver. 

### Executing the server
Go to the root directory (with `docker-compose.yml` file and execute the following command

`docker compose build --no-cache` (build)

`docker compose up` (run)

The server will be running on port `:8080`

### Cleaning up the server
On the root directory

`docker compose down` (stop)

`sudo rm -r tmp` (remove temp files)

`docker volume rm $(docker volume ls -q)` (remove volumes *CAREFUL THIS WILL DELETE ALL VOLUME*)

### Troubleshooting

- Docker failed to start postgre, redis, or webserver
 
 Make sure the port `:8080`, `:5433` and `:6379` is not currently used. You can modify `docker-compose.yml` if you want to change the host port.
 
- Host not found when docker is running the webserver

Try running `docker compose up` again if you encounter this issue. It seems like when docker is initiating, 
the db shuts down once and gets assigned new IP when restarting. Which cause the webserver to not be able to find the DB IP.

## Using the server

### Endpoints

#### Get all Products
`/`

any http method

Gets all products. 10 seconds cache is implemented. If you don't see the value changing, wait 10 seconds for the cache to expire.

Optional parameter:
`sortby`

with possible values:
- `newest`
- `cheapest`
- `mostexpensive`
- `nameascending`
- `namedescending`

example: `localhost:8080?sortby=newest`

#### Insert product
`/insert`

any http method

Insert product based on request body

Body:

content type: `application/json`

```
{
  "name": string, (required)
  "price: number, (required)
  "description": string, (required)
  "quantity": number (required)
}
```

example:
```
{
  "name": "susu anak",
  "price": 100000,
  "description": "susu anak sehat 100 gram",
  "quantity": 123
}
```
Not using all the required param in the body will cause the request to fail.

## Architecture

The general architecture of this service is there are 3 server running, a Golang webserver, a Postgres DB and a Redis Cache. The execution
of these 3 servers is managed using docker compose.

I split the Golang webserver into several layers.
- Main
- Server Layer
- Usecase Layer
- Resource Layer

Main Layer is the `main.go` file, this layer will deal with all the main initialization, is the main function caller of other layers,
and registers the handler for this server.

Server layer contains the handler logic for this webserver. This layer processes the parameter passed by the main layer, calls the usecase layer 
to call the logic it wants to execute, and returns the http response.

Usecacse layer is the layer where the main logic of this application resides. We can combine multiple resources in this layer.

Resource layer is the lowest layer of this app. The `db` and `cache` package is included in the resource layer. This layer handles the 
logic of interacting with external resources such as database and cache. We format the response from the external resources into something our 
app can handle before passing it back to the Usecase layer.

The reason behind this layer splitting is so that I can isolate the responsibilities of each layers, to ease the current and future development.
By separating the responsibilities of each layer, we don't end up with a massive file containing all the logics needed for a feature. 
By separating it into layers we also improves modularity between layers. Components from one layer can be used by many layers above it.

I also implemented Dependency injection in this application. Each layer has an interface that needs to be fulfilled when calling the `New()` method.
The reason to use dependency injection is to prevent hard dependency between layers. By using an interface, we can abstract our requirements, and it 
becomes possible to swap out the actual implementation for each layer (for example switching from postgre to mysql) as long as the new implementation 
satisfies our interface. Futher improving the modularity of the system

The package structure of each layer will look like so, these are just the possible files that a package can contain. it does not mean all package 
have these files:
- `packagename.go` (responsible to define interface, create new object with dependency fulfilled, and define object where logic methods will be written to)
- `types.go` (contains all type declaration required by the package)
- `queries.go` (contains constans of queries used by the package)
- `init.go` (contains initialization logic for that layer. We can modify/replace the actual implementation inside this file)
- `logicname.go` (we separate the actual logic of each layer by its main purpose for ease of reading. This file contains all the main logic of each layer)

## Possible improvements

By the nature of the limited time I have to develop this application, there are many things that can be improved further.
- Currently postgre and redis parameter are hardcoded. I want to make it possible to pass parameter from the outside in case there are configuration 
changes or when we use secret management tools. Centralizing the configuration will also make maintaining this application easier
- Further data validation. Currently the data passed into the server's insert product endpoint is not validated gracefully. This might cause an error. 
If possible I want to implement data validation so that the user can know which part of their request is wrong.
- Docker is a bit unstable. Running the app sometimes fail. I don't have time to debug the issue, so currently the experience of running this application
is not as smooth as i expected it to be
- Implement handlers that differentiates http methods. I want to be able to distinguish request by its http method (GET, POST, etc.)
