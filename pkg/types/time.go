package types

import "time"

type TimeUnix uint64

func (t TimeUnix) Time() time.Time {
	return time.Unix(0, t.Int64())
}

func (t TimeUnix) Int64() int64 {
	return int64(t)
}
