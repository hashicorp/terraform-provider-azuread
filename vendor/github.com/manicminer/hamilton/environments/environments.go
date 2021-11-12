package environments

import (
	"fmt"
	"strings"
)

// Environment represents a set of API configurations for a particular cloud.
type Environment struct {
	// The Azure AD endpoint for acquiring access tokens.
	AzureADEndpoint AzureADEndpoint

	// The Microsoft Graph configuration for an environment.
	MsGraph Api

	// The Azure Active Directory Graph configuration for an environment.
	AadGraph Api
}

var (
	Global = Environment{
		AzureADEndpoint: AzureADGlobal,
		MsGraph:         MsGraphGlobal,
		AadGraph:        AadGraphGlobal,
	}

	Germany = Environment{
		AzureADEndpoint: AzureADGermany,
		MsGraph:         MsGraphGermany,
		AadGraph:        AadGraphGermany,
	}

	China = Environment{
		AzureADEndpoint: AzureADChina,
		MsGraph:         MsGraphChina,
		AadGraph:        AadGraphChina,
	}

	USGovernmentL4 = Environment{
		AzureADEndpoint: AzureADUSGov,
		MsGraph:         MsGraphUSGovL4,
		AadGraph:        AadGraphUSGov,
	}

	USGovernmentL5 = Environment{
		AzureADEndpoint: AzureADUSGov,
		MsGraph:         MsGraphUSGovL5,
		AadGraph:        AadGraphUSGov,
	}

	Canary = Environment{
		AzureADEndpoint: AzureADGlobal,
		MsGraph:         MsGraphCanary,
	}
)

func EnvironmentFromString(env string) (Environment, error) {
	switch strings.ToLower(env) {
	case "", "public", "global":
		return Global, nil
	case "usgovernment", "usgovernmentl4":
		return USGovernmentL4, nil
	case "dod", "usgovernmentl5":
		return USGovernmentL5, nil
	case "canary":
		return Canary, nil
	case "china":
		return China, nil
	case "germany":
		return Germany, nil
	}

	return Environment{}, fmt.Errorf("invalid environment specified: %s", env)
}
