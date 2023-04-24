#!/bin/bash
docker network create nt
docker pull "postgres:latest"

docker run --name exDB --network nt -e POSTGRES_PASSWORD=password -d postgres
sleep 5

docker exec -it exDB psql -U postgres -c "CREATE USER user1 WITH PASSWORD 'password1'"
docker exec -it exDB psql -U postgres -c "CREATE DATABASE thesaurus"
docker exec -it exDB psql -U postgres -c "\c thesaurus"
docker exec -it exDB psql -U postgres -c "CREATE TABLE synonyms (
   word CHAR(50),
   synonym TEXT[],
   similarity INTEGER[]
);"

docker exec -it exDB psql -U postgres -c "INSERT INTO synonyms (word, synonym, similarity) VALUES ('apple', '{fruit, red, delicious}', '{9, 5, 4}')"
docker exec -it exDB psql -U postgres -c "INSERT INTO synonyms (word, synonym, similarity) VALUES ('happy', '{delighted, fun, excited}', '{9, 8, 8}')"
docker exec -it exDB psql -U postgres -c "INSERT INTO synonyms (word, synonym, similarity) VALUES ('sad', '{frustrated, furious, mad}', '{9, 8, 7}')"