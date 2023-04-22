package custom

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalDate marshals time into gql Date type
func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(`"` + t.Format(time.RFC3339) + `"`))
	})
}

// UnmarshalDate unmarshals time from gql Date type
func UnmarshalDate(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	case int:
		return time.Unix(0, int64(v)*1000000), nil
	default:
		return time.Time{}, fmt.Errorf("%T is not a valid time", v)
	}
}
