# REST service
Rest Service for the Multiply Logic ([learn more](https://github.com/kenesparta/multiplyLogic))

# 1. Requirements

| Software         | Version | Importance                   |
| ---------------- | ------- | ---------------------------- |
| 🐳 Docker         | 20.10.9 | Required                     |
| 🐙 Docker Compose | 1.29.2  | Required                     |
| 🐃 GNU Make       | 4.2.1   | Optional                     |


# 2. Execute the service

The service runs over the port `8084`

## 2.1 Using makefile

Run `make l/up`

## 2.1 Using a docker commands

Run these commands:

```shell
docker-compose down --remove-orphans --rmi all
docker-compose up --detach --remove-orphans --force-recreate
```

# 3. Running the test

- Execute the command `make l/tco` or `make l/tch` (to get HTML report).
- If you want HTML report execute the command:
```shell
cd ./src/ ; go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out ; rm coverage.out
```
