# Study Blockchain

## Installation

### golang

`brew install go`

or

```
asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
asdf install golang x.xx.xx
```

### environment variables

(if you use [direnv](https://github.com/direnv/direnv))

- `brew install direnv`
- add hook into the shell.


finally, (see .envrc.sample)

| name | description | sample |
| --- | --- | --- |
| GOPATH | go workspace location | `(project-root)/go`

### check your Installation

```
cd go/src/hello
go build
./hello
# hello, world
```
