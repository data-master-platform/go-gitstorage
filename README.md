# gitstorage

[![CodeFactor](https://www.codefactor.io/repository/github/data-master-platform/go-gitstorage/badge)](https://www.codefactor.io/repository/github/data-master-platform/go-gitstorage) <a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-85%25-brightgreen.svg?longCache=true&style=flat)</a>
[![Shields](https://img.shields.io/github/go-mod/go-version/data-master-platform/go-gitstorage)](https://img.shields.io/github/go-mod/go-version/data-master-platform/go-gitstorage) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=data-master-platform_go-gitstorage&metric=code_smells)](https://sonarcloud.io/dashboard?id=data-master-platform_go-gitstorage)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=data-master-platform_go-gitstorage&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=data-master-platform_go-gitstorage)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=data-master-platform_go-gitstorage&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=data-master-platform_go-gitstorage)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=data-master-platform_go-gitstorage&metric=security_rating)](https://sonarcloud.io/dashboard?id=data-master-platform_go-gitstorage)

Simple wrapper around [go-git](https://github.com/go-git/go-git) which allows you to easily apply _CRUD_ operations in a similar way as pushing data in a blob fashion.

Create file to a path e.g. `/app/data/file.txt` will do `git add`, `git commit` and `git push` sequentially.


## Usage

Simple example containing all functionality

```go
package main

import (
	"log"

	"github.com/data-master-platform/gitstorage"
)

const (
	repo     = "http://docker.for.mac.localhost:3000/gogs/db-storage"
	branch   = "refs/heads/master"
	username = "gogs"
	password = "admin"
)

var testFileName = "file.txt"
var testDataFile = "hello"
var testDataFileUpdated = "hello world"

func main() {
	cl := gitstorage.New(username, password, repo, branch)
	err := cl.Create(testFileName, testDataFile)
	if err != nil {
		log.Println(err)
	}
	err = cl.Update(testFileName, testDataFileUpdated)
	if err != nil {
		log.Println(err)
	}
	data, err := cl.Read(testFileName)
	if err != nil {
		log.Println(err)
	}

	// Output of the data is equal to the changed data
	log.Println(data)

	err = cl.Delete(testFileName)
	if err != nil {
		log.Println(err)
	}
}
```

## Tests

Unit tests can be ran with `$ make test` and runs all without dependencies. When you'd like to run integration tests `$ make test-integration` would run them, but you need to have the dependencies running.

For convencience `$ make test-in-docker` will create the environment for you and run all the tests together and cleans up afterwards. If a test fails it will return an `exit 2` otherwise will have `exit 0`.

### In code

When you would like to test the code you can simply take use of the Mock as being used in `mock_test.go`.


*Note:*
Only tested on [gogs](https://github.com/gogs/gogs) and on [github](https://github.com)