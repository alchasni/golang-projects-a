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

![Link](https://lh3.googleusercontent.com/fife/AAWUweXGvtheiUYwNldIwG06a7SaWyIZ5EQdexn-21AohCFaBgthkHMGvRaklExeqM0RSg8VNOuoJSnpsfFQBw5vqBZIK0zRNxDObd41w7qvZ2FSg5hM13wMC8ziBQQyW84amTQ5q9qRBmcT_E_Q-quFvJ_OR28NvI-PSRNvDZLbPnNIvvOFy27_2OkRpwvs5W3_wMBdfUy_v4FIGpse0hLNP8wlFc8DM1mYz3B2AQNkUIINYiAVPq3HfycZsMZ6FjQ1MHjlsqIQqOO0HrvYv43ktPseL1WTpuu5P4qdfAhb3L7zVuNqQD4CbpjMHh4nBetZdUHUzMqwIgc9FyacBWTkn5iqtXIyT8c66FZmq1FugyeCELWcZf6xgDNAq-ACjnpekTaP-7bdf6L-7we19z7vMM59QijvbYf2j6BcdQ0K_ZOTukCe_q7jrZoCgbgHI8F-6Tn5Z--BuLND-Ihg4m8lHBMdNOvBzyTyhL7lm_fTCNDtijjO7Eceg2NE_hMOATa4DwM6ddxA-d7Ono4IRYj-x2oO8yuHgCc6UIwvvPFZOTj6Cyz0UyEsuMCUIw1LzFhSMKtBRxOBBCiG8u1y4Xqd6LlFwH1Sl1HY6UXeVWaAHqt5VDBu4RPzJ7QGsC3vPr8rnvIku3TwKTj9DCbi55iStZb1hFyfNKPJygLkScSjVekkFqPVgMl7xbLwgeqHecob2te_hL4DKoKvIu1WucDPAvi5hokBCYsh_PrW9SGDL_uoh78Hhv3FRPw1z8-DuI3PjqvfJUiNNJxhw9BMnGm1yZlzuHaibFEM8UAsv0t09Lr5KscIY6bQxwti=w1366-h663)

API Blueprints:

- [Link]()

## Setup

### Prerequisite

1. Golang 1.17
2. MongoDB
3. Docker-Compose

### Setup and Development Guide

1. install required dependencies
```bash
  go mod
  go mod tidy
```

2. Setup docker

3. Setup environment and database
```bash
  cp env.sample .env
  # Then modify these .env variables according to your mysql config:
  # DATABASE_USERNAME, DATABASE_PASSWORD, DATABASE_TEST_USERNAME, DATABASE_TEST_PASSWORD
  migarte ???
```

### Run

```bash
  go run
```
