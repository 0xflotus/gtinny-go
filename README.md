# gtinny-go

It's a simple GTIN Validator in Go.

### Development

```sh
go get
go run .
```

### Build

```sh
go build
```

### Usage

```sh
> gtinny-go "97350053850012"
> echo $? # 0

> gtinny-go "12345678901234"
> echo $? # 1
```