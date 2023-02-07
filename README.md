
# Go Hexagolan Architecture

A project demonstrating the use of pure architecture in the Golang language using a hexagonal structure. This is a design pattern using encapsulation of the application core with an application API. We add left and right sides to the API part

The right side to be responsible for applications that are part of the application such as a database, cloud solution provider. 

The left side is responsible for the user interaction and it is the user who calls the action and it, the action of the right side is automatic and forced when needed.

Below is a picture showing an overview of the application built in hexagonal architecture

![App Screenshot](https://github.com/kleczynski/GolangHexArch/blob/main/images/hexarch.png)

On the left-hand side of the project is the system responsible for the fastest communication currently possible in the systems. The gRPC system. On the right we have a database whose operation will be imposed by user interaction 


In the project, in the section responsible for communications, there is a file with tests. These tests are built in a way that resembles the main file responsible for starting the application. We need to create an imitation of a connection and test the communication with the imposed values. 

## Building and testing - Applications

Using command ```docker-compose -d up``` aplication will be build and entire test for communication will be run. In docker we should see logs same as on the picture belowe. 

![App Screenshot](https://github.com/kleczynski/GolangHexArch/blob/main/images/testimage.png)
