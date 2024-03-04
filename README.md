```markdown
# Architecture

The API is built using Clean Architecture. The principles of DIP (Dependency Inversion Principle) and DI (Dependency Injection) are utilized. The error handling structure is designed to provide maximum flexibility at the API level and convenience for the client.

## Libs

- gin
- gorm
- zap
- cleanenv
- testify
- mockery
- swaggo

## Makefile Commands

- `make up` - to start the API using docker-compose
- `make test` - to run tests
- `make docs` - to generate swagger documentation
- `make mocs` - to generate mocks for testing

## Documentation

The documentation is available at the following link: http://localhost:8080/api/v1/docs/swagger/index.html

## Possible Improvements

1. **Use DTOs for Data Transfer:** Utilize DTOs for data transfer to the storage layer, as currently, database details are leaking into the entity (gorm tags). However, this might significantly complicate the code without providing substantial benefits for such a small API, which is why it hasn't been implemented.

2. **Add Health Checks:** Implement health checks for the database and server, and invoke them periodically when the API is running.

3. **Enhance Error Checking in Tests:** In all negative test cases, do not only check for the presence of an error (using expectErr) but also ensure that the error type is verified (whether it is expected or not, and if expected, what exactly it is).

4. **Add Controller-Level Tests:** Implement controller-level tests using mocks for the service layer.
```
