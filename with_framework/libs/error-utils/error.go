package errorUtils

import "fmt"

type applicationError struct {
	error
	Code int
}

func WrapWithCode(err error, code int) error {
	return &applicationError{wrap(err, ""), code}
}

func UnWrapWithCode(err error) (*applicationError, bool) {
	fmt.Println(err)
	for err != nil {
		switch err.(type) {
		case *applicationError:
			return err.(*applicationError), true
		}
	}
	return nil, false
}
