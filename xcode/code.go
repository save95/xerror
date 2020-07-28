package xcode

type xCode struct {
	httpStatus int
	code       int
	message    string
}

func (c xCode) Code() int {
	return c.code
}

func (c xCode) HttpStatus() int {
	return c.httpStatus
}

func (c xCode) String() string {
	return c.message
}
