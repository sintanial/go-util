package timeutil

import (
	"time"
	"strconv"
	"encoding/json"
	"reflect"
	"strings"
)

type Timestamp struct {
	time.Time
}

func (self Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(self.Unix(), 10)), nil
}
func (self *Timestamp) UnmarshalJSON(b []byte) error {
	tm, err := strconv.ParseInt(string(b), 10, 0)
	if err != nil {
		return &json.UnmarshalTypeError{
			Value:  "number " + string(b),
			Type:   reflect.TypeOf(0),
			Offset: 0,
		}
	}
	self.Time = time.Unix(tm, 0)

	return nil
}

func NowTimestamp() Timestamp {
	return Timestamp{time.Now()}
}

type Timerfc struct {
	time.Time
}

func NowTimerfc() Timerfc {
	return Timerfc{time.Now()}
}

func (self Timerfc) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(self.Unix(), 10)), nil
}
func (self *Timerfc) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	self.Time, _ = time.Parse(time.RFC3339, s)
	return nil
}
