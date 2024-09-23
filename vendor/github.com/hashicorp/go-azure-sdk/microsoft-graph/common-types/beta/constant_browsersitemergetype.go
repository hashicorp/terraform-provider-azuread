package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteMergeType string

const (
	BrowserSiteMergeType_Default BrowserSiteMergeType = "default"
	BrowserSiteMergeType_NoMerge BrowserSiteMergeType = "noMerge"
)

func PossibleValuesForBrowserSiteMergeType() []string {
	return []string{
		string(BrowserSiteMergeType_Default),
		string(BrowserSiteMergeType_NoMerge),
	}
}

func (s *BrowserSiteMergeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteMergeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteMergeType(input string) (*BrowserSiteMergeType, error) {
	vals := map[string]BrowserSiteMergeType{
		"default": BrowserSiteMergeType_Default,
		"nomerge": BrowserSiteMergeType_NoMerge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteMergeType(input)
	return &out, nil
}
