package entity_test

import (
	"testing"
)

func TestNewEmptyCommit(t *testing.T) {
	t.Run("NewEmptyCommit", func(t *testing.T) {
		t.Parallel()
		t.Run("should return a new empty commit", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestNewCommit(t *testing.T) {
	t.Run("NewCommit", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an error when the commit type is empty", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return an error when the comment is empty", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return an error when the commit type is empty", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return a new commit", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestCommit_EnableEmoji(t *testing.T) {
	t.Run("Commit.EnableEmoji", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an array of booleans", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestCommit_Validate(t *testing.T) {
	t.Run("Commit.Validate", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an error when the subject is empty", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return an error when the commit type is unknown", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return an error when the commit type is empty", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("should return nil when the commit is valid", func(t *testing.T) {
			t.Parallel()
		})
	})
}

func TestCommit_Commit(t *testing.T) {
	t.Run("Commit.Commit", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an error when the commit is empty", func(t *testing.T) {
			t.Parallel()
		})
	})
}
