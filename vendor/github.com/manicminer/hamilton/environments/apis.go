package environments

type ApiEndpoint string

const (
	AadGraphGlobalEndpoint  ApiEndpoint = "https://graph.windows.net"
	AadGraphGermanyEndpoint ApiEndpoint = "https://graph.cloudapi.de"
	AadGraphChinaEndpoint   ApiEndpoint = "https://graph.chinacloudapi.cn"
	AadGraphUSGovEndpoint   ApiEndpoint = "https://graph.microsoftazure.us"
	MsGraphGlobalEndpoint   ApiEndpoint = "https://graph.microsoft.com"
	MsGraphGermanyEndpoint  ApiEndpoint = "https://graph.microsoft.de"
	MsGraphChinaEndpoint    ApiEndpoint = "https://microsoftgraph.chinacloudapi.cn"
	MsGraphUSGovL4Endpoint  ApiEndpoint = "https://graph.microsoft.us"
	MsGraphUSGovL5Endpoint  ApiEndpoint = "https://dod-graph.microsoft.us"
	MsGraphCanaryEndpoint   ApiEndpoint = "https://canary.graph.microsoft.com"
)

type ApiCliName string

const (
	AadGraphCliName ApiCliName = "aad-graph"
	MsGraphCliName  ApiCliName = "ms-graph"
)

// API represent an API configuration for Microsoft Graph or Azure Active Directory Graph.
type Api struct {
	// The Application ID for the API.
	AppId    ApiAppId

	// The Azure CLI codename for the API. Used with `az account get-access-token`.
	CliName  ApiCliName

	// The endpoint for the API, including scheme.
	Endpoint ApiEndpoint
}

var (
	MsGraphGlobal = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		CliName:  MsGraphCliName,
		Endpoint: MsGraphGlobalEndpoint,
	}

	MsGraphGermany = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		CliName:  MsGraphCliName,
		Endpoint: MsGraphGermanyEndpoint,
	}

	MsGraphChina = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		CliName:  MsGraphCliName,
		Endpoint: MsGraphChinaEndpoint,
	}

	MsGraphUSGovL4 = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		CliName:  MsGraphCliName,
		Endpoint: MsGraphUSGovL4Endpoint,
	}

	MsGraphUSGovL5 = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		CliName:  MsGraphCliName,
		Endpoint: MsGraphUSGovL5Endpoint,
	}

	MsGraphCanary = Api{
		AppId:    PublishedApis["MicrosoftGraph"],
		CliName:  MsGraphCliName,
		Endpoint: MsGraphCanaryEndpoint,
	}

	AadGraphGlobal = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		CliName:  AadGraphCliName,
		Endpoint: AadGraphGlobalEndpoint,
	}

	AadGraphGermany = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		CliName:  AadGraphCliName,
		Endpoint: AadGraphGermanyEndpoint,
	}

	AadGraphChina = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		CliName:  AadGraphCliName,
		Endpoint: AadGraphChinaEndpoint,
	}

	AadGraphUSGov = Api{
		AppId:    PublishedApis["AzureActiveDirectoryGraph"],
		CliName:  AadGraphCliName,
		Endpoint: AadGraphUSGovEndpoint,
	}
)
