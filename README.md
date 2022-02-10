# lexis-backend-services-api

This repository contains a middleware service in the LEXIS PORTAL, which is back-end to the front-end, and front-end to user-org-service, dataset-management-interface, etc.

## Acknowledgement

This code repository is a result / contains results of the LEXIS project. The project has received funding from the European Union’s Horizon 2020 Research and Innovation programme (2014-2020) under grant agreement No. 825532.

## LEXIS PORTAL - API (proxy) Service

The API (proxy) service is meant to serve as a proxy for all the other services interfaces that conform the LEXIS PORTAL.
So far, this service exposes the following services:
- LEXIS UserOrg
- WP3 Datasets
- A4C Workflows
- Cyclops Accounting
- HPC Resources Approval


## How to run the Server

To run the server, enter the server directory and run
    go run main.go config.go


## Capabilities

The service implements a RESTful API using the golang implementation of Swagger, to be precise we're using the Stratoscale implementation.

The service is able to check for valid keycloak tokens for its usage.

## Building

The building of the service is carried by a multistage docker build, we build everything in an image containing everything needed to build Golang applications
and then the executable is moved to a new image that contains the bareminimum starting from the scratch image.

Within the folder build at the root of the repo there's a script call start.sh, it's invokation admits "Dev" as parameter.
Running the script with the parameter will use the local version in the repository as the base for the building of the service's docker image,
however doing it without providing it will make the building of the service taking everything from sources.

```
./start.sh [Dev]
```

The initial branch used when building from sources is "master", it can be changed with a few other parameters by editing the script.

## Running

Within the folder run at the root of the repo there's a docker-compose sample file and a config.toml sample file.
Once configured with appropriate data to start the service just issue the following comand:

```
docker-compose up -d
```

Within the folder Utils there's also a couple of scripts to generate self-signed certificates and keycloak tokens, but they might need some work out.
Same goes for the system-tester (WIP: some more tests should be created) script which only contains a basic test of the system.


## Useful links

- Fast overview of golang: https://learnxinyminutes.com/docs/go/
- Swagger generation tool: https://github.com/Stratoscale/swagger
- Golang swagger documentation: https://goswagger.io/

## How to generate model/server/client

### Official docs:

- https://goswagger.io/generate/model.html
- https://goswagger.io/generate/client.html
- https://goswagger.io/generate/server.html

### It uses swagger generation tool:

```
docker run --rm -e GOPATH=/home/davidhruby/go:/go -v ${HOME}:${HOME} -w $(pwd) -u $(id -u):$(id -g) quay.io/goswagger/swagger:v0.25.0 generate model --template=stratoscale
docker run --rm -e GOPATH=/home/davidhruby/go:/go -v ${HOME}:${HOME} -w $(pwd) -u $(id -u):$(id -g) quay.io/goswagger/swagger:v0.25.0 generate server --template=stratoscale
``` 

### If you use Fish shell then run it in following format:

```
docker run --rm -e GOPATH=/Users/hrubos/go:/go -v {$HOME}:{$HOME} -w (pwd) -u (id -u):(id -g) quay.io/goswagger/swagger:v0.25.0 generate model --template=stratoscale
docker run --rm -e GOPATH=/Users/hrubos/go:/go -v {$HOME}:{$HOME} -w (pwd) -u (id -u):(id -g) quay.io/goswagger/swagger:v0.25.0 generate server --template=stratoscale
```
