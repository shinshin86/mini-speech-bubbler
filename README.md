# mini-speech-bubbler



## Usage

Store the avatar image in the static directory with the name avatar.png before running.
Currently only the png format is supported.
(I will soon be supporting `jpeg` format as well.)

Execute this command to access the `http://localhost:3000`.

```bash
go run main.go
```



## Development

I use [this package (shogo82148/assets-life)](https://github.com/shogo82148/assets-life) to be able to use it as a single binary.



To generate the package, I'm running this command first.

(You don't have to run this command, but I'm writing it for my own history.)

```bash
assets-life static/ static
```

To regenerate the package, run this command.

```bash
go generate ./static
```

Running then access to `http://localhost:3000`.

```bash
go run main.go
```

