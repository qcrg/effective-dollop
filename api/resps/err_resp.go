package resps

const (
	ErrGeneric = -1
)

type SectionError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

type ResponseErr struct {
	Error SectionError `json:"error"`
}

func MakeErr(code int, reason string) ResponseErr {
	return ResponseErr{SectionError{
		Code:   code,
		Reason: reason,
	}}
}

func MakeGenericErr(reason string) ResponseErr {
	return MakeErr(ErrGeneric, reason)
}
