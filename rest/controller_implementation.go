package rest

import "io"

// ActiveResponseController implementation of the ActiveResponseController interface
type ActiveResponseController struct {
	*ClientWithResponses
}

// RunCommandWithBody calls the ActiveResponse controller´s function
func (c *ActiveResponseController) RunCommandWithBody(params *ActiveResponseControllerRunCommandParams, contentType string, body io.Reader) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ActiveResponseControllerRunCommandWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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
func (c *ActiveResponseController) RunCommand(params *ActiveResponseControllerRunCommandParams, activeResponseControllerRunCommandJSONRequestBody ActiveResponseControllerRunCommandJSONRequestBody) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ActiveResponseControllerRunCommandWithResponse(c.ClientInterface.(*Client).ctx, params, activeResponseControllerRunCommandJSONRequestBody))
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
func (c *CiscatController) GetAgentsCiscatResults(agentID AgentId, params *CiscatControllerGetAgentsCiscatResultsParams) (*AllItemsResponseCiscatResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CiscatControllerGetAgentsCiscatResultsWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// ClusterController implementation of the ClusterController interface
type ClusterController struct {
	*ClientWithResponses
}

// DeleteFilesNode calls the Cluster controller´s function
func (c *ClusterController) DeleteFilesNode(nodeID NodeId, params *ClusterControllerDeleteFilesNodeParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerDeleteFilesNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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

// GetAPIConfig calls the Cluster controller´s function
func (c *ClusterController) GetAPIConfig(params *ClusterControllerGetApiConfigParams) (*struct {
	AdditionalProperties map[string]interface{} "json:\"-\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetApiConfigWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ClusterController) GetClusterNode(params *ClusterControllerGetClusterNodeParams) (*struct {
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
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetClusterNodeWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ClusterController) GetClusterNodes(params *ClusterControllerGetClusterNodesParams) (*AllItemsResponseClusterNodes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetClusterNodesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ClusterController) GetConfValidation(params *ClusterControllerGetConfValidationParams) (*AllItemsResponseValidationStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfValidationWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ClusterController) GetConfig(params *ClusterControllerGetConfigParams) (*struct {
	BindAddr *string   "json:\"bind_addr,omitempty\""
	Disabled *bool     "json:\"disabled,omitempty\""
	Hidden   *string   "json:\"hidden,omitempty\""
	Key      *string   "json:\"key,omitempty\""
	Name     *string   "json:\"name,omitempty\""
	NodeName *string   "json:\"node_name,omitempty\""
	NodeType *string   "json:\"node_type,omitempty\""
	Nodes    *[]string "json:\"nodes,omitempty\""
	Port     *int      "json:\"port,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfigWithResponse(c.ClientInterface.(*Client).ctx, params))
	if err != nil {
		return nil, err
	}
	// convert to *struct { BindAddr *string "json:\"bind_addr,omitempty\""; Disabled *bool "json:\"disabled,omitempty\""; Hidden *string "json:\"hidden,omitempty\""; Key *string "json:\"key,omitempty\""; Name *string "json:\"name,omitempty\""; NodeName *string "json:\"node_name,omitempty\""; NodeType *string "json:\"node_type,omitempty\""; Nodes *[]string "json:\"nodes,omitempty\""; Port *int "json:\"port,omitempty\"" }
	if i, ok := r.(*struct {
		BindAddr *string   "json:\"bind_addr,omitempty\""
		Disabled *bool     "json:\"disabled,omitempty\""
		Hidden   *string   "json:\"hidden,omitempty\""
		Key      *string   "json:\"key,omitempty\""
		Name     *string   "json:\"name,omitempty\""
		NodeName *string   "json:\"node_name,omitempty\""
		NodeType *string   "json:\"node_type,omitempty\""
		Nodes    *[]string "json:\"nodes,omitempty\""
		Port     *int      "json:\"port,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetConfigurationNode calls the Cluster controller´s function
func (c *ClusterController) GetConfigurationNode(nodeID NodeId, params *ClusterControllerGetConfigurationNodeParams) (*WazuhMangerConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetConfigurationNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhMangerConfiguration
	if i, ok := r.(*WazuhMangerConfiguration); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetFilesNode calls the Cluster controller´s function
func (c *ClusterController) GetFilesNode(nodeID NodeId, params *ClusterControllerGetFilesNodeParams) (*struct {
	Contents *string "json:\"contents,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetFilesNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Contents *string "json:\"contents,omitempty\"" }
	if i, ok := r.(*struct {
		Contents *string "json:\"contents,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetHealthcheck calls the Cluster controller´s function
func (c *ClusterController) GetHealthcheck(params *ClusterControllerGetHealthcheckParams) (*AllItemsResponseNodeHealthcheck, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetHealthcheckWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ClusterController) GetInfoNode(nodeID NodeId, params *ClusterControllerGetInfoNodeParams) (*WazuhInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetInfoNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetLogNode(nodeID NodeId, params *ClusterControllerGetLogNodeParams) (*AllItemsResponseWazuhLogs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetLogNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetLogSummaryNode(nodeID NodeId, params *ClusterControllerGetLogSummaryNodeParams) (*WazuhLogsSummary, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetLogSummaryNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetNodeConfig(nodeID NodeId, component Component, configuration Configuration, params *ClusterControllerGetNodeConfigParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetNodeConfigWithResponse(c.ClientInterface.(*Client).ctx, nodeID, component, configuration, params))
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
func (c *ClusterController) GetStatsAnalysisdNode(nodeID NodeId, params *ClusterControllerGetStatsAnalysisdNodeParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsAnalysisdNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetStatsHourlyNode(nodeID NodeId, params *ClusterControllerGetStatsHourlyNodeParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsHourlyNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetStatsNode(nodeID NodeId, params *ClusterControllerGetStatsNodeParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetStatsRemotedNode(nodeID NodeId, params *ClusterControllerGetStatsRemotedNodeParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsRemotedNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetStatsWeeklyNode(nodeID NodeId, params *ClusterControllerGetStatsWeeklyNodeParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatsWeeklyNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetStatusNode(nodeID NodeId, params *ClusterControllerGetStatusNodeParams) (*WazuhDaemonsStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatusNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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
func (c *ClusterController) GetStatus(params *ClusterControllerGetStatusParams) (*struct {
	Enabled *string "json:\"enabled,omitempty\""
	Running *string "json:\"running,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerGetStatusWithResponse(c.ClientInterface.(*Client).ctx, params))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Enabled *string "json:\"enabled,omitempty\""; Running *string "json:\"running,omitempty\"" }
	if i, ok := r.(*struct {
		Enabled *string "json:\"enabled,omitempty\""
		Running *string "json:\"running,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PutFilesNodeWithBody calls the Cluster controller´s function
func (c *ClusterController) PutFilesNodeWithBody(nodeID NodeId, params *ClusterControllerPutFilesNodeParams, contentType string, body io.Reader) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerPutFilesNodeWithBodyWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params, contentType, body))
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

// PutRestart calls the Cluster controller´s function
func (c *ClusterController) PutRestart(params *ClusterControllerPutRestartParams) (*AllItemsResponseNodeIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ClusterControllerPutRestartWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// DecodersController implementation of the DecodersController interface
type DecodersController struct {
	*ClientWithResponses
}

// GetDecodersFiles calls the Decoders controller´s function
func (c *DecodersController) GetDecodersFiles(params *DecodersControllerGetDecodersFilesParams) (*AllItemsResponseDecodersFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecodersControllerGetDecodersFilesWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetDecodersParents calls the Decoders controller´s function
func (c *DecodersController) GetDecodersParents(params *DecodersControllerGetDecodersParentsParams) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecodersControllerGetDecodersParentsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetDecoders calls the Decoders controller´s function
func (c *DecodersController) GetDecoders(params *DecodersControllerGetDecodersParams) (*AllItemsResponseDecoders, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecodersControllerGetDecodersWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetDownloadFile calls the Decoders controller´s function
func (c *DecodersController) GetDownloadFile(downloadFile DownloadFile, params *DecodersControllerGetDownloadFileParams) (*DecodersControllerGetDownloadFileResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DecodersControllerGetDownloadFileWithResponse(c.ClientInterface.(*Client).ctx, downloadFile, params))
	if err != nil {
		return nil, err
	}
	// convert to *DecodersControllerGetDownloadFileResponse
	if i, ok := r.(*DecodersControllerGetDownloadFileResponse); ok {
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
func (c *DefaultController) DefaultInfo(params *DefaultControllerDefaultInfoParams) (*BasicInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.DefaultControllerDefaultInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// RulesController implementation of the RulesController interface
type RulesController struct {
	*ClientWithResponses
}

// GetDownloadFile calls the Rules controller´s function
func (c *RulesController) GetDownloadFile(downloadFile DownloadFile, params *RulesControllerGetDownloadFileParams) (*RulesControllerGetDownloadFileResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RulesControllerGetDownloadFileWithResponse(c.ClientInterface.(*Client).ctx, downloadFile, params))
	if err != nil {
		return nil, err
	}
	// convert to *RulesControllerGetDownloadFileResponse
	if i, ok := r.(*RulesControllerGetDownloadFileResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetRulesFiles calls the Rules controller´s function
func (c *RulesController) GetRulesFiles(params *RulesControllerGetRulesFilesParams) (*AllItemsResponseRulesFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RulesControllerGetRulesFilesWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetRulesGroups calls the Rules controller´s function
func (c *RulesController) GetRulesGroups(params *RulesControllerGetRulesGroupsParams) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RulesControllerGetRulesGroupsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetRulesRequirement calls the Rules controller´s function
func (c *RulesController) GetRulesRequirement(ruleRequirement RuleRequirement, params *RulesControllerGetRulesRequirementParams) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RulesControllerGetRulesRequirementWithResponse(c.ClientInterface.(*Client).ctx, ruleRequirement, params))
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

// GetRules calls the Rules controller´s function
func (c *RulesController) GetRules(params *RulesControllerGetRulesParams) (*AllItemsResponseRules, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.RulesControllerGetRulesWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// SecurityController implementation of the SecurityController interface
type SecurityController struct {
	*ClientWithResponses
}

// AddPolicyWithBody calls the Security controller´s function
func (c *SecurityController) AddPolicyWithBody(params *SecurityControllerAddPolicyParams, contentType string, body io.Reader) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddPolicyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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
func (c *SecurityController) AddPolicy(params *SecurityControllerAddPolicyParams, securityControllerAddPolicyJSONRequestBody SecurityControllerAddPolicyJSONRequestBody) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddPolicyWithResponse(c.ClientInterface.(*Client).ctx, params, securityControllerAddPolicyJSONRequestBody))
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
func (c *SecurityController) AddRoleWithBody(params *SecurityControllerAddRoleParams, contentType string, body io.Reader) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRoleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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
func (c *SecurityController) AddRole(params *SecurityControllerAddRoleParams, securityControllerAddRoleJSONRequestBody SecurityControllerAddRoleJSONRequestBody) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRoleWithResponse(c.ClientInterface.(*Client).ctx, params, securityControllerAddRoleJSONRequestBody))
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
func (c *SecurityController) AddRuleWithBody(params *SecurityControllerAddRuleParams, contentType string, body io.Reader) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRuleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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
func (c *SecurityController) AddRule(params *SecurityControllerAddRuleParams, securityControllerAddRuleJSONRequestBody SecurityControllerAddRuleJSONRequestBody) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerAddRuleWithResponse(c.ClientInterface.(*Client).ctx, params, securityControllerAddRuleJSONRequestBody))
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
func (c *SecurityController) CreateUserWithBody(params *SecurityControllerCreateUserParams, contentType string, body io.Reader) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerCreateUserWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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
func (c *SecurityController) CreateUser(params *SecurityControllerCreateUserParams, securityControllerCreateUserJSONRequestBody SecurityControllerCreateUserJSONRequestBody) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerCreateUserWithResponse(c.ClientInterface.(*Client).ctx, params, securityControllerCreateUserJSONRequestBody))
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
func (c *SecurityController) DeleteSecurityConfig(params *SecurityControllerDeleteSecurityConfigParams) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerDeleteSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) DeleteUsers(params *SecurityControllerDeleteUsersParams) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerDeleteUsersWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetPolicies(params *SecurityControllerGetPoliciesParams) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetPoliciesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetRbacActions(params *SecurityControllerGetRbacActionsParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRbacActionsWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetRbacResources(params *SecurityControllerGetRbacResourcesParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRbacResourcesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetRoles(params *SecurityControllerGetRolesParams) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRolesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetRules(params *SecurityControllerGetRulesParams) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetRulesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetSecurityConfig(params *SecurityControllerGetSecurityConfigParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetUserMePolicies(params *SecurityControllerGetUserMePoliciesParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUserMePoliciesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetUserMe(params *SecurityControllerGetUserMeParams) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUserMeWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) GetUsers(params *SecurityControllerGetUsersParams) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerGetUsersWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// LoginUserRunAsWithBody calls the Security controller´s function
func (c *SecurityController) LoginUserRunAsWithBody(params *SecurityControllerLoginUserParams, contentType string, body io.Reader) (*struct {
	Token *string "json:\"token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLoginUserRunAsWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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

// LoginUserRunAs calls the Security controller´s function
func (c *SecurityController) LoginUserRunAs(params *SecurityControllerLoginUserParams, securityControllerLoginUserJSONRequestBody SecurityControllerLoginUserJSONRequestBody) (*struct {
	Token *string "json:\"token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLoginUserRunAsWithResponse(c.ClientInterface.(*Client).ctx, params, securityControllerLoginUserJSONRequestBody))
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

// LoginUser calls the Security controller´s function
func (c *SecurityController) LoginUser(params *SecurityControllerLoginUserParams) (*Token, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLoginUserWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) LogoutUser() (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerLogoutUserWithResponse(c.ClientInterface.(*Client).ctx))
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
func (c *SecurityController) PutSecurityConfigWithBody(params *SecurityControllerPutSecurityConfigParams, contentType string, body io.Reader) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerPutSecurityConfigWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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
func (c *SecurityController) PutSecurityConfig(params *SecurityControllerPutSecurityConfigParams, securityControllerPutSecurityConfigJSONRequestBody SecurityControllerPutSecurityConfigJSONRequestBody) (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerPutSecurityConfigWithResponse(c.ClientInterface.(*Client).ctx, params, securityControllerPutSecurityConfigJSONRequestBody))
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
func (c *SecurityController) RemovePolicies(params *SecurityControllerRemovePoliciesParams) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemovePoliciesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) RemoveRolePolicy(roleID RoleId, params *SecurityControllerRemoveRolePolicyParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRolePolicyWithResponse(c.ClientInterface.(*Client).ctx, roleID, params))
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
func (c *SecurityController) RemoveRoleRule(roleID RoleId, params *SecurityControllerRemoveRoleRuleParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRoleRuleWithResponse(c.ClientInterface.(*Client).ctx, roleID, params))
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
func (c *SecurityController) RemoveRoles(params *SecurityControllerRemoveRolesParams) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRolesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) RemoveRules(params *SecurityControllerRemoveRulesParams) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveRulesWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *SecurityController) RemoveUserRole(userIDRequired UserIdRequired, params *SecurityControllerRemoveUserRoleParams) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRemoveUserRoleWithResponse(c.ClientInterface.(*Client).ctx, userIDRequired, params))
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
func (c *SecurityController) RevokeAllTokens() (*map[string]interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerRevokeAllTokensWithResponse(c.ClientInterface.(*Client).ctx))
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

// SetRolePolicy calls the Security controller´s function
func (c *SecurityController) SetRolePolicy(roleID RoleId, params *SecurityControllerSetRolePolicyParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetRolePolicyWithResponse(c.ClientInterface.(*Client).ctx, roleID, params))
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
func (c *SecurityController) SetRoleRule(roleID RoleId, params *SecurityControllerSetRoleRuleParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetRoleRuleWithResponse(c.ClientInterface.(*Client).ctx, roleID, params))
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
func (c *SecurityController) SetUserRole(userIDRequired UserIdRequired, params *SecurityControllerSetUserRoleParams) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerSetUserRoleWithResponse(c.ClientInterface.(*Client).ctx, userIDRequired, params))
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
func (c *SecurityController) UpdatePolicyWithBody(policyIDRbac PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, contentType string, body io.Reader) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdatePolicyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, policyIDRbac, params, contentType, body))
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
func (c *SecurityController) UpdatePolicy(policyIDRbac PolicyIdRbac, params *SecurityControllerUpdatePolicyParams, securityControllerUpdatePolicyJSONRequestBody SecurityControllerUpdatePolicyJSONRequestBody) (*AllItemsResponsePolicies, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdatePolicyWithResponse(c.ClientInterface.(*Client).ctx, policyIDRbac, params, securityControllerUpdatePolicyJSONRequestBody))
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
func (c *SecurityController) UpdateRoleWithBody(roleID RoleId, params *SecurityControllerUpdateRoleParams, contentType string, body io.Reader) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRoleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, roleID, params, contentType, body))
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
func (c *SecurityController) UpdateRole(roleID RoleId, params *SecurityControllerUpdateRoleParams, securityControllerUpdateRoleJSONRequestBody SecurityControllerUpdateRoleJSONRequestBody) (*AllItemsResponseRoles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRoleWithResponse(c.ClientInterface.(*Client).ctx, roleID, params, securityControllerUpdateRoleJSONRequestBody))
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
func (c *SecurityController) UpdateRuleWithBody(securityRuleID SecurityRuleId, params *SecurityControllerUpdateRuleParams, contentType string, body io.Reader) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRuleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, securityRuleID, params, contentType, body))
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
func (c *SecurityController) UpdateRule(securityRuleID SecurityRuleId, params *SecurityControllerUpdateRuleParams, securityControllerUpdateRuleJSONRequestBody SecurityControllerUpdateRuleJSONRequestBody) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateRuleWithResponse(c.ClientInterface.(*Client).ctx, securityRuleID, params, securityControllerUpdateRuleJSONRequestBody))
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
func (c *SecurityController) UpdateUserWithBody(userIDRequired UserIdRequired, params *SecurityControllerUpdateUserParams, contentType string, body io.Reader) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateUserWithBodyWithResponse(c.ClientInterface.(*Client).ctx, userIDRequired, params, contentType, body))
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
func (c *SecurityController) UpdateUser(userIDRequired UserIdRequired, params *SecurityControllerUpdateUserParams, securityControllerUpdateUserJSONRequestBody SecurityControllerUpdateUserJSONRequestBody) (*AllItemsResponseUsers, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SecurityControllerUpdateUserWithResponse(c.ClientInterface.(*Client).ctx, userIDRequired, params, securityControllerUpdateUserJSONRequestBody))
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

// SyscheckController implementation of the SyscheckController interface
type SyscheckController struct {
	*ClientWithResponses
}

// DeleteSyscheckAgent calls the Syscheck controller´s function
func (c *SyscheckController) DeleteSyscheckAgent(agentID AgentId, params *SyscheckControllerDeleteSyscheckAgentParams) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerDeleteSyscheckAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscheckController) GetLastScanAgent(agentID AgentId, params *SyscheckControllerGetLastScanAgentParams) (*AllItemsResponseSyscheckLastScan, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerGetLastScanAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscheckController) GetSyscheckAgent(agentID AgentId, params *SyscheckControllerGetSyscheckAgentParams) (*AllItemsResponseSyscheckResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerGetSyscheckAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscheckController) PutSyscheck(params *SyscheckControllerPutSyscheckParams) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscheckControllerPutSyscheckWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// DeleteFiles calls the Manager controller´s function
func (c *ManagerController) DeleteFiles(params *ManagerControllerDeleteFilesParams) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerDeleteFilesWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetAPIConfig calls the Manager controller´s function
func (c *ManagerController) GetAPIConfig(params *ManagerControllerGetApiConfigParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetApiConfigWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetConfValidation(params *ManagerControllerGetConfValidationParams) (*ConfigurationValidation, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetConfValidationWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetConfiguration(params *ManagerControllerGetConfigurationParams) (*WazuhMangerConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetConfigurationWithResponse(c.ClientInterface.(*Client).ctx, params))
	if err != nil {
		return nil, err
	}
	// convert to *WazuhMangerConfiguration
	if i, ok := r.(*WazuhMangerConfiguration); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetFiles calls the Manager controller´s function
func (c *ManagerController) GetFiles(params *ManagerControllerGetFilesParams) (*struct {
	Contents *string "json:\"contents,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetFilesWithResponse(c.ClientInterface.(*Client).ctx, params))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Contents *string "json:\"contents,omitempty\"" }
	if i, ok := r.(*struct {
		Contents *string "json:\"contents,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetInfo calls the Manager controller´s function
func (c *ManagerController) GetInfo(params *ManagerControllerGetInfoParams) (*WazuhInfo, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetLogSummary(params *ManagerControllerGetLogSummaryParams) (*WazuhLogsSummary, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetLogSummaryWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetLog(params *ManagerControllerGetLogParams) (*AllItemsResponseWazuhLogs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetLogWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetManagerConfigOndemand(component Component, configuration Configuration, params *ManagerControllerGetManagerConfigOndemandParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetManagerConfigOndemandWithResponse(c.ClientInterface.(*Client).ctx, component, configuration, params))
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
func (c *ManagerController) GetStatsAnalysisd(params *ManagerControllerGetStatsAnalysisdParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsAnalysisdWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetStatsHourly(params *ManagerControllerGetStatsHourlyParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsHourlyWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetStatsRemoted(params *ManagerControllerGetStatsRemotedParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsRemotedWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetStatsWeekly(params *ManagerControllerGetStatsWeeklyParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsWeeklyWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetStats(params *ManagerControllerGetStatsParams) (*AllItemsResponseWazuhStats, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatsWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ManagerController) GetStatus(params *ManagerControllerGetStatusParams) (*WazuhDaemonsStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerGetStatusWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// PutFilesWithBody calls the Manager controller´s function
func (c *ManagerController) PutFilesWithBody(params *ManagerControllerPutFilesParams, contentType string, body io.Reader) (*struct {
	ApiResponse         "yaml:\",inline\""
	ConfirmationMessage "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerPutFilesWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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

// PutRestart calls the Manager controller´s function
func (c *ManagerController) PutRestart(params *ManagerControllerPutRestartParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ManagerControllerPutRestartWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *OverviewController) GetOverviewAgents(params *OverviewControllerGetOverviewAgentsParams) (*OverviewAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.OverviewControllerGetOverviewAgentsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// ScaController implementation of the ScaController interface
type ScaController struct {
	*ClientWithResponses
}

// GetScaAgent calls the Sca controller´s function
func (c *ScaController) GetScaAgent(agentID AgentId, params *ScaControllerGetScaAgentParams) (*AllItemsResponseSCADatabase, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScaControllerGetScaAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *ScaController) GetScaChecks(agentID AgentId, policyID PolicyId, params *ScaControllerGetScaChecksParams) (*AllItemsResponseSCAChecks, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScaControllerGetScaChecksWithResponse(c.ClientInterface.(*Client).ctx, agentID, policyID, params))
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

// AgentsController implementation of the AgentsController interface
type AgentsController struct {
	*ClientWithResponses
}

// AddAgentWithBody calls the Agents controller´s function
func (c *AgentsController) AddAgentWithBody(params *AgentsControllerAddAgentParams, contentType string, body io.Reader) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerAddAgentWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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

// AddAgent calls the Agents controller´s function
func (c *AgentsController) AddAgent(params *AgentsControllerAddAgentParams, agentsControllerAddAgentJSONRequestBody AgentsControllerAddAgentJSONRequestBody) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerAddAgentWithResponse(c.ClientInterface.(*Client).ctx, params, agentsControllerAddAgentJSONRequestBody))
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

// DeleteAgents calls the Agents controller´s function
func (c *AgentsController) DeleteAgents(params *AgentsControllerDeleteAgentsParams) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
	OlderThan                *string "json:\"older_than,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerDeleteAgentsWithResponse(c.ClientInterface.(*Client).ctx, params))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllItemsResponseAgentIDs "yaml:\",inline\""; OlderThan *string "json:\"older_than,omitempty\"" }
	if i, ok := r.(*struct {
		AllItemsResponseAgentIDs "yaml:\",inline\""
		OlderThan                *string "json:\"older_than,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteGroups calls the Agents controller´s function
func (c *AgentsController) DeleteGroups(params *AgentsControllerDeleteGroupsParams) (*struct {
	AllItemsResponseGroupIDs "yaml:\",inline\""
	AgentGroupDeleted        "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerDeleteGroupsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// DeleteMultipleAgentSingleGroup calls the Agents controller´s function
func (c *AgentsController) DeleteMultipleAgentSingleGroup(params *AgentsControllerDeleteMultipleAgentSingleGroupParams) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerDeleteMultipleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// DeleteSingleAgentMultipleGroups calls the Agents controller´s function
func (c *AgentsController) DeleteSingleAgentMultipleGroups(agentID AgentId, params *AgentsControllerDeleteSingleAgentMultipleGroupsParams) (*struct {
	AllItemsResponseGroupIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerDeleteSingleAgentMultipleGroupsWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// DeleteSingleAgentSingleGroup calls the Agents controller´s function
func (c *AgentsController) DeleteSingleAgentSingleGroup(agentID AgentId, groupID GroupId, params *AgentsControllerDeleteSingleAgentSingleGroupParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerDeleteSingleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, agentID, groupID, params))
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

// GetAgentConfig calls the Agents controller´s function
func (c *AgentsController) GetAgentConfig(agentID AgentId, component Component, configuration Configuration, params *AgentsControllerGetAgentConfigParams) (*AgentConfiguration, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentConfigWithResponse(c.ClientInterface.(*Client).ctx, agentID, component, configuration, params))
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

// GetAgentFields calls the Agents controller´s function
func (c *AgentsController) GetAgentFields(params *AgentsControllerGetAgentFieldsParams) (*AllItemsResponseAgentsDistinct, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentFieldsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetAgentKey calls the Agents controller´s function
func (c *AgentsController) GetAgentKey(agentID AgentId, params *AgentsControllerGetAgentKeyParams) (*AllItemsResponseAgentsKeys, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentKeyWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// GetAgentNoGroup calls the Agents controller´s function
func (c *AgentsController) GetAgentNoGroup(params *AgentsControllerGetAgentNoGroupParams) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentNoGroupWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetAgentOutdated calls the Agents controller´s function
func (c *AgentsController) GetAgentOutdated(params *AgentsControllerGetAgentOutdatedParams) (*AllItemsResponseAgentsSimple, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentOutdatedWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetAgentSummaryOs calls the Agents controller´s function
func (c *AgentsController) GetAgentSummaryOs(params *AgentsControllerGetAgentSummaryOsParams) (*struct {
	ApiResponse "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentSummaryOsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetAgentSummaryStatus calls the Agents controller´s function
func (c *AgentsController) GetAgentSummaryStatus(params *AgentsControllerGetAgentSummaryStatusParams) (*AgentsSummaryStatus, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentSummaryStatusWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetAgentUpgrade calls the Agents controller´s function
func (c *AgentsController) GetAgentUpgrade(agentID AgentId, params *AgentsControllerGetAgentUpgradeParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentUpgradeWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// GetAgentsInGroup calls the Agents controller´s function
func (c *AgentsController) GetAgentsInGroup(groupID GroupId, params *AgentsControllerGetAgentsInGroupParams) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentsInGroupWithResponse(c.ClientInterface.(*Client).ctx, groupID, params))
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

// GetAgents calls the Agents controller´s function
func (c *AgentsController) GetAgents(params *AgentsControllerGetAgentsParams) (*AllItemsResponseAgents, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetAgentsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetGroupConfig calls the Agents controller´s function
func (c *AgentsController) GetGroupConfig(groupID GroupId, params *AgentsControllerGetGroupConfigParams) (*struct {
	AffectedItems      *[]GroupConfiguration "json:\"affected_items,omitempty\""
	TotalAffectedItems *int32                "json:\"total_affected_items,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetGroupConfigWithResponse(c.ClientInterface.(*Client).ctx, groupID, params))
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

// GetGroupFileJSON calls the Agents controller´s function
func (c *AgentsController) GetGroupFileJSON(groupID GroupId, fileName FileName, params *AgentsControllerGetGroupFileJsonParams) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetGroupFileJsonWithResponse(c.ClientInterface.(*Client).ctx, groupID, fileName, params))
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

// GetGroupFileXML calls the Agents controller´s function
func (c *AgentsController) GetGroupFileXML(groupID GroupId, fileName FileName, params *AgentsControllerGetGroupFileXmlParams) (*AgentsControllerGetGroupFileXmlResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetGroupFileXmlWithResponse(c.ClientInterface.(*Client).ctx, groupID, fileName, params))
	if err != nil {
		return nil, err
	}
	// convert to *AgentsControllerGetGroupFileXmlResponse
	if i, ok := r.(*AgentsControllerGetGroupFileXmlResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetGroupFiles calls the Agents controller´s function
func (c *AgentsController) GetGroupFiles(groupID GroupId, params *AgentsControllerGetGroupFilesParams) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetGroupFilesWithResponse(c.ClientInterface.(*Client).ctx, groupID, params))
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

// GetListGroup calls the Agents controller´s function
func (c *AgentsController) GetListGroup(params *AgentsControllerGetListGroupParams) (*AllItemsResponseGroups, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetListGroupWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetSyncAgent calls the Agents controller´s function
func (c *AgentsController) GetSyncAgent(agentID AgentId, params *AgentsControllerGetSyncAgentParams) (*AllItemsResponseAgentsSynced, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerGetSyncAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// InsertAgentWithBody calls the Agents controller´s function
func (c *AgentsController) InsertAgentWithBody(params *AgentsControllerInsertAgentParams, contentType string, body io.Reader) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerInsertAgentWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, contentType, body))
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

// InsertAgent calls the Agents controller´s function
func (c *AgentsController) InsertAgent(params *AgentsControllerInsertAgentParams, agentsControllerInsertAgentJSONRequestBody AgentsControllerInsertAgentJSONRequestBody) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerInsertAgentWithResponse(c.ClientInterface.(*Client).ctx, params, agentsControllerInsertAgentJSONRequestBody))
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

// PostGroup calls the Agents controller´s function
func (c *AgentsController) PostGroup(params *AgentsControllerPostGroupParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPostGroupWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// PostNewAgent calls the Agents controller´s function
func (c *AgentsController) PostNewAgent(params *AgentsControllerPostNewAgentParams) (*AgentIdKey, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPostNewAgentWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// PutAgentSingleGroup calls the Agents controller´s function
func (c *AgentsController) PutAgentSingleGroup(agentID AgentId, groupID GroupId, params *AgentsControllerPutAgentSingleGroupParams) (*AllItemsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPutAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, agentID, groupID, params))
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

// PutGroupConfigWithBody calls the Agents controller´s function
func (c *AgentsController) PutGroupConfigWithBody(groupID GroupId, params *AgentsControllerPutGroupConfigParams, contentType string, body io.Reader) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPutGroupConfigWithBodyWithResponse(c.ClientInterface.(*Client).ctx, groupID, params, contentType, body))
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

// PutMultipleAgentSingleGroup calls the Agents controller´s function
func (c *AgentsController) PutMultipleAgentSingleGroup(params *AgentsControllerPutMultipleAgentSingleGroupParams) (*struct {
	AllItemsResponseAgentIDs "yaml:\",inline\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPutMultipleAgentSingleGroupWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// PutUpgradeAgent calls the Agents controller´s function
func (c *AgentsController) PutUpgradeAgent(agentID AgentId, params *AgentsControllerPutUpgradeAgentParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPutUpgradeAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// PutUpgradeCustomAgent calls the Agents controller´s function
func (c *AgentsController) PutUpgradeCustomAgent(agentID AgentId, params *AgentsControllerPutUpgradeCustomAgentParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerPutUpgradeCustomAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// RestartAgent calls the Agents controller´s function
func (c *AgentsController) RestartAgent(agentID AgentId, params *AgentsControllerRestartAgentParams) (*ItemAffected, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerRestartAgentWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// RestartAgentsByGroup calls the Agents controller´s function
func (c *AgentsController) RestartAgentsByGroup(groupID GroupId, params *AgentsControllerRestartAgentsByGroupParams) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerRestartAgentsByGroupWithResponse(c.ClientInterface.(*Client).ctx, groupID, params))
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

// RestartAgentsByNode calls the Agents controller´s function
func (c *AgentsController) RestartAgentsByNode(nodeID NodeId, params *AgentsControllerRestartAgentsByNodeParams) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerRestartAgentsByNodeWithResponse(c.ClientInterface.(*Client).ctx, nodeID, params))
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

// RestartAgents calls the Agents controller´s function
func (c *AgentsController) RestartAgents(params *AgentsControllerRestartAgentsParams) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentsControllerRestartAgentsWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// ExperimentalController implementation of the ExperimentalController interface
type ExperimentalController struct {
	*ClientWithResponses
}

// ClearSyscheckDatabase calls the Experimental controller´s function
func (c *ExperimentalController) ClearSyscheckDatabase(params *ExperimentalControllerClearSyscheckDatabaseParams) (*AllItemsResponseAgentIDs, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerClearSyscheckDatabaseWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetCisCatResults(params *ExperimentalControllerGetCisCatResultsParams) (*AllItemsResponseCiscatResult, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetCisCatResultsWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetHardwareInfo(params *ExperimentalControllerGetHardwareInfoParams) (*AllItemsResponseSyscollectorHardware, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetHardwareInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetHotfixesInfo(params *ExperimentalControllerGetHotfixesInfoParams) (*AllItemsResponseSyscollectorHotfixes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetHotfixesInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetNetworkAddressInfo(params *ExperimentalControllerGetNetworkAddressInfoParams) (*AllItemsResponseSyscollectorNetwork, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkAddressInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetNetworkInterfaceInfo(params *ExperimentalControllerGetNetworkInterfaceInfoParams) (*AllItemsResponseSyscollectorInterface, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkInterfaceInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetNetworkProtocolInfo(params *ExperimentalControllerGetNetworkProtocolInfoParams) (*AllItemsResponseSyscollectorProtocol, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetNetworkProtocolInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetOsInfo(params *ExperimentalControllerGetOsInfoParams) (*AllItemsResponseSyscollectorOS, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetOsInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetPackagesInfo(params *ExperimentalControllerGetPackagesInfoParams) (*AllItemsResponseSyscollectorPackages, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetPackagesInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetPortsInfo(params *ExperimentalControllerGetPortsInfoParams) (*AllItemsResponseSyscollectorPorts, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetPortsInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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
func (c *ExperimentalController) GetProcessesInfo(params *ExperimentalControllerGetProcessesInfoParams) (*AllItemsResponseSyscollectorProcesses, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExperimentalControllerGetProcessesInfoWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// MitreController implementation of the MitreController interface
type MitreController struct {
	*ClientWithResponses
}

// GetAttack calls the Mitre controller´s function
func (c *MitreController) GetAttack(params *MitreControllerGetAttackParams) (*ApiResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.MitreControllerGetAttackWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// SyscollectorController implementation of the SyscollectorController interface
type SyscollectorController struct {
	*ClientWithResponses
}

// GetHardwareInfo calls the Syscollector controller´s function
func (c *SyscollectorController) GetHardwareInfo(agentID AgentId, params *SyscollectorControllerGetHardwareInfoParams) (*AllItemsResponseSyscollectorHardware, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetHardwareInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetHotfixInfo(agentID AgentId, params *SyscollectorControllerGetHotfixInfoParams) (*AllItemsResponseSyscollectorHotfixes, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetHotfixInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetNetworkAddressInfo(agentID AgentId, params *SyscollectorControllerGetNetworkAddressInfoParams) (*AllItemsResponseSyscollectorNetwork, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkAddressInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetNetworkInterfaceInfo(agentID AgentId, params *SyscollectorControllerGetNetworkInterfaceInfoParams) (*AllItemsResponseSyscollectorInterface, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkInterfaceInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetNetworkProtocolInfo(agentID AgentId, params *SyscollectorControllerGetNetworkProtocolInfoParams) (*AllItemsResponseSyscollectorProtocol, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetNetworkProtocolInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetOsInfo(agentID AgentId, params *SyscollectorControllerGetOsInfoParams) (*AllItemsResponseSyscollectorOS, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetOsInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetPackagesInfo(agentID AgentId, params *SyscollectorControllerGetPackagesInfoParams) (*AllItemsResponseSyscollectorPackages, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetPackagesInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetPortsInfo(agentID AgentId, params *SyscollectorControllerGetPortsInfoParams) (*AllItemsResponseSyscollectorPorts, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetPortsInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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
func (c *SyscollectorController) GetProcessesInfo(agentID AgentId, params *SyscollectorControllerGetProcessesInfoParams) (*AllItemsResponseSyscollectorProcesses, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.SyscollectorControllerGetProcessesInfoWithResponse(c.ClientInterface.(*Client).ctx, agentID, params))
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

// ListsController implementation of the ListsController interface
type ListsController struct {
	*ClientWithResponses
}

// GetListsFiles calls the Lists controller´s function
func (c *ListsController) GetListsFiles(params *ListsControllerGetListsFilesParams) (*AllItemsResponseListsFiles, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ListsControllerGetListsFilesWithResponse(c.ClientInterface.(*Client).ctx, params))
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

// GetLists calls the Lists controller´s function
func (c *ListsController) GetLists(params *ListsControllerGetListsParams) (*AllItemsResponseLists, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ListsControllerGetListsWithResponse(c.ClientInterface.(*Client).ctx, params))
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
