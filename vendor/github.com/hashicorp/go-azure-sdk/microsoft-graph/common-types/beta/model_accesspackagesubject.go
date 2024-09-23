package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageSubject{}

type AccessPackageSubject struct {
	// Not Supported.
	AltSecId nullable.Type[string] `json:"altSecId,omitempty"`

	// The date and time the subject is marked to be blocked from sign in or deleted. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time.
	CleanupScheduledDateTime nullable.Type[string] `json:"cleanupScheduledDateTime,omitempty"`

	// The connected organization of the subject. Read-only. Nullable.
	ConnectedOrganization *ConnectedOrganization `json:"connectedOrganization,omitempty"`

	// The identifier of the connected organization of the subject.
	ConnectedOrganizationId nullable.Type[string] `json:"connectedOrganizationId,omitempty"`

	// The display name of the subject.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The email address of the subject.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The object identifier of the subject. null if the subject isn't yet a user in the tenant. Alternate key.
	ObjectId nullable.Type[string] `json:"objectId,omitempty"`

	// A string representation of the principal's security identifier, if known, or null if the subject doesn't have a
	// security identifier.
	OnPremisesSecurityIdentifier nullable.Type[string] `json:"onPremisesSecurityIdentifier,omitempty"`

	// The principal name, if known, of the subject.
	PrincipalName nullable.Type[string] `json:"principalName,omitempty"`

	// The lifecycle of the subject user, if a guest. The possible values are: notDefined, notGoverned, governed,
	// unknownFutureValue.
	SubjectLifecycle *AccessPackageSubjectLifecycle `json:"subjectLifecycle,omitempty"`

	// The resource type of the subject.
	Type nullable.Type[string] `json:"type,omitempty"`

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

func (s AccessPackageSubject) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageSubject{}

func (s AccessPackageSubject) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageSubject
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageSubject: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageSubject: %+v", err)
	}

	delete(decoded, "connectedOrganization")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageSubject"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageSubject: %+v", err)
	}

	return encoded, nil
}
