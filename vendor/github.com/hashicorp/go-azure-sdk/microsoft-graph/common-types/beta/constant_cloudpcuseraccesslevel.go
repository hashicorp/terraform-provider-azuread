package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCUserAccessLevel string

const (
	CloudPCUserAccessLevel_Restricted   CloudPCUserAccessLevel = "restricted"
	CloudPCUserAccessLevel_Unrestricted CloudPCUserAccessLevel = "unrestricted"
)

func PossibleValuesForCloudPCUserAccessLevel() []string {
	return []string{
		string(CloudPCUserAccessLevel_Restricted),
		string(CloudPCUserAccessLevel_Unrestricted),
	}
}

func (s *CloudPCUserAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCUserAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCUserAccessLevel(input string) (*CloudPCUserAccessLevel, error) {
	vals := map[string]CloudPCUserAccessLevel{
		"restricted":   CloudPCUserAccessLevel_Restricted,
		"unrestricted": CloudPCUserAccessLevel_Unrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCUserAccessLevel(input)
	return &out, nil
}
