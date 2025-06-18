package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityDiscoveredCloudAppDetail = SecurityEndpointDiscoveredCloudAppDetail{}

type SecurityEndpointDiscoveredCloudAppDetail struct {
	// The number of devices that accessed the discovered app.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// Represents the devices that access the discovered apps.
	Devices *[]SecurityDiscoveredCloudAppDevice `json:"devices,omitempty"`

	// Fields inherited from SecurityDiscoveredCloudAppDetail

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

func (s SecurityEndpointDiscoveredCloudAppDetail) SecurityDiscoveredCloudAppDetail() BaseSecurityDiscoveredCloudAppDetailImpl {
	return BaseSecurityDiscoveredCloudAppDetailImpl{
		AppInfo:                       s.AppInfo,
		Category:                      s.Category,
		Description:                   s.Description,
		DisplayName:                   s.DisplayName,
		Domains:                       s.Domains,
		DownloadNetworkTrafficInBytes: s.DownloadNetworkTrafficInBytes,
		FirstSeenDateTime:             s.FirstSeenDateTime,
		IPAddressCount:                s.IPAddressCount,
		IPAddresses:                   s.IPAddresses,
		LastSeenDateTime:              s.LastSeenDateTime,
		RiskScore:                     s.RiskScore,
		Tags:                          s.Tags,
		TransactionCount:              s.TransactionCount,
		UploadNetworkTrafficInBytes:   s.UploadNetworkTrafficInBytes,
		UserCount:                     s.UserCount,
		Users:                         s.Users,
		Id:                            s.Id,
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
	}
}

func (s SecurityEndpointDiscoveredCloudAppDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEndpointDiscoveredCloudAppDetail{}

func (s SecurityEndpointDiscoveredCloudAppDetail) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEndpointDiscoveredCloudAppDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEndpointDiscoveredCloudAppDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEndpointDiscoveredCloudAppDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.endpointDiscoveredCloudAppDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEndpointDiscoveredCloudAppDetail: %+v", err)
	}

	return encoded, nil
}
