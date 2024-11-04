# Product Catalog 

Product Catalog API is responsible for feed recommendations

## How to Run

To run this application locally you will need:
- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-started/get-docker/)

### Commands
- `make up`: run the application in a docker container, triggering rebuild each time you run it.
- `make down`: shutdown the application containers and database volume.
- `make run_integration`: run the integration tests in a isolated database(`teardown_integration` runs before running the containers)
- `make teardown_integration`: shutdown the test containers and database volume.
- `make run_tests`: run unit tests.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

