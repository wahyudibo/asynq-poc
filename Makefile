GENERATED_DOCKER_IMAGES := $(shell docker images | grep "asynq-poc" | awk '{print $$1}')

svc-up:
	docker compose up -d

svc-down:
	docker compose down -v --remove-orphans

delete-images:
ifneq ($(strip $(GENERATED_DOCKER_IMAGES)),)
	docker rmi -f $(GENERATED_DOCKER_IMAGES)
	@echo "generated docker images succesfully deleted"
else
	@echo "generated docker images are not found"
endif