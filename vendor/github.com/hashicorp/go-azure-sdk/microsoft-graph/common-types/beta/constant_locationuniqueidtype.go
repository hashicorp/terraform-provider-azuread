package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationUniqueIdType string

const (
	LocationUniqueIdType_Bing          LocationUniqueIdType = "bing"
	LocationUniqueIdType_Directory     LocationUniqueIdType = "directory"
	LocationUniqueIdType_LocationStore LocationUniqueIdType = "locationStore"
	LocationUniqueIdType_Private       LocationUniqueIdType = "private"
	LocationUniqueIdType_Unknown       LocationUniqueIdType = "unknown"
)

func PossibleValuesForLocationUniqueIdType() []string {
	return []string{
		string(LocationUniqueIdType_Bing),
		string(LocationUniqueIdType_Directory),
		string(LocationUniqueIdType_LocationStore),
		string(LocationUniqueIdType_Private),
		string(LocationUniqueIdType_Unknown),
	}
}

func (s *LocationUniqueIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationUniqueIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationUniqueIdType(input string) (*LocationUniqueIdType, error) {
	vals := map[string]LocationUniqueIdType{
		"bing":          LocationUniqueIdType_Bing,
		"directory":     LocationUniqueIdType_Directory,
		"locationstore": LocationUniqueIdType_LocationStore,
		"private":       LocationUniqueIdType_Private,
		"unknown":       LocationUniqueIdType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationUniqueIdType(input)
	return &out, nil
}
