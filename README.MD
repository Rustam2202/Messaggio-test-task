<!-- POSTGRES_DSN=host=db 
user=user 
password=password 
dbname=messages 
port=5432 
sslmode=disable
KAFKA_BROKER=kafka:9092
# KAFKA_ADVERTISED_LISTENERS=localhost:9092 -->


# Messaggio Messages Processer

REST API server for store text messages in PostgreSQL database and using Apache Kafka for message processing.


## Deployment
Create .env file for set configs, for example:
```
POSTGRES_DSN=host=db 
user=user 
password=password 
dbname=messages 
port=5432 
sslmode=disable
KAFKA_BROKER=localhost:9092
LOG_LEVEL=debug
```
Use `makefile` commands:  
`make up` for up docker compose container.  
`make down` for shutdown and remove it.  

## API Endpoints

### Add message
URL: /api/messages  
Method: POST  
Request Body:
```
{
  "content": "My Message"
}
```
Response:
```
[
  {
    "id": 1,
    "content": "My Message",
    "processed": false
    "created_at": "2024-07-28T12:34:56Z"
  },
]
```

### Get All Messages  
URL: /api/messages  
Method: GET  
Response:
```
[
  {
    "id": 1,
    "content": "My Message",
    "processed": true
    "created_at": "2024-07-28T12:34:56Z"
  },
  ...
]
```

### Get processed messages count
URL: /api/messages/count  
Method: GET  
Response:
```
{
    "count": 1
}
```
