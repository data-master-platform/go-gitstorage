package gitstorage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	hashValue  = "hash_value"
	textInFile = "text_in_file"
	testFile   = "test_file.txt"
)

func TestNewMock(t *testing.T) {
	m := NewMock("", nil)
	var h Abstractor = (m)
	assert.Equal(t, "*gitstorage.Mock", reflect.TypeOf(h).String())
}

func TestMockCreate(t *testing.T) {
	m := Mock{ReturnError: nil}
	err := m.Create(testFile, textInFile)
	assert.NoError(t, err)
}

func TestMockRead(t *testing.T) {
	m := Mock{ReturnString: "hello", ReturnError: nil}
	str, err := m.Read(testFile)
	assert.NoError(t, err)
	assert.Equal(t, "hello", str)
}

func TestMockUpdate(t *testing.T) {
	m := Mock{ReturnError: nil}
	err := m.Update(testFile, textInFile)
	assert.NoError(t, err)
}

func TestMockDelete(t *testing.T) {
	m := Mock{ReturnError: nil}
	err := m.Delete(testFile)
	assert.NoError(t, err)
}

func TestMockgetHead(t *testing.T) {
	m := Mock{ReturnError: nil, ReturnString: hashValue}
	str, err := m.getHead()
	assert.NoError(t, err)
	assert.Equal(t, hashValue, str)
}

func TestMockpush(t *testing.T) {
	m := Mock{ReturnError: nil, ReturnString: hashValue}
	err := m.push(testFile)
	assert.NoError(t, err)
}

func TestMockcreateFile(t *testing.T) {
	m := Mock{ReturnError: nil}
	err := m.createFile(testFile, textInFile)
	assert.NoError(t, err)
}

func TestMockadd(t *testing.T) {
	m := Mock{ReturnError: nil, ReturnString: hashValue}
	str, err := m.add(testFile)
	assert.NoError(t, err)
	assert.Equal(t, hashValue, str)
}

func TestMockget(t *testing.T) {
	m := Mock{ReturnError: nil, ReturnString: textInFile}
	str, err := m.get(testFile)
	assert.NoError(t, err)
	assert.Equal(t, textInFile, str)
}

func TestMockcommit(t *testing.T) {
	m := Mock{ReturnError: nil, ReturnString: hashValue}
	str, err := m.commit()
	assert.NoError(t, err)
	assert.Equal(t, hashValue, str)
}

func TestMockdeleteFile(t *testing.T) {
	m := Mock{ReturnError: nil}
	err := m.deleteFile(fileName)
	assert.NoError(t, err)
}
