help: ## Show this help
	@printf "***\nUsage: Make {target}\nAvailable targets:\n\n"
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

set-project: ## set the correct project
	@gcloud config set project send-this-42ae

deploy: deploy-backend deploy-frontend  ## deploy frontend and backend
deploy-backend: deploy-geturls deploy-downurls  ## deploy backend cloud functions

deploy-geturls: set-project ## deploy geturls cloud function
	@cd backend/functions;gcloud functions deploy geturls --entry-point Geturls --runtime go116 --trigger-http --allow-unauthenticated > current-geturls-deploy.log 2>&1 &

deploy-downurls: set-project ## deploy downurls cloud function
	@cd backend/functions;gcloud functions deploy downurl --entry-point DownloadUrl --runtime go116 --trigger-http --allow-unauthenticated > current-downurl-deploy.log 2>&1 &

deploy-frontend: ## deploy the frontend angular project to firebase
	@firebase deploy


