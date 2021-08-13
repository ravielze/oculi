package auth

import (
	"encoding/json"
	"reflect"

	"github.com/danhper/structomap"
)

type StandardCredentials struct {
	ID       uint64      `json:"id"`
	Metadata interface{} `json:"metadata"`
}

// If metadata is a struct, return snakecased key and its value.
// The key needs to be exported.
//
// If metadata is a map, it will be converted to map[string]interface{}.
//
// If metadata is not a struct or map, it will return a map with single key called "metadata" and put its value on there.
func (s *StandardCredentials) MapMetadata() (map[string]interface{}, error) {
	emptyMap := map[string]interface{}{}
	if s.Metadata == nil {
		return emptyMap, nil
	}
	metadataType := reflect.ValueOf(s.Metadata).Kind()
	if metadataType == reflect.Struct {
		base := structomap.New().UseSnakeCase().PickAll().Transform(s.Metadata)
		return base, nil
	} else if metadataType == reflect.Map {
		bytes, err := json.Marshal(s.Metadata)
		if err != nil {
			return emptyMap, err
		}
		var x map[string]interface{}
		err = json.Unmarshal(bytes, &x)
		if err != nil {
			return emptyMap, err
		}
		return x, nil
	}
	return map[string]interface{}{"metadata": s.Metadata}, nil
}
