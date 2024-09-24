package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCStatus string

const (
	CloudPCStatus_Deprovisioning          CloudPCStatus = "deprovisioning"
	CloudPCStatus_Failed                  CloudPCStatus = "failed"
	CloudPCStatus_InGracePeriod           CloudPCStatus = "inGracePeriod"
	CloudPCStatus_ModifyingSingleSignOn   CloudPCStatus = "modifyingSingleSignOn"
	CloudPCStatus_MovingRegion            CloudPCStatus = "movingRegion"
	CloudPCStatus_NotProvisioned          CloudPCStatus = "notProvisioned"
	CloudPCStatus_PendingProvision        CloudPCStatus = "pendingProvision"
	CloudPCStatus_Provisioned             CloudPCStatus = "provisioned"
	CloudPCStatus_ProvisionedWithWarnings CloudPCStatus = "provisionedWithWarnings"
	CloudPCStatus_Provisioning            CloudPCStatus = "provisioning"
	CloudPCStatus_ResizePendingLicense    CloudPCStatus = "resizePendingLicense"
	CloudPCStatus_Resizing                CloudPCStatus = "resizing"
	CloudPCStatus_Restoring               CloudPCStatus = "restoring"
	CloudPCStatus_UpdatingSingleSignOn    CloudPCStatus = "updatingSingleSignOn"
)

func PossibleValuesForCloudPCStatus() []string {
	return []string{
		string(CloudPCStatus_Deprovisioning),
		string(CloudPCStatus_Failed),
		string(CloudPCStatus_InGracePeriod),
		string(CloudPCStatus_ModifyingSingleSignOn),
		string(CloudPCStatus_MovingRegion),
		string(CloudPCStatus_NotProvisioned),
		string(CloudPCStatus_PendingProvision),
		string(CloudPCStatus_Provisioned),
		string(CloudPCStatus_ProvisionedWithWarnings),
		string(CloudPCStatus_Provisioning),
		string(CloudPCStatus_ResizePendingLicense),
		string(CloudPCStatus_Resizing),
		string(CloudPCStatus_Restoring),
		string(CloudPCStatus_UpdatingSingleSignOn),
	}
}

func (s *CloudPCStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCStatus(input string) (*CloudPCStatus, error) {
	vals := map[string]CloudPCStatus{
		"deprovisioning":          CloudPCStatus_Deprovisioning,
		"failed":                  CloudPCStatus_Failed,
		"ingraceperiod":           CloudPCStatus_InGracePeriod,
		"modifyingsinglesignon":   CloudPCStatus_ModifyingSingleSignOn,
		"movingregion":            CloudPCStatus_MovingRegion,
		"notprovisioned":          CloudPCStatus_NotProvisioned,
		"pendingprovision":        CloudPCStatus_PendingProvision,
		"provisioned":             CloudPCStatus_Provisioned,
		"provisionedwithwarnings": CloudPCStatus_ProvisionedWithWarnings,
		"provisioning":            CloudPCStatus_Provisioning,
		"resizependinglicense":    CloudPCStatus_ResizePendingLicense,
		"resizing":                CloudPCStatus_Resizing,
		"restoring":               CloudPCStatus_Restoring,
		"updatingsinglesignon":    CloudPCStatus_UpdatingSingleSignOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCStatus(input)
	return &out, nil
}
