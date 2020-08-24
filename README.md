# GO-MICROS

###
Go-Micros, short for Golang Microservices, is an orchestrated pair of docker containers. One is a golang RESTful API and the other is a MySQL server with a simple database. Using docker-compose commands, you can spin up the API and connect to the database.

## REST Endpoints
````
/* Return all users */
GET localhost:8080/users
Ex. curl -X -G http://127.0.0.1:8080/users

/* Return user by ID */
GET localhost:8080/users/:user_id
@params id
Ex. curl -X -G http://127.0.0.1:8080/users/3

/* Create and insert user into table */
POST localhost:8080/users
@params first_name, last_name, email, username, password
Ex. curl -X -P http://127.0.0.1:8080/users?first_name=Hashim&last_name=Jacobs&email=hashiejizzle4whaat@mindspring.com&username=hashjaco&password=Iamthefuckinrootuser187!
````