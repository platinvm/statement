# Statement

## Overview

Statement is a personal project that aims to create a simple, human-readable configuration language designed for structured data representation. It supports nested objects, various data types, and comments.

## Syntax Basics
### Comments
- Single-line: `// comment`
- Multi-line: `/* comment */`
### Data Types
- Integer: `10`
- Float: `3.141`
- Hex: `0xFF`
- Binary: `0b1001`
- String: `"hello, world!"`
- Multiline String: `"""multiline!"""`
- Boolean: `true`, `false`
- MAC Address: `ad:d2:5b:f1:17:b1`
- IPv4 Address: `242.205.159.106`
- IPv6 Address: `601e:bd66:e1fd:5437:fa41:d25a:a5c2:122f`
- Color: `#ffffff`
- Duration: `1h`, `10m`, `2.5s`, `50ms`
- Timestamp: `2023-04-05T12:34:56+07:00`
- Percentage: `50%`
- Version: `1.0.0-alpha`

### Todo
- refactor the lex to allow reusability of fragments and improve readability

## Structure
- Arrays are wrapped in `[ ]`
- Objects are wrapped in `{ }`
- Key-value pairs separated by `:`
- Items separated by `,` (optional for last item)

## Examples

Watch out the [examples directory](./src/test/resources/examples/) for some examples of the grammar.

## Test

To test the code run 
```sh
mvn clean test # "clean" is optional 
```

## Generate targets

To generate sources run 
```sh
mvn clean generate-sources # "clean" is optional
ls target/generated-sources/
# java
# ... other target defined in pom.xml
```