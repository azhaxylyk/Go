# go-reloaded

This Go project provides a versatile text processing utility capable of handling various transformations such as punctuation correction, converting binary and hexadecimal numbers to decimal, case conversion based on specific markers, and automatic article adjustment before words starting with vowel.

## Getting Started

Navigate to the project directory:
```bash
cd project/go-reloaded/try
```

## Usage

```bash
go run . sample.txt result.txt
```
Replace sample.txt with the input file containing the text you want to process and result.txt with the desired output file name. This command executes the text processing utility, applying various transformations to the input text and saving the result to the specified output file.

## Running Tests

To ensure the functionality works as expected, you can run the included tests with:

```bash
go test ./...
```
This command runs all tests in the project directory and subdirectories, providing a report on their success or failure. It's a good practice to run tests to verify that the implemented features are working correctly and to catch any potential issues early in the development process.
