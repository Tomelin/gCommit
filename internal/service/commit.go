package service

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Tomelin/go-convetional-commit/internal/entity"
)

type ServiceCommit struct{}

func NewServiceCommit() entity.CommitInterface {
	return &ServiceCommit{}
}

func (sc *ServiceCommit) Reset(count uint) error {
	var cmd *exec.Cmd
	count--
	if count == 0 {
		cmd = exec.Command("git", "reset", "HEAD~")
	} else {
		cmd = exec.Command("git", "reset", fmt.Sprintf("HEAD~%d", count))
	}

	return cmd.Run()

}
func (sc *ServiceCommit) Push() error {

	cmd := exec.Command("git", "push")
	return cmd.Run()
}

func (sc *ServiceCommit) Commit(cm *entity.Commit) error {

	var body []string

	if cm.Emoji != "" {
		body = append(body, fmt.Sprintf("%s: %s %s\n\n", cm.Option.String(), cm.Option.Emoji(), cm.Subject))
	} else {
		body = append(body, fmt.Sprintf("%s: %s\n\n", cm.Option.String(), cm.Subject))
	}

	if cm.TaskID != "" && (cm.StatusType != entity.UnknownStatus && cm.StatusType != entity.Ignored) {

		body = append(body, fmt.Sprintf("%s\n\n%s: %s", cm.Comment, cm.StatusType.String(), cm.TaskID))

	} else if cm.TaskID == "" && (cm.StatusType != entity.UnknownStatus && cm.StatusType != entity.Ignored) {
		body = append(body, fmt.Sprintf("%s\n\nStatus: %s", cm.Comment, cm.StatusType.String()))
	} else if cm.TaskID != "" && (cm.StatusType == entity.UnknownStatus || cm.StatusType == entity.Ignored) {
		body = append(body, fmt.Sprintf("%s\n\nTask: %s", cm.Comment, cm.TaskID))
	} else {
		body = append(body, cm.Comment)
	}

	commitMessage := strings.Join(body, "\n")

	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}
	if out.Len() == 0 {
		return fmt.Errorf("no changes to commit")
	}
	cmd = exec.Command("git", "commit", "-m", commitMessage)
	return cmd.Run()

}
