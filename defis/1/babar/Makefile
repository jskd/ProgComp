all:
	@cargo build

run:
	@cargo run ressources/data.csv ressources/user.txt ressources/view0.csv ressources/changes.txt

build:
	@docker build -t ws .
rund:
	@docker run ws
