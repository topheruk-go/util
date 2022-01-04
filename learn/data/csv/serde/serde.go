package serde

type Serde interface {
	Decode(v interface{}) error
	Encode(v interface{}) error
	Flush() error

	Map(f func(field string, col string, v interface{}) string)
}

type ModelSerde interface {
	Set(v interface{})
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
