package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieStatus string

const (
	BrowserSharedCookieStatus_PendingAdd    BrowserSharedCookieStatus = "pendingAdd"
	BrowserSharedCookieStatus_PendingDelete BrowserSharedCookieStatus = "pendingDelete"
	BrowserSharedCookieStatus_PendingEdit   BrowserSharedCookieStatus = "pendingEdit"
	BrowserSharedCookieStatus_Published     BrowserSharedCookieStatus = "published"
)

func PossibleValuesForBrowserSharedCookieStatus() []string {
	return []string{
		string(BrowserSharedCookieStatus_PendingAdd),
		string(BrowserSharedCookieStatus_PendingDelete),
		string(BrowserSharedCookieStatus_PendingEdit),
		string(BrowserSharedCookieStatus_Published),
	}
}

func (s *BrowserSharedCookieStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSharedCookieStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSharedCookieStatus(input string) (*BrowserSharedCookieStatus, error) {
	vals := map[string]BrowserSharedCookieStatus{
		"pendingadd":    BrowserSharedCookieStatus_PendingAdd,
		"pendingdelete": BrowserSharedCookieStatus_PendingDelete,
		"pendingedit":   BrowserSharedCookieStatus_PendingEdit,
		"published":     BrowserSharedCookieStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSharedCookieStatus(input)
	return &out, nil
}
