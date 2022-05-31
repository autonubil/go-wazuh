package rest

import "io"

// CdbListControllerInterface contains all methods for the wazuh controller api
type CdbListControllerInterface interface {
	DeleteFile(arg1 ListFilenamePath, params *CdbListControllerDeleteFileParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
	GetFile(arg1 ListFilenamePath, params *CdbListControllerGetFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponseListsFiles, error)
	GetListsFiles(params *CdbListControllerGetListsFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseListsFiles, error)
	GetLists(params *CdbListControllerGetListsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseLists, error)
	PutFileWithBody(arg1 ListFilenamePath, params *CdbListControllerPutFileParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
}

// ClusterControllerInterface contains all methods for the wazuh controller api
type ClusterControllerInterface interface {
	GetAPIConfig(params *ClusterControllerGetApiConfigParams, reqEditors ...RequestEditorFn) (*struct {
		AdditionalProperties map[string]interface{} "json:\"-\""
	}, error)
	GetClusterNode(params *ClusterControllerGetClusterNodeParams, reqEditors ...RequestEditorFn) (*struct {
		Cluster *string "json:\"cluster,omitempty\""
		Node    *string "json:\"node,omitempty\""
		Type    *string "json:\"type,omitempty\""
	}, error)
	GetClusterNodes(params *ClusterControllerGetClusterNodesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseClusterNodes, error)
	GetConfValidation(params *ClusterControllerGetConfValidationParams, reqEditors ...RequestEditorFn) (*AllItemsResponseValidationStatus, error)
	GetConfig(params *ClusterControllerGetConfigParams, reqEditors ...RequestEditorFn) (*struct {
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
	GetConfigurationNode(arg1 NodeId, params *ClusterControllerGetConfigurationNodeParams, reqEditors ...RequestEditorFn) (*WazuhManagerConfiguration, error)
	GetHealthcheck(params *ClusterControllerGetHealthcheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponseNodeHealthcheck, error)
	GetInfoNode(arg1 NodeId, params *ClusterControllerGetInfoNodeParams, reqEditors ...RequestEditorFn) (*WazuhInfo, error)
	GetLogNode(arg1 NodeId, params *ClusterControllerGetLogNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error)
	GetLogSummaryNode(arg1 NodeId, params *ClusterControllerGetLogSummaryNodeParams, reqEditors ...RequestEditorFn) (*WazuhLogsSummary, error)
	GetNodeConfig(arg1 NodeId, arg2 ClusterControllerGetNodeConfigParamsComponent, arg3 ClusterControllerGetNodeConfigParamsConfiguration, params *ClusterControllerGetNodeConfigParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetStatsAnalysisdNode(arg1 NodeId, params *ClusterControllerGetStatsAnalysisdNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsHourlyNode(arg1 NodeId, params *ClusterControllerGetStatsHourlyNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsNode(arg1 NodeId, params *ClusterControllerGetStatsNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsRemotedNode(arg1 NodeId, params *ClusterControllerGetStatsRemotedNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsWeeklyNode(arg1 NodeId, params *ClusterControllerGetStatsWeeklyNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatusNode(arg1 NodeId, params *ClusterControllerGetStatusNodeParams, reqEditors ...RequestEditorFn) (*WazuhDaemonsStatus, error)
	GetStatus(params *ClusterControllerGetStatusParams, reqEditors ...RequestEditorFn) (*struct {
		Enabled *N200DataEnabled "json:\"enabled,omitempty\""
		Running *N200DataRunning "json:\"running,omitempty\""
	}, error)
	PutRestart(params *ClusterControllerPutRestartParams, reqEditors ...RequestEditorFn) (*AllItemsResponseNodeIDs, error)
	UpdateConfigurationWithBody(arg1 NodeId, params *ClusterControllerUpdateConfigurationParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
}

// DefaultControllerInterface contains all methods for the wazuh controller api
type DefaultControllerInterface interface {
	DefaultInfo(params *DefaultControllerDefaultInfoParams, reqEditors ...RequestEditorFn) (*BasicInfo, error)
}

// LogtestControllerInterface contains all methods for the wazuh controller api
type LogtestControllerInterface interface {
	EndLogtestSession(arg1 LogtestToken, params *LogtestControllerEndLogtestSessionParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	RunLogtestToolWithBody(params *LogtestControllerRunLogtestToolParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	RunLogtestTool(params *LogtestControllerRunLogtestToolParams, arg2 LogtestControllerRunLogtestToolJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiResponse, error)
}

// MitreControllerInterface contains all methods for the wazuh controller api
type MitreControllerInterface interface {
	GetGroups(params *MitreControllerGetGroupsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetMetadata(params *MitreControllerGetMetadataParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetMitigations(params *MitreControllerGetMitigationsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetReferences(params *MitreControllerGetReferencesParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetSoftware(params *MitreControllerGetSoftwareParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetTactics(params *MitreControllerGetTacticsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetTechniques(params *MitreControllerGetTechniquesParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
}

// RuleControllerInterface contains all methods for the wazuh controller api
type RuleControllerInterface interface {
	DeleteFile(arg1 XmlFilenamePath, params *RuleControllerDeleteFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetFile(arg1 XmlFilenamePath, params *RuleControllerGetFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetRulesFiles(params *RuleControllerGetRulesFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRulesFiles, error)
	GetRulesGroups(params *RuleControllerGetRulesGroupsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetRulesRequirement(arg1 RuleControllerGetRulesRequirementParamsRequirement, params *RuleControllerGetRulesRequirementParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetRules(params *RuleControllerGetRulesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRules, error)
	PutFileWithBody(arg1 XmlFilenamePath, params *RuleControllerPutFileParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
}

// CiscatControllerInterface contains all methods for the wazuh controller api
type CiscatControllerInterface interface {
	GetAgentsCiscatResults(arg1 AgentId, params *CiscatControllerGetAgentsCiscatResultsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseCiscatResult, error)
}

// RootcheckControllerInterface contains all methods for the wazuh controller api
type RootcheckControllerInterface interface {
	DeleteRootcheck(arg1 AgentId, params *RootcheckControllerDeleteRootcheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetLastScanAgent(arg1 AgentId, params *RootcheckControllerGetLastScanAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetRootcheckAgent(arg1 AgentId, params *RootcheckControllerGetRootcheckAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	PutRootcheck(params *RootcheckControllerPutRootcheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
}

// ScaControllerInterface contains all methods for the wazuh controller api
type ScaControllerInterface interface {
	GetScaAgent(arg1 AgentId, params *ScaControllerGetScaAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSCADatabase, error)
	GetScaChecks(arg1 AgentId, arg2 PolicyId, params *ScaControllerGetScaChecksParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSCAChecks, error)
}

// SyscheckControllerInterface contains all methods for the wazuh controller api
type SyscheckControllerInterface interface {
	DeleteSyscheckAgent(arg1 AgentId, params *SyscheckControllerDeleteSyscheckAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetLastScanAgent(arg1 AgentId, params *SyscheckControllerGetLastScanAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseLastScan, error)
	GetSyscheckAgent(arg1 AgentId, params *SyscheckControllerGetSyscheckAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscheckResult, error)
	PutSyscheck(params *SyscheckControllerPutSyscheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
}

// SyscollectorControllerInterface contains all methods for the wazuh controller api
type SyscollectorControllerInterface interface {
	GetHardwareInfo(arg1 AgentId, params *SyscollectorControllerGetHardwareInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error)
	GetHotfixInfo(arg1 AgentId, params *SyscollectorControllerGetHotfixInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error)
	GetNetworkAddressInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkAddressInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error)
	GetNetworkInterfaceInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkInterfaceInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error)
	GetNetworkProtocolInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkProtocolInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error)
	GetOsInfo(arg1 AgentId, params *SyscollectorControllerGetOsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error)
	GetPackagesInfo(arg1 AgentId, params *SyscollectorControllerGetPackagesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error)
	GetPortsInfo(arg1 AgentId, params *SyscollectorControllerGetPortsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error)
	GetProcessesInfo(arg1 AgentId, params *SyscollectorControllerGetProcessesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error)
}

// ActiveResponseControllerInterface contains all methods for the wazuh controller api
type ActiveResponseControllerInterface interface {
	RunCommandWithBody(params *ActiveResponseControllerRunCommandParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	RunCommand(params *ActiveResponseControllerRunCommandParams, arg2 ActiveResponseControllerRunCommandJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiResponse, error)
}

// DecoderControllerInterface contains all methods for the wazuh controller api
type DecoderControllerInterface interface {
	DeleteFile(arg1 XmlFilenamePath, params *DecoderControllerDeleteFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetDecodersFiles(params *DecoderControllerGetDecodersFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseDecodersFiles, error)
	GetDecodersParents(params *DecoderControllerGetDecodersParentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetDecoders(params *DecoderControllerGetDecodersParams, reqEditors ...RequestEditorFn) (*AllItemsResponseDecoders, error)
	GetFile(arg1 XmlFilenamePath, params *DecoderControllerGetFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	PutFileWithBody(arg1 XmlFilenamePath, params *DecoderControllerPutFileParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
}

// ExperimentalControllerInterface contains all methods for the wazuh controller api
type ExperimentalControllerInterface interface {
	ClearRootcheckDatabase(params *ExperimentalControllerClearRootcheckDatabaseParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	ClearSyscheckDatabase(params *ExperimentalControllerClearSyscheckDatabaseParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	GetCisCatResults(params *ExperimentalControllerGetCisCatResultsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseCiscatResult, error)
	GetHardwareInfo(params *ExperimentalControllerGetHardwareInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error)
	GetHotfixesInfo(params *ExperimentalControllerGetHotfixesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error)
	GetNetworkAddressInfo(params *ExperimentalControllerGetNetworkAddressInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error)
	GetNetworkInterfaceInfo(params *ExperimentalControllerGetNetworkInterfaceInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error)
	GetNetworkProtocolInfo(params *ExperimentalControllerGetNetworkProtocolInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error)
	GetOsInfo(params *ExperimentalControllerGetOsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error)
	GetPackagesInfo(params *ExperimentalControllerGetPackagesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error)
	GetPortsInfo(params *ExperimentalControllerGetPortsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error)
	GetProcessesInfo(params *ExperimentalControllerGetProcessesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error)
}

// OverviewControllerInterface contains all methods for the wazuh controller api
type OverviewControllerInterface interface {
	GetOverviewAgents(params *OverviewControllerGetOverviewAgentsParams, reqEditors ...RequestEditorFn) (*OverviewAgents, error)
}

// AgentControllerInterface contains all methods for the wazuh controller api
type AgentControllerInterface interface {
	AddAgentWithBody(params *AgentControllerAddAgentParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AgentIdKey, error)
	AddAgent(params *AgentControllerAddAgentParams, arg2 AgentControllerAddAgentJSONRequestBody, reqEditors ...RequestEditorFn) (*AgentIdKey, error)
	DeleteAgents(params *AgentControllerDeleteAgentsParams, reqEditors ...RequestEditorFn) (*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}, error)
	DeleteGroups(params *AgentControllerDeleteGroupsParams, reqEditors ...RequestEditorFn) (*struct {
		AgentGroupDeleted "yaml:\",inline\""
	}, error)
	DeleteMultipleAgentSingleGroup(params *AgentControllerDeleteMultipleAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}, error)
	DeleteSingleAgentMultipleGroups(arg1 AgentId, params *AgentControllerDeleteSingleAgentMultipleGroupsParams, reqEditors ...RequestEditorFn) (*struct {
		AllItemsResponseGroupIDs "yaml:\",inline\""
	}, error)
	DeleteSingleAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerDeleteSingleAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetAgentConfig(arg1 AgentId, arg2 AgentControllerGetAgentConfigParamsComponent, arg3 AgentControllerGetAgentConfigParamsConfiguration, params *AgentControllerGetAgentConfigParams, reqEditors ...RequestEditorFn) (*AgentConfiguration, error)
	GetAgentFields(params *AgentControllerGetAgentFieldsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsDistinct, error)
	GetAgentKey(arg1 AgentId, params *AgentControllerGetAgentKeyParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsKeys, error)
	GetAgentNoGroup(params *AgentControllerGetAgentNoGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgents, error)
	GetAgentOutdated(params *AgentControllerGetAgentOutdatedParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsSimple, error)
	GetAgentSummaryOs(params *AgentControllerGetAgentSummaryOsParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetAgentSummaryStatus(params *AgentControllerGetAgentSummaryStatusParams, reqEditors ...RequestEditorFn) (*AgentsSummaryStatus, error)
	GetAgentUpgrade(params *AgentControllerGetAgentUpgradeParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetAgentsInGroup(arg1 GroupId, params *AgentControllerGetAgentsInGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgents, error)
	GetAgents(params *AgentControllerGetAgentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgents, error)
	GetComponentStats(arg1 AgentId, arg2 AgentControllerGetComponentStatsParamsComponent, params *AgentControllerGetComponentStatsParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetGroupConfig(arg1 GroupId, params *AgentControllerGetGroupConfigParams, reqEditors ...RequestEditorFn) (*struct {
		AffectedItems      *[]GroupConfiguration "json:\"affected_items,omitempty\""
		TotalAffectedItems *int32                "json:\"total_affected_items,omitempty\""
	}, error)
	GetGroupFileJSON(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileJsonParams, reqEditors ...RequestEditorFn) (*interface{}, error)
	GetGroupFileXML(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileXmlParams, reqEditors ...RequestEditorFn) (*AgentControllerGetGroupFileXmlResponse, error)
	GetGroupFiles(arg1 GroupId, params *AgentControllerGetGroupFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error)
	GetListGroup(params *AgentControllerGetListGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseGroups, error)
	GetSyncAgent(arg1 AgentId, params *AgentControllerGetSyncAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsSynced, error)
	InsertAgentWithBody(params *AgentControllerInsertAgentParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AgentIdKey, error)
	InsertAgent(params *AgentControllerInsertAgentParams, arg2 AgentControllerInsertAgentJSONRequestBody, reqEditors ...RequestEditorFn) (*AgentIdKey, error)
	PostGroupWithBody(params *AgentControllerPostGroupParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	PostGroup(params *AgentControllerPostGroupParams, arg2 AgentControllerPostGroupJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	PostNewAgent(params *AgentControllerPostNewAgentParams, reqEditors ...RequestEditorFn) (*AgentIdKey, error)
	PutAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerPutAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	PutGroupConfigWithBody(arg1 GroupId, params *AgentControllerPutGroupConfigParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	PutMultipleAgentSingleGroup(params *AgentControllerPutMultipleAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}, error)
	PutUpgradeAgents(params *AgentControllerPutUpgradeAgentsParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	PutUpgradeCustomAgents(params *AgentControllerPutUpgradeCustomAgentsParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	ReconnectAgents(params *AgentControllerReconnectAgentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	RestartAgent(arg1 AgentId, params *AgentControllerRestartAgentParams, reqEditors ...RequestEditorFn) (*ItemAffected, error)
	RestartAgentsByGroup(arg1 GroupId, params *AgentControllerRestartAgentsByGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	RestartAgentsByNode(arg1 NodeId, params *AgentControllerRestartAgentsByNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
	RestartAgents(params *AgentControllerRestartAgentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error)
}

// ManagerControllerInterface contains all methods for the wazuh controller api
type ManagerControllerInterface interface {
	GetAPIConfig(params *ManagerControllerGetApiConfigParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetConfValidation(params *ManagerControllerGetConfValidationParams, reqEditors ...RequestEditorFn) (*ConfigurationValidation, error)
	GetConfiguration(params *ManagerControllerGetConfigurationParams, reqEditors ...RequestEditorFn) (*WazuhManagerConfiguration, error)
	GetInfo(params *ManagerControllerGetInfoParams, reqEditors ...RequestEditorFn) (*WazuhInfo, error)
	GetLogSummary(params *ManagerControllerGetLogSummaryParams, reqEditors ...RequestEditorFn) (*WazuhLogsSummary, error)
	GetLog(params *ManagerControllerGetLogParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error)
	GetManagerConfigOndemand(arg1 ManagerControllerGetManagerConfigOndemandParamsComponent, arg2 ManagerControllerGetManagerConfigOndemandParamsConfiguration, params *ManagerControllerGetManagerConfigOndemandParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetStatsAnalysisd(params *ManagerControllerGetStatsAnalysisdParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsHourly(params *ManagerControllerGetStatsHourlyParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsRemoted(params *ManagerControllerGetStatsRemotedParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatsWeekly(params *ManagerControllerGetStatsWeeklyParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStats(params *ManagerControllerGetStatsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error)
	GetStatus(params *ManagerControllerGetStatusParams, reqEditors ...RequestEditorFn) (*WazuhDaemonsStatus, error)
	PutRestart(params *ManagerControllerPutRestartParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	UpdateConfigurationWithBody(params *ManagerControllerUpdateConfigurationParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}, error)
}

// SecurityControllerInterface contains all methods for the wazuh controller api
type SecurityControllerInterface interface {
	AddPolicyWithBody(params *SecurityControllerAddPolicyParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	AddPolicy(params *SecurityControllerAddPolicyParams, arg2 SecurityControllerAddPolicyJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	AddRoleWithBody(params *SecurityControllerAddRoleParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	AddRole(params *SecurityControllerAddRoleParams, arg2 SecurityControllerAddRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	AddRuleWithBody(params *SecurityControllerAddRuleParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	AddRule(params *SecurityControllerAddRuleParams, arg2 SecurityControllerAddRuleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	CreateUserWithBody(params *SecurityControllerCreateUserParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	CreateUser(params *SecurityControllerCreateUserParams, arg2 SecurityControllerCreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	DeleteSecurityConfig(params *SecurityControllerDeleteSecurityConfigParams, reqEditors ...RequestEditorFn) (*map[string]interface{}, error)
	DeleteUsers(params *SecurityControllerDeleteUsersParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	EditRunAs(arg1 UserIdRequired, params *SecurityControllerEditRunAsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	GetPolicies(params *SecurityControllerGetPoliciesParams, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	GetRbacActions(params *SecurityControllerGetRbacActionsParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetRbacResources(params *SecurityControllerGetRbacResourcesParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	GetRoles(params *SecurityControllerGetRolesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	GetRules(params *SecurityControllerGetRulesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	GetSecurityConfig(params *SecurityControllerGetSecurityConfigParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetUserMePolicies(params *SecurityControllerGetUserMePoliciesParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetUserMe(params *SecurityControllerGetUserMeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	GetUsers(params *SecurityControllerGetUsersParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	LoginUser(params *SecurityControllerLoginUserParams, reqEditors ...RequestEditorFn) (*Token, error)
	LogoutUser(reqEditors ...RequestEditorFn) (*ApiResponse, error)
	PutSecurityConfigWithBody(params *SecurityControllerPutSecurityConfigParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*map[string]interface{}, error)
	PutSecurityConfig(params *SecurityControllerPutSecurityConfigParams, arg2 SecurityControllerPutSecurityConfigJSONRequestBody, reqEditors ...RequestEditorFn) (*map[string]interface{}, error)
	RemovePolicies(params *SecurityControllerRemovePoliciesParams, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	RemoveRolePolicy(arg1 RoleId, params *SecurityControllerRemoveRolePolicyParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	RemoveRoleRule(arg1 RoleId, params *SecurityControllerRemoveRoleRuleParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	RemoveRoles(params *SecurityControllerRemoveRolesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	RemoveRules(params *SecurityControllerRemoveRulesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	RemoveUserRole(arg1 UserIdRequired, params *SecurityControllerRemoveUserRoleParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	RevokeAllTokens(reqEditors ...RequestEditorFn) (*map[string]interface{}, error)
	RunAsLoginWithBody(params *SecurityControllerRunAsLoginParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Token *string "json:\"token,omitempty\""
	}, error)
	RunAsLogin(params *SecurityControllerRunAsLoginParams, arg2 SecurityControllerRunAsLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Token *string "json:\"token,omitempty\""
	}, error)
	SetRolePolicy(arg1 RoleId, params *SecurityControllerSetRolePolicyParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	SetRoleRule(arg1 RoleId, params *SecurityControllerSetRoleRuleParams, reqEditors ...RequestEditorFn) (*struct {
		ApiResponse "yaml:\",inline\""
	}, error)
	SetUserRole(arg1 UserIdRequired, params *SecurityControllerSetUserRoleParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdatePolicyWithBody(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	UpdatePolicy(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 SecurityControllerUpdatePolicyJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error)
	UpdateRoleWithBody(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	UpdateRole(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 SecurityControllerUpdateRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error)
	UpdateRuleWithBody(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdateRule(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 SecurityControllerUpdateRuleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdateUserWithBody(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
	UpdateUser(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 SecurityControllerUpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error)
}

// TaskControllerInterface contains all methods for the wazuh controller api
type TaskControllerInterface interface {
	GetTasksStatus(params *TaskControllerGetTasksStatusParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
}

// VulnerabilityControllerInterface contains all methods for the wazuh controller api
type VulnerabilityControllerInterface interface {
	GetLastScanAgent(arg1 AgentId, params *VulnerabilityControllerGetLastScanAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseLastScan, error)
	GetVulnerabilitiesFieldSummary(arg1 AgentId, arg2 VulnerabilityControllerGetVulnerabilitiesFieldSummaryParamsField, params *VulnerabilityControllerGetVulnerabilitiesFieldSummaryParams, reqEditors ...RequestEditorFn) (*ApiResponse, error)
	GetVulnerabilityAgent(arg1 AgentId, params *VulnerabilityControllerGetVulnerabilityAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseVulnerabilities, error)
}
