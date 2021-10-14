# data-points
Simple system for saving and calculating data

## Requirements
Docker & Docker compose

## Run with
`docker-compose up --build'

## API

### Save items into db
url: `http://localhost:8000/api`
method: POST
payload: json
example: `curl_requests/add_data.sh`

### Get basic statistics
url: `http://localhost:8000/api/<name>`
method: GET
path variable: data point name
query arguments:
- `from` time range start (epoch)
- `to` time range stop (epoch)
example: `curl_requests/get_stats.sh`

## Tech stack
API service
- python/django

Calculetions service
- go

DB
- postgres

## What is missing
- Healthcheck endpoint
- Open API spec
- Better setup scripting
- Code documentation 
