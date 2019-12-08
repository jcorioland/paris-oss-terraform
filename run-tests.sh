#!/bin/bash

set -e

# ensure dependencies
dep ensure

# set environment variables
export TF_VAR_service_principal_client_id=$SERVICE_PRINCIPAL_CLIENT_ID
export TF_VAR_service_principal_client_secret=$SERVICE_PRINCIPAL_CLIENT_SECRET
export AZURE_SUBSCRIPTION_ID=$ARM_SUBSCRIPTION_ID
export AZURE_TENANT_ID=$ARM_TENANT_ID
export AZURE_CLIENT_ID=$ARM_CLIENT_ID
export AZURE_CLIENT_SECRET=$ARM_CLIENT_SECRET

# run test
go test -v ./tests/ -timeout 30m | tee test_output.log
terratest_log_parser -testlog test_output.log -outputdir test_output