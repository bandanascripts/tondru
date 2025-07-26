
# Tondru

Tondru is a lightweight microservice for generating and validating JWTs, built with Go, Gin, and Redis. It’s designed to handle centralized token logic for secure backend systems.

## Features 

- Generates Access and Refresh JWTs with support for key rotation

- Validates tokens by extracting and verifying the kid from the token header

- Logs token usage and user authentication history into Redis.

## Tech Stack

- Golang 
- Gin (HTTP Framework)
- Redis
- JWTs

## Getting started 
```bash
git clone https://github.com/bandanascripts/tondru.git
cd tondru 
go build
```

🟥 Important: Make sure Redis is running locally on the default port (6379).
You can start Redis using Docker or a local install before running the service.


## API Endpoints

| Method | Route                   | Description                                       |
| ------ | ----------------------- | ------------------------------------------------- |
| POST   | `/tondru/generatetoken` | Generates access and refresh tokens               |
| GET    | `/tondru/inspect`       | Validates a token and returns the user's claims   |
| GET    | `/tondru/userhistory`   | Returns the user's token usage history from Redis |

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.
