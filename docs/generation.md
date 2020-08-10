# Generation



### Generated Code

New cluster and action subcommands can be generated so that developers only have
to write the business logic they actually care about. In order to generate a new
action subcommand for an existing cluster definition you need to execute two
commands. The first command generates the code for the action. The second
command generates the code for the cluster. This recursive procedure ensures
everything is wired and ready for use. Further see the resulting diff below.

```
awscnfm generate action -c cl001 -a ac003
```

```
awscnfm generate cluster -c cl001
```

```
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
like shown below.

```
awscnfm generate action -c cl001 -a all
```



### Business Logic

All you as developer now have to do is to implement your business logic
according to the provided `Executor` and `Explainer` interfaces. Based on the
example above, any changes made in the files below will not be overwritten on
consecutive code generation attempts.

```
pkg/action/cl001/ac003/executor.go
pkg/action/cl001/ac003/explainer.go
```
