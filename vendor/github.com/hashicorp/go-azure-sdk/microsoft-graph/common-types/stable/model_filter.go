package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Filter struct {
	// *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this
	// object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
	CategoryFilterGroups *[]FilterGroup `json:"categoryFilterGroups,omitempty"`

	// Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be
	// used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was
	// changed so that filter isn't satisfied any longer, such object will get deprovisioned'. An object is considered in
	// scope if ANY of the groups in the collection is evaluated to true.
	Groups *[]FilterGroup `json:"groups,omitempty"`

	// *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If
	// an object doesn't satisfy this filter, then it will not be processed further. Important to understand is that if an
	// object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is
	// no longer satisfied, such object will NOT get deprovisioned. An object is considered in scope if ANY of the groups in
	// the collection is evaluated to true.
	InputFilterGroups *[]FilterGroup `json:"inputFilterGroups,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
