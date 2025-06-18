package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDiscoveredCloudAppDetail interface {
	Entity
	SecurityDiscoveredCloudAppDetail() BaseSecurityDiscoveredCloudAppDetailImpl
}

var _ SecurityDiscoveredCloudAppDetail = BaseSecurityDiscoveredCloudAppDetailImpl{}

type BaseSecurityDiscoveredCloudAppDetailImpl struct {
	// The application information.
	AppInfo *SecurityDiscoveredCloudAppInfo `json:"appInfo,omitempty"`

	Category    *SecurityAppCategory  `json:"category,omitempty"`
	Description nullable.Type[string] `json:"description,omitempty"`

	// The app name.
	DisplayName *string `json:"displayName,omitempty"`

	// The domain.
	Domains *[]string `json:"domains,omitempty"`

	// The download traffic size.
	DownloadNetworkTrafficInBytes *int64 `json:"downloadNetworkTrafficInBytes,omitempty"`

	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`

	// The IP address.
	IPAddressCount *int64 `json:"ipAddressCount,omitempty"`

	// The list of IP addresses accessed by the app.
	IPAddresses *[]SecurityDiscoveredCloudAppIPAddress `json:"ipAddresses,omitempty"`

	// The last seen date of the discovered app. The Timestamp represents date and time information using ISO 8601 format
	// and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`

	// The risk score of the app.
	RiskScore nullable.Type[int64] `json:"riskScore,omitempty"`

	// The tags applied to an app. Possible values include Unsanctioned, Sanctioned, Monitored, or a custom value.
	Tags *[]string `json:"tags,omitempty"`

	// The app transaction count.
	TransactionCount *int64 `json:"transactionCount,omitempty"`

	// The app upload traffic size, in bytes.
	UploadNetworkTrafficInBytes *int64 `json:"uploadNetworkTrafficInBytes,omitempty"`

	// The count of users who use the app.
	UserCount *int64 `json:"userCount,omitempty"`

	// The list of users who access the app.
	Users *[]SecurityDiscoveredCloudAppUser `json:"users,omitempty"`

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

func (s BaseSecurityDiscoveredCloudAppDetailImpl) SecurityDiscoveredCloudAppDetail() BaseSecurityDiscoveredCloudAppDetailImpl {
	return s
}

func (s BaseSecurityDiscoveredCloudAppDetailImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityDiscoveredCloudAppDetail = RawSecurityDiscoveredCloudAppDetailImpl{}

// RawSecurityDiscoveredCloudAppDetailImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityDiscoveredCloudAppDetailImpl struct {
	securityDiscoveredCloudAppDetail BaseSecurityDiscoveredCloudAppDetailImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawSecurityDiscoveredCloudAppDetailImpl) SecurityDiscoveredCloudAppDetail() BaseSecurityDiscoveredCloudAppDetailImpl {
	return s.securityDiscoveredCloudAppDetail
}

func (s RawSecurityDiscoveredCloudAppDetailImpl) Entity() BaseEntityImpl {
	return s.securityDiscoveredCloudAppDetail.Entity()
}

var _ json.Marshaler = BaseSecurityDiscoveredCloudAppDetailImpl{}

func (s BaseSecurityDiscoveredCloudAppDetailImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityDiscoveredCloudAppDetailImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityDiscoveredCloudAppDetailImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityDiscoveredCloudAppDetailImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.discoveredCloudAppDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityDiscoveredCloudAppDetailImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityDiscoveredCloudAppDetailImplementation(input []byte) (SecurityDiscoveredCloudAppDetail, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDiscoveredCloudAppDetail into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.endpointDiscoveredCloudAppDetail") {
		var out SecurityEndpointDiscoveredCloudAppDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEndpointDiscoveredCloudAppDetail: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityDiscoveredCloudAppDetailImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityDiscoveredCloudAppDetailImpl: %+v", err)
	}

	return RawSecurityDiscoveredCloudAppDetailImpl{
		securityDiscoveredCloudAppDetail: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
