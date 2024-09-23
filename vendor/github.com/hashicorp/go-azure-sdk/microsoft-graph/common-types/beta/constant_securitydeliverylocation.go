package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeliveryLocation string

const (
	SecurityDeliveryLocation_DeletedFolder  SecurityDeliveryLocation = "deletedFolder"
	SecurityDeliveryLocation_Dropped        SecurityDeliveryLocation = "dropped"
	SecurityDeliveryLocation_Failed         SecurityDeliveryLocation = "failed"
	SecurityDeliveryLocation_Inboxfolder    SecurityDeliveryLocation = "inbox_folder"
	SecurityDeliveryLocation_JunkFolder     SecurityDeliveryLocation = "junkFolder"
	SecurityDeliveryLocation_Onpremexternal SecurityDeliveryLocation = "onprem_external"
	SecurityDeliveryLocation_Others         SecurityDeliveryLocation = "others"
	SecurityDeliveryLocation_Quarantine     SecurityDeliveryLocation = "quarantine"
	SecurityDeliveryLocation_Unknown        SecurityDeliveryLocation = "unknown"
)

func PossibleValuesForSecurityDeliveryLocation() []string {
	return []string{
		string(SecurityDeliveryLocation_DeletedFolder),
		string(SecurityDeliveryLocation_Dropped),
		string(SecurityDeliveryLocation_Failed),
		string(SecurityDeliveryLocation_Inboxfolder),
		string(SecurityDeliveryLocation_JunkFolder),
		string(SecurityDeliveryLocation_Onpremexternal),
		string(SecurityDeliveryLocation_Others),
		string(SecurityDeliveryLocation_Quarantine),
		string(SecurityDeliveryLocation_Unknown),
	}
}

func (s *SecurityDeliveryLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeliveryLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeliveryLocation(input string) (*SecurityDeliveryLocation, error) {
	vals := map[string]SecurityDeliveryLocation{
		"deletedfolder":   SecurityDeliveryLocation_DeletedFolder,
		"dropped":         SecurityDeliveryLocation_Dropped,
		"failed":          SecurityDeliveryLocation_Failed,
		"inbox_folder":    SecurityDeliveryLocation_Inboxfolder,
		"junkfolder":      SecurityDeliveryLocation_JunkFolder,
		"onprem_external": SecurityDeliveryLocation_Onpremexternal,
		"others":          SecurityDeliveryLocation_Others,
		"quarantine":      SecurityDeliveryLocation_Quarantine,
		"unknown":         SecurityDeliveryLocation_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeliveryLocation(input)
	return &out, nil
}
