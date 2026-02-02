package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CampaignStatus string

const (
	CampaignStatus_Cancelled  CampaignStatus = "cancelled"
	CampaignStatus_Completed  CampaignStatus = "completed"
	CampaignStatus_Deleted    CampaignStatus = "deleted"
	CampaignStatus_Draft      CampaignStatus = "draft"
	CampaignStatus_Excluded   CampaignStatus = "excluded"
	CampaignStatus_Failed     CampaignStatus = "failed"
	CampaignStatus_InProgress CampaignStatus = "inProgress"
	CampaignStatus_Scheduled  CampaignStatus = "scheduled"
	CampaignStatus_Unknown    CampaignStatus = "unknown"
)

func PossibleValuesForCampaignStatus() []string {
	return []string{
		string(CampaignStatus_Cancelled),
		string(CampaignStatus_Completed),
		string(CampaignStatus_Deleted),
		string(CampaignStatus_Draft),
		string(CampaignStatus_Excluded),
		string(CampaignStatus_Failed),
		string(CampaignStatus_InProgress),
		string(CampaignStatus_Scheduled),
		string(CampaignStatus_Unknown),
	}
}

func (s *CampaignStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCampaignStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCampaignStatus(input string) (*CampaignStatus, error) {
	vals := map[string]CampaignStatus{
		"cancelled":  CampaignStatus_Cancelled,
		"completed":  CampaignStatus_Completed,
		"deleted":    CampaignStatus_Deleted,
		"draft":      CampaignStatus_Draft,
		"excluded":   CampaignStatus_Excluded,
		"failed":     CampaignStatus_Failed,
		"inprogress": CampaignStatus_InProgress,
		"scheduled":  CampaignStatus_Scheduled,
		"unknown":    CampaignStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CampaignStatus(input)
	return &out, nil
}
