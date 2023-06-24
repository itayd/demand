**Demand** is a tool to check if executables are in path and in the correct version.

Examples:

```
$ demand examples/jq.json examples/unobtainium1.json
{
  "ok": true,
  "executable": "jq",
  "full_path": "/opt/homebrew/bin/jq",
  "checks": [
    {
      "ok": true,
      "args": [
        "--version"
      ],
      "capture": "1.6",
      "test": {
        "ok": true,
        "name": "semver",
        "args": [
          ">= 1.6"
        ]
      }
    }
  ]
}
{
  "ok": false,
  "executable": "unobtainium"
}
```

Fail if incompatible and list incompabilities:

```
$ ./bin/demand -q -l -f examples/*; echo $?
unobtainium1 awk
9
```

Show detailed results of only incompatabilities:

```
$ ./bin/demand -o examples/*; echo $?    
{
  "ok": false,
  "executable": "unobtainium1"
}
{
  "ok": false,
  "executable": "awk",
  "full_path": "/usr/bin/awk",
  "checks": [
    {
      "ok": false,
      "args": [
        "--version"
      ],
      "capture": "20200816",
      "test": {
        "ok": false,
        "name": "semver",
        "args": [
          "> 30000000"
        ]
      }
    }
  ]
}
0
```

One liner:

```
$ ./bin/demand -s "go version semver >=9.0.0" -qlf ; echo $?
go
9
```