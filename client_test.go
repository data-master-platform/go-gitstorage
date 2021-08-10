package gitstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	repo     = "http://localhost:3000/gogs/db-storage"
	branch   = "refs/heads/master"
	username = "gogs"
	password = "admin"

	msgSkipTest = "skipping integration test"
)

var cl Abstractor
var testFileName = "file.txt"
var testDataFile = "hello"
var testDataFileUpdated = "hello world"

func TestIntegrationNew(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	cl = New(username, password, repo, branch)
}

func TestIntegrationCreateAssertNoPath(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	err := cl.Create("", testDataFile)

	assert.Error(t, err)
}

func TestIntegrationUpdateAssertNoPath(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	err := cl.Update("", testDataFile)

	assert.Error(t, err)
}

func TestIntegrationDeleteAssertNoPath(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	err := cl.Delete("")

	assert.Error(t, err)
}

func TestIntegrationReadAssertNoPath(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	str, err := cl.Read("")

	assert.Error(t, err)
	assert.Equal(t, "", str)
}

func TestIntegrationCreate(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	err := cl.Create(testFileName, testDataFile)

	assert.NoError(t, err)
}

func TestIntegrationList(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	str, err := cl.List()

	var present = false
	for _, v := range str {
		if v == "/"+testFileName {
			present = true
		}
	}

	assert.True(t, present)
	assert.NoError(t, err)
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	err := cl.Update(testFileName, testDataFileUpdated)

	assert.NoError(t, err)
}

func TestIntegrationRead(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	str, err := cl.Read(testFileName)

	assert.Equal(t, testDataFileUpdated, str)
	assert.NoError(t, err)
}

func TestIntegrationDelete(t *testing.T) {
	if testing.Short() {
		t.Skip(msgSkipTest)
	}
	err := cl.Delete(testFileName)

	assert.NoError(t, err)
}
