package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryLegalHoldStatus string

const (
	EdiscoveryLegalHoldStatus_Error   EdiscoveryLegalHoldStatus = "Error"
	EdiscoveryLegalHoldStatus_Pending EdiscoveryLegalHoldStatus = "Pending"
	EdiscoveryLegalHoldStatus_Success EdiscoveryLegalHoldStatus = "Success"
)

func PossibleValuesForEdiscoveryLegalHoldStatus() []string {
	return []string{
		string(EdiscoveryLegalHoldStatus_Error),
		string(EdiscoveryLegalHoldStatus_Pending),
		string(EdiscoveryLegalHoldStatus_Success),
	}
}

func (s *EdiscoveryLegalHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryLegalHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryLegalHoldStatus(input string) (*EdiscoveryLegalHoldStatus, error) {
	vals := map[string]EdiscoveryLegalHoldStatus{
		"error":   EdiscoveryLegalHoldStatus_Error,
		"pending": EdiscoveryLegalHoldStatus_Pending,
		"success": EdiscoveryLegalHoldStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryLegalHoldStatus(input)
	return &out, nil
}
