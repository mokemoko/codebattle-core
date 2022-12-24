build-api:
	docker build -t cdbt-api -f Dockerfile.api ..

run-api:
	docker run --rm --env-file env.api \
	  -v ${PWD}/../sql/db.sqlite3:/sql/db.sqlite3 \
	  -p 8083:8080 \
	  cdbt-api --debug

build-batch:
	docker build -t cdbt-batch -f Dockerfile.batch ..

run-batch-%:
	docker run --rm \
	  -v ${PWD}/../sql/db.sqlite3:/sql/db.sqlite3 \
	  cdbt-batch --debug --contestId 90090783-4d64-4d02-ba26-cc62b8111775 \
	  --command ${@:run-batch-%=%}
