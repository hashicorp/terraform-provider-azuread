package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendDtmfCompletionReason string

const (
	SendDtmfCompletionReason_CompletedSuccessfully  SendDtmfCompletionReason = "completedSuccessfully"
	SendDtmfCompletionReason_MediaOperationCanceled SendDtmfCompletionReason = "mediaOperationCanceled"
	SendDtmfCompletionReason_Unknown                SendDtmfCompletionReason = "unknown"
)

func PossibleValuesForSendDtmfCompletionReason() []string {
	return []string{
		string(SendDtmfCompletionReason_CompletedSuccessfully),
		string(SendDtmfCompletionReason_MediaOperationCanceled),
		string(SendDtmfCompletionReason_Unknown),
	}
}

func (s *SendDtmfCompletionReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSendDtmfCompletionReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSendDtmfCompletionReason(input string) (*SendDtmfCompletionReason, error) {
	vals := map[string]SendDtmfCompletionReason{
		"completedsuccessfully":  SendDtmfCompletionReason_CompletedSuccessfully,
		"mediaoperationcanceled": SendDtmfCompletionReason_MediaOperationCanceled,
		"unknown":                SendDtmfCompletionReason_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SendDtmfCompletionReason(input)
	return &out, nil
}
