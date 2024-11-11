# Greetings App

A simple Go application that runs a web server and greets the user with a message. This project is containerized with Docker and can be easily deployed on platforms like Heroku or any Docker-compatible environment.

## Table of Contents

- [Overview](#overview)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Docker](#docker)


## Overview

This application uses a basic HTTP server written in Go. When deployed, it responds with a simple greeting message at the root URL. The app is designed for deployment in cloud environments and includes configurations for Docker and Heroku.

## Requirements

- **Go**: Version 1.20 or later, currently using tag 'latest' on Docker
- **Docker**: Optional, for containerized deployment
- **Heroku CLI**: Optional, for deployment to Heroku

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/username/greetings-app.git
   cd greetings-app
   ```

2. **Set up dependencies**:

   ```bash
   go mod tidy
   ```

## Usage

To run the app locally, execute the following command:

```bash
go run greetings.go
```

## Docker
**Build the Docker Image**
To build a Docker image for this application, use:

```bash
docker build -t esedgarcia/greetings .
```

And, to run the docker container use:

```bash
docker run -p 8080:8080 esedgarcia/greetings
```
