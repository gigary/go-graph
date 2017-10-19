# Graph Algorithms Implemented in Go (Golang)

## Test

### Prepare a test file as in `test.data`
Each line of the file is one test case that comes with a graph definition and 10 tests:
1. Calculate distance along a valid route
2. Calculate distance along a valid route
3. Calculate distance along a valid route
4. Calculate distance along a valid route
5. Throw error if calculate distance with an invalid route
6. Number of routes between 2 nodes with a maximum stops limitation
7. Number of routes between 2 nodes with exact stops limitation
8. Calculate shortest distance between 2 nodes 
9. Calculate shortest distance between 2 nodes 
10. Calculate routes between 2 nodes with a maximum distance limitation

### Run tests against the provided test file
Run `go run main.go "test.data"` (assumed that `test.data` is the test file) in the root directory of the repo

### Run unit tests
Run `go test -v ./...` (argument `-v` is optional) in the root directory of the repo

## Author
[Michael Bui](mailto:mf.michaelbui@gmail.com) from [GigaryTeam](https://gigary.com)
