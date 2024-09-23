package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictionTrigger string

const (
	RestrictionTrigger_CloudEgress          RestrictionTrigger = "cloudEgress"
	RestrictionTrigger_CopyPaste            RestrictionTrigger = "copyPaste"
	RestrictionTrigger_CopyToNetworkShare   RestrictionTrigger = "copyToNetworkShare"
	RestrictionTrigger_CopyToRemovableMedia RestrictionTrigger = "copyToRemovableMedia"
	RestrictionTrigger_Print                RestrictionTrigger = "print"
	RestrictionTrigger_ScreenCapture        RestrictionTrigger = "screenCapture"
	RestrictionTrigger_UnallowedApps        RestrictionTrigger = "unallowedApps"
)

func PossibleValuesForRestrictionTrigger() []string {
	return []string{
		string(RestrictionTrigger_CloudEgress),
		string(RestrictionTrigger_CopyPaste),
		string(RestrictionTrigger_CopyToNetworkShare),
		string(RestrictionTrigger_CopyToRemovableMedia),
		string(RestrictionTrigger_Print),
		string(RestrictionTrigger_ScreenCapture),
		string(RestrictionTrigger_UnallowedApps),
	}
}

func (s *RestrictionTrigger) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestrictionTrigger(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestrictionTrigger(input string) (*RestrictionTrigger, error) {
	vals := map[string]RestrictionTrigger{
		"cloudegress":          RestrictionTrigger_CloudEgress,
		"copypaste":            RestrictionTrigger_CopyPaste,
		"copytonetworkshare":   RestrictionTrigger_CopyToNetworkShare,
		"copytoremovablemedia": RestrictionTrigger_CopyToRemovableMedia,
		"print":                RestrictionTrigger_Print,
		"screencapture":        RestrictionTrigger_ScreenCapture,
		"unallowedapps":        RestrictionTrigger_UnallowedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictionTrigger(input)
	return &out, nil
}
