# Introduction

`awscnfm` is a conformance test tool designed to define conformity of Giant
Swarm tenant clusters managed on AWS.

Check the recording in which we show how
[the project structure] looks like and [how code generation works]:
https://drive.google.com/file/d/1qGGoTOkTOW0pt4boPlOqeS-TG9n39yaP/view?usp=sharing.

## What is a "cluster scope"?

A cluster scope represents a very specific **tenant cluster configuration**.

Each cluster scope has a unique name, e.g. `cl001`.

For example, one cluster scope might want to test single master clusters, where another
cluster scope might want to test a HA Masters cluster.

In addition, a cluster scope defines a
**list of actions** you can execute against a tenant cluster aligning with the
defininition of the cluster scope.

Note that you cannot expect an action of one
cluster scope to successfully run against a tenant cluster of another cluster
scope.

Also note that most code for a cluster scope is generated. You can find more
information about this in our docs explaining [how code generation works].

```nohighlight
$ awscnfm cl001
Conformance tests for cluster scope cl001.

Usage:
  awscnfm cl001 [flags]
  awscnfm cl001 [command]

Available Commands:
  ac000       Action ac000 for cluster 001.
  ac001       Action ac001 for cluster 001.
  ac002       Action ac002 for cluster 001.
  ac003       Action ac003 for cluster 001.
  ac004       Action ac004 for cluster 001.
  ac005       Action ac005 for cluster 001.

Flags:
  -h, --help   help for cl001

Use "awscnfm cl001 [command] --help" for more information about a command.
```

## Actions

An action is simply a declarative piece of behavior bound to a specific cluster
scope, e.g. `awscnfm cl001 ac005`.

An action can:

- implement behavior to setup some test, e.g. cluster creation or cluster deletion
- implement the testing behavior itself, e.g. checking how many k8s nodes are ready

You decide what an action does and how actions are designed per cluster scope.

By convention

- `ac000` is the test plan executing all actions of a cluster scope
- `ac001` is the action creating a cluster based on a custom
definition.
- every action implements the same two interfaces according to their intentional use case:
  - the `Executer` to executes the business logic
  - the `Explainer`to explain what the action does.

Note that most code for an action is generated. You can find more information
about this in our docs explaining [how code generation works].

Example output when calling an action in a cluster scope, not specifying the `execute` or `explain` subcommand:

```nohighlight
awscnfm cl001 ac005
Action ac005 for cluster 001.

Usage:
  awscnfm cl001 ac005 [flags]
  awscnfm cl001 ac005 [command]

Available Commands:
  execute     Execute action ac005 for cluster cl001.
  explain     Explain action ac005 for cluster cl001.

Flags:
  -h, --help   help for ac005

Use "awscnfm cl001 ac005 [command] --help" for more information about a command.
```

## Test plans

A test plan by convention is defined by action `ac000`. This special action
within a cluster scope defines the list of actions being executed sequentially.
Executing the test plan without any errors means that the tenant cluster defined
for the executed cluster scope is conform to our currently implemented
definition. Note that a test plan takes hours to execute due to the nature of
infrastructure and several cluster transitions we want to go through.

```nohighlight
$ awscnfm cl001 ac000 explain
Execute the conformance test plan of this cluster scope. Actions below are
executed in order. A tenant cluster is conform if the plan executes without
errors. Plan execution might take up to 1h54m30s.

ACTION  RETRY  WAIT     COMMENT
ac001   2s     10s      create cluster
ac002   3m0s   24m0s    check cluster access
ac003   2s     10s      check master count
ac004   2s     10s      check worker count
ac005   9m0s   1h30m0s  delete cluster
```


[the project structure]: structure.md
[how code generation works]: generation.md
