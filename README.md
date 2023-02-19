# orders
The project is mainly focus on order management system. In this project we performed three activities.
1.Create a order into Database.
2.Get the all orders details.
3.Update the existing order.

To run this project : go run sellerapp_order.go

Docker Commands: 
Build: docker build -t app:v1 -f Dockerfile .
Run: docker run -d -p 9088:9088 -it app:v1
If you want to enter into docker container: docker exec -t -i (containername) /bin/bash

Technologies:
1.Go
It is an open-source language. It is a statically typed, compiled language designed to be simple, efficient, and reliable. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
Go is designed to be fast, scalable, and productive.

2.Mysql:
MySQL is a relational database management system (RDBMS) based on Structured Query Language (SQL). It is an open-source software.
It can be used to store, retrieve, modify, and delete data from databases

3.Json:
Json is JavaScript Object Notation. It is primarily used to transmit data between a server and web applications.

4. REST APIs:
REST (Representational State Transfer) is an architectural style for building web services. It is based on the principles of resource-oriented design and makes use of HTTP, which is the most popular protocol used for communication over the internet. RESTful APIs use HTTP requests to perform operations such as creating, reading, updating and deleting data from a web service. The operations are typically referred to as CRUD operations (Create, Read, Update, Delete).
