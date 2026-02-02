package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCustomExtensionHandlerStatus string

const (
	AccessPackageCustomExtensionHandlerStatus_RequestReceived AccessPackageCustomExtensionHandlerStatus = "requestReceived"
	AccessPackageCustomExtensionHandlerStatus_RequestSent     AccessPackageCustomExtensionHandlerStatus = "requestSent"
)

func PossibleValuesForAccessPackageCustomExtensionHandlerStatus() []string {
	return []string{
		string(AccessPackageCustomExtensionHandlerStatus_RequestReceived),
		string(AccessPackageCustomExtensionHandlerStatus_RequestSent),
	}
}

func (s *AccessPackageCustomExtensionHandlerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageCustomExtensionHandlerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageCustomExtensionHandlerStatus(input string) (*AccessPackageCustomExtensionHandlerStatus, error) {
	vals := map[string]AccessPackageCustomExtensionHandlerStatus{
		"requestreceived": AccessPackageCustomExtensionHandlerStatus_RequestReceived,
		"requestsent":     AccessPackageCustomExtensionHandlerStatus_RequestSent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCustomExtensionHandlerStatus(input)
	return &out, nil
}
