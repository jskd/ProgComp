.PHONY: all run build rund mrproper clean

build:
	@docker build -t ws .

all:
	@cargo build

mrproper:
	@docker run ws make clean

clean:
	@cargo clean


