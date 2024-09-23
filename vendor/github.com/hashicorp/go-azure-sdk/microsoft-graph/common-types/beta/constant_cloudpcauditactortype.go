package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditActorType string

const (
	CloudPCAuditActorType_Application CloudPCAuditActorType = "application"
	CloudPCAuditActorType_ItPro       CloudPCAuditActorType = "itPro"
	CloudPCAuditActorType_Partner     CloudPCAuditActorType = "partner"
)

func PossibleValuesForCloudPCAuditActorType() []string {
	return []string{
		string(CloudPCAuditActorType_Application),
		string(CloudPCAuditActorType_ItPro),
		string(CloudPCAuditActorType_Partner),
	}
}

func (s *CloudPCAuditActorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditActorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditActorType(input string) (*CloudPCAuditActorType, error) {
	vals := map[string]CloudPCAuditActorType{
		"application": CloudPCAuditActorType_Application,
		"itpro":       CloudPCAuditActorType_ItPro,
		"partner":     CloudPCAuditActorType_Partner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditActorType(input)
	return &out, nil
}
