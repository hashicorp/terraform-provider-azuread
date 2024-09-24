package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingCapabilities string

const (
	SharingCapabilities_Disabled                        SharingCapabilities = "disabled"
	SharingCapabilities_ExistingExternalUserSharingOnly SharingCapabilities = "existingExternalUserSharingOnly"
	SharingCapabilities_ExternalUserAndGuestSharing     SharingCapabilities = "externalUserAndGuestSharing"
	SharingCapabilities_ExternalUserSharingOnly         SharingCapabilities = "externalUserSharingOnly"
)

func PossibleValuesForSharingCapabilities() []string {
	return []string{
		string(SharingCapabilities_Disabled),
		string(SharingCapabilities_ExistingExternalUserSharingOnly),
		string(SharingCapabilities_ExternalUserAndGuestSharing),
		string(SharingCapabilities_ExternalUserSharingOnly),
	}
}

func (s *SharingCapabilities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharingCapabilities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharingCapabilities(input string) (*SharingCapabilities, error) {
	vals := map[string]SharingCapabilities{
		"disabled":                        SharingCapabilities_Disabled,
		"existingexternalusersharingonly": SharingCapabilities_ExistingExternalUserSharingOnly,
		"externaluserandguestsharing":     SharingCapabilities_ExternalUserAndGuestSharing,
		"externalusersharingonly":         SharingCapabilities_ExternalUserSharingOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharingCapabilities(input)
	return &out, nil
}
