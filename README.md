To pull this from the current private repo, add the following to `~/.gitconfig`

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

And add the following to your .bashrc

```
export GOPRIVATE="`go env GOPRIVATE`,github.com/koinos/koinos-types-golang"
```
