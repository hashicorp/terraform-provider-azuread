package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenActionFailureReason string

const (
	VppTokenActionFailureReason_AppleFailure                            VppTokenActionFailureReason = "appleFailure"
	VppTokenActionFailureReason_ExpiredApplePushNotificationCertificate VppTokenActionFailureReason = "expiredApplePushNotificationCertificate"
	VppTokenActionFailureReason_ExpiredVppToken                         VppTokenActionFailureReason = "expiredVppToken"
	VppTokenActionFailureReason_InternalError                           VppTokenActionFailureReason = "internalError"
	VppTokenActionFailureReason_None                                    VppTokenActionFailureReason = "none"
)

func PossibleValuesForVppTokenActionFailureReason() []string {
	return []string{
		string(VppTokenActionFailureReason_AppleFailure),
		string(VppTokenActionFailureReason_ExpiredApplePushNotificationCertificate),
		string(VppTokenActionFailureReason_ExpiredVppToken),
		string(VppTokenActionFailureReason_InternalError),
		string(VppTokenActionFailureReason_None),
	}
}

func (s *VppTokenActionFailureReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenActionFailureReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenActionFailureReason(input string) (*VppTokenActionFailureReason, error) {
	vals := map[string]VppTokenActionFailureReason{
		"applefailure": VppTokenActionFailureReason_AppleFailure,
		"expiredapplepushnotificationcertificate": VppTokenActionFailureReason_ExpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         VppTokenActionFailureReason_ExpiredVppToken,
		"internalerror":                           VppTokenActionFailureReason_InternalError,
		"none":                                    VppTokenActionFailureReason_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenActionFailureReason(input)
	return &out, nil
}
