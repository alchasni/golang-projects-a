package validatoradapter

import (
	"fmt"
	"strings"
)

type tag string

func Tag(v ...string) tag {
	t := tag("")
	for _, value := range v {
		t = t.append(value)
	}

	return t
}

func (t tag) String() string {
	return string(t)
}

func (t tag) append(ext string) tag {
	str := t.String()
	if str != "" {
		str += ","
	}

	return tag(str + ext)
}

func (t tag) Required() tag {
	return t.append("required")
}

func (t tag) Omitempty() tag {
	return t.append("omitempty")
}

func (t tag) Numeric() tag {
	return t.append("numeric")
}

func (t tag) Alpha() tag {
	return t.append("alpha")
}

func (t tag) AlphaNum() tag {
	return t.append("alphanum")
}

func (t tag) Email() tag {
	return t.append("email")
}

func (t tag) URL() tag {
	return t.append("url")
}

func (t tag) Max(v int) tag {
	return t.append(fmt.Sprintf("max=%d", v))
}

func (t tag) Min(v int) tag {
	return t.append(fmt.Sprintf("min=%d", v))
}

func (t tag) OneOf(v ...interface{}) tag {
	values := make([]string, len(v))
	for index, value := range v {
		values[index] = fmt.Sprintf("%v", value)
	}

	return t.append(fmt.Sprintf("oneof=%s", strings.Join(values, " ")))
}

func (t tag) Gt(v int) tag {
	return t.append(fmt.Sprintf("gt=%d", v))
}

func (t tag) Gte(v int) tag {
	return t.append(fmt.Sprintf("gte=%d", v))
}

func (t tag) Lt(v int) tag {
	return t.append(fmt.Sprintf("lt=%d", v))
}

func (t tag) Lte(v int) tag {
	return t.append(fmt.Sprintf("lte=%d", v))
}
