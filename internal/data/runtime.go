package data

import (
	"fmt"
	"strconv"
)

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
