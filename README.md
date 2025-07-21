# Glass Photo Backend API

This repository contains the backend API for the Glass Photo application. This README provides instructions on how to set up and run the application using Docker, and how to create the necessary local PostgreSQL database.

## Table of Contents

1.  [Prerequisites](#1-prerequisites)
2.  [Database Setup (Local PostgreSQL)](#2-database-setup-local-postgresql)
    - [Installation (if not already installed)](#installation-if-not-already-installed)
    - [Configuration](#configuration)
    - [Database and Table Creation](#database-and-table-creation)
3.  [Running the Application with Docker](#3-running-the-application-with-docker)
    - [Pulling the Docker Image](#pulling-the-docker-image)
    - [Environment Variables](#environment-variables)
    - [Running the Docker Container](#running-the-docker-container)
4.  [API Endpoints](#4-api-endpoints) (Optional: Replace with actual API docs or link to them)
5.  [Contributing](#5-contributing) (Optional)
6.  [License](#6-license) (Optional)

---

## 1. Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Docker Desktop** (for Windows/macOS) or **Docker Engine** (for Linux)
  - [Install Docker](https://docs.docker.com/get-docker/)
- **PostgreSQL** (if you plan to run a local database, which is recommended for development)
  - [Download PostgreSQL](https://www.postgresql.org/download/)
- **`psql` client** (command-line client for PostgreSQL, usually installed with PostgreSQL)

---

## 2. Database Setup (Local PostgreSQL)

This section guides you through setting up a local PostgreSQL database instance and creating the necessary tables.

### Installation (if not already installed)

Follow the official PostgreSQL installation guide for your operating system.

- **Ubuntu/Debian:**
  ```bash
  sudo apt update
  sudo apt install postgresql postgresql-contrib
  sudo systemctl enable postgresql
  sudo systemctl start postgresql
  ```
- **Arch Linux:**
  ```bash
  sudo pacman -S postgresql
  sudo systemctl enable postgresql
  sudo systemctl start postgresql
  ```
- **macOS (using Homebrew):**
  ```bash
  brew install postgresql
  brew services start postgresql
  ```
- **Windows:** Download the installer from the [official PostgreSQL website](https://www.postgresql.org/download/windows/).

### Configuration

For your Docker container to connect to your local PostgreSQL instance, PostgreSQL must be configured to accept connections from the Docker network.

1.  **Edit `postgresql.conf`:**
    Locate your `postgresql.conf` file (common paths: `/etc/postgresql/{version}/main/postgresql.conf` on Linux, or in your data directory).
    Find `listen_addresses` and change it to:

    ```
    listen_addresses = '*'
    ```

    This allows PostgreSQL to listen on all network interfaces.

2.  **Edit `pg_hba.conf`:**
    Locate your `pg_hba.conf` file (usually in the same directory as `postgresql.conf`).
    Add the following line to allow connections from the Docker bridge network (Docker's default network usually uses the `172.17.0.0/16` subnet):

    ```
    # TYPE  DATABASE        USER            ADDRESS                 METHOD
    host    all             all             172.17.0.0/16           md5
    ```

    - `md5` specifies password-based authentication. If you need a simpler (less secure for production) setup for local development, you could temporarily use `trust` (e.g., `host all all 172.17.0.0/16 trust`).
    - **Important:** Ensure this line is placed **above** any more restrictive rules that might override it.

3.  **Restart PostgreSQL:**
    Apply the changes by restarting the PostgreSQL service.
    - **Linux:**
      ```bash
      sudo systemctl restart postgresql
      ```
    - **macOS (Homebrew):**
      ```bash
      brew services restart postgresql
      ```
    - **Windows:** Restart the PostgreSQL service via Services Manager.

### Database and Table Creation

1.  **Access PostgreSQL terminal (`psql`):**
    Switch to the `postgres` user (if on Linux) or simply open `psql`.

    ```bash
    # On Linux, switch to postgres user
    sudo -i -u postgres psql
    # On other OS, just run psql if it's in your PATH
    psql -U postgres # You might be prompted for a password
    ```

    (If you don't have a password for the `postgres` user, you might want to set one: `ALTER USER postgres WITH PASSWORD 'your_password';`)

2.  **Create the database:**

    ```sql
    CREATE DATABASE "glass-photo";
    ```

3.  **Connect to the new database:**

    ```sql
    \c glass-photo;
    ```

4.  **Create the tables:**
    Paste the following SQL commands into the `psql` prompt:

    ```sql
    CREATE TABLE users (
        id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        user_img TEXT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        description VARCHAR(128),
        created_at TIMESTAMPTZ DEFAULT now()
    );

    CREATE TABLE posts (
        id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        post_img TEXT NOT NULL,
        title TEXT NOT NULL,
        description TEXT,
        created_at TIMESTAMPTZ DEFAULT now()
    );

    CREATE TABLE comments (
        id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
        content TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT now()
    );

    CREATE TABLE likes (
        id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
        created_at TIMESTAMPTZ DEFAULT now(),
        UNIQUE (user_id, post_id)
    );
    ```

5.  **Exit `psql`:**
    ```sql
    \q
    ```

Your local database is now ready!

---

## 3. Running the Application with Docker

This application is provided as a Docker image, making it easy to run.

### Pulling the Docker Image

First, pull the latest version of the application's Docker image from Docker Hub.

```bash
docker pull h1mky/adelevnikita:latest
```

```
docker run -d \
  -e DB_CONNECT="postgresql://postgres:your_password@<HOST_IP_ADDRESS>:5432/glass-photo?sslmode=disable" \
  -e PORT="3000" \
  -e JWT_SECRET="your_super_secret_jwt_key" \
  -p 3000:3000 \
  --name glass-photo \
  h1mky/adelevnikita:latest
```
