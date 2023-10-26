docker build --no-cache -t food-api:latest -f deploy/Dockerfile .

docker create --name food-api -p 8080:8080 food-api:latest

docker start food-api