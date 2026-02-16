package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationLinkedObjects struct {
	Manager *SynchronizationJobSubject `json:"manager,omitempty"`

	// All group members that you would like to provision.
	Members *[]SynchronizationJobSubject `json:"members,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Owners *[]SynchronizationJobSubject `json:"owners,omitempty"`
}
