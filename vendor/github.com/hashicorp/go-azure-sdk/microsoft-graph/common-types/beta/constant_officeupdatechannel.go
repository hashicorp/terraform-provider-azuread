package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeUpdateChannel string

const (
	OfficeUpdateChannel_Current              OfficeUpdateChannel = "current"
	OfficeUpdateChannel_Deferred             OfficeUpdateChannel = "deferred"
	OfficeUpdateChannel_FirstReleaseCurrent  OfficeUpdateChannel = "firstReleaseCurrent"
	OfficeUpdateChannel_FirstReleaseDeferred OfficeUpdateChannel = "firstReleaseDeferred"
	OfficeUpdateChannel_MonthlyEnterprise    OfficeUpdateChannel = "monthlyEnterprise"
	OfficeUpdateChannel_None                 OfficeUpdateChannel = "none"
)

func PossibleValuesForOfficeUpdateChannel() []string {
	return []string{
		string(OfficeUpdateChannel_Current),
		string(OfficeUpdateChannel_Deferred),
		string(OfficeUpdateChannel_FirstReleaseCurrent),
		string(OfficeUpdateChannel_FirstReleaseDeferred),
		string(OfficeUpdateChannel_MonthlyEnterprise),
		string(OfficeUpdateChannel_None),
	}
}

func (s *OfficeUpdateChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeUpdateChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeUpdateChannel(input string) (*OfficeUpdateChannel, error) {
	vals := map[string]OfficeUpdateChannel{
		"current":              OfficeUpdateChannel_Current,
		"deferred":             OfficeUpdateChannel_Deferred,
		"firstreleasecurrent":  OfficeUpdateChannel_FirstReleaseCurrent,
		"firstreleasedeferred": OfficeUpdateChannel_FirstReleaseDeferred,
		"monthlyenterprise":    OfficeUpdateChannel_MonthlyEnterprise,
		"none":                 OfficeUpdateChannel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeUpdateChannel(input)
	return &out, nil
}
