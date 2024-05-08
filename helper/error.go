package helper

import "github.com/corysakti/cats-social-go/exception"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicThrowNotFound(err error) {
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
}

func PanicThrowAlreadyExist(err error) {
	if err != nil {
		panic(exception.NewAlreadyExistError(err.Error()))
	}
}

func PanicPasswordIsWrong(err error) {
	if err != nil {
		panic(exception.NewPasswordIsWrongError(err.Error()))
	}
}
