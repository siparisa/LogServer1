# Service - Services

## Introduction

`LogServer` is responsible for providing the domain of *System's Logs*.

## Contents

* [API Documentation](#API-Documentation)
    * [Installation]()
    * [Endpoints]()
* [Development](#development)
    * [Prerequisites](#prerequisites)
    * [Make Commands](#make-commands)
    * [Running the Service Locally](#running-the-service-locally)
    * [Project Layout](#project-layout)
        * [Layers and Folder Structure](#layers-and-folder-structure)


## API Documentation

### Installation
to start the service check `Prerequisites` of this file then run `make run` from the root directory of the project and run below endpoints.

example : `http://localhost:8080/logs?filename=system.log&keyword=log&n=5`

### Endpoints

all Endpoints are located in main.go to call related handlers.



**1.Get Log Lines**
Endpoint: `GET /logs`

Description: Retrieve log lines from a log file based on specified query parameters. The endpoint supports additional query parameters for filtering and pagination.

Query Parameters:
- **filename** (string, required): The name of the log file to retrieve log lines from.
- **keyword** (string, optional): Filter log lines by a specific keyword or text.
- **n** (integer, optional): Retrieve the last 'n' log lines from the file.

Example Request:
```
GET /logs?filename=example.log&keyword=error&n=50
```
Please note that the `keyword` and `n` parameters are optional. If `keyword` is provided, only log lines containing the specified keyword will be returned. If `n` is provided, the response will include the last 'n' log lines.

### 2. List Log Files
Endpoint: `GET /log-files`

Description: Retrieve a list of available log files in the `/var/log` directory.

Example Request:

```
GET /log-files
```
## Development

### Prerequisites

The following table lists _hard_ dependencies you will need to use this project.

| Name                                                       | Version  | Notes                                    |
|------------------------------------------------------------|----------|------------------------------------------|
| [Go](https://golang.org/doc/)                              | 1.17+    | Required to build and spin up service    |



### Make Commands

| Command  | Description                                                                                                                       |
|----------|-----------------------------------------------------------------------------------------------------------------------------------|
| run      | Starts the service                                                                                                                |
| test     | starts running unit tests                                                                                                         |



### Running the Service Locally

1. Run ```make run``` from project root it starts the service on port 8080.
2. Run ```make test```  from project root it starts the unit tests.


### Project Layout

This project roughly followed the layout of Go projects as described at
[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout).

| Directory     | Description                                                                                    |
|---------------|------------------------------------------------------------------------------------------------|
| `cmd/`        | This Go package is where `main` is used for the executables of the project                     |
| `internal/`   | Application specific Go packages, e.g., they cannot be shared and are specific to this service |
| `tests/`      | tests for the service are located in here.                                                     |

#### Layers and Folder Structure

There are 2 main layers in this repo, including `Controller` and `Service`. The only way for these layers
to interact with each other should be through their interfaces. The lower layers do not have any knowledge about
the upper layers.

The `entity` is the entities that represents the model in the database.

`internal/controller/helper` contains the models for every request and response.
The `helper` should not be used in `Service` or other layers.

