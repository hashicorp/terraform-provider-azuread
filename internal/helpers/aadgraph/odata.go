package aadgraph

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

// NewOdataError attempts to unmarshal an autorest.Response to *graphrbac.ODataError
func NewOdataError(resp autorest.Response) (*graphrbac.OdataError, error) {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	dest := struct {
		Odata graphrbac.OdataError `json:"odata.error"`
	}{}
	if err := json.Unmarshal(body, &dest); err != nil {
		return nil, err
	}
	return &dest.Odata, nil
}

// OdataErrorContains checks whether an OData error message contains a specified substring
func OdataErrorContains(odata *graphrbac.OdataError, search string) bool {
	if odata.ErrorMessage == nil || odata.ErrorMessage.Message == nil {
		return false
	}
	if strings.Contains(*odata.ErrorMessage.Message, search) {
		return true
	}
	return false
}
