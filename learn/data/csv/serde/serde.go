package serde

type Serde interface {
	Decode(v interface{}) error
	Encode(v interface{}) error
	Flush() error
}

type Decoder interface {
	Decode(v interface{}) error
}

type Encoder interface {
	Encode(v interface{}) error
	Flush() error
}

type EncodeTyp int

const (
	Encode EncodeTyp = iota
	Decode
)
