package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Anna"
	want := regexp.MustCompile(`\b` + name + `\b`)
	mess, err := Hello("Anna")
	if !want.MatchString(mess) || err != nil {
		t.Fatalf(`Hello("Anna") = %q, %v, want match for %#q, nil`, mess, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
