package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

/*
We are implementing a MarshalJSON method on Runtime so that it satisfies the json.Marshaler interface.
This let's us customaize the json field, like {"runtime": "118 mins"}.
No need to modify r, that's why Runtime and not *Runtime

refer ch 3.5, pg 62 (Advanced JSON customization)
'When Go is encoding a particular type to JSON, it looks to see if the type has a MarshalJSON()
method implemented on it. If it has, then Go will call this method to determine how to encode
it.'
And we are basically creating a new type and implementing that method
*/
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// wrapping it in double quotes, to use it in JSON
	quotedValue := strconv.Quote(jsonValue)
	return []byte(quotedValue), nil
}

// *Runtime cause we need to modify it. int -> string
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJsonValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// parsing it to int type
	l, err := strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(l)
	return nil
}
