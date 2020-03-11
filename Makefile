NAME=visket-playground
VERSION=0.0.1
TAG=aratanvm/$(NAME):$(VERSION)

docker/build:
	docker build -t $(TAG) .

docker/run: docker/build
	docker run -it --name $(NAME) $(TAG) /bin/ash

docker/stop:
	docker stop $(NAME)

docker/rm:
	docker rm $(NAME)

.PHONY: build run test clean
