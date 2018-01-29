.PHONY: all run build rund mrproper

build:
	@docker build -t ws .

all:
	@cargo build

run:
	@cargo run ressources/data.csv ressources/user.txt ressources/view0.csv ressources/changes.txt

rund:
	@docker run ws

mrproper: build
