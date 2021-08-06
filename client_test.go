package gitstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	repo     = "http://docker.for.mac.localhost:3000/gogs/db-storage"
	branch   = "refs/heads/master"
	username = "gogs"
	password = "admin"
)

var cl Abstractor
var testFileName = "file.txt"
var testDataFile = "hello"
var testDataFileUpdated = "hello world"

func TestIntegrationNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	cl = New(username, password, repo, branch)
}

func TestIntegrationCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.Create(testFileName, testDataFile)

	assert.NoError(t, err)
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.Update(testFileName, testDataFileUpdated)

	assert.NoError(t, err)
}

func TestIntegrationRead(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	str, err := cl.Read(testFileName)

	assert.Equal(t, testDataFileUpdated, str)
	assert.NoError(t, err)
}

func TestIntegrationDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.Delete(testFileName)

	assert.NoError(t, err)
}
