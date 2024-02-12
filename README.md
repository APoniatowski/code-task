# Customer Email Microservice

This microservice is designed to store customer information and send them emails based on mailing ID. It is written in Go and uses PostgreSQL as the database. Both the database and the application are containerized and ready for deployment using an orchestration tool of your choice.

## Getting Started

Follow these instructions to set up and run the microservice locally.

## Prerequisites

Make sure you have Docker installed on your machine. You can download and install Docker from [https://www.docker.com/get-started](https://www.docker.com/get-started).

## Installation

1. Clone the repository to your local machine:
```bash
git clone https://github.com/APoniatowski/code-task
```

2. Navigate to the project directory:
```bash
cd code-task
```
## Running the Microservice

To run the microservice, execute the following command:
```bash
docker-compose up
```

This command will build the Docker images for the application and the PostgreSQL database, and then start the containers.

## API Endpoints

### Add Customer Messages

To add customer messages to the database, use the following `curl` command:
```bash
curl -X POST localhost:8080/api/messages -d '{"email":"jan.kowalski@example.com","title":"Interview","content":"simple text","mailing_id":1, "insert_time": "2020-04-24T05:42:38.725412916Z"}'
```
Replace the values in the JSON payload with the relevant customer information.

### Send Email to Customers

To send emails to customers with a specific mailing ID, use the following `curl` command:
```bash
curl -X POST localhost:8080/api/messages/send -d '{"mailing_id":1}'
```

This command will send a mocked message to customers with the specified mailing ID and delete those customers from the database.

### Delete Customer Entry

To delete a customer entry from the database, use the following `curl` command:
```bash
curl -X DELETE localhost:8080/api/messages/{id}
```

Replace `{id}` with the ID of the customer entry you want to delete.

## Cleanup

To delete all customer entries older than 5 minutes from the database, the microservice includes an automatic cleanup process.

## Testing

Unit tests have been provided for one function/method. To run the tests, execute the following command:
```bash
go test ./...
```

