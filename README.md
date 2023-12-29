# Project cobra-api

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

---

## Cobra runner
This project updates the `go-blueprint` to use [Cobra](https://github.com/spf13/cobra) as an entrypoint mechanism to allow applications to handle multiple concerns.

This is useful projects that don't do a single task or need to be able to scale activities in a more controlled manner.

Using this you would be able to:
* Start cron jobs in a separate container/process
* Run your DB migrations for ORMS like Gorm as a command, i.e: `go run migrate`
* Add a queue consumer process that can be scaled up/down independently of the REST API
* Allow for more controlled implementations and complete projects

### Running API
```BASH
make run api
```

### Running Cron Jobs
```BASH
make run cron
```

### Starting both in containers
```BASH
make docker-run
```