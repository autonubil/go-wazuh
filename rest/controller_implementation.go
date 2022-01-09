package rest

import "io"

// RootcheckController implementation of the RootcheckController interface
type RootcheckController struct {
	*ClientWithResponses
}

// DeleteRootcheck calls the Rootcheck controller´s function
func (c *RootcheckController) DeleteRootcheck(params *RootcheckControllerDeleteRootcheckParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerDeleteRootcheckWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLastScanAgent calls the Rootcheck controller´s function
func (c *RootcheckController) GetLastScanAgent(arg1 AgentId, params *RootcheckControllerGetLastScanAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerGetLastScanAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRootcheckAgent calls the Rootcheck controller´s function
func (c *RootcheckController) GetRootcheckAgent(arg1 AgentId, params *RootcheckControllerGetRootcheckAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerGetRootcheckAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutRootcheck calls the Rootcheck controller´s function
func (c *RootcheckController) PutRootcheck(params *RootcheckControllerPutRootcheckParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerPutRootcheckWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ActiveResponseController implementation of the ActiveResponseController interface
type ActiveResponseController struct {
	*ClientWithResponses
}

// RunCommandWithBody calls the ActiveResponse controller´s function
func (c *ActiveResponseController) RunCommandWithBody(params *ActiveResponseControllerRunCommandParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ActiveResponseControllerRunCommandWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RunCommand calls the ActiveResponse controller´s function
func (c *ActiveResponseController) RunCommand(params *ActiveResponseControllerRunCommandParams, arg2 ActiveResponseControllerRunCommandJSONRequestBody, arg3 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ActiveResponseControllerRunCommandWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CiscatController implementation of the CiscatController interface
type CiscatController struct {
	*ClientWithResponses
}

// GetAgentsCiscatResults calls the Ciscat controller´s function
func (c *CiscatController) GetAgentsCiscatResults(arg1 AgentId, params *CiscatControllerGetAgentsCiscatResultsParams, arg3 ...RequestEditorFn) (*AllItemsResponseCiscatResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CiscatControllerGetAgentsCiscatResultsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseCiscatResult
	if i, ok := r.(*AllItemsResponseCiscatResult); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DefaultController implementation of the DefaultController interface
type DefaultController struct {
	*ClientWithResponses
}

// DefaultInfo calls the Default controller´s function
func (c *DefaultController) DefaultInfo(params *DefaultControllerDefaultInfoParams, arg2 ...RequestEditorFn) (*BasicInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DefaultControllerDefaultInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *BasicInfo
	if i, ok := r.(*BasicInfo); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// LogtestController implementation of the LogtestController interface
type LogtestController struct {
	*ClientWithResponses
}

// EndLogtestSession calls the Logtest controller´s function
func (c *LogtestController) EndLogtestSession(arg1 LogtestToken, params *LogtestControllerEndLogtestSessionParams, arg3 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.LogtestControllerEndLogtestSessionWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RunLogtestToolWithBody calls the Logtest controller´s function
func (c *LogtestController) RunLogtestToolWithBody(params *LogtestControllerRunLogtestToolParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.LogtestControllerRunLogtestToolWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RunLogtestTool calls the Logtest controller´s function
func (c *LogtestController) RunLogtestTool(params *LogtestControllerRunLogtestToolParams, arg2 LogtestControllerRunLogtestToolJSONRequestBody, arg3 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.LogtestControllerRunLogtestToolWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ManagerController implementation of the ManagerController interface
type ManagerController struct {
	*ClientWithResponses
}

// GetAPIConfig calls the Manager controller´s function
func (c *ManagerController) GetAPIConfig(params *ManagerControllerGetApiConfigParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetApiConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetConfValidation calls the Manager controller´s function
func (c *ManagerController) GetConfValidation(params *ManagerControllerGetConfValidationParams, arg2 ...RequestEditorFn) (*ConfigurationValidation, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetConfValidationWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ConfigurationValidation
	if i, ok := r.(*ConfigurationValidation); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetConfiguration calls the Manager controller´s function
func (c *ManagerController) GetConfiguration(params *ManagerControllerGetConfigurationParams, arg2 ...RequestEditorFn) (*WazuhManagerConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetConfigurationWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhManagerConfiguration
	if i, ok := r.(*WazuhManagerConfiguration); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetInfo calls the Manager controller´s function
func (c *ManagerController) GetInfo(params *ManagerControllerGetInfoParams, arg2 ...RequestEditorFn) (*WazuhInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhInfo
	if i, ok := r.(*WazuhInfo); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLogSummary calls the Manager controller´s function
func (c *ManagerController) GetLogSummary(params *ManagerControllerGetLogSummaryParams, arg2 ...RequestEditorFn) (*WazuhLogsSummary, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetLogSummaryWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhLogsSummary
	if i, ok := r.(*WazuhLogsSummary); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLog calls the Manager controller´s function
func (c *ManagerController) GetLog(params *ManagerControllerGetLogParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetLogWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhLogs
	if i, ok := r.(*AllItemsResponseWazuhLogs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetManagerConfigOndemand calls the Manager controller´s function
func (c *ManagerController) GetManagerConfigOndemand(arg1 ManagerControllerGetManagerConfigOndemandParamsComponent, arg2 ManagerControllerGetManagerConfigOndemandParamsConfiguration, params *ManagerControllerGetManagerConfigOndemandParams, arg4 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetManagerConfigOndemandWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsAnalysisd calls the Manager controller´s function
func (c *ManagerController) GetStatsAnalysisd(params *ManagerControllerGetStatsAnalysisdParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsAnalysisdWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsHourly calls the Manager controller´s function
func (c *ManagerController) GetStatsHourly(params *ManagerControllerGetStatsHourlyParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsHourlyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsRemoted calls the Manager controller´s function
func (c *ManagerController) GetStatsRemoted(params *ManagerControllerGetStatsRemotedParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsRemotedWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsWeekly calls the Manager controller´s function
func (c *ManagerController) GetStatsWeekly(params *ManagerControllerGetStatsWeeklyParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsWeeklyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStats calls the Manager controller´s function
func (c *ManagerController) GetStats(params *ManagerControllerGetStatsParams, arg2 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatus calls the Manager controller´s function
func (c *ManagerController) GetStatus(params *ManagerControllerGetStatusParams, arg2 ...RequestEditorFn) (*WazuhDaemonsStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatusWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhDaemonsStatus
	if i, ok := r.(*WazuhDaemonsStatus); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutRestart calls the Manager controller´s function
func (c *ManagerController) PutRestart(params *ManagerControllerPutRestartParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerPutRestartWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateConfigurationWithBody calls the Manager controller´s function
func (c *ManagerController) UpdateConfigurationWithBody(params *ManagerControllerUpdateConfigurationParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerUpdateConfigurationWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\""; ConfirmationMessage "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// MitreController implementation of the MitreController interface
type MitreController struct {
	*ClientWithResponses
}

// GetAttack calls the Mitre controller´s function
func (c *MitreController) GetAttack(params *MitreControllerGetAttackParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetAttackWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// OverviewController implementation of the OverviewController interface
type OverviewController struct {
	*ClientWithResponses
}

// GetOverviewAgents calls the Overview controller´s function
func (c *OverviewController) GetOverviewAgents(params *OverviewControllerGetOverviewAgentsParams, arg2 ...RequestEditorFn) (*OverviewAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.OverviewControllerGetOverviewAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *OverviewAgents
	if i, ok := r.(*OverviewAgents); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// SyscheckController implementation of the SyscheckController interface
type SyscheckController struct {
	*ClientWithResponses
}

// DeleteSyscheckAgent calls the Syscheck controller´s function
func (c *SyscheckController) DeleteSyscheckAgent(arg1 AgentId, params *SyscheckControllerDeleteSyscheckAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerDeleteSyscheckAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLastScanAgent calls the Syscheck controller´s function
func (c *SyscheckController) GetLastScanAgent(arg1 AgentId, params *SyscheckControllerGetLastScanAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscheckLastScan, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerGetLastScanAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscheckLastScan
	if i, ok := r.(*AllItemsResponseSyscheckLastScan); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetSyscheckAgent calls the Syscheck controller´s function
func (c *SyscheckController) GetSyscheckAgent(arg1 AgentId, params *SyscheckControllerGetSyscheckAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscheckResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerGetSyscheckAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscheckResult
	if i, ok := r.(*AllItemsResponseSyscheckResult); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutSyscheck calls the Syscheck controller´s function
func (c *SyscheckController) PutSyscheck(params *SyscheckControllerPutSyscheckParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerPutSyscheckWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentIDs
	if i, ok := r.(*AllItemsResponseAgentIDs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// SyscollectorController implementation of the SyscollectorController interface
type SyscollectorController struct {
	*ClientWithResponses
}

// GetHardwareInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetHardwareInfo(arg1 AgentId, params *SyscollectorControllerGetHardwareInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetHardwareInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorHardware
	if i, ok := r.(*AllItemsResponseSyscollectorHardware); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetHotfixInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetHotfixInfo(arg1 AgentId, params *SyscollectorControllerGetHotfixInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetHotfixInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorHotfixes
	if i, ok := r.(*AllItemsResponseSyscollectorHotfixes); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNetworkAddressInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetNetworkAddressInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkAddressInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkAddressInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorNetwork
	if i, ok := r.(*AllItemsResponseSyscollectorNetwork); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNetworkInterfaceInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetNetworkInterfaceInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkInterfaceInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkInterfaceInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorInterface
	if i, ok := r.(*AllItemsResponseSyscollectorInterface); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNetworkProtocolInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetNetworkProtocolInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkProtocolInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkProtocolInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorProtocol
	if i, ok := r.(*AllItemsResponseSyscollectorProtocol); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetOsInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetOsInfo(arg1 AgentId, params *SyscollectorControllerGetOsInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetOsInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorOS
	if i, ok := r.(*AllItemsResponseSyscollectorOS); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetPackagesInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetPackagesInfo(arg1 AgentId, params *SyscollectorControllerGetPackagesInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetPackagesInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorPackages
	if i, ok := r.(*AllItemsResponseSyscollectorPackages); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetPortsInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetPortsInfo(arg1 AgentId, params *SyscollectorControllerGetPortsInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetPortsInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorPorts
	if i, ok := r.(*AllItemsResponseSyscollectorPorts); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetProcessesInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetProcessesInfo(arg1 AgentId, params *SyscollectorControllerGetProcessesInfoParams, arg3 ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetProcessesInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorProcesses
	if i, ok := r.(*AllItemsResponseSyscollectorProcesses); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ClusterController implementation of the ClusterController interface
type ClusterController struct {
	*ClientWithResponses
}

// GetAPIConfig calls the Cluster controller´s function
func (c *ClusterController) GetAPIConfig(params *ClusterControllerGetApiConfigParams, arg2 ...RequestEditorFn) (*struct {
	AdditionalProperties map[string]interface{} "json:\"-\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetApiConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AdditionalProperties map[string]interface {} "json:\"-\"" }
	if i, ok := r.(*struct {
		AdditionalProperties map[string]interface{} "json:\"-\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetClusterNode calls the Cluster controller´s function
func (c *ClusterController) GetClusterNode(params *ClusterControllerGetClusterNodeParams, arg2 ...RequestEditorFn) (*struct {
	Cluster *string "json:\"cluster,omitempty\""
	Node    *string "json:\"node,omitempty\""
	Type    *string "json:\"type,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetClusterNodeWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Cluster *string "json:\"cluster,omitempty\""; Node *string "json:\"node,omitempty\""; Type *string "json:\"type,omitempty\"" }
	if i, ok := r.(*struct {
		Cluster *string "json:\"cluster,omitempty\""
		Node    *string "json:\"node,omitempty\""
		Type    *string "json:\"type,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetClusterNodes calls the Cluster controller´s function
func (c *ClusterController) GetClusterNodes(params *ClusterControllerGetClusterNodesParams, arg2 ...RequestEditorFn) (*AllItemsResponseClusterNodes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetClusterNodesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseClusterNodes
	if i, ok := r.(*AllItemsResponseClusterNodes); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetConfValidation calls the Cluster controller´s function
func (c *ClusterController) GetConfValidation(params *ClusterControllerGetConfValidationParams, arg2 ...RequestEditorFn) (*AllItemsResponseValidationStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfValidationWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseValidationStatus
	if i, ok := r.(*AllItemsResponseValidationStatus); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetConfig calls the Cluster controller´s function
func (c *ClusterController) GetConfig(params *ClusterControllerGetConfigParams, arg2 ...RequestEditorFn) (*struct {
	BindAddr *string           "json:\"bind_addr,omitempty\""
	Disabled *bool             "json:\"disabled,omitempty\""
	Hidden   *string           "json:\"hidden,omitempty\""
	Key      *string           "json:\"key,omitempty\""
	Name     *string           "json:\"name,omitempty\""
	NodeName *string           "json:\"node_name,omitempty\""
	NodeType *N200DataNodeType "json:\"node_type,omitempty\""
	Nodes    *[]string         "json:\"nodes,omitempty\""
	Port     *int              "json:\"port,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { BindAddr *string "json:\"bind_addr,omitempty\""; Disabled *bool "json:\"disabled,omitempty\""; Hidden *string "json:\"hidden,omitempty\""; Key *string "json:\"key,omitempty\""; Name *string "json:\"name,omitempty\""; NodeName *string "json:\"node_name,omitempty\""; NodeType *N200DataNodeType "json:\"node_type,omitempty\""; Nodes *[]string "json:\"nodes,omitempty\""; Port *int "json:\"port,omitempty\"" }
	if i, ok := r.(*struct {
		BindAddr *string           "json:\"bind_addr,omitempty\""
		Disabled *bool             "json:\"disabled,omitempty\""
		Hidden   *string           "json:\"hidden,omitempty\""
		Key      *string           "json:\"key,omitempty\""
		Name     *string           "json:\"name,omitempty\""
		NodeName *string           "json:\"node_name,omitempty\""
		NodeType *N200DataNodeType "json:\"node_type,omitempty\""
		Nodes    *[]string         "json:\"nodes,omitempty\""
		Port     *int              "json:\"port,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetConfigurationNode calls the Cluster controller´s function
func (c *ClusterController) GetConfigurationNode(arg1 NodeId, params *ClusterControllerGetConfigurationNodeParams, arg3 ...RequestEditorFn) (*WazuhManagerConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfigurationNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhManagerConfiguration
	if i, ok := r.(*WazuhManagerConfiguration); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetHealthcheck calls the Cluster controller´s function
func (c *ClusterController) GetHealthcheck(params *ClusterControllerGetHealthcheckParams, arg2 ...RequestEditorFn) (*AllItemsResponseNodeHealthcheck, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetHealthcheckWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseNodeHealthcheck
	if i, ok := r.(*AllItemsResponseNodeHealthcheck); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetInfoNode calls the Cluster controller´s function
func (c *ClusterController) GetInfoNode(arg1 NodeId, params *ClusterControllerGetInfoNodeParams, arg3 ...RequestEditorFn) (*WazuhInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetInfoNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhInfo
	if i, ok := r.(*WazuhInfo); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLogNode calls the Cluster controller´s function
func (c *ClusterController) GetLogNode(arg1 NodeId, params *ClusterControllerGetLogNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetLogNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhLogs
	if i, ok := r.(*AllItemsResponseWazuhLogs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLogSummaryNode calls the Cluster controller´s function
func (c *ClusterController) GetLogSummaryNode(arg1 NodeId, params *ClusterControllerGetLogSummaryNodeParams, arg3 ...RequestEditorFn) (*WazuhLogsSummary, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetLogSummaryNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhLogsSummary
	if i, ok := r.(*WazuhLogsSummary); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNodeConfig calls the Cluster controller´s function
func (c *ClusterController) GetNodeConfig(arg1 NodeId, arg2 ClusterControllerGetNodeConfigParamsComponent, arg3 ClusterControllerGetNodeConfigParamsConfiguration, params *ClusterControllerGetNodeConfigParams, arg5 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetNodeConfigWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, params, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsAnalysisdNode calls the Cluster controller´s function
func (c *ClusterController) GetStatsAnalysisdNode(arg1 NodeId, params *ClusterControllerGetStatsAnalysisdNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsAnalysisdNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsHourlyNode calls the Cluster controller´s function
func (c *ClusterController) GetStatsHourlyNode(arg1 NodeId, params *ClusterControllerGetStatsHourlyNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsHourlyNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsNode calls the Cluster controller´s function
func (c *ClusterController) GetStatsNode(arg1 NodeId, params *ClusterControllerGetStatsNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsRemotedNode calls the Cluster controller´s function
func (c *ClusterController) GetStatsRemotedNode(arg1 NodeId, params *ClusterControllerGetStatsRemotedNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsRemotedNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatsWeeklyNode calls the Cluster controller´s function
func (c *ClusterController) GetStatsWeeklyNode(arg1 NodeId, params *ClusterControllerGetStatsWeeklyNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsWeeklyNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseWazuhStats
	if i, ok := r.(*AllItemsResponseWazuhStats); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatusNode calls the Cluster controller´s function
func (c *ClusterController) GetStatusNode(arg1 NodeId, params *ClusterControllerGetStatusNodeParams, arg3 ...RequestEditorFn) (*WazuhDaemonsStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatusNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhDaemonsStatus
	if i, ok := r.(*WazuhDaemonsStatus); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetStatus calls the Cluster controller´s function
func (c *ClusterController) GetStatus(params *ClusterControllerGetStatusParams, arg2 ...RequestEditorFn) (*struct {
	Enabled *N200DataEnabled "json:\"enabled,omitempty\""
	Running *N200DataRunning "json:\"running,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatusWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Enabled *N200DataEnabled "json:\"enabled,omitempty\""; Running *N200DataRunning "json:\"running,omitempty\"" }
	if i, ok := r.(*struct {
		Enabled *N200DataEnabled "json:\"enabled,omitempty\""
		Running *N200DataRunning "json:\"running,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutRestart calls the Cluster controller´s function
func (c *ClusterController) PutRestart(params *ClusterControllerPutRestartParams, arg2 ...RequestEditorFn) (*AllItemsResponseNodeIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerPutRestartWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseNodeIDs
	if i, ok := r.(*AllItemsResponseNodeIDs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateConfigurationWithBody calls the Cluster controller´s function
func (c *ClusterController) UpdateConfigurationWithBody(arg1 NodeId, params *ClusterControllerUpdateConfigurationParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerUpdateConfigurationWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\""; ConfirmationMessage "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExperimentalController implementation of the ExperimentalController interface
type ExperimentalController struct {
	*ClientWithResponses
}

// ClearSyscheckDatabase calls the Experimental controller´s function
func (c *ExperimentalController) ClearSyscheckDatabase(params *ExperimentalControllerClearSyscheckDatabaseParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerClearSyscheckDatabaseWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentIDs
	if i, ok := r.(*AllItemsResponseAgentIDs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetCisCatResults calls the Experimental controller´s function
func (c *ExperimentalController) GetCisCatResults(params *ExperimentalControllerGetCisCatResultsParams, arg2 ...RequestEditorFn) (*AllItemsResponseCiscatResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetCisCatResultsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseCiscatResult
	if i, ok := r.(*AllItemsResponseCiscatResult); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetHardwareInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetHardwareInfo(params *ExperimentalControllerGetHardwareInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetHardwareInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorHardware
	if i, ok := r.(*AllItemsResponseSyscollectorHardware); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetHotfixesInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetHotfixesInfo(params *ExperimentalControllerGetHotfixesInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetHotfixesInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorHotfixes
	if i, ok := r.(*AllItemsResponseSyscollectorHotfixes); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNetworkAddressInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetNetworkAddressInfo(params *ExperimentalControllerGetNetworkAddressInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkAddressInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorNetwork
	if i, ok := r.(*AllItemsResponseSyscollectorNetwork); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNetworkInterfaceInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetNetworkInterfaceInfo(params *ExperimentalControllerGetNetworkInterfaceInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkInterfaceInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorInterface
	if i, ok := r.(*AllItemsResponseSyscollectorInterface); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetNetworkProtocolInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetNetworkProtocolInfo(params *ExperimentalControllerGetNetworkProtocolInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkProtocolInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorProtocol
	if i, ok := r.(*AllItemsResponseSyscollectorProtocol); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetOsInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetOsInfo(params *ExperimentalControllerGetOsInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetOsInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorOS
	if i, ok := r.(*AllItemsResponseSyscollectorOS); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetPackagesInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetPackagesInfo(params *ExperimentalControllerGetPackagesInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetPackagesInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorPackages
	if i, ok := r.(*AllItemsResponseSyscollectorPackages); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetPortsInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetPortsInfo(params *ExperimentalControllerGetPortsInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetPortsInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorPorts
	if i, ok := r.(*AllItemsResponseSyscollectorPorts); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetProcessesInfo calls the Experimental controller´s function
func (c *ExperimentalController) GetProcessesInfo(params *ExperimentalControllerGetProcessesInfoParams, arg2 ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetProcessesInfoWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSyscollectorProcesses
	if i, ok := r.(*AllItemsResponseSyscollectorProcesses); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ScaController implementation of the ScaController interface
type ScaController struct {
	*ClientWithResponses
}

// GetScaAgent calls the Sca controller´s function
func (c *ScaController) GetScaAgent(arg1 AgentId, params *ScaControllerGetScaAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseSCADatabase, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScaControllerGetScaAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSCADatabase
	if i, ok := r.(*AllItemsResponseSCADatabase); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetScaChecks calls the Sca controller´s function
func (c *ScaController) GetScaChecks(arg1 AgentId, arg2 PolicyId, params *ScaControllerGetScaChecksParams, arg4 ...RequestEditorFn) (*AllItemsResponseSCAChecks, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScaControllerGetScaChecksWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseSCAChecks
	if i, ok := r.(*AllItemsResponseSCAChecks); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// VulnerabilityController implementation of the VulnerabilityController interface
type VulnerabilityController struct {
	*ClientWithResponses
}

// GetVulnerabilityAgent calls the Vulnerability controller´s function
func (c *VulnerabilityController) GetVulnerabilityAgent(arg1 AgentId, params *VulnerabilityControllerGetVulnerabilityAgentParams, arg3 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilityControllerGetVulnerabilityAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AgentController implementation of the AgentController interface
type AgentController struct {
	*ClientWithResponses
}

// AddAgentWithBody calls the Agent controller´s function
func (c *AgentController) AddAgentWithBody(params *AgentControllerAddAgentParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerAddAgentWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentIdKey
	if i, ok := r.(*AgentIdKey); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddAgent calls the Agent controller´s function
func (c *AgentController) AddAgent(params *AgentControllerAddAgentParams, arg2 AgentControllerAddAgentJSONRequestBody, arg3 ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerAddAgentWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentIdKey
	if i, ok := r.(*AgentIdKey); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteAgents calls the Agent controller´s function
func (c *AgentController) DeleteAgents(params *AgentControllerDeleteAgentsParams, arg2 ...RequestEditorFn) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllItemsResponseAgentIDs "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteGroups calls the Agent controller´s function
func (c *AgentController) DeleteGroups(params *AgentControllerDeleteGroupsParams, arg2 ...RequestEditorFn) (*struct {
	AllItemsResponseGroupIDs "yaml:\",inline\""
	AgentGroupDeleted        "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteGroupsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllItemsResponseGroupIDs "yaml:\",inline\""; AgentGroupDeleted "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		AllItemsResponseGroupIDs "yaml:\",inline\""
		AgentGroupDeleted        "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteMultipleAgentSingleGroup calls the Agent controller´s function
func (c *AgentController) DeleteMultipleAgentSingleGroup(params *AgentControllerDeleteMultipleAgentSingleGroupParams, arg2 ...RequestEditorFn) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteMultipleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllItemsResponseAgentIDs "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteSingleAgentMultipleGroups calls the Agent controller´s function
func (c *AgentController) DeleteSingleAgentMultipleGroups(arg1 AgentId, params *AgentControllerDeleteSingleAgentMultipleGroupsParams, arg3 ...RequestEditorFn) (*struct {
	AllItemsResponseGroupIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteSingleAgentMultipleGroupsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllItemsResponseGroupIDs "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		AllItemsResponseGroupIDs "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteSingleAgentSingleGroup calls the Agent controller´s function
func (c *AgentController) DeleteSingleAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerDeleteSingleAgentSingleGroupParams, arg4 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteSingleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentConfig calls the Agent controller´s function
func (c *AgentController) GetAgentConfig(arg1 AgentId, arg2 AgentControllerGetAgentConfigParamsComponent, arg3 AgentControllerGetAgentConfigParamsConfiguration, params *AgentControllerGetAgentConfigParams, arg5 ...RequestEditorFn) (*AgentConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentConfigWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, params, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentConfiguration
	if i, ok := r.(*AgentConfiguration); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentFields calls the Agent controller´s function
func (c *AgentController) GetAgentFields(params *AgentControllerGetAgentFieldsParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentsDistinct, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentFieldsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentsDistinct
	if i, ok := r.(*AllItemsResponseAgentsDistinct); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentKey calls the Agent controller´s function
func (c *AgentController) GetAgentKey(arg1 AgentId, params *AgentControllerGetAgentKeyParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentsKeys, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentKeyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentsKeys
	if i, ok := r.(*AllItemsResponseAgentsKeys); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentNoGroup calls the Agent controller´s function
func (c *AgentController) GetAgentNoGroup(params *AgentControllerGetAgentNoGroupParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentNoGroupWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgents
	if i, ok := r.(*AllItemsResponseAgents); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentOutdated calls the Agent controller´s function
func (c *AgentController) GetAgentOutdated(params *AgentControllerGetAgentOutdatedParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentsSimple, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentOutdatedWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentsSimple
	if i, ok := r.(*AllItemsResponseAgentsSimple); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentSummaryOs calls the Agent controller´s function
func (c *AgentController) GetAgentSummaryOs(params *AgentControllerGetAgentSummaryOsParams, arg2 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentSummaryOsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentSummaryStatus calls the Agent controller´s function
func (c *AgentController) GetAgentSummaryStatus(params *AgentControllerGetAgentSummaryStatusParams, arg2 ...RequestEditorFn) (*AgentsSummaryStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentSummaryStatusWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentsSummaryStatus
	if i, ok := r.(*AgentsSummaryStatus); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentUpgrade calls the Agent controller´s function
func (c *AgentController) GetAgentUpgrade(params *AgentControllerGetAgentUpgradeParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentUpgradeWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgentsInGroup calls the Agent controller´s function
func (c *AgentController) GetAgentsInGroup(arg1 GroupId, params *AgentControllerGetAgentsInGroupParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentsInGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgents
	if i, ok := r.(*AllItemsResponseAgents); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAgents calls the Agent controller´s function
func (c *AgentController) GetAgents(params *AgentControllerGetAgentsParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgents
	if i, ok := r.(*AllItemsResponseAgents); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetComponentStats calls the Agent controller´s function
func (c *AgentController) GetComponentStats(arg1 AgentId, arg2 AgentControllerGetComponentStatsParamsComponent, params *AgentControllerGetComponentStatsParams, arg4 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetComponentStatsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetGroupConfig calls the Agent controller´s function
func (c *AgentController) GetGroupConfig(arg1 GroupId, params *AgentControllerGetGroupConfigParams, arg3 ...RequestEditorFn) (*struct {
	AffectedItems      *[]GroupConfiguration "json:\"affected_items,omitempty\""
	TotalAffectedItems *int32                "json:\"total_affected_items,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupConfigWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AffectedItems *[]GroupConfiguration "json:\"affected_items,omitempty\""; TotalAffectedItems *int32 "json:\"total_affected_items,omitempty\"" }
	if i, ok := r.(*struct {
		AffectedItems      *[]GroupConfiguration "json:\"affected_items,omitempty\""
		TotalAffectedItems *int32                "json:\"total_affected_items,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetGroupFileJSON calls the Agent controller´s function
func (c *AgentController) GetGroupFileJSON(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileJsonParams, arg4 ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupFileJsonWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetGroupFileXML calls the Agent controller´s function
func (c *AgentController) GetGroupFileXML(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileXmlParams, arg4 ...RequestEditorFn) (*AgentControllerGetGroupFileXmlResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupFileXmlWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentControllerGetGroupFileXmlResponse
	if i, ok := r.(*AgentControllerGetGroupFileXmlResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetGroupFiles calls the Agent controller´s function
func (c *AgentController) GetGroupFiles(arg1 GroupId, params *AgentControllerGetGroupFilesParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupFilesWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetListGroup calls the Agent controller´s function
func (c *AgentController) GetListGroup(params *AgentControllerGetListGroupParams, arg2 ...RequestEditorFn) (*AllItemsResponseGroups, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetListGroupWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseGroups
	if i, ok := r.(*AllItemsResponseGroups); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetSyncAgent calls the Agent controller´s function
func (c *AgentController) GetSyncAgent(arg1 AgentId, params *AgentControllerGetSyncAgentParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentsSynced, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetSyncAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentsSynced
	if i, ok := r.(*AllItemsResponseAgentsSynced); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// InsertAgentWithBody calls the Agent controller´s function
func (c *AgentController) InsertAgentWithBody(params *AgentControllerInsertAgentParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerInsertAgentWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentIdKey
	if i, ok := r.(*AgentIdKey); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// InsertAgent calls the Agent controller´s function
func (c *AgentController) InsertAgent(params *AgentControllerInsertAgentParams, arg2 AgentControllerInsertAgentJSONRequestBody, arg3 ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerInsertAgentWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentIdKey
	if i, ok := r.(*AgentIdKey); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PostGroupWithBody calls the Agent controller´s function
func (c *AgentController) PostGroupWithBody(params *AgentControllerPostGroupParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPostGroupWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PostGroup calls the Agent controller´s function
func (c *AgentController) PostGroup(params *AgentControllerPostGroupParams, arg2 AgentControllerPostGroupJSONRequestBody, arg3 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPostGroupWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PostNewAgent calls the Agent controller´s function
func (c *AgentController) PostNewAgent(params *AgentControllerPostNewAgentParams, arg2 ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPostNewAgentWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AgentIdKey
	if i, ok := r.(*AgentIdKey); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutAgentSingleGroup calls the Agent controller´s function
func (c *AgentController) PutAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerPutAgentSingleGroupParams, arg4 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutGroupConfigWithBody calls the Agent controller´s function
func (c *AgentController) PutGroupConfigWithBody(arg1 GroupId, params *AgentControllerPutGroupConfigParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutGroupConfigWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutMultipleAgentSingleGroup calls the Agent controller´s function
func (c *AgentController) PutMultipleAgentSingleGroup(params *AgentControllerPutMultipleAgentSingleGroupParams, arg2 ...RequestEditorFn) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutMultipleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllItemsResponseAgentIDs "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutUpgradeAgents calls the Agent controller´s function
func (c *AgentController) PutUpgradeAgents(params *AgentControllerPutUpgradeAgentsParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutUpgradeAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutUpgradeCustomAgents calls the Agent controller´s function
func (c *AgentController) PutUpgradeCustomAgents(params *AgentControllerPutUpgradeCustomAgentsParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutUpgradeCustomAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// rtAgent calls the Agent controller´s function
func (c *AgentController) rtAgent(arg1 AgentId, params *AgentControllerRestartAgentParams, arg3 ...RequestEditorFn) (*ItemAffected, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *ItemAffected
	if i, ok := r.(*ItemAffected); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// rtAgentsByGroup calls the Agent controller´s function
func (c *AgentController) rtAgentsByGroup(arg1 GroupId, params *AgentControllerRestartAgentsByGroupParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentsByGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentIDs
	if i, ok := r.(*AllItemsResponseAgentIDs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// rtAgentsByNode calls the Agent controller´s function
func (c *AgentController) rtAgentsByNode(arg1 NodeId, params *AgentControllerRestartAgentsByNodeParams, arg3 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentsByNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentIDs
	if i, ok := r.(*AllItemsResponseAgentIDs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// rtAgents calls the Agent controller´s function
func (c *AgentController) rtAgents(params *AgentControllerRestartAgentsParams, arg2 ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseAgentIDs
	if i, ok := r.(*AllItemsResponseAgentIDs); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DecoderController implementation of the DecoderController interface
type DecoderController struct {
	*ClientWithResponses
}

// DeleteFile calls the Decoder controller´s function
func (c *DecoderController) DeleteFile(arg1 XmlFilenamePath, params *DecoderControllerDeleteFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerDeleteFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetDecodersFiles calls the Decoder controller´s function
func (c *DecoderController) GetDecodersFiles(params *DecoderControllerGetDecodersFilesParams, arg2 ...RequestEditorFn) (*AllItemsResponseDecodersFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetDecodersFilesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseDecodersFiles
	if i, ok := r.(*AllItemsResponseDecodersFiles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetDecodersParents calls the Decoder controller´s function
func (c *DecoderController) GetDecodersParents(params *DecoderControllerGetDecodersParentsParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetDecodersParentsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetDecoders calls the Decoder controller´s function
func (c *DecoderController) GetDecoders(params *DecoderControllerGetDecodersParams, arg2 ...RequestEditorFn) (*AllItemsResponseDecoders, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetDecodersWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseDecoders
	if i, ok := r.(*AllItemsResponseDecoders); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetFile calls the Decoder controller´s function
func (c *DecoderController) GetFile(arg1 XmlFilenamePath, params *DecoderControllerGetFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutFileWithBody calls the Decoder controller´s function
func (c *DecoderController) PutFileWithBody(arg1 XmlFilenamePath, params *DecoderControllerPutFileParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerPutFileWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CdbListController implementation of the CdbListController interface
type CdbListController struct {
	*ClientWithResponses
}

// DeleteFile calls the CdbList controller´s function
func (c *CdbListController) DeleteFile(arg1 ListFilenamePath, params *CdbListControllerDeleteFileParams, arg3 ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerDeleteFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\""; ConfirmationMessage "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetFile calls the CdbList controller´s function
func (c *CdbListController) GetFile(arg1 ListFilenamePath, params *CdbListControllerGetFileParams, arg3 ...RequestEditorFn) (*AllItemsResponseListsFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerGetFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseListsFiles
	if i, ok := r.(*AllItemsResponseListsFiles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetListsFiles calls the CdbList controller´s function
func (c *CdbListController) GetListsFiles(params *CdbListControllerGetListsFilesParams, arg2 ...RequestEditorFn) (*AllItemsResponseListsFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerGetListsFilesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseListsFiles
	if i, ok := r.(*AllItemsResponseListsFiles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLists calls the CdbList controller´s function
func (c *CdbListController) GetLists(params *CdbListControllerGetListsParams, arg2 ...RequestEditorFn) (*AllItemsResponseLists, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerGetListsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseLists
	if i, ok := r.(*AllItemsResponseLists); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutFileWithBody calls the CdbList controller´s function
func (c *CdbListController) PutFileWithBody(arg1 ListFilenamePath, params *CdbListControllerPutFileParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerPutFileWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\""; ConfirmationMessage "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse         "yaml:\",inline\""
		ConfirmationMessage "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RuleController implementation of the RuleController interface
type RuleController struct {
	*ClientWithResponses
}

// DeleteFile calls the Rule controller´s function
func (c *RuleController) DeleteFile(arg1 XmlFilenamePath, params *RuleControllerDeleteFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerDeleteFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetFile calls the Rule controller´s function
func (c *RuleController) GetFile(arg1 XmlFilenamePath, params *RuleControllerGetFileParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRulesFiles calls the Rule controller´s function
func (c *RuleController) GetRulesFiles(params *RuleControllerGetRulesFilesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRulesFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesFilesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRulesFiles
	if i, ok := r.(*AllItemsResponseRulesFiles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRulesGroups calls the Rule controller´s function
func (c *RuleController) GetRulesGroups(params *RuleControllerGetRulesGroupsParams, arg2 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesGroupsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRulesRequirement calls the Rule controller´s function
func (c *RuleController) GetRulesRequirement(arg1 RuleControllerGetRulesRequirementParamsRequirement, params *RuleControllerGetRulesRequirementParams, arg3 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesRequirementWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRules calls the Rule controller´s function
func (c *RuleController) GetRules(params *RuleControllerGetRulesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRules, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRules
	if i, ok := r.(*AllItemsResponseRules); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutFileWithBody calls the Rule controller´s function
func (c *RuleController) PutFileWithBody(arg1 XmlFilenamePath, params *RuleControllerPutFileParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerPutFileWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponse
	if i, ok := r.(*AllItemsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// SecurityController implementation of the SecurityController interface
type SecurityController struct {
	*ClientWithResponses
}

// AddPolicyWithBody calls the Security controller´s function
func (c *SecurityController) AddPolicyWithBody(params *SecurityControllerAddPolicyParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddPolicyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponsePolicies
	if i, ok := r.(*AllItemsResponsePolicies); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddPolicy calls the Security controller´s function
func (c *SecurityController) AddPolicy(params *SecurityControllerAddPolicyParams, arg2 SecurityControllerAddPolicyJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddPolicyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponsePolicies
	if i, ok := r.(*AllItemsResponsePolicies); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddRoleWithBody calls the Security controller´s function
func (c *SecurityController) AddRoleWithBody(params *SecurityControllerAddRoleParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRoleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddRole calls the Security controller´s function
func (c *SecurityController) AddRole(params *SecurityControllerAddRoleParams, arg2 SecurityControllerAddRoleJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRoleWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddRuleWithBody calls the Security controller´s function
func (c *SecurityController) AddRuleWithBody(params *SecurityControllerAddRuleParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRuleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddRule calls the Security controller´s function
func (c *SecurityController) AddRule(params *SecurityControllerAddRuleParams, arg2 SecurityControllerAddRuleJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRuleWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateUserWithBody calls the Security controller´s function
func (c *SecurityController) CreateUserWithBody(params *SecurityControllerCreateUserParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerCreateUserWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateUser calls the Security controller´s function
func (c *SecurityController) CreateUser(params *SecurityControllerCreateUserParams, arg2 SecurityControllerCreateUserJSONRequestBody, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerCreateUserWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteSecurityConfig calls the Security controller´s function
func (c *SecurityController) DeleteSecurityConfig(params *SecurityControllerDeleteSecurityConfigParams, arg2 ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerDeleteSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *map[string]interface {}
	if i, ok := r.(*map[string]interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteUsers calls the Security controller´s function
func (c *SecurityController) DeleteUsers(params *SecurityControllerDeleteUsersParams, arg2 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerDeleteUsersWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditRunAs calls the Security controller´s function
func (c *SecurityController) EditRunAs(arg1 UserIdRequired, params *SecurityControllerEditRunAsParams, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerEditRunAsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetPolicies calls the Security controller´s function
func (c *SecurityController) GetPolicies(params *SecurityControllerGetPoliciesParams, arg2 ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetPoliciesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponsePolicies
	if i, ok := r.(*AllItemsResponsePolicies); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRbacActions calls the Security controller´s function
func (c *SecurityController) GetRbacActions(params *SecurityControllerGetRbacActionsParams, arg2 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRbacActionsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRbacResources calls the Security controller´s function
func (c *SecurityController) GetRbacResources(params *SecurityControllerGetRbacResourcesParams, arg2 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRbacResourcesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRoles calls the Security controller´s function
func (c *SecurityController) GetRoles(params *SecurityControllerGetRolesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRolesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRules calls the Security controller´s function
func (c *SecurityController) GetRules(params *SecurityControllerGetRulesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRulesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetSecurityConfig calls the Security controller´s function
func (c *SecurityController) GetSecurityConfig(params *SecurityControllerGetSecurityConfigParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetUserMePolicies calls the Security controller´s function
func (c *SecurityController) GetUserMePolicies(params *SecurityControllerGetUserMePoliciesParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUserMePoliciesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetUserMe calls the Security controller´s function
func (c *SecurityController) GetUserMe(params *SecurityControllerGetUserMeParams, arg2 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUserMeWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetUsers calls the Security controller´s function
func (c *SecurityController) GetUsers(params *SecurityControllerGetUsersParams, arg2 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUsersWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// LoginUser calls the Security controller´s function
func (c *SecurityController) LoginUser(params *SecurityControllerLoginUserParams, arg2 ...RequestEditorFn) (*Token, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLoginUserWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *Token
	if i, ok := r.(*Token); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// LogoutUser calls the Security controller´s function
func (c *SecurityController) LogoutUser(arg1 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLogoutUserWithResponse(c.ClientInterface.(*Client).ctx, arg1...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutSecurityConfigWithBody calls the Security controller´s function
func (c *SecurityController) PutSecurityConfigWithBody(params *SecurityControllerPutSecurityConfigParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerPutSecurityConfigWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *map[string]interface {}
	if i, ok := r.(*map[string]interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutSecurityConfig calls the Security controller´s function
func (c *SecurityController) PutSecurityConfig(params *SecurityControllerPutSecurityConfigParams, arg2 SecurityControllerPutSecurityConfigJSONRequestBody, arg3 ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerPutSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *map[string]interface {}
	if i, ok := r.(*map[string]interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemovePolicies calls the Security controller´s function
func (c *SecurityController) RemovePolicies(params *SecurityControllerRemovePoliciesParams, arg2 ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemovePoliciesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponsePolicies
	if i, ok := r.(*AllItemsResponsePolicies); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveRolePolicy calls the Security controller´s function
func (c *SecurityController) RemoveRolePolicy(arg1 RoleId, params *SecurityControllerRemoveRolePolicyParams, arg3 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRolePolicyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveRoleRule calls the Security controller´s function
func (c *SecurityController) RemoveRoleRule(arg1 RoleId, params *SecurityControllerRemoveRoleRuleParams, arg3 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRoleRuleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveRoles calls the Security controller´s function
func (c *SecurityController) RemoveRoles(params *SecurityControllerRemoveRolesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRolesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveRules calls the Security controller´s function
func (c *SecurityController) RemoveRules(params *SecurityControllerRemoveRulesParams, arg2 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRulesWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveUserRole calls the Security controller´s function
func (c *SecurityController) RemoveUserRole(arg1 UserIdRequired, params *SecurityControllerRemoveUserRoleParams, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveUserRoleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RevokeAllTokens calls the Security controller´s function
func (c *SecurityController) RevokeAllTokens(arg1 ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRevokeAllTokensWithResponse(c.ClientInterface.(*Client).ctx, arg1...))
	if err != nil {
		return nil, err
	}
	// convert to *map[string]interface {}
	if i, ok := r.(*map[string]interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RunAsLoginWithBody calls the Security controller´s function
func (c *SecurityController) RunAsLoginWithBody(params *SecurityControllerRunAsLoginParams, arg2 string, body io.Reader, arg4 ...RequestEditorFn) (*struct {
	Token *string "json:\"token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRunAsLoginWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Token *string "json:\"token,omitempty\"" }
	if i, ok := r.(*struct {
		Token *string "json:\"token,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RunAsLogin calls the Security controller´s function
func (c *SecurityController) RunAsLogin(params *SecurityControllerRunAsLoginParams, arg2 SecurityControllerRunAsLoginJSONRequestBody, arg3 ...RequestEditorFn) (*struct {
	Token *string "json:\"token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRunAsLoginWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Token *string "json:\"token,omitempty\"" }
	if i, ok := r.(*struct {
		Token *string "json:\"token,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// SetRolePolicy calls the Security controller´s function
func (c *SecurityController) SetRolePolicy(arg1 RoleId, params *SecurityControllerSetRolePolicyParams, arg3 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetRolePolicyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// SetRoleRule calls the Security controller´s function
func (c *SecurityController) SetRoleRule(arg1 RoleId, params *SecurityControllerSetRoleRuleParams, arg3 ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetRoleRuleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ApiResponse "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		ApiResponse "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// SetUserRole calls the Security controller´s function
func (c *SecurityController) SetUserRole(arg1 UserIdRequired, params *SecurityControllerSetUserRoleParams, arg3 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetUserRoleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdatePolicyWithBody calls the Security controller´s function
func (c *SecurityController) UpdatePolicyWithBody(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdatePolicyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponsePolicies
	if i, ok := r.(*AllItemsResponsePolicies); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdatePolicy calls the Security controller´s function
func (c *SecurityController) UpdatePolicy(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 SecurityControllerUpdatePolicyJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdatePolicyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponsePolicies
	if i, ok := r.(*AllItemsResponsePolicies); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateRoleWithBody calls the Security controller´s function
func (c *SecurityController) UpdateRoleWithBody(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRoleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateRole calls the Security controller´s function
func (c *SecurityController) UpdateRole(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 SecurityControllerUpdateRoleJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRoleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseRoles
	if i, ok := r.(*AllItemsResponseRoles); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateRuleWithBody calls the Security controller´s function
func (c *SecurityController) UpdateRuleWithBody(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRuleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateRule calls the Security controller´s function
func (c *SecurityController) UpdateRule(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 SecurityControllerUpdateRuleJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRuleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateUserWithBody calls the Security controller´s function
func (c *SecurityController) UpdateUserWithBody(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 string, body io.Reader, arg5 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateUserWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, arg5...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateUser calls the Security controller´s function
func (c *SecurityController) UpdateUser(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 SecurityControllerUpdateUserJSONRequestBody, arg4 ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateUserWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, arg4...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseUsers
	if i, ok := r.(*AllItemsResponseUsers); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TaskController implementation of the TaskController interface
type TaskController struct {
	*ClientWithResponses
}

// GetTasksStatus calls the Task controller´s function
func (c *TaskController) GetTasksStatus(params *TaskControllerGetTasksStatusParams, arg2 ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TaskControllerGetTasksStatusWithResponse(c.ClientInterface.(*Client).ctx, params, arg2...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*ApiResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

/*
	DefaultController DefaultControllerInterface
	LogtestController LogtestControllerInterface
	ManagerController ManagerControllerInterface
	MitreController MitreControllerInterface
	OverviewController OverviewControllerInterface
	RootcheckController RootcheckControllerInterface
	ActiveResponseController ActiveResponseControllerInterface
	CiscatController CiscatControllerInterface
	SyscheckController SyscheckControllerInterface
	SyscollectorController SyscollectorControllerInterface
	ScaController ScaControllerInterface
	VulnerabilityController VulnerabilityControllerInterface
	ClusterController ClusterControllerInterface
	ExperimentalController ExperimentalControllerInterface
	AgentController AgentControllerInterface
	DecoderController DecoderControllerInterface
	SecurityController SecurityControllerInterface
	TaskController TaskControllerInterface
	CdbListController CdbListControllerInterface
	RuleController RuleControllerInterface
	ExperimentalController:  &ExperimentalController{clientWithResponses},
	ScaController:  &ScaController{clientWithResponses},
	VulnerabilityController:  &VulnerabilityController{clientWithResponses},
	ClusterController:  &ClusterController{clientWithResponses},
	DecoderController:  &DecoderController{clientWithResponses},
	AgentController:  &AgentController{clientWithResponses},
	RuleController:  &RuleController{clientWithResponses},
	SecurityController:  &SecurityController{clientWithResponses},
	TaskController:  &TaskController{clientWithResponses},
	CdbListController:  &CdbListController{clientWithResponses},
	CiscatController:  &CiscatController{clientWithResponses},
	DefaultController:  &DefaultController{clientWithResponses},
	LogtestController:  &LogtestController{clientWithResponses},
	ManagerController:  &ManagerController{clientWithResponses},
	MitreController:  &MitreController{clientWithResponses},
	OverviewController:  &OverviewController{clientWithResponses},
	RootcheckController:  &RootcheckController{clientWithResponses},
	ActiveResponseController:  &ActiveResponseController{clientWithResponses},
	SyscollectorController:  &SyscollectorController{clientWithResponses},
	SyscheckController:  &SyscheckController{clientWithResponses},

*/
