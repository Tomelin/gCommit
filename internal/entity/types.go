package entity

import "sort"

// Option represents the type of commit
type Option uint8

const (
	Unknown Option = iota
	Feat
	Fix
	Chore
	Docs
	Style
	Refactor
	Perf
	Test
	CI
	Build
	Revert
)

// String returns the string representation of the Option
func (o *Option) String() string {
	switch *o {
	case Feat:
		return "feat"
	case Fix:
		return "fix"
	case Chore:
		return "chore"
	case Docs:
		return "docs"
	case Style:
		return "style"
	case Refactor:
		return "refactor"
	case Perf:
		return "perf"
	case Test:
		return "test"
	case CI:
		return "ci"
	default:
		return "unknown"
	}
}

// FromString sets the Option from a string
func (o *Option) FromString(s string) Option {
	switch s {
	case "feat":
		*o = Feat
		return Feat
	case "fix":
		*o = Fix
		return Fix
	case "chore":
		*o = Chore
		return Chore
	case "docs":
		*o = Docs
		return Docs
	case "style":
		*o = Style
		return Style
	case "refactor":
		*o = Refactor
		return Refactor
	case "perf":
		*o = Perf
		return Perf
	case "test":
		*o = Test
		return Test
	case "ci":
		*o = CI
		return CI
	default:
		*o = Unknown
		return Unknown
	}
}

// Emoji returns the emoji representation of the Option
// If the Option is unknown, it returns an empty string
// Examples:
//
//	Option(Feat).Emoji() => ":sparkles:"
//	Option(Fix).Emoji() => ":bug:"
//	Option(Unknown).Emoji() => ""
func (o *Option) Emoji() string {
	switch *o {
	case Feat:
		return ":sparkles:"
	case Fix:
		return ":bug:"
	case Chore:
		return ":wrench:"
	case Docs:
		return ":books:"
	case Style:
		return ":gem:"
	case Refactor:
		return ":hammer:"
	case Perf:
		return ":rocket:"
	case Test:
		return ":rotating_light:"
	case CI:
		return ":construction_worker:"
	case Build:
		return ":package:"
	case Revert:
		return ":rewind:"
	default:
		return ""
	}
}

func (o *Option) OptionsList() []string {

	options := []string{"feat", "fix", "chore", "docs", "style", "refactor", "perf", "test", "ci", "build", "revert"}
	sort.Strings(options)
	return options
}

// Status represents the type of commit
type Status uint8

const (
	UnknownStatus Status = iota
	Ignored
	Ready
	NotReady
	InProgress
	Resolved
	Backlog
	Done
	Removed
)

// String returns the string representation of the Status
func (s *Status) String() string {
	switch *s {
	case Ignored:
		return "ignored"
	case Ready:
		return "ready"
	case NotReady:
		return "notReady"
	case InProgress:
		return "inProgress"
	case Resolved:
		return "resolved"
	case Backlog:
		return "backlog"
	case Done:
		return "done"
	case Removed:
		return "removed"
	default:
		return "unknown"
	}
}

// FromString sets the Status from a string
func (s *Status) FromString(status string) Status {
	switch status {
	case "ignored":
		*s = Ignored
		return Ignored
	case "ready":
		*s = Ready
		return Ready
	case "notReady":
		*s = NotReady
		return NotReady
	case "inProgress":
		*s = InProgress
		return InProgress
	case "resolved":
		*s = Resolved
		return Resolved
	case "backlog":
		*s = Backlog
		return Backlog
	case "done":
		*s = Done
		return Done
	case "removed":
		*s = Removed
		return Removed
	default:
		*s = UnknownStatus
		return UnknownStatus

	}
}

func (s *Status) StatusList() []string {

	status := []string{"ignored", "ready", "notReady", "inProgress", "resolved", "backlog", "done", "removed"}
	sort.Strings(status)
	return status
}
