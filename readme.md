# Booking App - Go
[Youtube course](https://www.youtube.com/watch?v=yyUHQIec83I)

## Creating the project
```bash
go mod init booking-app
```

## Running
```bash
go run main.go
```

## Constants and variables
- We can use the short way to declare variables
    ```go
    conferenceName := "Go Conference"
    ```
- We cannot use it
    - for constants
    - for explicity data type like `variableName uint := 2`

## Data Types

### Int
| Go | Java |
| -- | -- |
| int8  | byte  |
| int16 | short |
| int32 | integer |
| int64 | long |

Use `uint8`, `uint16`, `uint32` and `uint64` for positive numbers