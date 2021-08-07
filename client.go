package gitstorage

import (
	"log"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
)

type Client struct {
	*git.Repository
	*git.Worktree
	*http.BasicAuth
	billy.Filesystem
}

type Abstractor interface {
	Create(string, string) error
	Read(string) (string, error)
	Update(string, string) error
	Delete(string) error
	Storer
}

type Storer interface {
	getHead() (string, error)
	push(string) error
	createFile(string, string) error
	add(string) (string, error)
	get(string) (string, error)
	commit() (string, error)
	deleteFile(string) error
}

func New(username, password, repository, branch string) Abstractor {
	auth := &http.BasicAuth{
		Username: username,
		Password: password,
	}
	fs := memfs.New()
	m := memory.NewStorage()

	r, err := git.Clone(m, fs, &git.CloneOptions{
		URL:           repository,
		Auth:          auth,
		ReferenceName: plumbing.ReferenceName(branch),
	})
	if err != nil {
		log.Fatalln(errors.Wrap(err, "unable to establish clone"))
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "unable to obtain worktree"))
	}

	return &Client{r, w, auth, fs}
}

func (c *Client) Create(name string, data string) error {
	err := c.createFile(name, data)
	if err != nil {
		return err
	}
	_, err = c.add(name)
	if err != nil {
		return err
	}
	_, err = c.commit()
	if err != nil {
		return err
	}
	err = c.push(name)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Read(name string) (string, error) {
	return c.get(name)
}

func (c *Client) Update(name string, data string) error {
	return c.Create(name, data)
}

func (c *Client) Delete(name string) error {
	err := c.deleteFile(name)
	if err != nil {
		return err
	}
	_, err = c.add(name)
	if err != nil {
		return err
	}
	_, err = c.commit()
	if err != nil {
		return err
	}
	err = c.push(name)
	if err != nil {
		return err
	}
	return nil
}
