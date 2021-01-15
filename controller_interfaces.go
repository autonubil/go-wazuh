package wazuh

import "io"

// SyscheckControllerInterface contains all methods for the wazuh controller api
type SyscheckControllerInterface interface {
	DeleteSyscheckAgent(agentID AgentId, params *SyscheckControllerDeleteSyscheckAgentParams) (*AllItemsResponse, error)
	GetLastScanAgent(agentID AgentId, params *SyscheckControllerGetLastScanAgentParams) (*AllItemsResponseSyscheckLastScan, error)
	GetSyscheckAgent(agentID AgentId, params *SyscheckControllerGetSyscheckAgentParams) (*AllItemsResponseSyscheckResult, error)
	PutSyscheck(params *SyscheckControllerPutSyscheckParams) (*AllItemsResponseAgentIDs, error)
}

// SyscollectorControllerInterface contains all methods for the wazuh controller api
type SyscollectorControllerInterface interface {
	GetHardwareInfo(agentID AgentId, params *SyscollectorControllerGetHardwareInfoParams) (*AllItemsResponseSyscollectorHardware, error)
	GetHotfixInfo(agentID AgentId, params *SyscollectorControllerGetHotfixInfoParams) (*AllItemsResponseSyscollectorHotfixes, error)
	GetNetworkAddressInfo(agentID AgentId, params *SyscollectorControllerGetNetworkAddressInfoParams) (*AllItemsResponseSyscollectorNetwork, error)
	GetNetworkInterfaceInfo(agentID AgentId, params *SyscollectorControllerGetNetworkInterfaceInfoParams) (*AllItemsResponseSyscollectorInterface, error)
	GetNetworkProtocolInfo(agentID AgentId, params *SyscollectorControllerGetNetworkProtocolInfoParams) (*AllItemsResponseSyscollectorProtocol, error)
	GetOsInfo(agentID AgentId, params *SyscollectorControllerGetOsInfoParams) (*AllItemsResponseSyscollectorOS, error)
	GetPackagesInfo(agentID AgentId, params *SyscollectorControllerGetPackagesInfoParams) (*AllItemsResponseSyscollectorPackages, error)
	GetPortsInfo(agentID AgentId, params *SyscollectorControllerGetPortsInfoParams) (*AllItemsResponseSyscollectorPorts, error)
	GetProcessesInfo(agentID AgentId, params *SyscollectorControllerGetProcessesInfoParams) (*AllItemsResponseSyscollectorProcesses, error)
}

// CiscatControllerInterface contains all methods for the wazuh controller api
type CiscatControllerInterface interface {
	GetAgentsCiscatResults(agentID AgentId, params *CiscatControllerGetAgentsCiscatResultsParams) (*AllItemsResponseCiscatResult, error)
}

// DecodersControllerInterface contains all methods for the wazuh controller api
type DecodersControllerInterface interface {
	GetDecodersFiles(params *DecodersControllerGetDecodersFilesParams) (*AllItemsResponseDecodersFiles, error)
	GetDecodersParents(params *DecodersControllerGetDecodersParentsParams) (*AllItemsResponse, error)
	GetDecoders(params *DecodersControllerGetDecodersParams) (*AllItemsResponseDecoders, error)
	GetDownloadFile(downloadFile DownloadFile, params *DecodersControllerGetDownloadFileParams) (*DecodersControllerGetDownloadFileResponse, error)
}

// DefaultControllerInterface contains all methods for the wazuh controller api
type DefaultControllerInterface interface {
	DefaultInfo(params *DefaultControllerDefaultInfoParams) (*BasicInfo, error)
}

