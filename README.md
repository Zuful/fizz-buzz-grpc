# README #

This document will guide you through the set up of the gRPC service.

### Dependency manager ###

* The dependency manager is Golang-dep (runs on Go 1.9 or higher)
* Download the binary that suits your operating system at this adress : https://github.com/golang/dep/releases
* Paste the binary at the root of the project directory
* Rename the binary file to "dep" (keep the extension if there is one)
* Install the project's dependencies with the following command : dep ensure
* Your dependencies should be installed in a directory called "vendor"


### Launching the service ###

* In a command line window enter in the "server" directory
* Type the following command to launch the gRPC server : go run main.go
* In a different command line window enter in the "client" directory
* Type the following command to call the gRPC server : go run main.go <string1> <string2> <int1> <int2> <limit>

