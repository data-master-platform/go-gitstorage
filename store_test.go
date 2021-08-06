package gitstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	s, err := cl.getHead()
	assert.NoError(t, err)
	assert.Greater(t, len(s), 5)
}

var fileName = "avro.avsc"
var testData = `{"namespace":"example.avro","type":"record","name":"User","fields":[{"name":"name","type":"string"},{"name":"favorite_number","type":["int","null"]},{"name":"favorite_color","type":["string","null"]}]}`

func TestIntegrationCreateFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.createFile(fileName, testData)
	assert.NoError(t, err)
}

func TestIntegrationAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	h, err := cl.add(fileName)
	assert.NoError(t, err)
	assert.Greater(t, len(h), 5)
}

func TestIntegrationCommit(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	h, err := cl.commit()
	assert.NoError(t, err)
	assert.Greater(t, len(h), 5)
}

func TestIntegrationPush(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.push(fileName)
	assert.NoError(t, err)
}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	str, err := cl.get(fileName)
	assert.NoError(t, err)
	assert.Equal(t, testData, str)
}

func TestIntegrationDeleteFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.deleteFile(fileName)
	assert.NoError(t, err)
}

func TestIntegrationAddSecond(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	h, err := cl.add(fileName)
	assert.NoError(t, err)
	assert.Greater(t, len(h), 5)
}

func TestIntegrationCommitSecond(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	h, err := cl.commit()
	assert.NoError(t, err)
	assert.Greater(t, len(h), 5)
}

func TestIntegrationPushSecond(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	err := cl.push(fileName)
	assert.NoError(t, err)
}
