# galaxy #

An visual simulation of gravity and celestial bodies

## Requirements

 - inotify-tools (optional)

## How to dev

Build the binary with

```shell
$ make build
```

This will produce a `galaxy` binary in the `/bin` folder

Use the project watcher to automatically rebuild the binary if a go file is modified

```shell
$ make watch
```

## How to run

```shell
$ make run

... or

$ /bin/galaxy
```
