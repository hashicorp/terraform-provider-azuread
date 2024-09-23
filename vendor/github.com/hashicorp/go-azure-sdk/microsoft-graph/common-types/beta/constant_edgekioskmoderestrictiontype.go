package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeKioskModeRestrictionType string

const (
	EdgeKioskModeRestrictionType_DigitalSignage          EdgeKioskModeRestrictionType = "digitalSignage"
	EdgeKioskModeRestrictionType_NormalMode              EdgeKioskModeRestrictionType = "normalMode"
	EdgeKioskModeRestrictionType_NotConfigured           EdgeKioskModeRestrictionType = "notConfigured"
	EdgeKioskModeRestrictionType_PublicBrowsingMultiApp  EdgeKioskModeRestrictionType = "publicBrowsingMultiApp"
	EdgeKioskModeRestrictionType_PublicBrowsingSingleApp EdgeKioskModeRestrictionType = "publicBrowsingSingleApp"
)

func PossibleValuesForEdgeKioskModeRestrictionType() []string {
	return []string{
		string(EdgeKioskModeRestrictionType_DigitalSignage),
		string(EdgeKioskModeRestrictionType_NormalMode),
		string(EdgeKioskModeRestrictionType_NotConfigured),
		string(EdgeKioskModeRestrictionType_PublicBrowsingMultiApp),
		string(EdgeKioskModeRestrictionType_PublicBrowsingSingleApp),
	}
}

func (s *EdgeKioskModeRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdgeKioskModeRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdgeKioskModeRestrictionType(input string) (*EdgeKioskModeRestrictionType, error) {
	vals := map[string]EdgeKioskModeRestrictionType{
		"digitalsignage":          EdgeKioskModeRestrictionType_DigitalSignage,
		"normalmode":              EdgeKioskModeRestrictionType_NormalMode,
		"notconfigured":           EdgeKioskModeRestrictionType_NotConfigured,
		"publicbrowsingmultiapp":  EdgeKioskModeRestrictionType_PublicBrowsingMultiApp,
		"publicbrowsingsingleapp": EdgeKioskModeRestrictionType_PublicBrowsingSingleApp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdgeKioskModeRestrictionType(input)
	return &out, nil
}
