# nadc-intro-to-grpc
The second workshop of Night Login App Development Community (NADC) - Intro to gRPC using Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Branches list
There are three branches in this repository, they are:
1. `master` : consists of basic code with no implementation on the methods. Used for code-along session.
3. `final` : consists of all implementations and ready for deployment.

### Prerequisites

1. Clone repository: `git clone git@github.com:ridwanakf/nadc-intro-to-grpc.git`
2. Run dep: `dep ensure -v`
3. Generate protobuf code: `make proto`
3. Install postgresql: [Postgresql]("https://www.postgresql.org/download/")
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
|-- internal                               # Go files in this folder represent the Big-Pictures and Contracts of the system
|-- migrations                             # Contains Database migration files or the system
|-- vendor                                 # Dependencies folder that's maintained by dep tool https://golang.github.io/dep/
|-- Gopkg.lock                             # https://golang.github.io/dep/docs/Gopkg.lock.html
|-- Gopkg.toml                             # https://golang.github.io/dep/docs/Gopkg.toml.html
```

## Tech Stacks

- Golang
- Postgresql