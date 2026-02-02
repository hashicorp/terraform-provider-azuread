package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MfaFailureReasonCode string

const (
	MfaFailureReasonCode_BadRequest    MfaFailureReasonCode = "badRequest"
	MfaFailureReasonCode_MfaDenied     MfaFailureReasonCode = "mfaDenied"
	MfaFailureReasonCode_MfaIncomplete MfaFailureReasonCode = "mfaIncomplete"
	MfaFailureReasonCode_Other         MfaFailureReasonCode = "other"
	MfaFailureReasonCode_SystemFailure MfaFailureReasonCode = "systemFailure"
)

func PossibleValuesForMfaFailureReasonCode() []string {
	return []string{
		string(MfaFailureReasonCode_BadRequest),
		string(MfaFailureReasonCode_MfaDenied),
		string(MfaFailureReasonCode_MfaIncomplete),
		string(MfaFailureReasonCode_Other),
		string(MfaFailureReasonCode_SystemFailure),
	}
}

func (s *MfaFailureReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMfaFailureReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMfaFailureReasonCode(input string) (*MfaFailureReasonCode, error) {
	vals := map[string]MfaFailureReasonCode{
		"badrequest":    MfaFailureReasonCode_BadRequest,
		"mfadenied":     MfaFailureReasonCode_MfaDenied,
		"mfaincomplete": MfaFailureReasonCode_MfaIncomplete,
		"other":         MfaFailureReasonCode_Other,
		"systemfailure": MfaFailureReasonCode_SystemFailure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MfaFailureReasonCode(input)
	return &out, nil
}
