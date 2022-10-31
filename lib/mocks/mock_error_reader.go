package mocks

import "errors"

type ErrReader int

func (ErrReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("an error")
}

func (ErrReader) Close() (err error) {
	return nil
}
