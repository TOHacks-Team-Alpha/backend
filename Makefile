build:
	DOCKER_BUILDKIT=1 docker build -t tohacks-backend .

pull:
	docker pull alphakilo07/tohacks-backend

push:
	docker tag tohacks-backend alphakilo07/tohacks-backend
	docker push alphakilo07/tohacks-backend

cloud:
	
	docker tag tohacks-backend gcr.io/vagon-abe86/tohacks-backend
	docker push gcr.io/vagon-abe86/tohacks-backend

run:
	docker run  --rm -d -p 8081:8081 -e PORT='8081' \
		--name tohacks-backend tohacks-backend

kill:
	docker kill tohacks-backend
	