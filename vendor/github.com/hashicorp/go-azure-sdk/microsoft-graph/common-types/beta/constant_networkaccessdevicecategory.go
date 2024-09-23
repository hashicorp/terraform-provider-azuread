package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceCategory string

const (
	NetworkaccessDeviceCategory_Branch        NetworkaccessDeviceCategory = "branch"
	NetworkaccessDeviceCategory_Client        NetworkaccessDeviceCategory = "client"
	NetworkaccessDeviceCategory_RemoteNetwork NetworkaccessDeviceCategory = "remoteNetwork"
)

func PossibleValuesForNetworkaccessDeviceCategory() []string {
	return []string{
		string(NetworkaccessDeviceCategory_Branch),
		string(NetworkaccessDeviceCategory_Client),
		string(NetworkaccessDeviceCategory_RemoteNetwork),
	}
}

func (s *NetworkaccessDeviceCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDeviceCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDeviceCategory(input string) (*NetworkaccessDeviceCategory, error) {
	vals := map[string]NetworkaccessDeviceCategory{
		"branch":        NetworkaccessDeviceCategory_Branch,
		"client":        NetworkaccessDeviceCategory_Client,
		"remotenetwork": NetworkaccessDeviceCategory_RemoteNetwork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDeviceCategory(input)
	return &out, nil
}
