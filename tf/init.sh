#!/bin/bash

TERRAFORM_STATE_STORAGE_KEY=$(az keyvault secret show --name terraformstoragekey --vault-name kv-oss-paris --query value -o tsv)

terraform init -backend-config="access_key=$TERRAFORM_STATE_STORAGE_KEY"
