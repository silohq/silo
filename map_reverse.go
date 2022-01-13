package silo

import (
	"fmt"
)

// implementation by: ChristopherThorpe
// Link to Gist: https://gist.github.com/ChristopherThorpe/fd3720efe2ba83c929bf4105719ee967
func mapreverse(m map[string]interface{}, value interface{}, ks []string) (rval interface{}, err error) {
	var ok bool
	if len(ks) == 0 { // degenerate input
		return nil, fmt.Errorf("NestedMapLookup needs at least one key")
	}
	if rval, ok = m[ks[0]]; !ok {
		return nil, fmt.Errorf("key not found; remaining keys: %v", ks)
	} else if len(ks) == 1 { // we've reached the final key
		m[ks[0]] = value
		return rval, nil
	} else if m, ok = rval.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("malformed structure at %#v", rval)
	} else { // 1+ morekeys
		return mapreverse(m, value, ks[1:])
	}
}
