# GoUrlShort

An API written in go, used to shorten urls. A test task for jMindSystems internship.

## Run Locally

- Start mysql, create database and table(`schema.sql`)
- `cp .env.example .env`
- Insert the credentials in the `.env` file
- `go build`
- `./gourlshort`

## Tasks

- [x] Write an API
- [x] Save in memory (Map)
- [x] Save to database (Used MySQL)
- [x] Unit tests
- [x] Cache redirect requests
- [ ] Docker deploy
