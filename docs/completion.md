# Completion

To make calling `awscnfm` actions easier and quicker, tab completion for shell
commands is supported.



## ZSH

If you are using Oh My ZSH you need to put the generated completion into the
plugins folder.

```nohighlight
mkdir -p ~/.oh-my-zsh/custom/plugins/awscnfm
awscnfm completion zsh > ~/.oh-my-zsh/custom/plugins/awscnfm/_awscnfm
```

Then add `awscnfm` to the list of plugins that should be loaded when sourcing
`~/.zshrc` in the oh-my-zsh config.

```nohighlight
plugins=(... awscnfm)
```
