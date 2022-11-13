generate-client:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/server/openapi.yaml -g typescript-fetch -o /local/front/src/generated/
generate-server:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/server/openapi.yaml -g go-gin-server -o /local/server/api/
mock:
	prism mock -p 8081 --cors server/openapi.yaml