// ExperimentalControllerInterface contains all methods for the wazuh controller api
type ExperimentalControllerInterface interface {
	ClearSyscheckDatabase(params *ExperimentalControllerClearSyscheckDatabaseParams) (*AllItemsResponseAgentIDs, error)
	GetCisCatResults(params *ExperimentalControllerGetCisCatResultsParams) (*AllItemsResponseCiscatResult, error)
	GetHardwareInfo(params *ExperimentalControllerGetHardwareInfoParams) (*AllItemsResponseSyscollectorHardware, error)
	GetHotfixesInfo(params *ExperimentalControllerGetHotfixesInfoParams) (*AllItemsResponseSyscollectorHotfixes, error)
	GetNetworkAddressInfo(params *ExperimentalControllerGetNetworkAddressInfoParams) (*AllItemsResponseSyscollectorNetwork, error)
	GetNetworkInterfaceInfo(params *ExperimentalControllerGetNetworkInterfaceInfoParams) (*AllItemsResponseSyscollectorInterface, error)
	GetNetworkProtocolInfo(params *ExperimentalControllerGetNetworkProtocolInfoParams) (*AllItemsResponseSyscollectorProtocol, error)
	GetOsInfo(params *ExperimentalControllerGetOsInfoParams) (*AllItemsResponseSyscollectorOS, error)
	GetPackagesInfo(params *ExperimentalControllerGetPackagesInfoParams) (*AllItemsResponseSyscollectorPackages, error)
	GetPortsInfo(params *ExperimentalControllerGetPortsInfoParams) (*AllItemsResponseSyscollectorPorts, error)
	GetProcessesInfo(params *ExperimentalControllerGetProcessesInfoParams) (*AllItemsResponseSyscollectorProcesses, error)
}

// ScaControllerInterface contains all methods for the wazuh controller api
type ScaControllerInterface interface {
	GetScaAgent(agentID AgentId, params *ScaControllerGetScaAgentParams) (*AllItemsResponseSCADatabase, error)
	GetScaChecks(agentID AgentId, policyID PolicyId, params *ScaControllerGetScaChecksParams) (*AllItemsResponseSCAChecks, error)
}

