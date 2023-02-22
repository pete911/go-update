# go-update
Simple tool to update go git projects. Run the tool one level up from the projects, for example if your project
structure looks like this:

```
github/
  my-org/
    project1/
    project2/
    projectX/
```

The tool needs to run inside `my-org` directory.

- build `go build`
- run `./go-update`

```
[INFO]  21:24:27 go project /abc/def/pete911/gh
[INFO]  21:24:27   git pull
[INFO]  21:24:29   Already up to date.

go.mod go 1.19 is different from current 1.20 version, update go.mod (y/N): y
[INFO]  21:24:31   go get -u -t ./...
[INFO]  21:24:32   go mod tidy
[INFO]  21:24:32   go mod vendor
gh project has changes, push to git (y/N): y
commit message (update dependencies):
[INFO]  21:24:36   git add .
[INFO]  21:24:36   git commit -m update dependencies
[INFO]  21:24:36   [main 581975d] update dependencies
 19 files changed, 425 insertions(+), 174 deletions(-)
 create mode 100644 vendor/golang.org/x/tools/internal/tokeninternal/tokeninternal.go

[INFO]  21:24:36   git push
```
