# Introduction

`awscnfm` is a conformance test tool designed to define conformity of Giant
Swarm tenant clusters managed on AWS. Check the recording in which we show how
[the project structure] looks like and [how code generation works]:
https://drive.google.com/file/d/1qGGoTOkTOW0pt4boPlOqeS-TG9n39yaP/view?usp=sharing.



# How to use awscnfm?

As one starting point you want to become familiar with [the project structure].
`awscnfm` can be used as standalone tool or in combination with other tools. You
can run a complete test plan or only execute single actions against an existing
tenant cluster.



# What is a "cluster scope"?

A cluster scope represents a very specific tenant cluster configuration, e.g.
`awscnfm cl001`. Here `cl001` is a cluster scope defining what kind of
conformity we want to ensure against which kind of tenant cluster separation.
One cluster scope might want to test Single Master clusters, where another
cluster scope might want to test a HA Masters cluster. A cluster scope defines a
list of actions you can execute against a tenant cluster aligning with the
defininition of the cluster scope. Note that you cannot expect an action of one
cluster scope to successfully run against a tenant cluster of another cluster
scope. Note that most code for a cluster scope is generated. You can find more
information about this in our docs explaining [how code generation works].

```
$ awscnfm cl001
Conformance tests for cluster cl001.

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



# What is an "action"

An action is simply a declarative piece of behaviour bound to a specific cluster
scope, e.g. `awscnfm cl001 ac005`. An action can implement behaviour to setup
some test, e.g. cluster creation or cluster deletion. An action can implement
the testing behaviour itself, e.g. checking how many k8s nodes are ready. You
decide what an action does and how actions are designed per cluster scope. By
convention `ac000` is the test plan executing all actions of a cluster scope. By
convention `ac001` is the action creating a cluster based on a custom
definition. By convention, every action implements the same two interfaces
according to their intentional use case. These two interfaces are the `Executer`
and the `Explainer`. One executes the business logic, the other explains it.
Note that most code for an action is generated. You can find more information
about this in our docs explaining [how code generation works].

```
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



# What is a "test plan"?

A test plan by convention is defined by action `ac000`. This special action
within a cluster scope defines the list of actions being executed sequentially.
Executing the test plan without any errors means that the tenant cluster defined
for the executed cluster scope is conform to our currently implemented
definition. Note that a test plan takes hours to execute due to the nature of
infrastructure and several cluster transitions we want to go through.

```
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



# How to write new tests?

New tests are implemented as actions. Here the design is important. As soon as
you notice you do two different things in one action you should ask yourself if
you should go for another design. The idea of actions is to have a composable
list of independent steps that work together when being executed after one
another. Let's say you want to check for ready k8s nodes in your tenant cluster.
You would then implement an action to check for master nodes and another
separate action to check for worker nodes. This separation makes it easier for
debugging later in case something is wrong with the tenant cluster and we have
to find out why either master or worker nodes are missing. Separation of concern
is king. Consider how we implement the action `ac001` to test cluster creation.
Initializing the cluster transition of creating a tenant cluster is as simple as
creating Custom Resources against a Kubernetes API. This action usually succeeds
within the blink of an eye and this is everything a single action should then
do. Initialize behaviour or wait for a transition or check information, but not
either of those things mixed together. For the example of cluster creation we
wait for API availability separately in action `ac002`. If in doubt, when
designing tests consult your favourite rubber duck and other team members.



[the project structure]: structure.md
[how code generation works]: generation.md
