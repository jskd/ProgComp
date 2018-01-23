FROM rust:1.23.0
WORKDIR /app
ADD . /app
RUN cargo install
CMD ["ws","res/data.csv", "res/user.txt", "res/view0.csv", "res/changes.txt"]
