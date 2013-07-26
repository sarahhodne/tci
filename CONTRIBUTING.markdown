# Contributing to tci

**First**: If you're unsure of afraid of *anything*, just ask or submit the
issue or pull request anyways. You won't be yelled at for giving your best
effort. Theworst that can happen is that you'll be politely asked to change
something. We appreciate any sort of contributions, and don't want a wall of
rules to get in the way of that.

However, for those individuals who want a bit more guidance on the best way to
contribute to the project, read on. This document will cover what we're looking
for. By addressing all the points we're looking for, it raises the chances we
can quickly merge or address your contributions.

## Issues

### Reporting an Issue

- Please test against the latest master commit. It is possible we already fixed
  the bug you're experiencing.
- Provide a reproducible test case. If a contributor can't reproduce an issue,
  then it dramatically lowers the chances it'll get fixed. And in some cases,
  the issue will eventually be closed.
- Respond promptly to any questions made to your issue. Stale issues will be
  closed.

## Setting up Go to work on tci

If you have never worked with Go before, you will have to complete the following
steps in order to be able to compile and test tci.

1. Install Go. On a Mac, you can `brew install go`.
2. Set and export the `GOPATH` environment variable. For example, you can add
   `export GOPATH=$HOME/Documents/goland` to your `.bash_profile`.
3. Download the tci source (and its dependencies) by running `go get
   github.com/henrikhodne/tci`. This will download the tci source to
   `$GOPATH/src/github.com/henrikhodne/tci`.
4. Make your changes to the tci source. You can run `go run tci.go [arguments]`
   from the main source directory to run your version of tci.
5. Test your changes by running `make test`.
6. If everything works well, and the tests pass, run `make format` before
   submitting a pull request.
