# mini-speech-bubbler

![Demo gif](./gif/demo.gif)



## Usage

Store the avatar image in the `static` directory with the name `avatar.png` before running.

**Note:**
Currently only the `png,jpeg(jpg)` format is supported.


Execute this command to access the `http://localhost:3000`.

```bash
go get github.com/shinshin86/mini-speech-bubbler
mini-speech-bubbler
```

If you want to use a single binary, please download from [this Release page](https://github.com/shinshin86/mini-speech-bubbler/releases).



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

