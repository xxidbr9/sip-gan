package shared

import (
	"fmt"
)

//InfrastructureError spawned if there are error related to infrastructure like database
type InfrastructureError struct {
	err error
}

func NewInfrastructureError(err error) *InfrastructureError {
	return &InfrastructureError{
		err: err,
	}
}

func (err *InfrastructureError) Error() string {
	return fmt.Sprintf("There is infrastructure error which is: %v", err.err.Error())
}
