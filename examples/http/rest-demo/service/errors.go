package service

import "errors"

var ErrTodo = errors.New("service: todo")
var ErrNoMatch = errors.New("sql: could not find a match")
