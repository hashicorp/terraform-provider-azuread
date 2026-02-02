package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteArchiveStatus string

const (
	SiteArchiveStatus_FullyArchived    SiteArchiveStatus = "fullyArchived"
	SiteArchiveStatus_Reactivating     SiteArchiveStatus = "reactivating"
	SiteArchiveStatus_RecentlyArchived SiteArchiveStatus = "recentlyArchived"
)

func PossibleValuesForSiteArchiveStatus() []string {
	return []string{
		string(SiteArchiveStatus_FullyArchived),
		string(SiteArchiveStatus_Reactivating),
		string(SiteArchiveStatus_RecentlyArchived),
	}
}

func (s *SiteArchiveStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteArchiveStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteArchiveStatus(input string) (*SiteArchiveStatus, error) {
	vals := map[string]SiteArchiveStatus{
		"fullyarchived":    SiteArchiveStatus_FullyArchived,
		"reactivating":     SiteArchiveStatus_Reactivating,
		"recentlyarchived": SiteArchiveStatus_RecentlyArchived,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteArchiveStatus(input)
	return &out, nil
}
