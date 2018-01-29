.PHONY: all run build rund mrproper clean

build:
	@docker build -t ws .

all:
	@cargo build

run:
	@cargo run ressources/data.csv ressources/user.txt ressources/view0.csv ressources/changes.txt

rund:
	@docker run ws

mrproper:
	@docker run ws make clean

clean:
	@cargo clean