// AgentsControllerInterface contains all methods for the wazuh controller api
type AgentsControllerInterface interface {
	AddAgentWithBody(params *AgentsControllerAddAgentParams, contentType string, body io.Reader) (*AgentIdKey, error)
	AddAgent(params *AgentsControllerAddAgentParams, agentsControllerAddAgentJSONRequestBody AgentsControllerAddAgentJSONRequestBody) (*AgentIdKey, error)
	DeleteAgents(params *AgentsControllerDeleteAgentsParams) (*struct {
		AllItemsResponseAgentIDs
		OlderThan *string
	}, error)
	DeleteGroups(params *AgentsControllerDeleteGroupsParams) (*struct {
		AllItemsResponseGroupIDs
		AgentGroupDeleted
	}, error)
	DeleteMultipleAgentSingleGroup(params *AgentsControllerDeleteMultipleAgentSingleGroupParams) (*struct{ AllItemsResponseAgentIDs }, error)
	DeleteSingleAgentMultipleGroups(agentID AgentId, params *AgentsControllerDeleteSingleAgentMultipleGroupsParams) (*struct{ AllItemsResponseGroupIDs }, error)
	DeleteSingleAgentSingleGroup(agentID AgentId, groupID GroupId, params *AgentsControllerDeleteSingleAgentSingleGroupParams) (*ApiResponse, error)
	GetAgentConfig(agentID AgentId, component Component, configuration Configuration, params *AgentsControllerGetAgentConfigParams) (*AgentConfiguration, error)
	GetAgentFields(params *AgentsControllerGetAgentFieldsParams) (*AllItemsResponseAgentsDistinct, error)
	GetAgentKey(agentID AgentId, params *AgentsControllerGetAgentKeyParams) (*AllItemsResponseAgentsKeys, error)
	GetAgentNoGroup(params *AgentsControllerGetAgentNoGroupParams) (*AllItemsResponseAgents, error)
	GetAgentOutdated(params *AgentsControllerGetAgentOutdatedParams) (*AllItemsResponseAgentsSimple, error)
	GetAgentSummaryOs(params *AgentsControllerGetAgentSummaryOsParams) (*struct{ ApiResponse }, error)
	GetAgentSummaryStatus(params *AgentsControllerGetAgentSummaryStatusParams) (*AgentsSummaryStatus, error)
	GetAgentUpgrade(agentID AgentId, params *AgentsControllerGetAgentUpgradeParams) (*ApiResponse, error)
	GetAgentsInGroup(groupID GroupId, params *AgentsControllerGetAgentsInGroupParams) (*AllItemsResponseAgents, error)
	GetAgents(params *AgentsControllerGetAgentsParams) (*AllItemsResponseAgents, error)
	GetGroupConfig(groupID GroupId, params *AgentsControllerGetGroupConfigParams) (*struct {
		AffectedItems      *[]GroupConfiguration
		TotalAffectedItems *int32
	}, error)
	GetGroupFileJSON(groupID GroupId, fileName FileName, params *AgentsControllerGetGroupFileJsonParams) (*interface{}, error)
	GetGroupFileXML(groupID GroupId, fileName FileName, params *AgentsControllerGetGroupFileXmlParams) (*AgentsControllerGetGroupFileXmlResponse, error)
	GetGroupFiles(groupID GroupId, params *AgentsControllerGetGroupFilesParams) (*AllItemsResponse, error)
	GetListGroup(params *AgentsControllerGetListGroupParams) (*AllItemsResponseGroups, error)
	GetSyncAgent(agentID AgentId, params *AgentsControllerGetSyncAgentParams) (*AllItemsResponseAgentsSynced, error)
	InsertAgentWithBody(params *AgentsControllerInsertAgentParams, contentType string, body io.Reader) (*AgentIdKey, error)
	InsertAgent(params *AgentsControllerInsertAgentParams, agentsControllerInsertAgentJSONRequestBody AgentsControllerInsertAgentJSONRequestBody) (*AgentIdKey, error)
	PostGroup(params *AgentsControllerPostGroupParams) (*ApiResponse, error)
	PostNewAgent(params *AgentsControllerPostNewAgentParams) (*AgentIdKey, error)
	PutAgentSingleGroup(agentID AgentId, groupID GroupId, params *AgentsControllerPutAgentSingleGroupParams) (*ApiResponse, error)
	PutGroupConfigWithBody(groupID GroupId, params *AgentsControllerPutGroupConfigParams, contentType string, body io.Reader) (*ApiResponse, error)
	PutMultipleAgentSingleGroup(params *AgentsControllerPutMultipleAgentSingleGroupParams) (*struct{ AllItemsResponseAgentIDs }, error)
	PutUpgradeAgent(agentID AgentId, params *AgentsControllerPutUpgradeAgentParams) (*ApiResponse, error)
	PutUpgradeCustomAgent(agentID AgentId, params *AgentsControllerPutUpgradeCustomAgentParams) (*ApiResponse, error)
	RestartAgent(agentID AgentId, params *AgentsControllerRestartAgentParams) (*ItemAffected, error)
	RestartAgentsByGroup(groupID GroupId, params *AgentsControllerRestartAgentsByGroupParams) (*AllItemsResponseAgentIDs, error)
	RestartAgentsByNode(nodeID NodeId, params *AgentsControllerRestartAgentsByNodeParams) (*AllItemsResponseAgentIDs, error)
	RestartAgents(params *AgentsControllerRestartAgentsParams) (*AllItemsResponseAgentIDs, error)
}

