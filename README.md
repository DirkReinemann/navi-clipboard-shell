# navi-clipboard-shell

Writes all arguments given starting from the third one to the clipboard.
Can be used as shell for the cheatsheet tool [navi](https://github.com/denisidoro/navi)
so that you have the command ready for use in any other tool.

This makes navi for me even more useful and valuable! Shout-outs to the navi developers!
Maybe I can bring this feature to the original tool once I'm more confident with Rust.
For now this is a quick and dirty alternative that works. So please enjoy ;-)!

## Installation

```shell
go build -ldflags "-w" -o ~/.local/bin/navi-clipboard-shell
```

## Usage

The second argument -c is unused and can be anything. You cannot leave it out so it must be be provided.
The program exits when the number of arguments is smaller than three including the program name itself.

In the following example **This is my command** is copied to the clipboard.

```shell
./navi-clipboard-shell -c This is my command
```

### navi 

Configure the navi-clipboard-shell as navi shell command.
The commands are now copied to the clipboard instead of executed or printed.
Make sure you include the installation directory in the path environment variable so that navi can access it.
You can also provide the absolute path to the binary when you don't want to modify your path.

**~/.zshrc** or **~/.bashrc**
```shell
export PATH="$PATH:$HOME/.local/bin"
```

**~/.config/navi/config.yaml**
```yaml
shell:
  command: navi-clipboard-shell
```
