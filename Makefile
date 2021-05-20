include .env

.PHONY: openapi
openapi:
	oapi-codegen --config api/openapi/todolist.cfg.yaml api/openapi/todolist.yaml