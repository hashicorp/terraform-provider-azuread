package parse

import "fmt"

type ApplicationExtensionPropertyId struct {
	ObjectSubResourceId
	ApplicationId       string
	ExtensionPropertyId string
}

func NewApplicationExtensionPropertyID(appId, extAttrId string) ApplicationExtensionPropertyId {
	return ApplicationExtensionPropertyId{
		ObjectSubResourceId: NewObjectSubResourceID(appId, "extensionProperty", extAttrId),
		ApplicationId:       appId,
		ExtensionPropertyId: extAttrId,
	}
}

func ApplicationExtensionPropertyID(idString string) (*ApplicationExtensionPropertyId, error) {
	id, err := ObjectSubResourceID(idString, "extensionProperty")
	if err != nil {
		return nil, fmt.Errorf("unable to parse ExtensionProperty ID: %v", err)
	}

	return &ApplicationExtensionPropertyId{
		ObjectSubResourceId: *id,
		ApplicationId:       id.objectId,
		ExtensionPropertyId: id.subId,
	}, nil
}
