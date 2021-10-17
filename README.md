# REST service

Rest Service for the Multiply Logic ([learn more](https://github.com/kenesparta/multiplyLogic))

# 1. Requirements

| Software         | Version | Importance                   |
| ---------------- | ------- | ---------------------------- |
| 🐳 Docker         | 20.10.9 | Required                     |
| 🐙 Docker Compose | 1.29.2  | Required                     |
| 🐃 GNU Make       | 4.2.1   | Optional                     |
| ‍🚀 Postman        | 9.0.7   | Optional                     |

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

# 4. Pulling from Github Docker registry

1.Pullt he image from

```shell
docker pull ghcr.io/kenesparta/tk_rest_service:latest
```

2. Run the container with the image

```shell
docker run --rm -d -p 8084:8084 --name rest_service ghcr.io/kenesparta/tk_rest_service
```

# 5. Import the API

- You should import the API using JSON postman collection in the `doc/postman` directory.
- Yu also view the documentation using the **OpenAPI** specification by coping the content from the
  file `doc/open-api/multiply.yaml` and paste it on the site `https://editor.swagger.io/`

# 6. Test the API

You shoud send a POST request to `lcoalhost:8084/v1/multiply` with the JSON payload:
```shell
{
  "first_number": 0,
  "second_number": 0
}
```