# Dog
Dog is a simple CLI program written in Go. It can display text file content in diffrent colors and can count characters, letters, words and lines of those files.

If you want to install prebuilt binaries of Dog (linux only), follow the instructions provided on [dog website](https://tekronet.github.io/programy/dog)

## Compiling from source
If you are using macOS or Windows you need to compile Dog from source. Here is how to do it.

1. Download/clone this repository to your system
2. Download go (1.16 or newer)
3. Download dependencies by running ```go get github.com/fatih/color```
4. Compile by running ```go build src/main.go``` in project's root directory.
