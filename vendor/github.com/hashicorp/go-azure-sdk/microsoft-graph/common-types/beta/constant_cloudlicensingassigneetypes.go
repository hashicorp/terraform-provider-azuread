package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudLicensingAssigneeTypes string

const (
	CloudLicensingAssigneeTypes_Device CloudLicensingAssigneeTypes = "device"
	CloudLicensingAssigneeTypes_Group  CloudLicensingAssigneeTypes = "group"
	CloudLicensingAssigneeTypes_None   CloudLicensingAssigneeTypes = "none"
	CloudLicensingAssigneeTypes_User   CloudLicensingAssigneeTypes = "user"
)

func PossibleValuesForCloudLicensingAssigneeTypes() []string {
	return []string{
		string(CloudLicensingAssigneeTypes_Device),
		string(CloudLicensingAssigneeTypes_Group),
		string(CloudLicensingAssigneeTypes_None),
		string(CloudLicensingAssigneeTypes_User),
	}
}

func (s *CloudLicensingAssigneeTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudLicensingAssigneeTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudLicensingAssigneeTypes(input string) (*CloudLicensingAssigneeTypes, error) {
	vals := map[string]CloudLicensingAssigneeTypes{
		"device": CloudLicensingAssigneeTypes_Device,
		"group":  CloudLicensingAssigneeTypes_Group,
		"none":   CloudLicensingAssigneeTypes_None,
		"user":   CloudLicensingAssigneeTypes_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudLicensingAssigneeTypes(input)
	return &out, nil
}
