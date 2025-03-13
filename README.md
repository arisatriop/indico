# Technical test 

## Overview

This project contains the technical test for the Backend Engineer (Golang) position at Indico.

## Tech Stack

- **Programming Language:** Go
- **Database:** PostgreSQL
- **Frameworks/Libraries:** Gin, GORM
- **Other Tools:** [Docker/Postman/etc.]

## Installation & Setup
##### **1. Clone the repository:**
   ```sh
   git clone https://github.com/arisatriop/indico.git
   cd indico
   ```



##### **2. Run .sql query in /docs/sql to create base table**
   ```sh
   location.sql 
   orders.sql
   products.sql 
   users.sql
   ```

##### **3. Configure environment variables:**
   ```sh
   DBMS=
   DB_USER=
   DB_PASSWORD=
   DB_HOST=
   DB_HOST=
   DB_PORT=
   DB_NAME=
   ```

##### **4. Run the application:**
   ```ssh
   go run main.go

   ```

## Running with Docker
Alternatively, you can run the application using Docker:
```
docker compose up
```


