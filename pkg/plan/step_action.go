package plan

import "strings"

type StepAction string

func (a StepAction) Split() []string {
	return strings.Split(string(a), "/")
}

func (a StepAction) String() string {
	return strings.ReplaceAll(string(a), "/", " ")
}
