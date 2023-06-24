**Demand** is a tool to check if executables are in path and in the correct version.

Examples:

```
$ demand examples/jq.json examples/un*
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
          "\u003e= 1.6"
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
unobtainium
9
```

Show detailed results of only incompatabilities:

```
$ ./bin/demand  -o examples/*; echo $?    
{
  "ok": false,
  "executable": "unobtainium"
}
9