# Belajar Golang Restful API

This repository contains a Go (Golang) project for building a RESTful API. It demonstrates the basic structure and components needed to create a RESTful API using Golang, including setting up a database, creating routes, middleware and implementing Unit Test.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

To get started with this project, follow these steps:

1. **Clone the repository:**

    ```sh
    git clone https://github.com/nugrohoarr/belajar-golang-restfulapi.git
    ```

2. **Navigate to the project directory:**

    ```sh
    cd belajar-golang-restfulapi
    ```

3. **Install the dependencies:**

    ```sh
    go mod tidy
    ```

4. **Set up the database:**

    Create a MySQL database named `db_golang_test` and ensure it is running on `localhost:3306`. You can modify the database connection settings in `test/setup_test.go` if needed.

5. **Run the application:**

    ```sh
    go run main.go
    ```

## Usage

To use the API, you can send HTTP requests to the endpoints defined in the project. You can use tools like `curl`, `Postman`, or any HTTP client of your choice to interact with the API.

## Project Structure

The project structure is organized as follows:

