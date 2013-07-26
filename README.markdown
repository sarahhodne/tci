# tci

tci is a command line tool for interacting with Travis CI.

It's currently a work in progress. Things will probably break, and there's a lot
missing.

## Installation

Make sure you have Go installed, your `GOPATH` set up correctly, and your `PATH`
set up to point to your `GOPATH`'s bin folder.

```
$ go install github.com/henrikhodne/tci
```

## Usage

### show

```
$ tci --repo henrikhodne/tci show
Build #3: Get dependencies when building on Travis
State:		failed
```

## License

tci is licensed under the MIT license. See the LICENSE file.
