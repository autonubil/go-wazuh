package rest

import "io"

// AgentController implementation of the AgentController interface
type AgentController struct {
	*ClientWithResponses
}

// AddAgentWithBody calls the Agent controller´s function
func (c *AgentController) AddAgentWithBody(params *AgentControllerAddAgentParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerAddAgentWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *AgentController) AddAgent(params *AgentControllerAddAgentParams, arg2 AgentControllerAddAgentJSONRequestBody, reqEditors ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerAddAgentWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *AgentController) DeleteAgents(params *AgentControllerDeleteAgentsParams, reqEditors ...RequestEditorFn) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) DeleteGroups(params *AgentControllerDeleteGroupsParams, reqEditors ...RequestEditorFn) (*struct {
	AgentGroupDeleted "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteGroupsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AgentGroupDeleted "yaml:\",inline\"" }
	if i, ok := r.(*struct {
		AgentGroupDeleted "yaml:\",inline\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteMultipleAgentSingleGroup calls the Agent controller´s function
func (c *AgentController) DeleteMultipleAgentSingleGroup(params *AgentControllerDeleteMultipleAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteMultipleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) DeleteSingleAgentMultipleGroups(arg1 AgentId, params *AgentControllerDeleteSingleAgentMultipleGroupsParams, reqEditors ...RequestEditorFn) (*struct {
	AllItemsResponseGroupIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteSingleAgentMultipleGroupsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *AgentController) DeleteSingleAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerDeleteSingleAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerDeleteSingleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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
func (c *AgentController) GetAgentConfig(arg1 AgentId, arg2 AgentControllerGetAgentConfigParamsComponent, arg3 AgentControllerGetAgentConfigParamsConfiguration, params *AgentControllerGetAgentConfigParams, reqEditors ...RequestEditorFn) (*AgentConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentConfigWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, params, reqEditors...))
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
func (c *AgentController) GetAgentFields(params *AgentControllerGetAgentFieldsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsDistinct, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentFieldsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetAgentKey(arg1 AgentId, params *AgentControllerGetAgentKeyParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsKeys, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentKeyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *AgentController) GetAgentNoGroup(params *AgentControllerGetAgentNoGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentNoGroupWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetAgentOutdated(params *AgentControllerGetAgentOutdatedParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsSimple, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentOutdatedWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetAgentSummaryOs(params *AgentControllerGetAgentSummaryOsParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentSummaryOsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetAgentSummaryStatus(params *AgentControllerGetAgentSummaryStatusParams, reqEditors ...RequestEditorFn) (*AgentsSummaryStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentSummaryStatusWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetAgentUpgrade(params *AgentControllerGetAgentUpgradeParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentUpgradeWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetAgentsInGroup(arg1 GroupId, params *AgentControllerGetAgentsInGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentsInGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *AgentController) GetAgents(params *AgentControllerGetAgentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetComponentStats(arg1 AgentId, arg2 AgentControllerGetComponentStatsParamsComponent, params *AgentControllerGetComponentStatsParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetComponentStatsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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
func (c *AgentController) GetGroupConfig(arg1 GroupId, params *AgentControllerGetGroupConfigParams, reqEditors ...RequestEditorFn) (*struct {
	AffectedItems      *[]GroupConfiguration "json:\"affected_items,omitempty\""
	TotalAffectedItems *int32                "json:\"total_affected_items,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupConfigWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *AgentController) GetGroupFileJSON(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileJsonParams, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupFileJsonWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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
func (c *AgentController) GetGroupFileXML(arg1 GroupId, arg2 FileName, params *AgentControllerGetGroupFileXmlParams, reqEditors ...RequestEditorFn) (*AgentControllerGetGroupFileXmlResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupFileXmlWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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
func (c *AgentController) GetGroupFiles(arg1 GroupId, params *AgentControllerGetGroupFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetGroupFilesWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *AgentController) GetListGroup(params *AgentControllerGetListGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseGroups, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetListGroupWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) GetSyncAgent(arg1 AgentId, params *AgentControllerGetSyncAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentsSynced, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerGetSyncAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *AgentController) InsertAgentWithBody(params *AgentControllerInsertAgentParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerInsertAgentWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *AgentController) InsertAgent(params *AgentControllerInsertAgentParams, arg2 AgentControllerInsertAgentJSONRequestBody, reqEditors ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerInsertAgentWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *AgentController) PostGroupWithBody(params *AgentControllerPostGroupParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPostGroupWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *AgentController) PostGroup(params *AgentControllerPostGroupParams, arg2 AgentControllerPostGroupJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPostGroupWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *AgentController) PostNewAgent(params *AgentControllerPostNewAgentParams, reqEditors ...RequestEditorFn) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPostNewAgentWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) PutAgentSingleGroup(arg1 AgentId, arg2 GroupId, params *AgentControllerPutAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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
func (c *AgentController) PutGroupConfigWithBody(arg1 GroupId, params *AgentControllerPutGroupConfigParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutGroupConfigWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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
func (c *AgentController) PutMultipleAgentSingleGroup(params *AgentControllerPutMultipleAgentSingleGroupParams, reqEditors ...RequestEditorFn) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutMultipleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) PutUpgradeAgents(params *AgentControllerPutUpgradeAgentsParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutUpgradeAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *AgentController) PutUpgradeCustomAgents(params *AgentControllerPutUpgradeCustomAgentsParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerPutUpgradeCustomAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// ReconnectAgents calls the Agent controller´s function
func (c *AgentController) ReconnectAgents(params *AgentControllerReconnectAgentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerReconnectAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// RestartAgent calls the Agent controller´s function
func (c *AgentController) RestartAgent(arg1 AgentId, params *AgentControllerRestartAgentParams, reqEditors ...RequestEditorFn) (*ItemAffected, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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

// RestartAgentsByGroup calls the Agent controller´s function
func (c *AgentController) RestartAgentsByGroup(arg1 GroupId, params *AgentControllerRestartAgentsByGroupParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentsByGroupWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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

// RestartAgentsByNode calls the Agent controller´s function
func (c *AgentController) RestartAgentsByNode(arg1 NodeId, params *AgentControllerRestartAgentsByNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentsByNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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

// RestartAgents calls the Agent controller´s function
func (c *AgentController) RestartAgents(params *AgentControllerRestartAgentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentControllerRestartAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// ManagerController implementation of the ManagerController interface
type ManagerController struct {
	*ClientWithResponses
}

// GetAPIConfig calls the Manager controller´s function
func (c *ManagerController) GetAPIConfig(params *ManagerControllerGetApiConfigParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetApiConfigWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetConfValidation(params *ManagerControllerGetConfValidationParams, reqEditors ...RequestEditorFn) (*ConfigurationValidation, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetConfValidationWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetConfiguration(params *ManagerControllerGetConfigurationParams, reqEditors ...RequestEditorFn) (*WazuhManagerConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetConfigurationWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetInfo(params *ManagerControllerGetInfoParams, reqEditors ...RequestEditorFn) (*WazuhInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetLogSummary(params *ManagerControllerGetLogSummaryParams, reqEditors ...RequestEditorFn) (*WazuhLogsSummary, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetLogSummaryWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetLog(params *ManagerControllerGetLogParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetLogWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetManagerConfigOndemand(arg1 ManagerControllerGetManagerConfigOndemandParamsComponent, arg2 ManagerControllerGetManagerConfigOndemandParamsConfiguration, params *ManagerControllerGetManagerConfigOndemandParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetManagerConfigOndemandWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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
func (c *ManagerController) GetStatsAnalysisd(params *ManagerControllerGetStatsAnalysisdParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsAnalysisdWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetStatsHourly(params *ManagerControllerGetStatsHourlyParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsHourlyWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetStatsRemoted(params *ManagerControllerGetStatsRemotedParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsRemotedWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetStatsWeekly(params *ManagerControllerGetStatsWeeklyParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsWeeklyWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetStats(params *ManagerControllerGetStatsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) GetStatus(params *ManagerControllerGetStatusParams, reqEditors ...RequestEditorFn) (*WazuhDaemonsStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatusWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) PutRestart(params *ManagerControllerPutRestartParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerPutRestartWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ManagerController) UpdateConfigurationWithBody(params *ManagerControllerUpdateConfigurationParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerUpdateConfigurationWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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

// SecurityController implementation of the SecurityController interface
type SecurityController struct {
	*ClientWithResponses
}

// AddPolicyWithBody calls the Security controller´s function
func (c *SecurityController) AddPolicyWithBody(params *SecurityControllerAddPolicyParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddPolicyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *SecurityController) AddPolicy(params *SecurityControllerAddPolicyParams, arg2 SecurityControllerAddPolicyJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddPolicyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *SecurityController) AddRoleWithBody(params *SecurityControllerAddRoleParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRoleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *SecurityController) AddRole(params *SecurityControllerAddRoleParams, arg2 SecurityControllerAddRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRoleWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *SecurityController) AddRuleWithBody(params *SecurityControllerAddRuleParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRuleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *SecurityController) AddRule(params *SecurityControllerAddRuleParams, arg2 SecurityControllerAddRuleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRuleWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *SecurityController) CreateUserWithBody(params *SecurityControllerCreateUserParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerCreateUserWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *SecurityController) CreateUser(params *SecurityControllerCreateUserParams, arg2 SecurityControllerCreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerCreateUserWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *SecurityController) DeleteSecurityConfig(params *SecurityControllerDeleteSecurityConfigParams, reqEditors ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerDeleteSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) DeleteUsers(params *SecurityControllerDeleteUsersParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerDeleteUsersWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) EditRunAs(arg1 UserIdRequired, params *SecurityControllerEditRunAsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerEditRunAsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) GetPolicies(params *SecurityControllerGetPoliciesParams, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetPoliciesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetRbacActions(params *SecurityControllerGetRbacActionsParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRbacActionsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetRbacResources(params *SecurityControllerGetRbacResourcesParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRbacResourcesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetRoles(params *SecurityControllerGetRolesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRolesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetRules(params *SecurityControllerGetRulesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRulesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetSecurityConfig(params *SecurityControllerGetSecurityConfigParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetUserMePolicies(params *SecurityControllerGetUserMePoliciesParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUserMePoliciesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetUserMe(params *SecurityControllerGetUserMeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUserMeWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) GetUsers(params *SecurityControllerGetUsersParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUsersWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) LoginUser(params *SecurityControllerLoginUserParams, reqEditors ...RequestEditorFn) (*Token, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLoginUserWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) LogoutUser(reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLogoutUserWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
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
func (c *SecurityController) PutSecurityConfigWithBody(params *SecurityControllerPutSecurityConfigParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerPutSecurityConfigWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *SecurityController) PutSecurityConfig(params *SecurityControllerPutSecurityConfigParams, arg2 SecurityControllerPutSecurityConfigJSONRequestBody, reqEditors ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerPutSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *SecurityController) RemovePolicies(params *SecurityControllerRemovePoliciesParams, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemovePoliciesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) RemoveRolePolicy(arg1 RoleId, params *SecurityControllerRemoveRolePolicyParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRolePolicyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) RemoveRoleRule(arg1 RoleId, params *SecurityControllerRemoveRoleRuleParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRoleRuleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) RemoveRoles(params *SecurityControllerRemoveRolesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRolesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) RemoveRules(params *SecurityControllerRemoveRulesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRulesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SecurityController) RemoveUserRole(arg1 UserIdRequired, params *SecurityControllerRemoveUserRoleParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveUserRoleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) RevokeAllTokens(reqEditors ...RequestEditorFn) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRevokeAllTokensWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
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
func (c *SecurityController) RunAsLoginWithBody(params *SecurityControllerRunAsLoginParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Token *string "json:\"token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRunAsLoginWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *SecurityController) RunAsLogin(params *SecurityControllerRunAsLoginParams, arg2 SecurityControllerRunAsLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Token *string "json:\"token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRunAsLoginWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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
func (c *SecurityController) SetRolePolicy(arg1 RoleId, params *SecurityControllerSetRolePolicyParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetRolePolicyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) SetRoleRule(arg1 RoleId, params *SecurityControllerSetRoleRuleParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetRoleRuleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) SetUserRole(arg1 UserIdRequired, params *SecurityControllerSetUserRoleParams, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetUserRoleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SecurityController) UpdatePolicyWithBody(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdatePolicyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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
func (c *SecurityController) UpdatePolicy(arg1 PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, arg3 SecurityControllerUpdatePolicyJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdatePolicyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, reqEditors...))
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
func (c *SecurityController) UpdateRoleWithBody(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRoleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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
func (c *SecurityController) UpdateRole(arg1 RoleId, params *SecurityControllerUpdateRoleParams, arg3 SecurityControllerUpdateRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRoleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, reqEditors...))
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
func (c *SecurityController) UpdateRuleWithBody(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRuleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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
func (c *SecurityController) UpdateRule(arg1 SecurityRuleId, params *SecurityControllerUpdateRuleParams, arg3 SecurityControllerUpdateRuleJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRuleWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, reqEditors...))
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
func (c *SecurityController) UpdateUserWithBody(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateUserWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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
func (c *SecurityController) UpdateUser(arg1 UserIdRequired, params *SecurityControllerUpdateUserParams, arg3 SecurityControllerUpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateUserWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, reqEditors...))
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
func (c *TaskController) GetTasksStatus(params *TaskControllerGetTasksStatusParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TaskControllerGetTasksStatusWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// VulnerabilityController implementation of the VulnerabilityController interface
type VulnerabilityController struct {
	*ClientWithResponses
}

// GetLastScanAgent calls the Vulnerability controller´s function
func (c *VulnerabilityController) GetLastScanAgent(arg1 AgentId, params *VulnerabilityControllerGetLastScanAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseLastScan, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilityControllerGetLastScanAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseLastScan
	if i, ok := r.(*AllItemsResponseLastScan); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetVulnerabilitiesFieldSummary calls the Vulnerability controller´s function
func (c *VulnerabilityController) GetVulnerabilitiesFieldSummary(arg1 AgentId, arg2 VulnerabilityControllerGetVulnerabilitiesFieldSummaryParamsField, params *VulnerabilityControllerGetVulnerabilitiesFieldSummaryParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilityControllerGetVulnerabilitiesFieldSummaryWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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

// GetVulnerabilityAgent calls the Vulnerability controller´s function
func (c *VulnerabilityController) GetVulnerabilityAgent(arg1 AgentId, params *VulnerabilityControllerGetVulnerabilityAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseVulnerability, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilityControllerGetVulnerabilityAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *ApiResponse
	if i, ok := r.(*AllItemsResponseVulnerability); ok {
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
func (c *CdbListController) DeleteFile(arg1 ListFilenamePath, params *CdbListControllerDeleteFileParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerDeleteFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *CdbListController) GetFile(arg1 ListFilenamePath, params *CdbListControllerGetFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponseListsFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerGetFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *CdbListController) GetListsFiles(params *CdbListControllerGetListsFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseListsFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerGetListsFilesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *CdbListController) GetLists(params *CdbListControllerGetListsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseLists, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerGetListsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *CdbListController) PutFileWithBody(arg1 ListFilenamePath, params *CdbListControllerPutFileParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CdbListControllerPutFileWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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

// ClusterController implementation of the ClusterController interface
type ClusterController struct {
	*ClientWithResponses
}

// GetAPIConfig calls the Cluster controller´s function
func (c *ClusterController) GetAPIConfig(params *ClusterControllerGetApiConfigParams, reqEditors ...RequestEditorFn) (*struct {
	AdditionalProperties map[string]interface{} "json:\"-\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetApiConfigWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) GetClusterNode(params *ClusterControllerGetClusterNodeParams, reqEditors ...RequestEditorFn) (*struct {
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
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetClusterNodeWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) GetClusterNodes(params *ClusterControllerGetClusterNodesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseClusterNodes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetClusterNodesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) GetConfValidation(params *ClusterControllerGetConfValidationParams, reqEditors ...RequestEditorFn) (*AllItemsResponseValidationStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfValidationWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) GetConfig(params *ClusterControllerGetConfigParams, reqEditors ...RequestEditorFn) (*struct {
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
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfigWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) GetConfigurationNode(arg1 NodeId, params *ClusterControllerGetConfigurationNodeParams, reqEditors ...RequestEditorFn) (*WazuhManagerConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfigurationNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetHealthcheck(params *ClusterControllerGetHealthcheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponseNodeHealthcheck, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetHealthcheckWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) GetInfoNode(arg1 NodeId, params *ClusterControllerGetInfoNodeParams, reqEditors ...RequestEditorFn) (*WazuhInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetInfoNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetLogNode(arg1 NodeId, params *ClusterControllerGetLogNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhLogs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetLogNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetLogSummaryNode(arg1 NodeId, params *ClusterControllerGetLogSummaryNodeParams, reqEditors ...RequestEditorFn) (*WazuhLogsSummary, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetLogSummaryNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetNodeConfig(arg1 NodeId, arg2 ClusterControllerGetNodeConfigParamsComponent, arg3 ClusterControllerGetNodeConfigParamsConfiguration, params *ClusterControllerGetNodeConfigParams, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetNodeConfigWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, params, reqEditors...))
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
func (c *ClusterController) GetStatsAnalysisdNode(arg1 NodeId, params *ClusterControllerGetStatsAnalysisdNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsAnalysisdNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetStatsHourlyNode(arg1 NodeId, params *ClusterControllerGetStatsHourlyNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsHourlyNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetStatsNode(arg1 NodeId, params *ClusterControllerGetStatsNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetStatsRemotedNode(arg1 NodeId, params *ClusterControllerGetStatsRemotedNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsRemotedNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetStatsWeeklyNode(arg1 NodeId, params *ClusterControllerGetStatsWeeklyNodeParams, reqEditors ...RequestEditorFn) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsWeeklyNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetStatusNode(arg1 NodeId, params *ClusterControllerGetStatusNodeParams, reqEditors ...RequestEditorFn) (*WazuhDaemonsStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatusNodeWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ClusterController) GetStatus(params *ClusterControllerGetStatusParams, reqEditors ...RequestEditorFn) (*struct {
	Enabled *N200DataEnabled "json:\"enabled,omitempty\""
	Running *N200DataRunning "json:\"running,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatusWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) PutRestart(params *ClusterControllerPutRestartParams, reqEditors ...RequestEditorFn) (*AllItemsResponseNodeIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerPutRestartWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ClusterController) UpdateConfigurationWithBody(arg1 NodeId, params *ClusterControllerUpdateConfigurationParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerUpdateConfigurationWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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

// DefaultController implementation of the DefaultController interface
type DefaultController struct {
	*ClientWithResponses
}

// DefaultInfo calls the Default controller´s function
func (c *DefaultController) DefaultInfo(params *DefaultControllerDefaultInfoParams, reqEditors ...RequestEditorFn) (*BasicInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DefaultControllerDefaultInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *LogtestController) EndLogtestSession(arg1 LogtestToken, params *LogtestControllerEndLogtestSessionParams, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.LogtestControllerEndLogtestSessionWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *LogtestController) RunLogtestToolWithBody(params *LogtestControllerRunLogtestToolParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.LogtestControllerRunLogtestToolWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *LogtestController) RunLogtestTool(params *LogtestControllerRunLogtestToolParams, arg2 LogtestControllerRunLogtestToolJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.LogtestControllerRunLogtestToolWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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

// MitreController implementation of the MitreController interface
type MitreController struct {
	*ClientWithResponses
}

// GetGroups calls the Mitre controller´s function
func (c *MitreController) GetGroups(params *MitreControllerGetGroupsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetGroupsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// GetMetadata calls the Mitre controller´s function
func (c *MitreController) GetMetadata(params *MitreControllerGetMetadataParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetMetadataWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// GetMitigations calls the Mitre controller´s function
func (c *MitreController) GetMitigations(params *MitreControllerGetMitigationsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetMitigationsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// GetReferences calls the Mitre controller´s function
func (c *MitreController) GetReferences(params *MitreControllerGetReferencesParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetReferencesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// GetSoftware calls the Mitre controller´s function
func (c *MitreController) GetSoftware(params *MitreControllerGetSoftwareParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetSoftwareWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// GetTactics calls the Mitre controller´s function
func (c *MitreController) GetTactics(params *MitreControllerGetTacticsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetTacticsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// GetTechniques calls the Mitre controller´s function
func (c *MitreController) GetTechniques(params *MitreControllerGetTechniquesParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetTechniquesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// RuleController implementation of the RuleController interface
type RuleController struct {
	*ClientWithResponses
}

// DeleteFile calls the Rule controller´s function
func (c *RuleController) DeleteFile(arg1 XmlFilenamePath, params *RuleControllerDeleteFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerDeleteFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *RuleController) GetFile(arg1 XmlFilenamePath, params *RuleControllerGetFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *RuleController) GetRulesFiles(params *RuleControllerGetRulesFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRulesFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesFilesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *RuleController) GetRulesGroups(params *RuleControllerGetRulesGroupsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesGroupsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *RuleController) GetRulesRequirement(arg1 RuleControllerGetRulesRequirementParamsRequirement, params *RuleControllerGetRulesRequirementParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesRequirementWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *RuleController) GetRules(params *RuleControllerGetRulesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseRules, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerGetRulesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *RuleController) PutFileWithBody(arg1 XmlFilenamePath, params *RuleControllerPutFileParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RuleControllerPutFileWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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

// CiscatController implementation of the CiscatController interface
type CiscatController struct {
	*ClientWithResponses
}

// GetAgentsCiscatResults calls the Ciscat controller´s function
func (c *CiscatController) GetAgentsCiscatResults(arg1 AgentId, params *CiscatControllerGetAgentsCiscatResultsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseCiscatResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CiscatControllerGetAgentsCiscatResultsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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

// RootcheckController implementation of the RootcheckController interface
type RootcheckController struct {
	*ClientWithResponses
}

// DeleteRootcheck calls the Rootcheck controller´s function
func (c *RootcheckController) DeleteRootcheck(arg1 AgentId, params *RootcheckControllerDeleteRootcheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerDeleteRootcheckWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *RootcheckController) GetLastScanAgent(arg1 AgentId, params *RootcheckControllerGetLastScanAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerGetLastScanAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *RootcheckController) GetRootcheckAgent(arg1 AgentId, params *RootcheckControllerGetRootcheckAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerGetRootcheckAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *RootcheckController) PutRootcheck(params *RootcheckControllerPutRootcheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RootcheckControllerPutRootcheckWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// ScaController implementation of the ScaController interface
type ScaController struct {
	*ClientWithResponses
}

// GetScaAgent calls the Sca controller´s function
func (c *ScaController) GetScaAgent(arg1 AgentId, params *ScaControllerGetScaAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSCADatabase, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScaControllerGetScaAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *ScaController) GetScaChecks(arg1 AgentId, arg2 PolicyId, params *ScaControllerGetScaChecksParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSCAChecks, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScaControllerGetScaChecksWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
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

// SyscheckController implementation of the SyscheckController interface
type SyscheckController struct {
	*ClientWithResponses
}

// DeleteSyscheckAgent calls the Syscheck controller´s function
func (c *SyscheckController) DeleteSyscheckAgent(arg1 AgentId, params *SyscheckControllerDeleteSyscheckAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerDeleteSyscheckAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscheckController) GetLastScanAgent(arg1 AgentId, params *SyscheckControllerGetLastScanAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseLastScan, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerGetLastScanAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AllItemsResponseLastScan
	if i, ok := r.(*AllItemsResponseLastScan); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetSyscheckAgent calls the Syscheck controller´s function
func (c *SyscheckController) GetSyscheckAgent(arg1 AgentId, params *SyscheckControllerGetSyscheckAgentParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscheckResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerGetSyscheckAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscheckController) PutSyscheck(params *SyscheckControllerPutSyscheckParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerPutSyscheckWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *SyscollectorController) GetHardwareInfo(arg1 AgentId, params *SyscollectorControllerGetHardwareInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetHardwareInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetHotfixInfo(arg1 AgentId, params *SyscollectorControllerGetHotfixInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetHotfixInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetNetworkAddressInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkAddressInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkAddressInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetNetworkInterfaceInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkInterfaceInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkInterfaceInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetNetworkProtocolInfo(arg1 AgentId, params *SyscollectorControllerGetNetworkProtocolInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkProtocolInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetOsInfo(arg1 AgentId, params *SyscollectorControllerGetOsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetOsInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetPackagesInfo(arg1 AgentId, params *SyscollectorControllerGetPackagesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetPackagesInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetPortsInfo(arg1 AgentId, params *SyscollectorControllerGetPortsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetPortsInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *SyscollectorController) GetProcessesInfo(arg1 AgentId, params *SyscollectorControllerGetProcessesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetProcessesInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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

// ActiveResponseController implementation of the ActiveResponseController interface
type ActiveResponseController struct {
	*ClientWithResponses
}

// RunCommandWithBody calls the ActiveResponse controller´s function
func (c *ActiveResponseController) RunCommandWithBody(params *ActiveResponseControllerRunCommandParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ActiveResponseControllerRunCommandWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
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
func (c *ActiveResponseController) RunCommand(params *ActiveResponseControllerRunCommandParams, arg2 ActiveResponseControllerRunCommandJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ActiveResponseControllerRunCommandWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
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

// DecoderController implementation of the DecoderController interface
type DecoderController struct {
	*ClientWithResponses
}

// DeleteFile calls the Decoder controller´s function
func (c *DecoderController) DeleteFile(arg1 XmlFilenamePath, params *DecoderControllerDeleteFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerDeleteFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *DecoderController) GetDecodersFiles(params *DecoderControllerGetDecodersFilesParams, reqEditors ...RequestEditorFn) (*AllItemsResponseDecodersFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetDecodersFilesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *DecoderController) GetDecodersParents(params *DecoderControllerGetDecodersParentsParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetDecodersParentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *DecoderController) GetDecoders(params *DecoderControllerGetDecodersParams, reqEditors ...RequestEditorFn) (*AllItemsResponseDecoders, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetDecodersWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *DecoderController) GetFile(arg1 XmlFilenamePath, params *DecoderControllerGetFileParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerGetFileWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
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
func (c *DecoderController) PutFileWithBody(arg1 XmlFilenamePath, params *DecoderControllerPutFileParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecoderControllerPutFileWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
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

// ExperimentalController implementation of the ExperimentalController interface
type ExperimentalController struct {
	*ClientWithResponses
}

// ClearRootcheckDatabase calls the Experimental controller´s function
func (c *ExperimentalController) ClearRootcheckDatabase(params *ExperimentalControllerClearRootcheckDatabaseParams, reqEditors ...RequestEditorFn) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerClearRootcheckDatabaseWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// ClearSyscheckDatabase calls the Experimental controller´s function
func (c *ExperimentalController) ClearSyscheckDatabase(params *ExperimentalControllerClearSyscheckDatabaseParams, reqEditors ...RequestEditorFn) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerClearSyscheckDatabaseWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetCisCatResults(params *ExperimentalControllerGetCisCatResultsParams, reqEditors ...RequestEditorFn) (*AllItemsResponseCiscatResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetCisCatResultsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetHardwareInfo(params *ExperimentalControllerGetHardwareInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHardware, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetHardwareInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetHotfixesInfo(params *ExperimentalControllerGetHotfixesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorHotfixes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetHotfixesInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetNetworkAddressInfo(params *ExperimentalControllerGetNetworkAddressInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorNetwork, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkAddressInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetNetworkInterfaceInfo(params *ExperimentalControllerGetNetworkInterfaceInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorInterface, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkInterfaceInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetNetworkProtocolInfo(params *ExperimentalControllerGetNetworkProtocolInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProtocol, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkProtocolInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetOsInfo(params *ExperimentalControllerGetOsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorOS, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetOsInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetPackagesInfo(params *ExperimentalControllerGetPackagesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPackages, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetPackagesInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetPortsInfo(params *ExperimentalControllerGetPortsInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorPorts, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetPortsInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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
func (c *ExperimentalController) GetProcessesInfo(params *ExperimentalControllerGetProcessesInfoParams, reqEditors ...RequestEditorFn) (*AllItemsResponseSyscollectorProcesses, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetProcessesInfoWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

// OverviewController implementation of the OverviewController interface
type OverviewController struct {
	*ClientWithResponses
}

// GetOverviewAgents calls the Overview controller´s function
func (c *OverviewController) GetOverviewAgents(params *OverviewControllerGetOverviewAgentsParams, reqEditors ...RequestEditorFn) (*OverviewAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.OverviewControllerGetOverviewAgentsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
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

/*
	SecurityController SecurityControllerInterface
	TaskController TaskControllerInterface
	VulnerabilityController VulnerabilityControllerInterface
	AgentController AgentControllerInterface
	ManagerController ManagerControllerInterface
	DefaultController DefaultControllerInterface
	LogtestController LogtestControllerInterface
	MitreController MitreControllerInterface
	RuleController RuleControllerInterface
	CdbListController CdbListControllerInterface
	ClusterController ClusterControllerInterface
	ScaController ScaControllerInterface
	SyscheckController SyscheckControllerInterface
	SyscollectorController SyscollectorControllerInterface
	CiscatController CiscatControllerInterface
	RootcheckController RootcheckControllerInterface
	ExperimentalController ExperimentalControllerInterface
	OverviewController OverviewControllerInterface
	ActiveResponseController ActiveResponseControllerInterface
	DecoderController DecoderControllerInterface

	OverviewController:  &OverviewController{clientWithResponses},
	ActiveResponseController:  &ActiveResponseController{clientWithResponses},
	DecoderController:  &DecoderController{clientWithResponses},
	ExperimentalController:  &ExperimentalController{clientWithResponses},
	TaskController:  &TaskController{clientWithResponses},
	VulnerabilityController:  &VulnerabilityController{clientWithResponses},
	AgentController:  &AgentController{clientWithResponses},
	ManagerController:  &ManagerController{clientWithResponses},
	SecurityController:  &SecurityController{clientWithResponses},
	LogtestController:  &LogtestController{clientWithResponses},
	MitreController:  &MitreController{clientWithResponses},
	RuleController:  &RuleController{clientWithResponses},
	CdbListController:  &CdbListController{clientWithResponses},
	ClusterController:  &ClusterController{clientWithResponses},
	DefaultController:  &DefaultController{clientWithResponses},
	SyscheckController:  &SyscheckController{clientWithResponses},
	SyscollectorController:  &SyscollectorController{clientWithResponses},
	CiscatController:  &CiscatController{clientWithResponses},
	RootcheckController:  &RootcheckController{clientWithResponses},
	ScaController:  &ScaController{clientWithResponses},

*/
