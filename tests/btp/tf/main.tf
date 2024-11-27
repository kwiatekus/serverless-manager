terraform {
  required_providers {
    btp = {
      source  = "SAP/btp"
    }
    http = {
      source = "hashicorp/http"
    }
    http-full = {
      source = "salrashid123/http-full"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.32.0"
    }   
  }
}


provider "http" {}
provider "http-full" {}
provider "btp" {
  globalaccount = var.BTP_GLOBAL_ACCOUNT
  cli_server_url = var.BTP_BACKEND_URL
  idp            = var.BTP_CUSTOM_IAS_TENANT
  username = var.BTP_BOT_USER
  password = var.BTP_BOT_PASSWORD
}

module "kyma" {
  source = "git::https://github.com/kyma-project/terraform-module.git?ref=v0.3.1"
  BTP_KYMA_PLAN = var.BTP_KYMA_PLAN
  BTP_NEW_SUBACCOUNT_NAME = var.BTP_NEW_SUBACCOUNT_NAME
  BTP_KYMA_REGION = var.BTP_KYMA_REGION
  BTP_NEW_SUBACCOUNT_REGION = var.BTP_NEW_SUBACCOUNT_REGION
}

output "subaccount_id" {
  value = module.kyma.subaccount_id
}

output "domain" {
  value = module.kyma.domain
}
