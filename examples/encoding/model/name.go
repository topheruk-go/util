package model

type Name string

func (n *Name) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return ErrTooShort
	}
	*n = Name(data)
	return nil
}
