package domainevent

import "time"
type DomainEvent interface {
	Name() *string
	Data() interface{}
	CreatedAt() time.Time
}