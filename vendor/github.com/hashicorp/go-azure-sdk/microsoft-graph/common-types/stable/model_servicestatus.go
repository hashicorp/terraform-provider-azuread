package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceStatus struct {
	// The type of consumer. The possible values are: unknown, firstparty, thirdparty, unknownFutureValue.
	BackupServiceConsumer *BackupServiceConsumer `json:"backupServiceConsumer,omitempty"`

	// The reason the service is disabled. The possible values are: none, controllerServiceAppDeleted,
	// invalidBillingProfile, userRequested, unknownFutureValue.
	DisableReason *DisableReason `json:"disableReason,omitempty"`

	// The expiration time of the grace period.
	GracePeriodDateTime nullable.Type[string] `json:"gracePeriodDateTime,omitempty"`

	// Identity of the person who last modified the entity.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification of the entity.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The expiration time of the restoration allowed period.
	RestoreAllowedTillDateTime nullable.Type[string] `json:"restoreAllowedTillDateTime,omitempty"`

	// Status of the service. This value indicates what capabilities can be used. The possible values are: disabled,
	// enabled, protectionChangeLocked, restoreLocked, unknownFutureValue.
	Status *BackupServiceStatus `json:"status,omitempty"`
}

var _ json.Unmarshaler = &ServiceStatus{}

func (s *ServiceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BackupServiceConsumer      *BackupServiceConsumer `json:"backupServiceConsumer,omitempty"`
		DisableReason              *DisableReason         `json:"disableReason,omitempty"`
		GracePeriodDateTime        nullable.Type[string]  `json:"gracePeriodDateTime,omitempty"`
		LastModifiedDateTime       nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`
		ODataId                    *string                `json:"@odata.id,omitempty"`
		ODataType                  *string                `json:"@odata.type,omitempty"`
		RestoreAllowedTillDateTime nullable.Type[string]  `json:"restoreAllowedTillDateTime,omitempty"`
		Status                     *BackupServiceStatus   `json:"status,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BackupServiceConsumer = decoded.BackupServiceConsumer
	s.DisableReason = decoded.DisableReason
	s.GracePeriodDateTime = decoded.GracePeriodDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RestoreAllowedTillDateTime = decoded.RestoreAllowedTillDateTime
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceStatus into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'ServiceStatus': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
