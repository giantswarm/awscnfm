# Writing Tests

Here you learn how to add tests for `awscnfm`, and what to consider when adding
tests.



## Design advice

New tests are implemented as actions. As soon as you notice that you do two
different things in one action, you should ask yourself if this could be broken
up into separate actions. The idea of actions is to have a composable list of
independent steps that work together when being executed after one another.

Let's say, for example, you want to check for K8s nodes in state `Ready` in your
tenant cluster. You would then implement an action to check for master nodes and
another separate action to check for worker nodes. This separation makes it
easier for debugging later in case something is wrong with the tenant cluster
and we have to find out why either master or worker nodes are missing.
Separation of concern is king.

To take another example, consider how we implement the action to test cluster
creation. Initializing the cluster transition of creating a Tenant Cluster is as
simple as creating Custom Resources against a Kubernetes API. This action
usually succeeds within the blink of an eye and this is everything a single
action should then do. Another action following up on the CR creation can then
check of the cluster got successfully created. Another indicator for separation
of concerns is the different lifecycles of these actions. Creating runtime
objects in Kubernetes is quick. Reconciling them often takes a lot of time. So
either initialize behavior _or_ wait for a transition _or_ check information,
but don't do all of those things within one action. If in doubt, when designing
tests, consult your favorite rubber duck and other team members.



## Actions

Actions in `awscnfm` are subcommands of the command line utility. They exist as
[go packages in the `cmd/action` directory]. Actions can be executed separately
like shown below.

```
awscnfm action verify cluster created
```

Adding new actions should simply follow the command structure of existing
actions where business logic is implemented in the command runner. When adding a
new action it can then also be added to an existing test plan. The design of
actions is as such that a single implementation of an action can be reused and
even be executed multiple times within a test plan.



## Plans

Plans in `awscnfm` are subcommands of the command line utility. They exist as
[go packages in the `cmd/plan` directory]. Plans can be executed separately.
Plans execute simply a list of actions.

```
awscnfm plan pl003
```

Adding new plans should simply follow the command structure of existing plans
where business logic is implemented in the command runner. Each plan defines a
list of steps which specifies which actions to run, in which order, and with
which backoff.



[go packages in the `cmd/action` directory]: https://github.com/giantswarm/awscnfm/tree/master/cmd/action)
[go packages in the `cmd/plan` directory]: https://github.com/giantswarm/awscnfm/tree/master/cmd/plan)
