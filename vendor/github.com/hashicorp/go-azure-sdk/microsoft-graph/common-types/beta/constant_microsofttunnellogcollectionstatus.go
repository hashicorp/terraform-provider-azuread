package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelLogCollectionStatus string

const (
	MicrosoftTunnelLogCollectionStatus_Completed MicrosoftTunnelLogCollectionStatus = "completed"
	MicrosoftTunnelLogCollectionStatus_Failed    MicrosoftTunnelLogCollectionStatus = "failed"
	MicrosoftTunnelLogCollectionStatus_Pending   MicrosoftTunnelLogCollectionStatus = "pending"
)

func PossibleValuesForMicrosoftTunnelLogCollectionStatus() []string {
	return []string{
		string(MicrosoftTunnelLogCollectionStatus_Completed),
		string(MicrosoftTunnelLogCollectionStatus_Failed),
		string(MicrosoftTunnelLogCollectionStatus_Pending),
	}
}

func (s *MicrosoftTunnelLogCollectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTunnelLogCollectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTunnelLogCollectionStatus(input string) (*MicrosoftTunnelLogCollectionStatus, error) {
	vals := map[string]MicrosoftTunnelLogCollectionStatus{
		"completed": MicrosoftTunnelLogCollectionStatus_Completed,
		"failed":    MicrosoftTunnelLogCollectionStatus_Failed,
		"pending":   MicrosoftTunnelLogCollectionStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTunnelLogCollectionStatus(input)
	return &out, nil
}
