package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingMetadata string

const (
	ObjectMappingMetadata_DisableMonitoringForChanges ObjectMappingMetadata = "DisableMonitoringForChanges"
	ObjectMappingMetadata_Disposition                 ObjectMappingMetadata = "Disposition"
	ObjectMappingMetadata_EscrowBehavior              ObjectMappingMetadata = "EscrowBehavior"
	ObjectMappingMetadata_ExcludeFromReporting        ObjectMappingMetadata = "ExcludeFromReporting"
	ObjectMappingMetadata_IsCustomerDefined           ObjectMappingMetadata = "IsCustomerDefined"
	ObjectMappingMetadata_OriginalJoiningProperty     ObjectMappingMetadata = "OriginalJoiningProperty"
	ObjectMappingMetadata_Unsynchronized              ObjectMappingMetadata = "Unsynchronized"
)

func PossibleValuesForObjectMappingMetadata() []string {
	return []string{
		string(ObjectMappingMetadata_DisableMonitoringForChanges),
		string(ObjectMappingMetadata_Disposition),
		string(ObjectMappingMetadata_EscrowBehavior),
		string(ObjectMappingMetadata_ExcludeFromReporting),
		string(ObjectMappingMetadata_IsCustomerDefined),
		string(ObjectMappingMetadata_OriginalJoiningProperty),
		string(ObjectMappingMetadata_Unsynchronized),
	}
}

func (s *ObjectMappingMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectMappingMetadata(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectMappingMetadata(input string) (*ObjectMappingMetadata, error) {
	vals := map[string]ObjectMappingMetadata{
		"disablemonitoringforchanges": ObjectMappingMetadata_DisableMonitoringForChanges,
		"disposition":                 ObjectMappingMetadata_Disposition,
		"escrowbehavior":              ObjectMappingMetadata_EscrowBehavior,
		"excludefromreporting":        ObjectMappingMetadata_ExcludeFromReporting,
		"iscustomerdefined":           ObjectMappingMetadata_IsCustomerDefined,
		"originaljoiningproperty":     ObjectMappingMetadata_OriginalJoiningProperty,
		"unsynchronized":              ObjectMappingMetadata_Unsynchronized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingMetadata(input)
	return &out, nil
}
