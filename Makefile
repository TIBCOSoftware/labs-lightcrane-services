SHELL := /bin/bash
SCRIPTS_PATH      := scripts

.PHONY: build-push-delete-lightcrane-service
build-push-delete-lightcrane-service: build-lightcrane-service push-image delete-local-image

.PHONY: build-lightcrane-service
build-lightcrane-service:
	@$(SCRIPTS_PATH)/build_lightcrane_service.sh ${SERVICE_NAME} ${APP_TYPE} ${IMAGE_NAME} ${IMAGE_TAG} ${IMAGE_URL} ${IMAGE_ARCH}
	
.PHONY: push-image
push-image:
	@$(SCRIPTS_PATH)/push_image.sh ${IMAGE_NAME} ${IMAGE_TAG} ${IMAGE_URL}

.PHONY: delete-local-image
delete-local-image:
	@$(SCRIPTS_PATH)/delete_local_image.sh ${IMAGE_NAME} ${IMAGE_TAG} ${IMAGE_URL}
	
.PHONY: build-installer
build-installer:
	@$(SCRIPTS_PATH)/build_installer.sh