// ClusterControllerInterface contains all methods for the wazuh controller api
type ClusterControllerInterface interface {
	DeleteFilesNode(nodeID NodeId, params *ClusterControllerDeleteFilesNodeParams) (*ApiResponse, error)
	GetAPIConfig(params *ClusterControllerGetApiConfigParams) (*struct{ AdditionalProperties map[string]interface{} }, error)
	GetClusterNode(params *ClusterControllerGetClusterNodeParams) (*struct {
		Cluster *string
		Node    *string
		Type    *string
	}, error)
	GetClusterNodes(params *ClusterControllerGetClusterNodesParams) (*AllItemsResponseClusterNodes, error)
	GetConfValidation(params *ClusterControllerGetConfValidationParams) (*AllItemsResponseValidationStatus, error)
	GetConfig(params *ClusterControllerGetConfigParams) (*struct {
		BindAddr *string
		Disabled *bool
		Hidden   *string
		Key      *string
		Name     *string
		NodeName *string
		NodeType *string
		Nodes    *[]string
		Port     *int
	}, error)
	GetConfigurationNode(nodeID NodeId, params *ClusterControllerGetConfigurationNodeParams) (*WazuhMangerConfiguration, error)
	GetFilesNode(nodeID NodeId, params *ClusterControllerGetFilesNodeParams) (*struct{ Contents *string }, error)
	GetHealthcheck(params *ClusterControllerGetHealthcheckParams) (*AllItemsResponseNodeHealthcheck, error)
	GetInfoNode(nodeID NodeId, params *ClusterControllerGetInfoNodeParams) (*WazuhInfo, error)
	GetLogNode(nodeID NodeId, params *ClusterControllerGetLogNodeParams) (*AllItemsResponseWazuhLogs, error)
	GetLogSummaryNode(nodeID NodeId, params *ClusterControllerGetLogSummaryNodeParams) (*WazuhLogsSummary, error)
	GetNodeConfig(nodeID NodeId, component Component, configuration Configuration, params *ClusterControllerGetNodeConfigParams) (*struct{ ApiResponse }, error)
	GetStatsAnalysisdNode(nodeID NodeId, params *ClusterControllerGetStatsAnalysisdNodeParams) (*AllItemsResponseWazuhStats, error)
	GetStatsHourlyNode(nodeID NodeId, params *ClusterControllerGetStatsHourlyNodeParams) (*AllItemsResponseWazuhStats, error)
	GetStatsNode(nodeID NodeId, params *ClusterControllerGetStatsNodeParams) (*AllItemsResponseWazuhStats, error)
	GetStatsRemotedNode(nodeID NodeId, params *ClusterControllerGetStatsRemotedNodeParams) (*AllItemsResponseWazuhStats, error)
	GetStatsWeeklyNode(nodeID NodeId, params *ClusterControllerGetStatsWeeklyNodeParams) (*AllItemsResponseWazuhStats, error)
	GetStatusNode(nodeID NodeId, params *ClusterControllerGetStatusNodeParams) (*WazuhDaemonsStatus, error)
	GetStatus(params *ClusterControllerGetStatusParams) (*struct {
		Enabled *string
		Running *string
	}, error)
	PutFilesNodeWithBody(nodeID NodeId, params *ClusterControllerPutFilesNodeParams, contentType string, body io.Reader) (*ApiResponse, error)
	PutRestart(params *ClusterControllerPutRestartParams) (*AllItemsResponseNodeIDs, error)
}

// ListsControllerInterface contains all methods for the wazuh controller api
type ListsControllerInterface interface {
	GetListsFiles(params *ListsControllerGetListsFilesParams) (*AllItemsResponseListsFiles, error)
	GetLists(params *ListsControllerGetListsParams) (*AllItemsResponseLists, error)
}

// ManagerControllerInterface contains all methods for the wazuh controller api
type ManagerControllerInterface interface {
	DeleteFiles(params *ManagerControllerDeleteFilesParams) (*struct {
		ApiResponse
		ConfirmationMessage
	}, error)
	GetAPIConfig(params *ManagerControllerGetApiConfigParams) (*ApiResponse, error)
	GetConfValidation(params *ManagerControllerGetConfValidationParams) (*ConfigurationValidation, error)
	GetConfiguration(params *ManagerControllerGetConfigurationParams) (*WazuhMangerConfiguration, error)
	GetFiles(params *ManagerControllerGetFilesParams) (*struct{ Contents *string }, error)
	GetInfo(params *ManagerControllerGetInfoParams) (*WazuhInfo, error)
	GetLogSummary(params *ManagerControllerGetLogSummaryParams) (*WazuhLogsSummary, error)
	GetLog(params *ManagerControllerGetLogParams) (*AllItemsResponseWazuhLogs, error)
	GetManagerConfigOndemand(component Component, configuration Configuration, params *ManagerControllerGetManagerConfigOndemandParams) (*struct{ ApiResponse }, error)
	GetStatsAnalysisd(params *ManagerControllerGetStatsAnalysisdParams) (*AllItemsResponseWazuhStats, error)
	GetStatsHourly(params *ManagerControllerGetStatsHourlyParams) (*AllItemsResponseWazuhStats, error)
	GetStatsRemoted(params *ManagerControllerGetStatsRemotedParams) (*AllItemsResponseWazuhStats, error)
	GetStatsWeekly(params *ManagerControllerGetStatsWeeklyParams) (*AllItemsResponseWazuhStats, error)
	GetStats(params *ManagerControllerGetStatsParams) (*AllItemsResponseWazuhStats, error)
	GetStatus(params *ManagerControllerGetStatusParams) (*WazuhDaemonsStatus, error)
	PutFilesWithBody(params *ManagerControllerPutFilesParams, contentType string, body io.Reader) (*struct {
		ApiResponse
		ConfirmationMessage
	}, error)
	PutRestart(params *ManagerControllerPutRestartParams) (*ApiResponse, error)
}

