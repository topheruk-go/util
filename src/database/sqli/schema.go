package sqli

type DataTyp int

const (
	Integer DataTyp = iota
	Text
	Real
	Blob
	Boolean
	DateTime
)
