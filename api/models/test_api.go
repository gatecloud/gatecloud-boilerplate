package models

import "time"

type TestAPI struct {
	Name   string
	Value  string
	From   time.Duration `json:",at_least_one_required"`
	To     time.Duration `json:",at_least_one_required"`
	Number string        `json:",read_only"`
}
