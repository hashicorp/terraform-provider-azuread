package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityStopAndQuarantineFileEntityIdentifier string

const (
	SecurityStopAndQuarantineFileEntityIdentifier_DeviceId              SecurityStopAndQuarantineFileEntityIdentifier = "deviceId"
	SecurityStopAndQuarantineFileEntityIdentifier_InitiatingProcessSHA1 SecurityStopAndQuarantineFileEntityIdentifier = "initiatingProcessSHA1"
	SecurityStopAndQuarantineFileEntityIdentifier_Sha1                  SecurityStopAndQuarantineFileEntityIdentifier = "sha1"
)

func PossibleValuesForSecurityStopAndQuarantineFileEntityIdentifier() []string {
	return []string{
		string(SecurityStopAndQuarantineFileEntityIdentifier_DeviceId),
		string(SecurityStopAndQuarantineFileEntityIdentifier_InitiatingProcessSHA1),
		string(SecurityStopAndQuarantineFileEntityIdentifier_Sha1),
	}
}

func (s *SecurityStopAndQuarantineFileEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityStopAndQuarantineFileEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityStopAndQuarantineFileEntityIdentifier(input string) (*SecurityStopAndQuarantineFileEntityIdentifier, error) {
	vals := map[string]SecurityStopAndQuarantineFileEntityIdentifier{
		"deviceid":              SecurityStopAndQuarantineFileEntityIdentifier_DeviceId,
		"initiatingprocesssha1": SecurityStopAndQuarantineFileEntityIdentifier_InitiatingProcessSHA1,
		"sha1":                  SecurityStopAndQuarantineFileEntityIdentifier_Sha1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityStopAndQuarantineFileEntityIdentifier(input)
	return &out, nil
}
