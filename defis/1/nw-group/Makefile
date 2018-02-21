IMG = nw-group-defi1

.PHONY: all build-docker-img mrproper

all: build-docker-img

build-docker-img:
	@docker build -t $(IMG) .

mrproper:
	echo "Do nothing"
	##@docker rmi -f $(IMG) 2> /dev/null
