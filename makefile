WEB_CONTAINER_NAME := bcp-web
FILE_CONTAINER_NAME := bcp-file

web/build:
	docker build -t $(WEB_CONTAINER_NAME) ./docker/web
web/up:
	$(MAKE) web/build
	docker run --rm --name $(WEB_CONTAINER_NAME) -v $(PWD)/docker/web/html:/usr/share/nginx/html -d -p 80:80 $(WEB_CONTAINER_NAME)
web/down:
	docker kill $(WEB_CONTAINER_NAME)
web/restart:
	$(MAKE) web/down
	$(MAKE) web/up
web/shell:
	docker exec -it $(WEB_CONTAINER_NAME) /bin/bash

file/build:
	docker build -t $(FILE_CONTAINER_NAME) ./docker/file
file/up:
	$(MAKE) file/build
	docker run --rm --name $(FILE_CONTAINER_NAME) -v $(PWD)/docker/file/html:/usr/share/nginx/html -v $(PWD)/docker/file/tmp:/usr/share/tmp -d -p 3000:80 $(FILE_CONTAINER_NAME)
file/down:
	docker kill $(FILE_CONTAINER_NAME)
file/restart:
	$(MAKE) file/down
	$(MAKE) file/up
file/shell:
	docker exec -it $(FILE_CONTAINER_NAME) /bin/bash
	