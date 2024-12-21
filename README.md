# GO BOILERPLATE

**NOTE**: This project is work in progress.

---

## OVERVIEW

This is a boilerplate project for building backend services using Go. The primary goal is to provide a solid foundation 
for backend applications, enabling developers to focus on implementing business logic rather than boilerplate setup.

---

## FEATURES

This project uses the current technologies stack:

- **PostgreSQL** as the database for robust and scalable data storage.
- **Firebase** Authentication to manage user authentication and generate secure JWT access tokens.
- **Wire** as dependency injection framework to promote clean architecture.
- **Ready** to use example CRUD APIs
- **Unit test** integration to ensure code reliability

---

## PROJECT STRUCTURE

```
.
├── cmd                             # Entry point for applications
│   ├── app                         # Main application code
│   └── migrations                  # Database migrations
├── pkg                             # Application logic
│   ├── api                         # API module
│   │   ├── handler                 # API handlers
│   │   └── middleware              # API middleware
│   ├── config                      # Configuration files (e.g., environment variables)
│   ├── db                          # Database specific logic and connections
│   ├── di                          # Dependency injection  
│   ├── domain                      # Domain layer
│   ├── repository                  # Repository layer  
│   │   └── interfaces              # Repository interfaces
│   └── usecase                     # Use cases  
│       └── interfaces              # Use cases interfaces 
├── tests                           # All test files go here
│   ├── api                         # API tests
│   ├── usecase                     # Use case tests
│   ├── repository                  # Repository tests
│   └── integration                 # Integration or end-to-end tests
├── go.mod                          # Go module file
├── go.sum                          # Dependency checksums
└── README.md                       # Project documentation
```

| Directory           | Description                               |
|---------------------|-------------------------------------------|
| `cmd/app`           | Entry points for running the application. |
| `cmd/migrations`    | Database migrations tasks.                |
| `pkg/api`           | API handlers and middleware.              |
| `pkg/config`        | Application configuration logic.          |
| `pkg/db`            | Database models and connection logic.     |
| `pkg/di`            | Dependency injection setup.               |
| `pkg/domain`        | Core domain models and business logic.    |
| `pkg/repository`    | Data persistence logic (DB operations).   |
| `pkg/usecase`       | Application-specific use cases.           |
| `tests/api`         | API tests.                                |
| `tests/usecase`     | Use case tests.                           |
| `tests/repository`  | Repository tests.                         |
| `tests/integration` | Integration or end-to-end tests.          |

---

## INSTALLATION

### Prerequisites:

- [Go](https://golang.org/dl/) (version 1.23+)
- [PostgreSQL](https://www.postgresql.org/)
- Firebase Admin SDK credentials

### Clone the repository

```bash
git clone https://github.com/Melom01/go-boilerplate.git
cd go-boilerplate
```

### Set up dependencies
Install Go modules:

```bash
go mod tidy
```

---

## CONFIGURATION

Copy the sample `.env.example` file and create an `.env` file:

```shell
cp .env.example .env
```

Update the `.env` file with your configuration:
- **DB_HOST**: Database host. For example, **localhost** or **127.0.0.1**.
- **DB_PORT**: Database port. For example, **5432**.
- **DB_NAME**: Database name.
- **DB_USER**: Database username.
- **DB_PASSWORD**: Your password to access the database.
- **FIREBASE_CREDENTIALS**: The path to your Firebase Admin SDK JSON file.

---

## USAGE

### Run the application

```shell
go run cmd/app/main.go
```

### Run migrations

```shell
go run cmd/migrations/main.go
```

### Run tests

```shell
go test ./...
```

### Sort imports
#### Making the `sort_imports.sh` script executable

Before running the script for the first time, you need to make it executable. You can do this by running the following command in your terminal:
```shell
chmod +x sort_imports.sh
```
This step is only required once.

#### Running the script
After making the script executable, you can sort all imports throughout the application by running:
```shell
./sort_imports.sh
```
The script will automatically process all Go files, excluding any specified patterns (e.g., files with `_gen.go` in the di folder).
Once completed, you will see a confirmation message.

---

## LICENSE

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
