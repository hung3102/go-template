# gcim-temporary

## Backend setup(by Docker)

1. Tool install & setup (first-time only)

```bash
make init
```

2. Login with Cloud SDK (first-time only)

```bash
docker compose run --rm cloud-sdk gcloud auth application-default login
```

3. Start up

```bash
# Start web api
docker compose up general-api

# Start batch api
docker compose up batch-api
```

Web API URL: http://localhost:27001/api/health  
Batch API URL: http://localhost:27002/api/health  
Emulator URL: http://localhost:4000/

## How to
- Debug: https://www.notion.so/backend-da60aab413974c0f9bddd0ad1850e860
- Generate code:
```bash
$ cd back
$ make generate
``` 
