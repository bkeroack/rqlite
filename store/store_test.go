package store

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	sql "github.com/otoolep/rqlite/db"
)

func Test_OpenStoreSingleNode(t *testing.T) {
	s := mustNewStore()
	defer os.RemoveAll(s.Path())

	if err := s.Open(true); err != nil {
		t.Fatalf("failed to open single-node store: %s", err.Error())
	}
}

func Test_OpenStoreCloseSingleNode(t *testing.T) {
	s := mustNewStore()
	defer os.RemoveAll(s.Path())

	if err := s.Open(true); err != nil {
		t.Fatalf("failed to open single-node store: %s", err.Error())
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close single-node store: %s", err.Error())
	}
}

func Test_SingleNodeExecuteQuery(t *testing.T) {
	s := mustNewStore()
	defer os.RemoveAll(s.Path())

	if err := s.Open(true); err != nil {
		t.Fatalf("failed to open single-node store: %s", err.Error())
	}
	defer s.Close()
	s.WaitForLeader(10 * time.Second)

	queries := []string{
		`CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT)`,
		`INSERT INTO foo(id, name) VALUES(1, "fiona")`,
	}
	_, err := s.Execute(queries, false)
	if err != nil {
		t.Fatalf("failed to execute on single node: %s", err.Error())
	}
}

func mustNewStore() *Store {
	path := mustTempDir()
	defer os.RemoveAll(path)

	s := New(newInMemoryConfig(), path, ":0")
	if s == nil {
		panic("failed to create new store")
	}
	return s
}

func mustTempDir() string {
	var err error
	path, err := ioutil.TempDir("", "rqlilte-test-")
	if err != nil {
		panic("failed to create temp dir")
	}
	return path
}

func newInMemoryConfig() *sql.Config {
	c := sql.NewConfig()
	c.Memory = true
	return c
}