# Project Overview

This project is a Go application that fetches and displays a user's public activity from GitHub. It takes a GitHub username as a command-line argument, makes a request to the GitHub API, and then prints a formatted table of the user's recent events to the console.

## Main Technologies

* **Language:** Go
* **Key Libraries:**
  * `net/http`: For making HTTP requests to the GitHub API.
  * `encoding/json`: For parsing the JSON response from the API.
  * `text/tabwriter`: For formatting the output into a clean, aligned table.

## Architecture

The application is composed of two main files:

* `main.go`: This file contains the core logic of the application. It handles command-line argument parsing, API request construction, JSON decoding, and formatted output.
* `models.go`: This file defines the Go structs that map to the JSON structure of the GitHub API's event response. This provides type safety and makes working with the API data easier.

# Building and Running

## Building the project

To build the project, you can use the `go build` command. This will create an executable file named `github-activity` in the current directory.

```bash
go build
```

## Running the application

To run the application, execute the compiled binary and provide a GitHub username as an argument:

```bash
./github-activity <github_username>
```

For example:

```bash
./github-activity luchrv
```

# Development Conventions

* **Error Handling:** The application uses standard Go error handling practices, checking for errors after operations that can fail (e.g., HTTP requests, JSON decoding) and printing an error message to the console if an error occurs.
* **Code Style:** The code follows standard Go formatting conventions.
* **Dependencies:** The project has no external dependencies outside of the Go standard library.
