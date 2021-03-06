# Xendit Software Engineer - Technical assesment

[![MONGO](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)](-)
[![GO](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](-)
[![GIT](https://img.shields.io/badge/GIT-E44C30?style=for-the-badge&logo=git&logoColor=white)](-)

## Tasks
### You've been asked to spin up an API on the following architecture:
- Application layer: Node.js with express or Golang 
- Database: MongoDB or Postgres
### The backend API, containerized to docker, should expose following functionality:
- POST requests to /orgs/<org-name>/comments should allow the user to persist
  comments (in a MongoDB collection or Postgres table) against a given github
  organization.
- GET requests to /orgs/<org-name>/comments/ should return an array of all the
  comments that have been registered against the organization.
- DELETE requests to /orgs/<org-name>/comments should soft delete all comments
  associated with a particular organization. We define a "soft delete" to mean that deleted
  items should not be returned in GET calls, but should remain in the database for
  emergency retrieval and audit purposes.
- GET requests to /orgs/<org-name>/members/ should return an array of members of an
  organization (with their login, avatar url, the numbers of followers they have, and the number of people they're following), sorted in descending order by the number
  of followers.

## System Documentations

Database Design

![Link](https://drive.google.com/uc?id=1TQ7nMonP0iwiG7qKoXu1t_trMl6ajrfh)

API Blueprints:

- [Comment APIs](https://documenter.getpostman.com/view/20770802/UyxdMVgk)
- [Organization APIs](https://documenter.getpostman.com/view/20770802/UyxdMps9)
- [User APIs](https://documenter.getpostman.com/view/20770802/UyxdMpsA)

## Setup

### Prerequisite

1. Golang 1.17
2. Docker-Compose

### Setup and Development Guide

1. install required dependencies

### Start
1. Run
```bash
  make all
```

### Run test
1. Run
```bash
  go test ./... -v
```
