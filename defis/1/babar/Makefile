all:
	@cargo build

run:
	@cargo run res/data.csv res/user.txt res/view0.csv res/changes.txt

build:
	@docker build -t ws .
rund:
	@docker run ws
