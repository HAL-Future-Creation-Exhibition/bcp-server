CONTAINER_NAME := bcp-nginx

build:
	docker build -t $(CONTAINER_NAME) ./docker
run:
	$(MAKE) build
	docker run --rm --name $(CONTAINER_NAME) -v $(PWD)/docker/html:/usr/share/nginx/html -d -p 80:80 $(CONTAINER_NAME)
stop:
	docker kill $(CONTAINER_NAME)
restart:
	$(MAKE) stop
	$(MAKE) run
shell:
	docker exec -it $(CONTAINER_NAME) /bin/bash