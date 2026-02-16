package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PermissionGrantConditionSet{}

type PermissionGrantConditionSet struct {
	// A list of appId values for the client applications to match with, or a list with the single value all to match any
	// client application. Default is the single value all.
	ClientApplicationIds *[]string `json:"clientApplicationIds,omitempty"`

	// A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the
	// single value all to match with client apps from any publisher. Default is the single value all.
	ClientApplicationPublisherIds *[]string `json:"clientApplicationPublisherIds,omitempty"`

	// A list of Microsoft Entra tenant IDs in which the client application is registered, or a list with the single value
	// all to match with client apps registered in any tenant. Default is the single value all.
	ClientApplicationTenantIds *[]string `json:"clientApplicationTenantIds,omitempty"`

	// Set to true to only match on client applications with a verified publisher. Set to false to match on any client app,
	// even if it doesn't have a verified publisher. Default is false.
	ClientApplicationsFromVerifiedPublisherOnly nullable.Type[bool] `json:"clientApplicationsFromVerifiedPublisherOnly,omitempty"`

	// The permission classification for the permission being granted, or all to match with any permission classification
	// (including permissions that aren't classified). Default is all.
	PermissionClassification nullable.Type[string] `json:"permissionClassification,omitempty"`

	// The permission type of the permission being granted. Possible values: application for application permissions (for
	// example app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated
	// permissions that haven't been configured by the API publisher to require admin consent—this value may be used in
	// built-in permission grant policies, but can't be used in custom permission grant policies. Required.
	PermissionType PermissionType `json:"permissionType"`

	// The list of id values for the specific permissions to match with, or a list with the single value all to match with
	// any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's
	// servicePrincipal object. The id of application permissions can be found in the appRoles property of the API's
	// servicePrincipal object. The id of resource-specific application permissions can be found in the
	// resourceSpecificApplicationPermissions property of the API's servicePrincipal object. Default is the single value
	// all.
	Permissions *[]string `json:"permissions,omitempty"`

	// The appId of the resource application (for example the API) for which a permission is being granted, or any to match
	// with any resource application or API. Default is any.
	ResourceApplication nullable.Type[string] `json:"resourceApplication,omitempty"`

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

func (s PermissionGrantConditionSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PermissionGrantConditionSet{}

func (s PermissionGrantConditionSet) MarshalJSON() ([]byte, error) {
	type wrapper PermissionGrantConditionSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PermissionGrantConditionSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionGrantConditionSet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.permissionGrantConditionSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PermissionGrantConditionSet: %+v", err)
	}

	return encoded, nil
}
