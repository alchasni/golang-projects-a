package types

type Code string

func (c Code) String() string {
	return string(c)
}

func (c Code) Empty() bool {
	return c.Equals("")
}

func (c Code) Equals(code Code) bool {
	return c == code
}

func (c Code) In(codes []Code) bool {
	for _, code := range codes {
		if c == code {
			return true
		}
	}

	return false
}

func (c Code) OneOf(codes ...Code) bool {
	return c.In(codes)
}
