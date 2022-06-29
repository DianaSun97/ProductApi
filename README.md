![gopher](https://camo.githubusercontent.com/20d197aca375375e6b8401364f51aad3e06b3e427b357480e81cc90e510c3bb7/68747470733a2f2f6873746f2e6f72672f776562742f69682f64732f66752f696864736675716e693561706a306d793138746e756b7a7a7477302e706e67 "Орк")

Product Api
=========== 
## Technical decisions
+ I created a database with a unique key and product data (id, name)
+ Connected the database to API
+ Set up MySQL in Docker
+ Configured routing
+ Added HTML pages (index and create)
+ Made a page redirect after sending a form
+ Implemented the task with standard library HTTP
+ Installed and rewrote task with the Gin framework

### Technologies used
+ Golang
+ Docker
+ HTML
+ Gin framework
+ MySQL

### Database Setup
Database used in this application is MySQL as per requirements.
Prior to starting application, it's necessary to configure the database.

I used Docker compose to run the database:
```
docker-compose up -d
```

After the database is started, you need to load the [schema.sql](./schema.sql).

Database uses the following credentials: `root:admin`.

### Start server

```
go run .
```

List of product is at [index](http://localhost:8181/), adding products is at [/create](http://localhost:8181/create),
API is at [/api/products](http://localhost:8181/api/products).

To list products:
```
GET /api/products
```
Response:
```
[{"Id":1,"Name":"Test"}]
```

To add product:
```
POST /api/products
{"Name":"Test"}
```
Response:
```
{"Id":1,"Name":"Test"}
```

### What could be better
+ Add changes and delete data
+ Write small unit tests
+ Add validation to the form and API
+ Set port on listen API
+ Give setting database address

If any questions about the task, you can send me a message to dianaart1997@gmail.com
