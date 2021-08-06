package gitstorage

import (
	"fmt"
	"io/ioutil"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
)

func (c *Client) getHead() (string, error) {
	pb, err := c.Repository.Head()
	if err != nil {
		return "", errors.Wrap(err, "unable to obtain worktree")
	}

	return pb.Hash().String(), nil
}

func (c *Client) createFile(path, data string) error {
	newFile, err := c.Filesystem.Create(path)
	if err != nil {
		return errors.Wrap(err, "unable to create file")
	}
	newFile.Write([]byte(data))
	newFile.Close()

	return nil
}

func (c *Client) deleteFile(path string) error {
	return c.Filesystem.Remove(path)
}

func (c *Client) add(path string) (string, error) {
	hash, err := c.Worktree.Add(path)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("unable to add path %s", path))
	}

	return hash.String(), nil
}

func (c *Client) commit() (string, error) {
	s := object.Signature{
		Name:  "test",
		Email: "gogs@gogs.com",
	}
	hash, err := c.Worktree.Commit("commit", &git.CommitOptions{Author: &s})
	if err != nil {
		return "", errors.Wrap(err, "unable to commit")
	}

	return hash.String(), nil
}

func (c *Client) get(path string) (string, error) {
	fs, err := c.Filesystem.Open(path)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("unable to open file with path %s", path))
	}

	b, err := ioutil.ReadAll(fs)
	if err != nil {
		return "", errors.Wrap(err, "unable to read content")
	}

	return string(b), nil
}

func (c *Client) push(path string) error {
	op := git.PushOptions{
		Auth: c.BasicAuth,
	}
	err := c.Repository.Push(&op)
	if err != nil {
		return errors.Wrap(err, "unable to push")
	}

	return nil
}
