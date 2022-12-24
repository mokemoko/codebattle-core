build-api:
	docker build -t cdbt-api -f Dockerfile.api ..

run-api:
	docker run --rm --env-file env.api \
	  -v ${PWD}/../sql/db.sqlite3:/sql/db.sqlite3 \
	  -p 8083:8080 \
	  cdbt-api --debug
