# Property agent

A simple go script for easy surfing of properties listed on domain.com.

## Setup:

1. Ensure you have go installed on your system.
2. Update parameters in the `config.toml` file. The parameters are self
explanatory.
3. The script can be run directly with:

```
$ go run main.go
```

or it can be compiled to a binary using:

```
$ go build
```

and then running:

```
$ ./property-agent
```
4. The generated files will be available in the `./temp/` directory. Accepted
listings will be in the `accepted-%` file and vice versa for rejected listings.
