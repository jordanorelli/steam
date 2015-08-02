package steam

import (
	"fmt"
)

type ClientError struct {
	msg    string
	parent error
}

func (c ClientError) Error() string {
	if c.parent == nil {
		return fmt.Sprintf("steam client error: %s", c.msg)
	}
	return fmt.Sprintf("steam client error: %s: %v", c.msg, c.parent)
}

func errorf(parent error, msg string, args ...interface{}) error {
	return ClientError{msg: fmt.Sprintf(msg, args...), parent: parent}
}
