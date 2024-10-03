package service_test

import (
	"testing"
)

func TestNewServiceCommit(t *testing.T) {
	t.Run("NewServiceCommit", func(t *testing.T) {
		t.Parallel()
		t.Run("should return a new service commit", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestServiceCommit_Reset(t *testing.T) {
	t.Run("ServiceCommit.Reset", func(t *testing.T) {
		t.Parallel()
		t.Run("should reset the last commit", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should reset the last 2 commits", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestServiceCommit_Push(t *testing.T) {
	t.Run("ServiceCommit.Push", func(t *testing.T) {
		t.Parallel()
		t.Run("should push the last commit", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestServiceCommit_Commit(t *testing.T) {
	t.Run("ServiceCommit.Commit", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an error when the commit message is empty", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return an error when the commit message is invalid", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should commit the changes", func(t *testing.T) {
			t.Parallel()
		})
	})
}
