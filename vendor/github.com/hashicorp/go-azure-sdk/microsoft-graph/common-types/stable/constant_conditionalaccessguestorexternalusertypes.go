package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGuestOrExternalUserTypes string

const (
	ConditionalAccessGuestOrExternalUserTypes_B2bCollaborationGuest  ConditionalAccessGuestOrExternalUserTypes = "b2bCollaborationGuest"
	ConditionalAccessGuestOrExternalUserTypes_B2bCollaborationMember ConditionalAccessGuestOrExternalUserTypes = "b2bCollaborationMember"
	ConditionalAccessGuestOrExternalUserTypes_B2bDirectConnectUser   ConditionalAccessGuestOrExternalUserTypes = "b2bDirectConnectUser"
	ConditionalAccessGuestOrExternalUserTypes_InternalGuest          ConditionalAccessGuestOrExternalUserTypes = "internalGuest"
	ConditionalAccessGuestOrExternalUserTypes_None                   ConditionalAccessGuestOrExternalUserTypes = "none"
	ConditionalAccessGuestOrExternalUserTypes_OtherExternalUser      ConditionalAccessGuestOrExternalUserTypes = "otherExternalUser"
	ConditionalAccessGuestOrExternalUserTypes_ServiceProvider        ConditionalAccessGuestOrExternalUserTypes = "serviceProvider"
)

func PossibleValuesForConditionalAccessGuestOrExternalUserTypes() []string {
	return []string{
		string(ConditionalAccessGuestOrExternalUserTypes_B2bCollaborationGuest),
		string(ConditionalAccessGuestOrExternalUserTypes_B2bCollaborationMember),
		string(ConditionalAccessGuestOrExternalUserTypes_B2bDirectConnectUser),
		string(ConditionalAccessGuestOrExternalUserTypes_InternalGuest),
		string(ConditionalAccessGuestOrExternalUserTypes_None),
		string(ConditionalAccessGuestOrExternalUserTypes_OtherExternalUser),
		string(ConditionalAccessGuestOrExternalUserTypes_ServiceProvider),
	}
}

func (s *ConditionalAccessGuestOrExternalUserTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessGuestOrExternalUserTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessGuestOrExternalUserTypes(input string) (*ConditionalAccessGuestOrExternalUserTypes, error) {
	vals := map[string]ConditionalAccessGuestOrExternalUserTypes{
		"b2bcollaborationguest":  ConditionalAccessGuestOrExternalUserTypes_B2bCollaborationGuest,
		"b2bcollaborationmember": ConditionalAccessGuestOrExternalUserTypes_B2bCollaborationMember,
		"b2bdirectconnectuser":   ConditionalAccessGuestOrExternalUserTypes_B2bDirectConnectUser,
		"internalguest":          ConditionalAccessGuestOrExternalUserTypes_InternalGuest,
		"none":                   ConditionalAccessGuestOrExternalUserTypes_None,
		"otherexternaluser":      ConditionalAccessGuestOrExternalUserTypes_OtherExternalUser,
		"serviceprovider":        ConditionalAccessGuestOrExternalUserTypes_ServiceProvider,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessGuestOrExternalUserTypes(input)
	return &out, nil
}
