package entity

import (
	"errors"
	"fmt"
)

type CommitInterface interface {
	Reset(count uint) error
	Push() error
	Commit(c *Commit) error
}

type Commit struct {
	Subject    string
	Option     Option
	StatusType Status
	Choice     string
	TaskID     string
	Comment    string
	Status     string
	Emoji      string
}

func NewEmptyCommit() *Commit {
	return &Commit{}
}

// NewCommit creates a new commit entity
// It requires a commit type and a comment
// It returns a pointer to a Commit entity and an error
// If the commit type is empty, it returns an error
// Examples:
//
//	NewCommit(nil, nil) => error
//	NewCommit(nil, "comment") => error
//	NewCommit("feat", nil) => error
//	NewCommit("feat", "comment") => &Commit{Option: Feat, Choice: "feat"}
func NewCommit(commitType, comment *string) (*Commit, error) {

	result := &Commit{
		Subject: *comment,
		Choice:  *commitType,
	}

	if err := result.Validate(); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Commit) EnableEmoji() []bool {
	return []bool{true, false}

}

func (c *Commit) Validate() error {

	if c.Subject == "" {
		return errors.New("subject is required")
	}

	if c.Option == Unknown || c.Option == Option(0) {
		return errors.New("commit type is required")
	}

	if c.Choice == "" {
		return errors.New("commit type is required")
	}

	o := Option(Unknown)
	c.Option = o.FromString(c.Choice)
	if c.Option == Unknown {
		options := []Option{Feat, Fix, Chore, Docs, Style, Refactor, Perf, Test, CI, Build, Revert}
		has := []string{}
		for _, option := range options {
			has = append(has, option.String())
		}

		return fmt.Errorf("invalid commit type. Available options: %v", has)
	}

	return nil
}
