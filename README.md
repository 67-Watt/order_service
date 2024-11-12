# Order Service

The **Order Service** is a microservice in the Order Management System responsible for managing customer orders, order statuses, and order details for multiple restaurants across various countries. This service allows for the seamless creation, tracking, and updating of orders, with support for region-specific configurations such as currency, tax, and other order-related details.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [API Documentation](#api-documentation)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Environment Variables](#environment-variables)
- [Usage](#usage)
- [Development](#development)
    - [Running Tests](#running-tests)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

The Order Service is designed to handle all order-related operations, including order creation, updates, order item management, and status tracking. It supports a multi-tenant, multi-country configuration, allowing for customized settings per restaurant and region, including local tax, currency, and other regional specifications.

## Features

- **Order Management**: Create, update, and delete orders with detailed item and modifier information.
- **Order Status Tracking**: Track order statuses with support for various states (e.g., Pending, Completed, Cancelled).
- **Multi-Store and Multi-Country Support**: Configure orders per restaurant, including regional settings like currency and tax.
- **Integration with Inventory Service**: Checks and adjusts inventory levels based on order items.
- **Event-Driven Notifications**: Publishes events for other services (e.g., Inventory Service, Customer Service) to react to order updates.

## Architecture

The Order Service is part of a larger microservices architecture. Key interactions include:

1. **Inventory Service**: Updates inventory levels based on order items.
2. **Tax Service**: Calculates taxes for orders based on country and location.
3. **Reservation and Table Management Service**: Coordinates dine-in orders with table reservations.
4. **Finance and Accounting Service**: Logs order transactions for financial reporting and compliance.

## API Documentation

Detailed API documentation is available using [Swagger](https://swagger.io/). After running the service, you can access the API documentation at: [http://localhost:8080/swagger/](http://localhost:8080/swagger/)


## Getting Started

### Prerequisites

- **Go**: Ensure Go is installed (version 1.16+).
- **PostgreSQL**: Used as the main database. Ensure it is installed and configured.
- **Docker** (optional): For containerized deployment.

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/order_service.git
    cd order_service
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

### Environment Variables

Configure the following environment variables in a `.env` file at the root of the project:

```plaintext
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=order_service_db

# Server
SERVER_PORT=8080

# Other configurations (e.g., for interacting with Inventory or Tax services)
INVENTORY_SERVICE_URL=http://inventory_service
TAX_SERVICE_URL=http://tax_service
```

## Usage

To start the service, run:

```bash
go run main.go
```
The server will start on the port specified in the environment variables (default: 8080).

## Development

### Running Tests

To run the tests for the Order Service, use:

```bash
go test ./...
```
This will run all unit tests and integration tests for the service.


## Deployment

### Docker Deployment

1. Build the Docker image:

    ```bash
    docker build -t order_service .
    ```

2. Run the container:

    ```bash
    docker run -d -p 8080:8080 --env-file .env order_service
    ```

### Kubernetes Deployment

Use the provided `order_service.yaml` file to deploy on Kubernetes:

```bash
kubectl apply -f order_service.yaml
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