// MitreControllerInterface contains all methods for the wazuh controller api
type MitreControllerInterface interface {
	GetAttack(params *MitreControllerGetAttackParams) (*ApiResponse, error)
}

// OverviewControllerInterface contains all methods for the wazuh controller api
type OverviewControllerInterface interface {
	GetOverviewAgents(params *OverviewControllerGetOverviewAgentsParams) (*OverviewAgents, error)
}

// RulesControllerInterface contains all methods for the wazuh controller api
type RulesControllerInterface interface {
	GetDownloadFile(downloadFile DownloadFile, params *RulesControllerGetDownloadFileParams) (*RulesControllerGetDownloadFileResponse, error)
	GetRulesFiles(params *RulesControllerGetRulesFilesParams) (*AllItemsResponseRulesFiles, error)
	GetRulesGroups(params *RulesControllerGetRulesGroupsParams) (*AllItemsResponse, error)
	GetRulesRequirement(ruleRequirement RuleRequirement, params *RulesControllerGetRulesRequirementParams) (*AllItemsResponse, error)
	GetRules(params *RulesControllerGetRulesParams) (*AllItemsResponseRules, error)
}

// SecurityControllerInterface contains all methods for the wazuh controller api
type SecurityControllerInterface interface {
	AddPolicyWithBody(params *SecurityControllerAddPolicyParams, contentType string, body io.Reader) (*AllItemsResponsePolicies, error)
	AddPolicy(params *SecurityControllerAddPolicyParams, securityControllerAddPolicyJSONRequestBody SecurityControllerAddPolicyJSONRequestBody) (*AllItemsResponsePolicies, error)
	AddRoleWithBody(params *SecurityControllerAddRoleParams, contentType string, body io.Reader) (*AllItemsResponseRoles, error)
	AddRole(params *SecurityControllerAddRoleParams, securityControllerAddRoleJSONRequestBody SecurityControllerAddRoleJSONRequestBody) (*AllItemsResponseRoles, error)
	AddRuleWithBody(params *SecurityControllerAddRuleParams, contentType string, body io.Reader) (*AllItemsResponseRoles, error)
	AddRule(params *SecurityControllerAddRuleParams, securityControllerAddRuleJSONRequestBody SecurityControllerAddRuleJSONRequestBody) (*AllItemsResponseRoles, error)
	CreateUserWithBody(params *SecurityControllerCreateUserParams, contentType string, body io.Reader) (*AllItemsResponseUsers, error)
	CreateUser(params *SecurityControllerCreateUserParams, securityControllerCreateUserJSONRequestBody SecurityControllerCreateUserJSONRequestBody) (*AllItemsResponseUsers, error)
	DeleteSecurityConfig(params *SecurityControllerDeleteSecurityConfigParams) (*map[string]interface{}, error)
	DeleteUsers(params *SecurityControllerDeleteUsersParams) (*AllItemsResponseUsers, error)
	GetPolicies(params *SecurityControllerGetPoliciesParams) (*AllItemsResponsePolicies, error)
	GetRbacActions(params *SecurityControllerGetRbacActionsParams) (*struct{ ApiResponse }, error)
	GetRbacResources(params *SecurityControllerGetRbacResourcesParams) (*struct{ ApiResponse }, error)
	GetRoles(params *SecurityControllerGetRolesParams) (*AllItemsResponseRoles, error)
	GetRules(params *SecurityControllerGetRulesParams) (*AllItemsResponseRoles, error)
	GetSecurityConfig(params *SecurityControllerGetSecurityConfigParams) (*ApiResponse, error)
	GetUserMePolicies(params *SecurityControllerGetUserMePoliciesParams) (*ApiResponse, error)
	GetUserMe(params *SecurityControllerGetUserMeParams) (*AllItemsResponseUsers, error)
	GetUsers(params *SecurityControllerGetUsersParams) (*AllItemsResponseUsers, error)
	LoginUserRunAsWithBody(params *SecurityControllerLoginUserParams, contentType string, body io.Reader) (*struct{ Token *string }, error)
	LoginUserRunAs(params *SecurityControllerLoginUserParams, securityControllerLoginUserJSONRequestBody SecurityControllerLoginUserJSONRequestBody) (*struct{ Token *string }, error)
	LoginUser(params *SecurityControllerLoginUserParams) (*Token, error)
	LogoutUser() (*ApiResponse, error)
	PutSecurityConfigWithBody(params *SecurityControllerPutSecurityConfigParams, contentType string, body io.Reader) (*map[string]interface{}, error)
	PutSecurityConfig(params *SecurityControllerPutSecurityConfigParams, securityControllerPutSecurityConfigJSONRequestBody SecurityControllerPutSecurityConfigJSONRequestBody) (*map[string]interface{}, error)
	RemovePolicies(params *SecurityControllerRemovePoliciesParams) (*AllItemsResponsePolicies, error)
	RemoveRolePolicy(roleID RoleId, params *SecurityControllerRemoveRolePolicyParams) (*struct{ ApiResponse }, error)
	RemoveRoleRule(roleID RoleId, params *SecurityControllerRemoveRoleRuleParams) (*struct{ ApiResponse }, error)
	RemoveRoles(params *SecurityControllerRemoveRolesParams) (*AllItemsResponseRoles, error)
	RemoveRules(params *SecurityControllerRemoveRulesParams) (*AllItemsResponseRoles, error)
	RemoveUserRole(userIDRequired UserIdRequired, params *SecurityControllerRemoveUserRoleParams) (*AllItemsResponseUsers, error)
	RevokeAllTokens() (*map[string]interface{}, error)
	SetRolePolicy(roleID RoleId, params *SecurityControllerSetRolePolicyParams) (*struct{ ApiResponse }, error)
	SetRoleRule(roleID RoleId, params *SecurityControllerSetRoleRuleParams) (*struct{ ApiResponse }, error)
	SetUserRole(userIDRequired UserIdRequired, params *SecurityControllerSetUserRoleParams) (*AllItemsResponseUsers, error)
	UpdatePolicyWithBody(policyIDRbac PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, contentType string, body io.Reader) (*AllItemsResponsePolicies, error)
	UpdatePolicy(policyIDRbac PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, securityControllerUpdatePolicyJSONRequestBody SecurityControllerUpdatePolicyJSONRequestBody) (*AllItemsResponsePolicies, error)
	UpdateRoleWithBody(roleID RoleId, params *SecurityControllerUpdateRoleParams, contentType string, body io.Reader) (*AllItemsResponseRoles, error)
	UpdateRole(roleID RoleId, params *SecurityControllerUpdateRoleParams, securityControllerUpdateRoleJSONRequestBody SecurityControllerUpdateRoleJSONRequestBody) (*AllItemsResponseRoles, error)
	UpdateRuleWithBody(securityRuleID SecurityRuleId, params *SecurityControllerUpdateRuleParams, contentType string, body io.Reader) (*AllItemsResponseUsers, error)
	UpdateRule(securityRuleID SecurityRuleId, params *SecurityControllerUpdateRuleParams, securityControllerUpdateRuleJSONRequestBody SecurityControllerUpdateRuleJSONRequestBody) (*AllItemsResponseUsers, error)
	UpdateUserWithBody(userIDRequired UserIdRequired, params *SecurityControllerUpdateUserParams, contentType string, body io.Reader) (*AllItemsResponseUsers, error)
	UpdateUser(userIDRequired UserIdRequired, params *SecurityControllerUpdateUserParams, securityControllerUpdateUserJSONRequestBody SecurityControllerUpdateUserJSONRequestBody) (*AllItemsResponseUsers, error)
}

// ActiveResponseControllerInterface contains all methods for the wazuh controller api
type ActiveResponseControllerInterface interface {
	RunCommandWithBody(params *ActiveResponseControllerRunCommandParams, contentType string, body io.Reader) (*ApiResponse, error)
	RunCommand(params *ActiveResponseControllerRunCommandParams, activeResponseControllerRunCommandJSONRequestBody ActiveResponseControllerRunCommandJSONRequestBody) (*ApiResponse, error)
}
