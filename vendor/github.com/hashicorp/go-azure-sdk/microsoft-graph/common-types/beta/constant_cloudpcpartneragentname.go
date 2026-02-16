package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentName string

const (
	CloudPCPartnerAgentName_Citrix CloudPCPartnerAgentName = "citrix"
	CloudPCPartnerAgentName_Hp     CloudPCPartnerAgentName = "hp"
	CloudPCPartnerAgentName_VMware CloudPCPartnerAgentName = "vMware"
)

func PossibleValuesForCloudPCPartnerAgentName() []string {
	return []string{
		string(CloudPCPartnerAgentName_Citrix),
		string(CloudPCPartnerAgentName_Hp),
		string(CloudPCPartnerAgentName_VMware),
	}
}

func (s *CloudPCPartnerAgentName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPartnerAgentName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPartnerAgentName(input string) (*CloudPCPartnerAgentName, error) {
	vals := map[string]CloudPCPartnerAgentName{
		"citrix": CloudPCPartnerAgentName_Citrix,
		"hp":     CloudPCPartnerAgentName_Hp,
		"vmware": CloudPCPartnerAgentName_VMware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentName(input)
	return &out, nil
}
