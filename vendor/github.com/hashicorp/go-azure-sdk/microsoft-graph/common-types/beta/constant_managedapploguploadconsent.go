package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppLogUploadConsent string

const (
	ManagedAppLogUploadConsent_Accepted ManagedAppLogUploadConsent = "accepted"
	ManagedAppLogUploadConsent_Declined ManagedAppLogUploadConsent = "declined"
	ManagedAppLogUploadConsent_Unknown  ManagedAppLogUploadConsent = "unknown"
)

func PossibleValuesForManagedAppLogUploadConsent() []string {
	return []string{
		string(ManagedAppLogUploadConsent_Accepted),
		string(ManagedAppLogUploadConsent_Declined),
		string(ManagedAppLogUploadConsent_Unknown),
	}
}

func (s *ManagedAppLogUploadConsent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppLogUploadConsent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppLogUploadConsent(input string) (*ManagedAppLogUploadConsent, error) {
	vals := map[string]ManagedAppLogUploadConsent{
		"accepted": ManagedAppLogUploadConsent_Accepted,
		"declined": ManagedAppLogUploadConsent_Declined,
		"unknown":  ManagedAppLogUploadConsent_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppLogUploadConsent(input)
	return &out, nil
}
