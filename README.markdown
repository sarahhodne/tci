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

### encrypt

```
$ tci encrypt FOO=bar
secure: "Tv7eONEJGIB/Mr563jIO06ffTWXccZ5KrEBuxGAysWnPR9hHX3rXCUSFTdAEZZ5IszG2aMAuDmYQPK/lS9gnjzKlBEZtcJ3hps/+RXmB/ooMYa4+RUiuRKZNeNFYKlOrMnPueXVXaW+9typQ0JAZQxf8B9K9tzH4Hq5CUQclED0="
```

### show

```
$ tci show
Build #6: Update API tests
State          passed
Type           push
Branch         master
Compare URL    https://github.com/henrikhodne/tci/compare/dc9524eb855d...d9bc220705b4
Duration       3m0s
Started        2013-07-26 15:48:56
Finished       2013-07-26 15:50:40
```

## License

tci is licensed under the MIT license. See the LICENSE file.
