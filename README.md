# MNC Wallet Service API
This is a wallet service API built with Go. It provides a set of features for managing users' wallets, including registration, login, logout, profile management, top-up, and fund transfers. The service is built with Goqu ORM for database management, PostgreSQL for storage, Redis for caching, and RabbitMQ for messaging. Docker Compose is used to simplify service setup and running the environment.

## Features
- **Register**: User can create a new account.
- **Login**: User can log into their account.
- **Logout**: User can log out from the system.
- **View Profile**: User can view their profile details.
- **Update Profile**: User can update their profile information.
- **Top-up**: User can add funds to their wallet.
- **Transfer**: User can transfer funds to another user's wallet.

## Prerequisites
To run this project, you need the following:

- **Go (Golang)**: Version 1.18 or higher.
- **Docker**: Installed on your machine for running PostgreSQL, Redis, and RabbitMQ.
- **Docker Compose**: For managing multi-container Docker applications.

## Getting Started

### 1. Clone the repository
Clone this repository to your local machine.

```bash
git clone https://github.com/yourusername/wallet-service.git

cd wallet-service
```

### 2. Configuration
The configuration for the application is located in the `config.json` file. Make sure the configurations are correct before starting the services.

Example `config.json`:


- **PostgreSQL**: Connection information for PostgreSQL.
- **Redis**: Connection details for the Redis server.
- **RabbitMQ**: URL for the RabbitMQ messaging system.

### 3. Set Up the Services
To set up all necessary services (PostgreSQL, Redis, RabbitMQ), you can use Docker Compose. This will automatically pull the required Docker images and start the services.

Run the following command to start all services:

```bash
docker-compose up
```

This will start the following services:

- **PostgreSQL**: Running on port `5432`.
- **Redis**: Running on port `6379`.
- **RabbitMQ**: Running on port `5672`.

### 4. Migrate Database
This project with auto migration. so you can skip part of migration.

### 5. Run the Service
To start the wallet service, run the following command:

``` bash
go run main.go
```

By default, the service will run on `localhost:8080`.

### 6. API Endpoints
The wallet service exposes the following API endpoints:

- **POST** `api/users/register`: Register a new user account.
- **POST** `api/users/login`: Login to obtain a session token.
- **POST** `api/users/logout`: Logout from the current session.
- **GET** `api/users/profile`: View the current user's profile.
- **PUT** `api/users/profile`: Update the user's profile information.
- **POST** `api/users/top-up`: Top-up the user's wallet balance.
- **POST** `api/users/transfer`: Transfer funds to another user.

### 7. Postman Docs
To make testing easier, a Postman collection with predefined API requests is provided. 
- **Link**: [Docs Postman](https://documenter.getpostman.com/view/36497926/2sAYXEEy7c)

### 8. Stop the Services
After you're done, you can stop all running services by using:

```bash
docker-compose down
```

This will stop and remove the containers.

## Environment Variables
The service configuration is managed through the `config.json` file. You can adjust the settings for PostgreSQL, Redis, and RabbitMQ by editing this file.

## Conclusion
This wallet service provides basic functionality for managing user accounts and wallet operations. It uses Docker Compose to simplify the setup of PostgreSQL, Redis, and RabbitMQ. The Goqu ORM is used for database interactions, and RabbitMQ is utilized for messaging.
Feel free to extend this service with additional features such as enhanced security, notifications, or transaction history. If you encounter any issues or have any questions, please open an issue or contribute to the repository.

### License
No License Yet.




