generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/server/openapi.yaml -g typescript-fetch -o /local/front/generated/
mock:
	prism mock -p 8081 --cors server/openapi.yaml
