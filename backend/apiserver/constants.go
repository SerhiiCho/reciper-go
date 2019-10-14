package apiserver

const (
	sessionName               = "reciper"
	contextKeyUser contextKey = iota
	contextKeyRequestID
)

type contextKey int8
