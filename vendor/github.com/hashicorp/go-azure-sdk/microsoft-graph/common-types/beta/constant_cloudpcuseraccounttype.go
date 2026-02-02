package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCUserAccountType string

const (
	CloudPCUserAccountType_Administrator CloudPCUserAccountType = "administrator"
	CloudPCUserAccountType_StandardUser  CloudPCUserAccountType = "standardUser"
)

func PossibleValuesForCloudPCUserAccountType() []string {
	return []string{
		string(CloudPCUserAccountType_Administrator),
		string(CloudPCUserAccountType_StandardUser),
	}
}

func (s *CloudPCUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCUserAccountType(input string) (*CloudPCUserAccountType, error) {
	vals := map[string]CloudPCUserAccountType{
		"administrator": CloudPCUserAccountType_Administrator,
		"standarduser":  CloudPCUserAccountType_StandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCUserAccountType(input)
	return &out, nil
}
