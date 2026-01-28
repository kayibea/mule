# Mule

A small CLI clipboard for files and directories.

## Usage

``` sh
mule add FILES...
mule list
mule copy
mule move
mule clear
```

## Shell aliases

Recommended shortcuts:
``` sh
alias mla='mule add'
alias mlc='mule copy'
alias mlv='mule move'
alias mll='mule list'
alias mlp='mule prune'
```

## Examples

Collect files from other tools, then act on them later.
``` sh
find src -type f -name '*.go' | xargs mule add
cd ~/somewhere/at/my/home
mule copy
```

Add modified files from git:
``` sh
git ls-files -m | xargs mule add
```

Copies or moves the stored files into the current directory.
