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
[INFO]  18:23:19 go version 1.20
go.mod go version 1.19 is different from current 1.20 version, update go.mod (y/N): y
[INFO]  18:23:23 running go [get -u -t ./...] in /abc/pete911/arpf
[INFO]  18:23:23 running go [mod tidy] in /abc/pete911/arpf
[INFO]  18:23:23 running go [mod vendor] in /abc/pete911/arpf
/abc/pete911/arpf has changes, push to git (y/N): y
commit message (update dependencies):
[INFO]  18:23:52 running git [add .] in /abc/pete911/arpf
[INFO]  18:23:52 running git [commit -m update dependencies] in /abc/pete911/arpf
[INFO]  18:23:52 [main 193d6fb] update dependencies
 33 files changed, 286 insertions(+), 77 deletions(-)
[INFO]  18:23:52 running git [push] in /abc/pete911/arpf
```
