package gitstorage

import (
	"log"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
)

type FileTreeReader struct {
	fs billy.Filesystem
}

type Client struct {
	*git.Repository
	*git.Worktree
	*http.BasicAuth
	FileTreeReader
}

type walkFunc = func(filepath string, isDir bool, err error) error

func (r FileTreeReader) Walk(walkFn walkFunc) error {
	if err := r.walk(string(filepath.Separator), true, walkFn); err != filepath.SkipDir {
		return err
	}
	return nil
}

func (r FileTreeReader) walk(path string, isDir bool, walkFn walkFunc) error {
	err := walkFn(path, isDir, nil)
	if err != nil {
		return err
	}
	if !isDir {
		return nil
	}

	infos, err := r.fs.ReadDir(path)
	if err != nil {
		return walkFn(path, isDir, err)
	}

	for _, info := range infos {
		filename := r.fs.Join(path, info.Name())
		err = r.walk(filename, info.IsDir(), walkFn)
		if err != nil {
			if !info.IsDir() || err != filepath.SkipDir {
				return err
			}
		}
	}
	return nil
}

type Abstractor interface {
	Create(string, string) error
	Read(string) (string, error)
	Update(string, string) error
	Delete(string) error
	List() ([]string, error)
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
	objects() ([]string, error)
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

	return &Client{r, w, auth, FileTreeReader{fs}}
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

func (c *Client) List() ([]string, error) {
	return c.objects()
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
