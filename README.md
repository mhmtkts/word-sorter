# word-sorter Project

A Go application that sorts words based on the number of "a" characters. The application sorts words in descending order based on the count of "a"s, and for words with equal counts, sorts them by their lengths.

## Project Structure

- `src/main.go`: Entry point of the application. Contains the `main` function that initiates the word sorting process.
- `src/sorter/sort.go`: Contains the function that sorts words based on "a" character count.
- `src/tests/sort_test.go`: Contains unit tests for functions in `sort.go`.
- `go.mod`: Contains project dependencies and module information.

## Implementation Plan

1. Create a custom sorting algorithm that:
   - Counts "a" characters in each word
   - Compares words based on "a" count (descending)
   - Uses word length as secondary sorting criteria
2. Implement unit tests
3. Create a simple CLI interface

## Usage

1. Clone or download the project
2. Navigate to the project directory
3. Run the application using:

```bash
go run src/main.go
```

## Tests

Use the following command to run the tests:

```bash
go test src/tests -v
```