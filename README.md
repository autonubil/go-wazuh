# go wazuh

go client for the wazuh [rest api](https://documentation.wazuh.com/4.0/user-manual/api/reference.html)

[![GitHub license](https://img.shields.io/github/license/xanzy/go-gitlab.svg)](https://github.com/autonubil/go-wazuh/blob/master/LICENSE)
[![Sourcegraph](https://sourcegraph.com/github.com/autonubil/go-wazuh/-/badge.svg)](https://sourcegraph.com/github.com/autonubil/go-wazuh?badge)
[![GoDoc](https://godoc.org/github.com/autonubil/go-wazuh?status.svg)](https://godoc.org/github.com/autonubil/go-wazuh)

it is generated from the OpenAPI 3.0 specifications. Thus it is not the most elegant API. Some effort has been put into an more go friendly interface by wrapping non successful results into errors and returning the `Data` objects instead of the raw result.

The list of supported controllers and their methods can be found in [controllerInterfaces.go](controllerInterfaces.go).

## Usage

```
import "github.com/autonubil/go-wazuh"
```

There are a few With... option functions that can be used to customize the API client:

- `WithBaseURL` custom base url
- `WithLogin` (username, password)
- `WithContext` (custom Context)
- `WithInsecure` allow insecure certificates
- `WithUserAgent` to set custom user agent

go-wazuh supports following environment variables for easy construction of a client:

- `WAZUH_URL`
- `WAZUH_USER`
- `WAZUH_PASSWORD`
- `WAZUH_INSECURE`

Construct a new Wazuh client, then use the various service on the client to access different parts of the wazuh API. For example, to list all agents:

```
c := NewAPIClient("https://localhost:55000", WithLogin("wazuh", "wazuh"), WithInsecure(true))
c.Authenticate()
agents := c.AgentsController.GetAgents(&AgentsControllerGetAgentsParams{})
fmt.Printf("Get Agents TotalAffectedItems %d\n", agents.AllItemsResponse.TotalAffectedItems)
for i, agent := range agents.AffectedItems {
    fmt.Printf(" %d: %s on %s\n", i, *agent.Id, *agent.NodeName)
}
```

Or use the environment to construct the client to get the server basic information:

```
c, err := NewClientFromEnvironment(WithInsecure(true))
if err != nil {
    panic(err)
}
// authenticate
err = c.Authenticate()
if err != nil {
    panic(err)
}

// call the DefaultInfo on the
status, err := c.Default.DefaultInfo(&DefaultControllerDefaultInfoParams{})
if err != nil {
    panic(err)
}
fmt.Printf("Connected to %s on %s\n", *status.Title, *status.Hostname)
```

## Testing

Prerequisite: <https://documentation.wazuh.com/4.0/docker/wazuh-container.html>
WAZUH\_\* environment variables must be configured.

Visual Studio Code launch configuration used for tests:

```
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch tests",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "program": "${workspaceFolder}",
      "env": {
        "WAZUH_URL": "https://localhost:55000",
        "WAZUH_USER": "wazuh",
        "WAZUH_PASSWORD": "wazuh",
        "WAZUH_INSECURE": true
      },
      "args": []
    }
  ]
}
```

## Compiling

you need zlib installed on the system


### MacOS (HomeBrew):

```bash
brew install zlib
brew install libdeflate
brew install pkg-config
```

### Ubuntu 

```bash
sudo apt-get install zlib1g zlib1g-dev
sudo apt-get install libdeflate0 libdeflate-dev
```

## ToDo

- more test cases

## Issues

- If you have an issue: report it on the [issue tracker](https://github.com/autonubil/go-wazuh/issues)

## Author

Carsten Zeumer (<carsten.zeumer@autonubil.net>)

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>
