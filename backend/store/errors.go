package store

import "errors"

// ErrRecordNotFound is returned when record has not been
// found in database or any other store
var ErrRecordNotFound = errors.New("store: record not found")
