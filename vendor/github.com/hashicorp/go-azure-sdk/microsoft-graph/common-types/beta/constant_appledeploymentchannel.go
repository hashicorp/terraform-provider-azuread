package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleDeploymentChannel string

const (
	AppleDeploymentChannel_DeviceChannel AppleDeploymentChannel = "deviceChannel"
	AppleDeploymentChannel_UserChannel   AppleDeploymentChannel = "userChannel"
)

func PossibleValuesForAppleDeploymentChannel() []string {
	return []string{
		string(AppleDeploymentChannel_DeviceChannel),
		string(AppleDeploymentChannel_UserChannel),
	}
}

func (s *AppleDeploymentChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleDeploymentChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleDeploymentChannel(input string) (*AppleDeploymentChannel, error) {
	vals := map[string]AppleDeploymentChannel{
		"devicechannel": AppleDeploymentChannel_DeviceChannel,
		"userchannel":   AppleDeploymentChannel_UserChannel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleDeploymentChannel(input)
	return &out, nil
}
