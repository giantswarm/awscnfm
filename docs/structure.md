# Structure



### Code Layout

The tool should be structured so that it can execute any action against our
Tenant Clusters. There might be many cluster definitions for which we have many
actions each. The command line tool layout looks like shown below. In the
schematic example `cl001` is the cluster tested for conformity based on a
specific definition. All actions within that subcommand structure are
exclusively designed for this particular cluster layout, so it can be specific
to a single feature we want to test, e.g. single master clusters. In the example
below `ac001` would provide subcommands to execute or explain the action itself.

```
├── cmd
│   ├── cl001
│       └── ac001
│           ├── execute
│           └── explain
└── pkg
    └── action
        └── ac001
```

```
awscnfm cl001 ac001 execute
awscnfm cl001 ac001 explain
awscnfm cl001 ac002 ...
```



### Execute Tests

```
awscnfm cl001 ac001 execute
```



### Explain Tests

```
awscnfm cl001 ac001 explain
```



### Tool Version

```
awscnfm version
```
