# GitHub Activity Viewer

A simple Go application to view a GitHub user's public activity directly in your terminal.

## Overview

This project provides a command-line tool that takes a GitHub username, fetches their recent public activity from the GitHub API, and displays it in a clean, readable table format.

## Main Technologies

* **Language:** Go
* **Key Libraries:**
  * `net/http`: For making HTTP requests to the GitHub API.
  * `encoding/json`: For parsing the JSON response from the API.
  * `text/tabwriter`: For formatting the output into a clean, aligned table.

## Building and Running

### Prerequisites

* Go (version 1.x or higher)

### Building

To build the project, clone the repository and run the `go build` command. This will create an executable file named `github-activity` in the project directory.

```bash
go build
```

### Running

To run the application, execute the compiled binary and provide a GitHub username as an argument:

```bash
./github-activity <github_username>
```

**Example:**

```bash
./github-activity luchrv
```

## Development Conventions

* **Error Handling:** The application uses standard Go error handling practices, checking for errors after operations that can fail (e.g., HTTP requests, JSON decoding) and printing an error message to the console if an error occurs.
* **Code Style:** The code follows standard Go formatting conventions.
* **Dependencies:** The project has no external dependencies outside of the Go standard library.
