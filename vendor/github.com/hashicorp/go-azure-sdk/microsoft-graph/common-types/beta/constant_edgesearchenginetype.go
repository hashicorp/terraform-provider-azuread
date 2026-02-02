package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeSearchEngineType string

const (
	EdgeSearchEngineType_Bing    EdgeSearchEngineType = "bing"
	EdgeSearchEngineType_Default EdgeSearchEngineType = "default"
)

func PossibleValuesForEdgeSearchEngineType() []string {
	return []string{
		string(EdgeSearchEngineType_Bing),
		string(EdgeSearchEngineType_Default),
	}
}

func (s *EdgeSearchEngineType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdgeSearchEngineType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdgeSearchEngineType(input string) (*EdgeSearchEngineType, error) {
	vals := map[string]EdgeSearchEngineType{
		"bing":    EdgeSearchEngineType_Bing,
		"default": EdgeSearchEngineType_Default,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdgeSearchEngineType(input)
	return &out, nil
}
