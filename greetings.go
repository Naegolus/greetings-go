
package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (error, string) {

	if name == "" {
		return errors.New("empty name"), ""
	}

	/* Could also be
	 * var msg string
	 * msg = fmt...
	 */
	msg := fmt.Sprintf(randomFormat(), name)

	return nil, msg
}

func Hellos(names []string) (error, map[string]string) {

	var msgs map[string]string

	msgs = make(map[string]string)

	for _, name := range names {
		err, msg := Hello(name)

		if err != nil {
			return err, nil
		}

		msgs[name] = msg
	}

	return nil, msgs
}

func randomFormat() string {

	var formats []string

	formats = []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

