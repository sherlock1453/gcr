package interfaces

import "time"

type Item interface {
	getParent() *Item
	getPath() string
	getSession() *Session
}

type Node interface {
	getParent() *Item
	getPath() string
	getSession() *Session
}

type Property interface {
	getParent() *Item
	getPath() string
	getSession() *Session
	getType() int
	getValue() Value
	getValues() []Value
}

type Value interface {
	getStringValue() (error, string)
	getFloatValue() (error, float64)
	getIntValue() (error, int32)
	getDateValue() (error, time.Time)
	getBinaryValue() (error, []byte)
}
