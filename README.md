# README #

This document will guide you through the set up of the service.

### Dependency manager ###

* The dependency manager is Golang-dep (runs on Go 1.9 or higher)
* Download the binary that suits your operating system at this adress : https://github.com/golang/dep/releases
* Paste the binary at the root of the project directory
* Rename the binary file to "dep" (keep the extension if there is one)
* Install the project's dependencies with the following command : dep ensure
* Your dependencies should be installed in a directory called "vendor"


### Launching the service ###

* In the command line go to the root of the project directory
* Type the following command : go run main.go fizzbuzz.handler.go
