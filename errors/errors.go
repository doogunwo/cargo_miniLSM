package errors

import (
	"errors"
	"fmt"

	"main/storage"
	"main/util"
)



//Common errors
var (
	ErrNotFound	=	new("leveldb :not found")
	ErrReleased =	util.ErrReleased
	ErrHasReleaser = util.ErrHasReleaser
)

// 

func New(text string) error {
	return errors.New(text)
}

type ErrCorrupted struct {
	Fd stroage.FildDesc
	Err error
}

func (e *ErrCorrupted) Error() string {
	if !e.Fd.Zero(){
		Sprintf("%v [fill=%v]", e.Err, e.Fd)
	}
	return e.Err.Error()
}