# pap

[![codebeat badge](https://codebeat.co/badges/95ce3938-9084-418c-b8fe-8093f6292d28)](https://codebeat.co/projects/github-com-talwat-pap-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/talwat/pap)](https://goreportcard.com/report/github.com/talwat/pap)
[![AUR version](https://img.shields.io/aur/version/pap)](https://aur.archlinux.org/packages/pap)
[![License](https://img.shields.io/github/license/talwat/pap)](https://github.com/talwat/pap/blob/main/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/talwat/pap)

A swiss army knife for minecraft server development.

## Table of contents

- [pap](#pap)
  - [Table of contents](#table-of-contents)
  - [Examples](#examples)
    - [Download the latest papermc jarfile](#download-the-latest-papermc-jarfile)
    - [Sign the EULA](#sign-the-eula)
    - [Generate a script to run the jarfile](#generate-a-script-to-run-the-jarfile)
    - [Turn off pvp](#turn-off-pvp)
    - [Install worldedit](#install-worldedit)
  - [Why though?](#why-though)
  - [Install](#install)
    - [Build Dependencies](#build-dependencies)
    - [Arch linux](#arch-linux)
    - [Unix](#unix)
      - [Unix - From Releases](#unix---from-releases)
        - [Unix - System wide from releases](#unix---system-wide-from-releases)
        - [Unix - Local from releases](#unix---local-from-releases)
      - [Unix - From Source](#unix---from-source)
        - [Unix - System wide from source](#unix---system-wide-from-source)
        - [Unix - Local from source](#unix---local-from-source)
    - [Windows](#windows)
      - [Windows - From Releases _(recommended)_](#windows---from-releases-recommended)
      - [Windows - From Source](#windows---from-source)
    - [Common issues](#common-issues)
      - [Local installation not found](#local-installation-not-found)
        - [Bash](#bash)
        - [Zsh](#zsh)
        - [Fish](#fish)
  - [Contributing](#contributing)
  - [Dependencies](#dependencies)

## Examples

### Download the latest papermc jarfile

`pap download paper`

### Sign the EULA

`pap sign`

### Generate a script to run the jarfile

`pap script --jar server.jar`

### Turn off pvp

`pap properties set pvp false`

### Install worldedit

`pap plugin install worldedit`

## Why though?

pap has a few purposes:

- To simplify some of the common tasks you need to do when creating or managing a server (such as when you download/update the server jar.)
- To easily and automatically verify the jars you download to avoid bad issues down the line.
- To provide an easy CLI to do common tasks like changing server.properties and signing EULA, for usage in scripts.
- To quickly install plugins directly from their sources.

## Install

### Build Dependencies

If you are obtaining pap from source, you will need these dependencies:

- [Go](https://go.dev/) 1.18 or later
- [Git](https://git-scm.com/)

### Arch linux

If you wish, pap can be installed from the AUR:

```sh
yay -S pap
```

### Unix

#### Unix - From Releases

You can go to the [latest release](https://github.com/talwat/pap/releases/latest)
and download the fitting binary for your system from there.

pap is available on most architectures and operating systems, so you will rarely need to compile it from source.

Simply mark the downloaded binary as executable and move it.

##### Unix - System wide from releases

```sh
sudo mv pap* /usr/local/bin/pap
sudo chmod +x /usr/local/bin/pap
```

##### Unix - Local from releases

```sh
mv pap* ~/.local/bin/pap
chmod +x ~/.local/bin/pap
```

#### Unix - From Source

First, clone pap:

```sh
git clone https://github.com/talwat/pap
cd pap
```

Switch to the latest tag _(optional)_:

```sh
git tag # get all tags
git checkout <tag>
```

And then build:

```sh
go build .
```

Finally, move it into your binary directory:

##### Unix - System wide from source

```sh
sudo mv pap /usr/local/bin/pap
sudo chmod +x /usr/local/bin/pap
```

##### Unix - Local from source

```sh
mv ~/.local/bin/pap
chmod +x ~/.local/bin/pap
```

### Windows

pap **does** work on windows, but the install steps listed are for unix-like systems.

#### Windows - From Releases _(recommended)_

If you want to download from releases, download the fitting windows exe and [put it into path](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows#:~:text=Go%20to%20%22My%20computer%20%2D%3E,exe%20's%20directory%20into%20path.).

#### Windows - From Source

First, clone pap:

```sh
git clone https://github.com/talwat/pap
cd pap
```

Switch to the latest tag _(optional)_:

```sh
git tag # get all tags
git checkout <tag>
```

And then build:

```sh
go build .
```

Finally, [put it into path](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows#:~:text=Go%20to%20%22My%20computer%20%2D%3E,exe%20's%20directory%20into%20path.).

### Common issues

#### Local installation not found

Usually this is because `~/.local/bin` is not in PATH.

You can add `~/.local/bin` to PATH through your shell:

##### Bash

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
```

##### Zsh

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
```

##### Fish

Look at the [fish docs](https://fishshell.com/docs/current/tutorial.html#path) for more detailed instructions.

```sh
fish_add_path $HOME/.local/bin
```

## Contributing

Anyone is welcome to contribute, and if someone can port pap to various package managers, it would be greatly appreciated.

If you want more info about how to contribute, take a look at [CONTRIBUTING.md](CONTRIBUTING.md).

If you would like to add a plugin to the repository, take a look at [PLUGINS.md](PLUGINS.md).

If you like pap, feel free to [star it on github](https://github.com/talwat/pap), or [vote for it on the AUR](https://aur.archlinux.org/packages/pap).

## Dependencies

- [schollz/progressbar](https://github.com/schollz/progressbar)
- [urfave/cli](https://github.com/urfave/cli)
