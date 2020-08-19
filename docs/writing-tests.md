# Writing new tests

Here you learn how to add tests for awscnfm, and what to consider when adding tests.

## Design advice

New tests are implemented as actions. As soon as
you notice that you do two different things in one action, you should ask yourself if
this could be broken up into separate actions.

The idea of actions is to have a composable
list of independent steps that work together when being executed after one
another.

Let's say, for example, you want to check for K8s nodes in state `Ready` in your tenant cluster.
You would then implement an action to check for master nodes and another
separate action to check for worker nodes. This separation makes it easier for
debugging later in case something is wrong with the tenant cluster and we have
to find out why either master or worker nodes are missing. Separation of concern
is king.

To take another example, consider how we implement the action `ac001` to test cluster creation.
Initializing the cluster transition of creating a tenant cluster is as simple as
creating Custom Resources against a Kubernetes API. This action usually succeeds
within the blink of an eye and this is everything a single action should then
do.

Either initialize behavior _or_ wait for a transition _or_ check information, but don't do all of those things within one action.

For the example, after the cluster creation action `ac001`, we
wait for API availability separately in action `ac002`.

If in doubt, when
designing tests, consult your favorite rubber duck and other team members.

## Actions are CLI subcommands

Actions in `awscnfm` are also subcommands of the command line utility, and they exist as Go packages in the [`/cmd`](https://github.com/giantswarm/awscnfm/tree/master/cmd) directory. In that directory, the first sub level represents the cluster scopes. Within those are the actions.

However actions are also Go packages available for re-use in the [`/pkg`](https://github.com/giantswarm/awscnfm/tree/master/pkg) path. Again, a hierarchy is formed where the first level is the cluster scope. Within each cluster scopes package, the individual action packages can be found.

Found out more about this in the [repository structure](structure.md) documentation.

To create the necessary files in both the `cmd` and the `pkg` directories, the utility provides code generation. We strongly recommend using this to create new actions. Read on to learn more.

## Generating action code

In order to generate a new
action subcommand for an existing cluster scope, you need to execute two
commands.

The first command generates the code for the action:

```nohighlight
awscnfm generate action -c cl001 -a ac003
```

The second
command generates the code for the cluster scope.

```nohighlight
awscnfm generate cluster -c cl001
```

This recursive procedure may seem counter-intuitive, but it ensures
everything is wired and ready for use.

The resulting diff below gives you an idea of what gets generated.

```nohighlight
$ git status
On branch generate-pkg-action
Your branch is up-to-date with 'origin/generate-pkg-action'.

Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

    new file:   cmd/cl001/ac003/execute/zz_generated.command.go
    new file:   cmd/cl001/ac003/execute/zz_generated.error.go
    new file:   cmd/cl001/ac003/execute/zz_generated.flag.go
    new file:   cmd/cl001/ac003/execute/zz_generated.runner.go
    new file:   cmd/cl001/ac003/explain/zz_generated.command.go
    new file:   cmd/cl001/ac003/explain/zz_generated.error.go
    new file:   cmd/cl001/ac003/explain/zz_generated.flag.go
    new file:   cmd/cl001/ac003/explain/zz_generated.runner.go
    new file:   cmd/cl001/ac003/zz_generated.command.go
    new file:   cmd/cl001/ac003/zz_generated.error.go
    new file:   cmd/cl001/ac003/zz_generated.flag.go
    new file:   cmd/cl001/ac003/zz_generated.runner.go
    modified:   cmd/cl001/zz_generated.command.go
    new file:   pkg/action/cl001/ac003/executor.go
    new file:   pkg/action/cl001/ac003/explainer.go
    new file:   pkg/action/cl001/ac003/zz_generated.error.go
    new file:   pkg/action/cl001/ac003/zz_generated.executor.go
    new file:   pkg/action/cl001/ac003/zz_generated.explainer.go
```

In some cases it might be useful to re-generate all actions, which you can do
like shown here:

```nohighlight
awscnfm generate action -c cl001 -a all
```

### Adding business Logic

All you as developer now have to do is to implement your business logic
according to the provided `Executor` and `Explainer` interfaces.

Based on the
example above, any changes made in the files below will **not** be overwritten on
consecutive code generation attempts.

```nohighlight
pkg/action/cl001/ac003/executor.go
pkg/action/cl001/ac003/explainer.go
```
