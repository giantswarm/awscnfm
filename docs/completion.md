# Completion

If you are using Oh My ZSH you need to put the generated completion into the
plugins folder.

```
mkdir -p ~/.oh-my-zsh/plugins/awscnfm && awscnfm completion zsh > ~/.oh-my-zsh/plugins/awscnfm/_awscnfm
```

Then add `awscnfm` to the list of plugins that should be loaded when sourcing
`~/.zshrc` in the oh-my-zsh config.

```
plugins=(... awscnfm)
```
