package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServicePort struct {
	// The application protocol for this port.
	AppProtocol nullable.Type[string] `json:"appProtocol,omitempty"`

	// The name of this port within the service.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The port on each node on which this service is exposed when the type is either NodePort or LoadBalancer.
	NodePort *int64 `json:"nodePort,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The port that this service exposes.
	Port *int64 `json:"port,omitempty"`

	// The protocol name. Possible values are: udp, tcp, sctp, unknownFutureValue.
	Protocol *SecurityContainerPortProtocol `json:"protocol,omitempty"`

	// The name or number of the port to access on the pods targeted by the service. The port number must be in the range 1
	// to 65535. The name must be an IANASVCNAME.
	TargetPort nullable.Type[string] `json:"targetPort,omitempty"`
}
