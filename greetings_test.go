package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("gladys")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("gladys") = %q,%v want match for %#q,nil`, msg, err, want)
	}
}

func TestEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q,%v want "", error`, msg, err)
	}

}
