FROM rust:1.23.0

#Avoids unnecessary recompilation
WORKDIR /app/
COPY . /app/
RUN ls -la /app/*
#Needed to run
RUN cargo install

CMD ["ws","ressources/data.csv", "ressources/user.txt", "ressources/view0.csv", "ressources/changes.txt"]

#docker run -v path_to_host:path_in_vm => repertoire dans host sera monte dans la vm
#faire un build
#le script ws doit lancer docker run, start le container, creer le volume
