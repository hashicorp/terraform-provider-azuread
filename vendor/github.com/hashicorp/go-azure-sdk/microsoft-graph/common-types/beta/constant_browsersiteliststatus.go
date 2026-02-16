package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteListStatus string

const (
	BrowserSiteListStatus_Draft     BrowserSiteListStatus = "draft"
	BrowserSiteListStatus_Pending   BrowserSiteListStatus = "pending"
	BrowserSiteListStatus_Published BrowserSiteListStatus = "published"
)

func PossibleValuesForBrowserSiteListStatus() []string {
	return []string{
		string(BrowserSiteListStatus_Draft),
		string(BrowserSiteListStatus_Pending),
		string(BrowserSiteListStatus_Published),
	}
}

func (s *BrowserSiteListStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteListStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteListStatus(input string) (*BrowserSiteListStatus, error) {
	vals := map[string]BrowserSiteListStatus{
		"draft":     BrowserSiteListStatus_Draft,
		"pending":   BrowserSiteListStatus_Pending,
		"published": BrowserSiteListStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteListStatus(input)
	return &out, nil
}
