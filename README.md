# falsename

### a simple cross-shell command aliaser

A CLI utility to make cross-shell command aliases, and access them from a simple TUI.

<img width="397" alt="Screenshot 2023-07-11 at 5 06 16 PM" src="https://github.com/blobbybilb/falsename/assets/58201828/60134ce0-b157-4b5a-b70c-05e82ee4ae4f">


## Status
- **V1:** Update to V2 (shouldn't break anything).
- **V2:** Adds custom config locations (for syncing, etc.), and passes arguments to command when running an alias by name. *Current version, complete.*

It's feature complete as far as I had planned, there might be a V3 if I think a new feature is necessary (feel free to suggest!).

## Download

Pre-built binaries for Windows, macOS, Linux, and FreeBSD are available in the build folder in the repository.

```sh
# macOS - Apple Silicon
curl -L https://github.com/blobbybilb/falsename/raw/main/build/mac-arm64/fn -O
# macOS - Intel
curl -L https://github.com/blobbybilb/falsename/raw/main/build/mac-amd64/fn -O

# Linux - amd64
curl -L https://github.com/blobbybilb/falsename/raw/main/build/linux-amd64/fn -O
# Linux - arm64
curl -L https://github.com/blobbybilb/falsename/raw/main/build/linux-arm64/fn -O

# FreeBSD - amd64
curl -L https://github.com/blobbybilb/falsename/raw/main/build/freebsd-amd64/fn -O
# FreeBSD - arm64
curl -L https://github.com/blobbybilb/falsename/raw/main/build/freebsd-arm64/fn -O

# Windows - amd64
curl -L https://github.com/blobbybilb/falsename/raw/main/build/windows-amd64/fn.exe -O
```

Then, make it executable and move it to a directory in your PATH.

```sh
# non-Windows (macOS, Linux, FreeBSD)
chmod +x fn
sudo mv fn /usr/local/bin
```

Finally, set your config directory (optional, default is `~/.config/falsename`) (useful if you have a directory synced across computers or the like).

```sh
fn config ~/new/config/directory
```

## Usage:

- `fn` -> choose an alias from a list of all aliases
- `fn <alias> [args]` -> run an alias with optional arguments
- `fn list` -> list all aliases
- `fn get <alias>` -> get the command for an alias
- `fn save <alias> <command>` -> save an alias
- `fn shell` -> get configured shell (default /bin/sh)
- `fn shell <shell>` -> set shell (not recommended unless you don't have /bin/sh)
- `fn delete <alias>` -> delete an alias
- `fn config` -> get the config directory
- `fn config <dir>` -> set the config directory

### Example:

```sh
fn save hi "echo hello"
fn hi # -> hello
fn # choose from a list of all aliases
fn list # list all aliases
fn get hi # get the command for an alias
fn delete hi # delete an alias
```

##### License: GNU GPLv3

*"run(ning) an alias" doesn't sound right...*