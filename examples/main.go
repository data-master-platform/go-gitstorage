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
