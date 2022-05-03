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

![Link](https://lh5.googleusercontent.com/6d4fj-4IyUe-uxeXz4kHVu-v-LBN8MZZOV5-o5P_9d0YPBmC5xbCp7xAvq1_KcYpGCsAXS18S1j1LA)

[//]: # (https://lh3.googleusercontent.com/fife/AAWUweWVOOqlTT6mTrk4SMv2apq56-MVDLOwLB6liCMVgdXuwJBKasNS_Vf04iQLx9zRkUAke4p9kdS0S1dcO9LvmsmBXBCxJLV-0pLVnNXbkZsSG0rlWWVAJ0Sz0Ufw2D8qMyfmCH-NeGvHBIfdSKXTNiVgaRpd9CDOtSGucBg_vg0HZX5wURhjnQxnEc30F1VUpcfbOQN-e7kM3fSLVYGkf6FbUEmBc1Gf9kKa796cBE5SSovG1p7RHj-IeAlriQ66BROfTz05m56ETLzrD480qPUEuXAK-Pq1zE5LIT5-XZsdY_VJ-C8BOu_N186WsP2hzqVFAGmNoa0IECS3AZDxLoqiDPGhYf3nYWA-LcAo2tKUIEAqRL4ZVcoCEAcJ5Bp3D945i70BZEifNfQm3vSSUi28IrTsGvOxW6LOEbwcaMrRogMpXgwDCYurVuw4kvyi9fp2qVTgshu2ixRDFUdRl1C36EeG5qagJRNtje1xVi1GUvzsV5Ck3jkVC0cPmYm22fl1MDmK-v0gZZNIlEcqjl1v_qZ0YgSImkZE9StEZYNg3biEZkSaR_2W605y-0PycLYZ4aVLeh3Hea15yp9wopAIDXKgNxK_ZWZfoZSavJGYQtJPU6feuvKctAD2koIs67j7uDQb0nsaQdBeRvqRY7-dnV6c5-Da4OxQh3kXkhogZbbB3l6DqVJnvueznfIL1n9E3Wht75hkKfOL6yY5scnGTeTJG6CiAdavDmCTZ88hnGUNLoWfbBOQKrWUsmTwvQZP8kKfEEcVEsR-2H84ZQK4rviHEUZ8ATdBli3VdjsYBIx70X0z8uhF=w1366-h663)
API Blueprints:

- [Link]()

## Setup

### Prerequisite

1. Golang 1.17
2. MongoDB
3. Docker-Compose

### Setup and Development Guide

1. install required dependencies
2. Run
```bash
  go mod
  go mod tidy
```

### Start
1. Run MongoDB (using docker-compose is recommended) in 127.0.0.1::27017
2. Run
```bash
  go run cmd/api/main.go
```
