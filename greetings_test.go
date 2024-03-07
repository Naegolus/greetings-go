
package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {

	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)

	err, msg := Hello("Gladys")
	if err != nil || !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %v, %q, want match for %#q, nil`, err, msg, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	
	err, msg := Hello("")
	if err == nil || msg != "" {
		t.Fatalf(`Hello("") = %v, %q, want "", error`, err, msg)
	}
}

