* (Resource | Go-fiber http backend api)[https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-fiber-version-4la0]

* The go.sum file contains all the dependency checksums, and is managed by the go tools. We don’t have to worry about it.

* configs is for modularizing project configuration files

* controllers is for modularizing application logics.

* models is for modularizing data and database logics.

* responses is for modularizing files describing the response we want our API to give. This will become clearer later on.

* routes is for modularizing URL pattern and handler information.

* why need to use context and cancellation in go-fiber controller: https://www.sohamkamani.com/golang/context-cancellation-and-values/

* [REST API Docs | Postman](https://www.postman.com/akshayrr27/workspace/akshay-r-r/collection/11715636-64adc057-9fc1-4948-9152-11eeb1558cda?ctx=documentation)