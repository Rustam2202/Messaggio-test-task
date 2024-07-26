up:
	make down
	docker compose up --build -d
down:
	docker compose down

zookeeper:
	docker run -d --name zookeeper -p 2181:2181 wurstmeister/zookeeper
run-kafka:
	docker run -d --name kafka -p 9092:9092 --link zookeeper:zookeeper \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
    -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
    wurstmeister/kafka
run-postgres:
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password postgres

	docker exec -it postgres psql -U postgres -c "CREATE DATABASE messages;"
	docker exec -it postgres psql -U postgres -d messages -c "CREATE TABLE messages (id SERIAL PRIMARY KEY, content TEXT NOT NULL, created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP);"
