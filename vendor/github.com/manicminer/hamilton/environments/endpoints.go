package environments

type AzureADEndpoint string

const (
	AzureADGlobal  AzureADEndpoint = "https://login.microsoftonline.com"
	AzureADUSGov   AzureADEndpoint = "https://login.microsoftonline.us"
	AzureADGermany AzureADEndpoint = "https://login.microsoftonline.de"
	AzureADChina   AzureADEndpoint = "https://login.chinacloudapi.cn"
)
