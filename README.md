# go-accessor-gen

This tool creates getter and setter methods for Go structs.

It only works on fields that are pointers (like `*string` or `*int`).

## How to use

Run it on your Go file:

```bash
go run main.go --source input.go
```

This makes a new file named `input.generated.go`.

## Options

- `--source`: The file to read (required). You can use this flag many times.
- `--suffix`: The ending of the new file. Default is `.generated.go`.
