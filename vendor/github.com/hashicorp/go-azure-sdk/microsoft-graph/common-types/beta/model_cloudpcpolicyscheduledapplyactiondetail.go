package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCPolicyScheduledApplyActionDetail{}

type CloudPCPolicyScheduledApplyActionDetail struct {
	// An expression that specifies the cron schedule. (For example, '0 0 0 20 ' means schedules a job to run at midnight on
	// the 20th of every month) Administrators can set a cron expression to define the scheduling rules for automatic
	// regular application. When auto provision is disabled, cronScheduleExpression is set to null, stopping the automatic
	// task scheduling. Read-Only.
	CronScheduleExpression nullable.Type[string] `json:"cronScheduleExpression,omitempty"`

	// Indicates IT Admins can set an end date to define the last scheduler run before this time. If not set, the scheduler
	// runs continuously. There is no time zone information at this time; it needs to be coordinated with timezone, for
	// example, '2025-02-01 00:00:00' with 'China Standard Time' means the scheduling rule takes effect before Feb 01 2025
	// 00:00:00 GMT+0800 (China Standard Time).
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Indicates IT Admins can see when the next automatic regular apply is executed. It needs to be coordinated with
	// timezone, for example, '2025-01-01 00:00:00' with 'China Standard Time' means the next task executes at Jan 01 2025
	// 00:00:00 GMT+0800 (China Standard Time). Read-Only.
	NextRunDateTime nullable.Type[string] `json:"nextRunDateTime,omitempty"`

	// The percentage of Cloud PCs to keep available. Administrators can set this property to a value from 0 to 99. Cloud
	// PCs are reprovisioned only when there are no active and connected Cloud PC users. Frontline shared only.
	ReservePercentage nullable.Type[int64] `json:"reservePercentage,omitempty"`

	// Indicates IT Admins can set a start date to define the first scheduler run after this time. If not set, the default
	// is the current time. There is no time zone information at this time, it needs to be coordinated with timezone, for
	// example, '2025-01-01 00:00:00' with 'China Standard Time' means the scheduling rule takes effect after Jan 01 2025
	// 00:00:00 GMT+0800 (China Standard Time).
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	Timezone *CloudPCPolicyTimezone `json:"timezone,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CloudPCPolicyScheduledApplyActionDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCPolicyScheduledApplyActionDetail{}

func (s CloudPCPolicyScheduledApplyActionDetail) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCPolicyScheduledApplyActionDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCPolicyScheduledApplyActionDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCPolicyScheduledApplyActionDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcPolicyScheduledApplyActionDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCPolicyScheduledApplyActionDetail: %+v", err)
	}

	return encoded, nil
}
