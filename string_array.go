package main

import (
	"strings"
)

type StringArray []string

func (a *StringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

func (a *StringArray) Get() interface{} {
	var s string
	for _, v := range *a {
		if v != "" {
			s += v
		}
	}
	if s != "" {
		return s
	}
	return nil
}

func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}
