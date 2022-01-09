package rest

import "io"

type CdbListControllerInterface interface {
	DeleteFile(arg1 ListFilenamePath, params *CdbListControllerDeleteFileParams, arg3 ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
	GetFile(arg1 ListFilenamePath, params *CdbListControllerGetFileParams, arg3 ...RequestEditorFn) (*AllItemsResponseListsFiles, error)
	GetListsFiles(params *CdbListControllerGetListsFilesParams, arg2 ...RequestEditorFn) (*AllItemsResponseListsFiles, error)
	GetLists(params *CdbListControllerGetListsParams, arg2 ...RequestEditorFn) (*AllItemsResponseLists, error)
	PutFileWithBody(arg1 ListFilenamePath, params *CdbListControllerPutFileParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
}

// RuleControllerInterface contains all methods for the wazuh controller api
type RuleControllerInterface interface {
	DeleteFile(arg1 XmlFilenamePath, params *RuleControllerDeleteFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetFile(arg1 XmlFilenamePath, params *RuleControllerGetFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetRulesFiles(params *RuleControllerGetRulesFilesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRulesFiles, error)
	GetRulesGroups(params *RuleControllerGetRulesGroupsParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error)
	GetRulesRequirement(arg1 RuleControllerGetRulesRequirementParamsRequirement, params *RuleControllerGetRulesRequirementParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetRules(params *RuleControllerGetRulesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRules, error)
	PutFileWithBody(arg1 XmlFilenamePath, params *RuleControllerPutFileParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponse, error)
}

// SecurityControllerInterface contains all methods for the wazuh controller api
type SecurityControllerInterface interface {
	AddPolicyWithBody(params *SecurityControllerAddPolicyParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	AddPolicy(params *SecurityControllerAddPolicyParams, arg2 SecurityControllerAddPolicyJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	AddRoleWithBody(params *SecurityControllerAddRoleParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	AddRole(params *SecurityControllerAddRoleParams, arg2 SecurityControllerAddRoleJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	AddRuleWithBody(params *SecurityControllerAddRuleParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	AddRule(params *SecurityControllerAddRuleParams, arg2 SecurityControllerAddRuleJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	CreateUserWithBody(params *SecurityControllerCreateUserParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	CreateUser(params *SecurityControllerCreateUserParams, arg2 SecurityControllerCreateUserJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	DeleteSecurityConfig(params *SecurityControllerDeleteSecurityConfigParams, arg2 ...RequestEditorFn) (*map[string]interface{}, error)
	DeleteUsers(params *SecurityControllerDeleteUsersParams, arg2 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	EditRunAs(arg1 UserIdRequired, params *SecurityControllerEditRunAsParams, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	GetPolicies(params *SecurityControllerGetPoliciesParams, arg2 ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	GetRbacActions(params *SecurityControllerGetRbacActionsParams, arg2 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetRbacResources(params *SecurityControllerGetRbacResourcesParams, arg2 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetRoles(params *SecurityControllerGetRolesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	GetRules(params *SecurityControllerGetRulesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	GetSecurityConfig(params *SecurityControllerGetSecurityConfigParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	GetUserMePolicies(params *SecurityControllerGetUserMePoliciesParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	GetUserMe(params *SecurityControllerGetUserMeParams, arg2 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	GetUsers(params *SecurityControllerGetUsersParams, arg2 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	LoginUser(params *SecurityControllerLoginUserParams, arg2 ...RequestEditorFn) (*Token, error)
	LogoutUser(arg1 ...RequestEditorFn) (*ApiResponse, error)
	PutSecurityConfigWithBody(params *SecurityControllerPutSecurityConfigParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*map[string]interface{}, error)
	PutSecurityConfig(params *SecurityControllerPutSecurityConfigParams, arg2 SecurityControllerPutSecurityConfigJSONRequestBody, arg3 ...RequestEditorFn) (*map[string]interface{}, error)
	RemovePolicies(params *SecurityControllerRemovePoliciesParams, arg2 ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	RemoveRolePolicy(arg1 RoleId, params *SecurityControllerRemoveRolePolicyParams, arg3 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	RemoveRoleRule(arg1 RoleId, params *SecurityControllerRemoveRoleRuleParams, arg3 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	RemoveRoles(params *SecurityControllerRemoveRolesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	RemoveRules(params *SecurityControllerRemoveRulesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	RemoveUserRole(arg1 UserIdRequired, params *SecurityControllerRemoveUserRoleParams, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	RevokeAllTokens(arg1 ...RequestEditorFn) (*map[string]interface{}, error)
	RunAsLoginWithBody(params *SecurityControllerRunAsLoginParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*struct {
		Token *string "json:\"token,omitempty\""
	}, error)
	RunAsLogin(params *SecurityControllerRunAsLoginParams, arg2 SecurityControllerRunAsLoginJSONRequestBody, arg3 ...RequestEditorFn) (*struct {
		Token *string "json:\"token,omitempty\""
	}, error)
	SetRolePolicy(arg1 RoleId, params *SecurityControllerSetRolePolicyParams, arg3 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	SetRoleRule(arg1 RoleId, params *SecurityControllerSetRoleRuleParams, arg3 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	SetUserRole(arg1 UserIdRequired, params *SecurityControllerSetUserRoleParams, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdatePolicyWithBody(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	UpdatePolicy(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 SecurityControllerUpdatePolicyJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	UpdateRoleWithBody(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	UpdateRole(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 SecurityControllerUpdateRoleJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponseRoles, error)
	UpdateRuleWithBody(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdateRule(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 SecurityControllerUpdateRuleJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdateUserWithBody(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdateUser(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 SecurityControllerUpdateUserJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponseUsers, error)
}

// TaskControllerInterface contains all methods for the wazuh controller api
type TaskControllerInterface interface {
	GetTasksStatus(params *TaskControllerGetTasksStatusParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
}

// ManagerControllerInterface contains all methods for the wazuh controller api
type ManagerControllerInterface interface {
	GetAPIConfig(params *ManagerControllerGetApiConfigParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	GetConfValidation(params *ManagerControllerGetConfValidationParams, arg2 ...RequestEditorFn) (*ConfigurationValidation, error)
	GetConfiguration(params *ManagerControllerGetConfigurationParams, arg2 ...RequestEditorFn) (*WazuhManagerConfiguration, error)
	GetInfo(params *ManagerControllerGetInfoParams, arg2 ...RequestEditorFn) (*WazuhInfo, error)
	GetLogSummary(params *ManagerControllerGetLogSummaryParams, arg2 ...RequestEditorFn) (*WazuhLogsSummary, error)
	GetLog(params *ManagerControllerGetLogParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error)
	GetManagerConfigOndemand(arg1 ManagerControllerGetManagerConfigOndemandParamsComponent, arg2 ManagerControllerGetManagerConfigOndemandParamsConfiguration, params *ManagerControllerGetManagerConfigOndemandParams, arg4 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetStatsAnalysisd(params *ManagerControllerGetStatsAnalysisdParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsHourly(params *ManagerControllerGetStatsHourlyParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsRemoted(params *ManagerControllerGetStatsRemotedParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsWeekly(params *ManagerControllerGetStatsWeeklyParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStats(params *ManagerControllerGetStatsParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatus(params *ManagerControllerGetStatusParams, arg2 ...RequestEditorFn) (*WazuhDaemonsStatus, error)
	PutRestart(params *ManagerControllerPutRestartParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	UpdateConfigurationWithBody(params *ManagerControllerUpdateConfigurationParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
}

// MitreControllerInterface contains all methods for the wazuh controller api
type MitreControllerInterface interface {
	GetAttack(params *MitreControllerGetAttackParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
}

// OverviewControllerInterface contains all methods for the wazuh controller api
type OverviewControllerInterface interface {
	GetOverviewAgents(params *OverviewControllerGetOverviewAgentsParams, arg2 ...RequestEditorFn) (*OverviewAgents, error)
}

// RootcheckControllerInterface contains all methods for the wazuh controller api
type RootcheckControllerInterface interface {
	DeleteRootcheck(params *RootcheckControllerDeleteRootcheckParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error)
	GetLastScanAgent(arg1 AgentId, params *RootcheckControllerGetLastScanAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetRootcheckAgent(arg1 AgentId, params *RootcheckControllerGetRootcheckAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	PutRootcheck(params *RootcheckControllerPutRootcheckParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error)
}

// ActiveResponseControllerInterface contains all methods for the wazuh controller api
type ActiveResponseControllerInterface interface {
	RunCommandWithBody(params *ActiveResponseControllerRunCommandParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*ApiResponse, error)
	RunCommand(params *ActiveResponseControllerRunCommandParams, arg2 ActiveResponseControllerRunCommandJSONRequestBody, arg3 ...RequestEditorFn) (*ApiResponse, error)
}

// CiscatControllerInterface contains all methods for the wazuh controller api
type CiscatControllerInterface interface {
	GetAgentsCiscatResults(arg1 AgentId, params *CiscatControllerGetAgentsCiscatResultsParams, arg3 ...RequestEditorFn) (*AllItemsResponseCiscatResult, error)
}

// DefaultControllerInterface contains all methods for the wazuh controller api
type DefaultControllerInterface interface {
	DefaultInfo(params *DefaultControllerDefaultInfoParams, arg2 ...RequestEditorFn) (*BasicInfo, error)
}

// LogtestControllerInterface contains all methods for the wazuh controller api
type LogtestControllerInterface interface {
	EndLogtestSession(arg1 LogtestToken, params *LogtestControllerEndLogtestSessionParams, arg3 ...RequestEditorFn) (*ApiResponse, error)
	RunLogtestToolWithBody(params *LogtestControllerRunLogtestToolParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*ApiResponse, error)
	RunLogtestTool(params *LogtestControllerRunLogtestToolParams, arg2 LogtestControllerRunLogtestToolJSONRequestBody, arg3 ...RequestEditorFn) (*ApiResponse, error)
}

// SyscheckControllerInterface contains all methods for the wazuh controller api
type SyscheckControllerInterface interface {
	DeleteSyscheckAgent(arg1 AgentId, params *SyscheckControllerDeleteSyscheckAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetLastScanAgent(arg1 AgentId, params *SyscheckControllerGetLastScanAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscheckLastScan, error)
	GetSyscheckAgent(arg1 AgentId, params *SyscheckControllerGetSyscheckAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscheckResult, error)
	PutSyscheck(params *SyscheckControllerPutSyscheckParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
}

// SyscollectorControllerInterface contains all methods for the wazuh controller api
type SyscollectorControllerInterface interface {
	GetHardwareInfo(arg1 AgentId, params *SyscollectorControllerGetHardwareInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error)
	GetHotfixInfo(arg1 AgentId, params *SyscollectorControllerGetHotfixInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error)
	GetNetworkAddressInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkAddressInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error)
	GetNetworkInterfaceInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkInterfaceInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error)
	GetNetworkProtocolInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkProtocolInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error)
	GetOsInfo(arg1 AgentId, params *SyscollectorControllerGetOsInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error)
	GetPackagesInfo(arg1 AgentId, params *SyscollectorControllerGetPackagesInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error)
	GetPortsInfo(arg1 AgentId, params *SyscollectorControllerGetPortsInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error)
	GetProcessesInfo(arg1 AgentId, params *SyscollectorControllerGetProcessesInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error)
}

// ClusterControllerInterface contains all methods for the wazuh controller api
type ClusterControllerInterface interface {
	GetAPIConfig(params *ClusterControllerGetApiConfigParams, arg2 ...RequestEditorFn) (*struct {
		AdditionalProperties map[string]interface{} "json:\"-\""
	}, error)
	GetClusterNode(params *ClusterControllerGetClusterNodeParams, arg2 ...RequestEditorFn) (*struct {
		Cluster *string "json:\"cluster,omitempty\""
		Node    *string "json:\"node,omitempty\""
		Type    *string "json:\"type,omitempty\""
	}, error)
	GetClusterNodes(params *ClusterControllerGetClusterNodesParams, arg2 ...RequestEditorFn) (*AllItemsResponseClusterNodes, error)
	GetConfValidation(params *ClusterControllerGetConfValidationParams, arg2 ...RequestEditorFn) (*AllItemsResponseValidationStatus, error)
	GetConfig(params *ClusterControllerGetConfigParams, arg2 ...RequestEditorFn) (*struct {
		BindAddr *string           "json:\"bind_addr,omitempty\""
		Disabled *bool             "json:\"disabled,omitempty\""
		Hidden   *string           "json:\"hidden,omitempty\""
		Key      *string           "json:\"key,omitempty\""
		Name     *string           "json:\"name,omitempty\""
		NodeName *string           "json:\"node_name,omitempty\""
		NodeType *N200DataNodeType "json:\"node_type,omitempty\""
		Nodes    *[]string         "json:\"nodes,omitempty\""
		Port     *int              "json:\"port,omitempty\""
	}, error)
	GetConfigurationNode(arg1 NodeId, params *ClusterControllerGetConfigurationNodeParams, arg3 ...RequestEditorFn) (*WazuhManagerConfiguration, error)
	GetHealthcheck(params *ClusterControllerGetHealthcheckParams, arg2 ...RequestEditorFn) (*AllItemsResponseNodeHealthcheck, error)
	GetInfoNode(arg1 NodeId, params *ClusterControllerGetInfoNodeParams, arg3 ...RequestEditorFn) (*WazuhInfo, error)
	GetLogNode(arg1 NodeId, params *ClusterControllerGetLogNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error)
	GetLogSummaryNode(arg1 NodeId, params *ClusterControllerGetLogSummaryNodeParams, arg3 ...RequestEditorFn) (*WazuhLogsSummary, error)
	GetNodeConfig(arg1 NodeId, arg2 ClusterControllerGetNodeConfigParamsComponent, arg3 ClusterControllerGetNodeConfigParamsConfiguration, params *ClusterControllerGetNodeConfigParams, arg5 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetStatsAnalysisdNode(arg1 NodeId, params *ClusterControllerGetStatsAnalysisdNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsHourlyNode(arg1 NodeId, params *ClusterControllerGetStatsHourlyNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsNode(arg1 NodeId, params *ClusterControllerGetStatsNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsRemotedNode(arg1 NodeId, params *ClusterControllerGetStatsRemotedNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsWeeklyNode(arg1 NodeId, params *ClusterControllerGetStatsWeeklyNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatusNode(arg1 NodeId, params *ClusterControllerGetStatusNodeParams, arg3 ...RequestEditorFn) (*WazuhDaemonsStatus, error)
	GetStatus(params *ClusterControllerGetStatusParams, arg2 ...RequestEditorFn) (*struct {
		Enabled *N200DataEnabled "json:\"enabled,omitempty\""
		Running *N200DataRunning "json:\"running,omitempty\""
	}, error)
	PutRestart(params *ClusterControllerPutRestartParams, arg2 ...RequestEditorFn) (*AllItemsResponseNodeIDs, error)
	UpdateConfigurationWithBody(arg1 NodeId, params *ClusterControllerUpdateConfigurationParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
}

// ExperimentalControllerInterface contains all methods for the wazuh controller api
type ExperimentalControllerInterface interface {
	ClearSyscheckDatabase(params *ExperimentalControllerClearSyscheckDatabaseParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	GetCisCatResults(params *ExperimentalControllerGetCisCatResultsParams, arg2 ...RequestEditorFn) (*AllItemsResponseCiscatResult, error)
	GetHardwareInfo(params *ExperimentalControllerGetHardwareInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error)
	GetHotfixesInfo(params *ExperimentalControllerGetHotfixesInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error)
	GetNetworkAddressInfo(params *ExperimentalControllerGetNetworkAddressInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error)
	GetNetworkInterfaceInfo(params *ExperimentalControllerGetNetworkInterfaceInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error)
	GetNetworkProtocolInfo(params *ExperimentalControllerGetNetworkProtocolInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error)
	GetOsInfo(params *ExperimentalControllerGetOsInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error)
	GetPackagesInfo(params *ExperimentalControllerGetPackagesInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error)
	GetPortsInfo(params *ExperimentalControllerGetPortsInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error)
	GetProcessesInfo(params *ExperimentalControllerGetProcessesInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error)
}

// ScaControllerInterface contains all methods for the wazuh controller api
type ScaControllerInterface interface {
	GetScaAgent(arg1 AgentId, params *ScaControllerGetScaAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseSCADatabase, error)
	GetScaChecks(arg1 AgentId, arg2 PolicyId, params *ScaControllerGetScaChecksParams, arg4 ...RequestEditorFn) (*AllItemsResponseSCAChecks, error)
}

// VulnerabilityControllerInterface contains all methods for the wazuh controller api
type VulnerabilityControllerInterface interface {
	GetVulnerabilityAgent(arg1 AgentId, params *VulnerabilityControllerGetVulnerabilityAgentParams, arg3 ...RequestEditorFn) (*ApiResponse, error)
}

// AgentControllerInterface contains all methods for the wazuh controller api
type AgentControllerInterface interface {
	AddAgentWithBody(params *AgentControllerAddAgentParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AgentIdKey, error)
	AddAgent(params *AgentControllerAddAgentParams, arg2 AgentControllerAddAgentJSONRequestBody, arg3 ...RequestEditorFn) (*AgentIdKey, error)
	DeleteAgents(params *AgentControllerDeleteAgentsParams, arg2 ...RequestEditorFn) (*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}, error)
	DeleteGroups(params *AgentControllerDeleteGroupsParams, arg2 ...RequestEditorFn) (*struct {
		AllItemsResponseGroupIDs "yaml:\",inline\""
		AgentGroupDeleted        "yaml:\",inline\""
	}, error)
	DeleteMultipleAgentSingleGroup(params *AgentControllerDeleteMultipleAgentSingleGroupParams, arg2 ...RequestEditorFn) (*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}, error)
	DeleteSingleAgentMultipleGroups(arg1 AgentId, params *AgentControllerDeleteSingleAgentMultipleGroupsParams, arg3 ...RequestEditorFn) (*struct {
		AllItemsResponseGroupIDs "yaml:\",inline\""
	}, error)
	DeleteSingleAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerDeleteSingleAgentSingleGroupParams, arg4 ...RequestEditorFn) (*ApiResponse, error)
	GetAgentConfig(arg1 AgentId, arg2 AgentControllerGetAgentConfigParamsComponent, arg3 AgentControllerGetAgentConfigParamsConfiguration, params *AgentControllerGetAgentConfigParams, arg5 ...RequestEditorFn) (*AgentConfiguration, error)
	GetAgentFields(params *AgentControllerGetAgentFieldsParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentsDistinct, error)
	GetAgentKey(arg1 AgentId, params *AgentControllerGetAgentKeyParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentsKeys, error)
	GetAgentNoGroup(params *AgentControllerGetAgentNoGroupParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgents, error)
	GetAgentOutdated(params *AgentControllerGetAgentOutdatedParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentsSimple, error)
	GetAgentSummaryOs(params *AgentControllerGetAgentSummaryOsParams, arg2 ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetAgentSummaryStatus(params *AgentControllerGetAgentSummaryStatusParams, arg2 ...RequestEditorFn) (*AgentsSummaryStatus, error)
	GetAgentUpgrade(params *AgentControllerGetAgentUpgradeParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	GetAgentsInGroup(arg1 GroupId, params *AgentControllerGetAgentsInGroupParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgents, error)
	GetAgents(params *AgentControllerGetAgentsParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgents, error)
	GetComponentStats(arg1 AgentId, arg2 AgentControllerGetComponentStatsParamsComponent, params *AgentControllerGetComponentStatsParams, arg4 ...RequestEditorFn) (*ApiResponse, error)
	GetGroupConfig(arg1 GroupId, params *AgentControllerGetGroupConfigParams, arg3 ...RequestEditorFn) (*struct {
		AffectedItems      *[]GroupConfiguration "json:\"affected_items,omitempty\""
		TotalAffectedItems *int32                "json:\"total_affected_items,omitempty\""
	}, error)
	GetGroupFileJSON(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileJsonParams, arg4 ...RequestEditorFn) (*interface{}, error)
	GetGroupFileXML(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileXmlParams, arg4 ...RequestEditorFn) (*AgentControllerGetGroupFileXmlResponse, error)
	GetGroupFiles(arg1 GroupId, params *AgentControllerGetGroupFilesParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetListGroup(params *AgentControllerGetListGroupParams, arg2 ...RequestEditorFn) (*AllItemsResponseGroups, error)
	GetSyncAgent(arg1 AgentId, params *AgentControllerGetSyncAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentsSynced, error)
	InsertAgentWithBody(params *AgentControllerInsertAgentParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AgentIdKey, error)
	InsertAgent(params *AgentControllerInsertAgentParams, arg2 AgentControllerInsertAgentJSONRequestBody, arg3 ...RequestEditorFn) (*AgentIdKey, error)
	PostGroupWithBody(params *AgentControllerPostGroupParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*ApiResponse, error)
	PostGroup(params *AgentControllerPostGroupParams, arg2 AgentControllerPostGroupJSONRequestBody, arg3 ...RequestEditorFn) (*ApiResponse, error)
	PostNewAgent(params *AgentControllerPostNewAgentParams, arg2 ...RequestEditorFn) (*AgentIdKey, error)
	PutAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerPutAgentSingleGroupParams, arg4 ...RequestEditorFn) (*ApiResponse, error)
	PutGroupConfigWithBody(arg1 GroupId, params *AgentControllerPutGroupConfigParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*ApiResponse, error)
	PutMultipleAgentSingleGroup(params *AgentControllerPutMultipleAgentSingleGroupParams, arg2 ...RequestEditorFn) (*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}, error)
	PutUpgradeAgents(params *AgentControllerPutUpgradeAgentsParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	PutUpgradeCustomAgents(params *AgentControllerPutUpgradeCustomAgentsParams, arg2 ...RequestEditorFn) (*ApiResponse, error)
	rtAgent(arg1 AgentId, params *AgentControllerRestartAgentParams, arg3 ...RequestEditorFn) (*ItemAffected, error)
	rtAgentsByGroup(arg1 GroupId, params *AgentControllerRestartAgentsByGroupParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	rtAgentsByNode(arg1 NodeId, params *AgentControllerRestartAgentsByNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	rtAgents(params *AgentControllerRestartAgentsParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
}

// DecoderControllerInterface contains all methods for the wazuh controller api
type DecoderControllerInterface interface {
	DeleteFile(arg1 XmlFilenamePath, params *DecoderControllerDeleteFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	GetDecodersFiles(params *DecoderControllerGetDecodersFilesParams, arg2 ...RequestEditorFn) (*AllItemsResponseDecodersFiles, error)
	GetDecodersParents(params *DecoderControllerGetDecodersParentsParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error)
	GetDecoders(params *DecoderControllerGetDecodersParams, arg2 ...RequestEditorFn) (*AllItemsResponseDecoders, error)
	GetFile(arg1 XmlFilenamePath, params *DecoderControllerGetFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error)
	PutFileWithBody(arg1 XmlFilenamePath, params *DecoderControllerPutFileParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponse, error)
}
