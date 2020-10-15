# GoUrlShort

An API written in go, used to shorten urls. A test task for jMindSystems internship.

## Run in Docker
- `docker-compose up --build`, should build and run the server

## Run Locally

- Create database(`gourlshort`) and table(`schema.sql`)
- `cp .env.example .env`
- Change the credentials in the `.env` file
- `go build`
- `./gourlshort`

## Tasks

- [x] Write an API
- [x] Save in memory (Map)
- [x] Save to database (Used MySQL)
- [x] Unit tests
- [x] Cache redirect requests
- [x] Docker deploy
