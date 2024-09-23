package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteStatus string

const (
	BrowserSiteStatus_PendingAdd    BrowserSiteStatus = "pendingAdd"
	BrowserSiteStatus_PendingDelete BrowserSiteStatus = "pendingDelete"
	BrowserSiteStatus_PendingEdit   BrowserSiteStatus = "pendingEdit"
	BrowserSiteStatus_Published     BrowserSiteStatus = "published"
)

func PossibleValuesForBrowserSiteStatus() []string {
	return []string{
		string(BrowserSiteStatus_PendingAdd),
		string(BrowserSiteStatus_PendingDelete),
		string(BrowserSiteStatus_PendingEdit),
		string(BrowserSiteStatus_Published),
	}
}

func (s *BrowserSiteStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteStatus(input string) (*BrowserSiteStatus, error) {
	vals := map[string]BrowserSiteStatus{
		"pendingadd":    BrowserSiteStatus_PendingAdd,
		"pendingdelete": BrowserSiteStatus_PendingDelete,
		"pendingedit":   BrowserSiteStatus_PendingEdit,
		"published":     BrowserSiteStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteStatus(input)
	return &out, nil
}
