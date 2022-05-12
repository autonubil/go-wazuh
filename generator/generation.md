## Generate Client

https://github.com/OpenAPITools/openapi-generator#1---installation

download https://raw.githubusercontent.com/wazuh/wazuh/v4.2.4/api/api/spec/spec.yaml and
for `/security/user/authenticate/run_as` rename operationId to: `api.controllers.security_controller.login_user_run_as`

```
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

rm -rf ./wazuh-generated/\* ; oapi-codegen -generate types -package api ./spec.yaml > ./api/types.go ; oapi-codegen -generate client -package api ./spec.yaml > ./api/client.go

```

you must fix the parameter in the `ApiControllersSecurityControllerLoginUserRunAs` functions to the generic (without RunAs) names
Regex Replace: `((params|body) \*?ApiControllersSecurityControllerLoginUser)RunAs` -> `$1`

Rename `Client` to `ApiClient` and `ClientOption` `ApiClientOption`

For clarity remove all `ApiControllers` prefixes

## Generate implementation

use "test" classes in `wazuh_gen_code_test.go`

Old

```
# curl -LO https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh
# chmod +x openapi-generator-cli.sh
# rm -rf ./wazuh-generated/\* ; ./openapi-generator-cli.sh generate -g go -i https://raw.githubusercontent.com/wazuh/wazuh/v4.2.4/api/api/spec/spec.yaml --http-user-agent "go-wazuh" -o ./# wazuh-generated --package-name wazuh --api-package api --model-package models --skip-validate-spec
```
