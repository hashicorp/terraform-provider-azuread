package environments

type Environment struct {
	AzureADEndpoint AzureADEndpoint
	MsGraph         Api
	AadGraph        Api
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
