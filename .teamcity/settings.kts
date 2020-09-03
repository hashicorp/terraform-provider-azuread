import AzureAD
import ClientConfiguration
import jetbrains.buildServer.configs.kotlin.v2019_2.*

version = "2020.1"

var clientId = DslContext.getParameter("clientId", "")
var clientSecret = DslContext.getParameter("clientSecret", "")
var tenantId = DslContext.getParameter("tenantId", "")
var environment = DslContext.getParameter("environment", "public")

var clientConfig = ClientConfiguration(clientId, clientSecret, tenantId)

project(AzureAD(environment, clientConfig))
