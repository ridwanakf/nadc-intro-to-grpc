# NADC - Intro to gRPC
The second workshop of Night Login App Development Community (NADC) - Intro to gRPC using Go.
This repo is intended to give an introduction to gRPC development using Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Branches list
There are three branches in this repository, they are:
1. `master` : consists of basic code with no implementation on the methods. Used for code-along session.
2. `final` : consists of all implementations and ready for deployment.
3. `final-clean-arch` : consists of all implementations and ready for deployment. Developed using Clean Architecture.

### Prerequisites

1. Clone repository: `git clone git@github.com:ridwanakf/nadc-intro-to-grpc.git`
2. Run dep: `dep ensure -v`
3. Generate protobuf code: `make proto`
3. Install postgresql: [Postgresql](https://www.postgresql.org/download/)
4. Run database migrations: refer below

### Run Project

Running project:

```$xslt
make run
```

### Migrations

When running in Local, you need to run the db-migrations to setup the app's database for your local machine.

1. Go to directory `nadc-intro-to-grpc/migrations`
2. Run `go run *.go up`

## Directory Structure

This repository is organized in the following directory structure.

```
nadc-intro-to-grpc
|-- cmd                                    # Contains executable codes and serves as entry point of the app
|   |-- main                               # entry point of the app
|-- config                                 # Configuration files needed for deployment
|-- constant                               # Collections of constants file for each module
|-- internal                               # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app                                # Contains constructions of the app and other app's related configs
|   |   |-- config                         # Configuration struct for the app
|   |-- delivery                           # Delivery layer of the app
|   |   |-- grpc                           # GRPC delivery of the app
|   |   |-- <other_delivery_mechanisms>    # Other delivery mechanisms of the app (eg. GRPC, Console, Web, etc.)
|   |
|   |-- entity                             # Enterprise Data structures
|   |   |-- book.go                        # Data structurefor Book module
|   |   |-- <other_entities>.go            # Other data structures, preferrably 1 struct 1 file
|   | 
|   |-- repo                               # Implementations of Repository-pattern to data-sources
|   |   |-- db                             # Implementations of the repositories with Postgres database
|   |   |-- <other_repos>                  # Other Repositories implementations based on interfaces on folder internal.
|   |
|   |-- usecase                            # Usecases implementations for Application Business Logic
|   |   |-- book.go                        # Other use-case implementations based on interfaces on folder internal.
|   |
|   |-- repo.go                            # Interfaces / Contracts of all the repositories (Repository Pattern)
|   |-- usecase.go                         # Interfaces / Contracts of all the use-cases (Application Business Logic)
|
|-- migrations                             # Contains Database migration files or the system
|-- vendor                                 # Dependencies folder that's maintained by dep tool https://golang.github.io/dep/
|-- Gopkg.lock                             # https://golang.github.io/dep/docs/Gopkg.lock.html
|-- Gopkg.toml                             # https://golang.github.io/dep/docs/Gopkg.toml.html

```

## Tech Stacks

- Golang
- Postgresql
- Protobuf
- gRPC