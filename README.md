# pap

[![codebeat badge](https://codebeat.co/badges/95ce3938-9084-418c-b8fe-8093f6292d28)](https://codebeat.co/projects/github-com-talwat-pap-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/talwat/pap)](https://goreportcard.com/report/github.com/talwat/pap)
[![License](https://img.shields.io/github/license/talwat/pap)](https://github.com/talwat/pap/blob/main/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/talwat/pap)

[![Packaging status](https://repology.org/badge/vertical-allrepos/pap.svg)](https://repology.org/project/pap/versions)

A swiss army knife for minecraft servers.

## pap is close to 1.0 🎉

pap is now feature complete (for now) and just needs testing & code reviewing.

If you want, try installing pap and messing around with it.

If you actually manage to break it, [open an issue](https://github.com/talwat/pap/issues).

Or, make a PR.

## Table of contents

- [pap](#pap)
  - [pap is close to 1.0 🎉](#pap-is-close-to-10-)
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
  - [Updating pap](#updating-pap)
  - [Uninstalling pap](#uninstalling-pap)
  - [Contributing](#contributing)
  - [Dependencies](#dependencies)
  - [Packaging](#packaging)

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
- [Make](https://en.wikipedia.org/wiki/Make_(software))

### Arch linux

> **Info**
> The AUR build might not have the latest version of pap, but it may be more stable.

If you wish, pap can be installed from the AUR:

```sh
yay -S pap
```

### Ubuntu

Install [Pacstall](https://github.com/pacstall/pacstall#installing), then run:
```bash
pacstall -I pap
```

### Unix

#### Unix - From Releases

You can go to the [latest release](https://github.com/talwat/pap/releases/latest)
and download the fitting binary for your system from there.

pap is available on most architectures and operating systems, so you will rarely need to compile it from source.

##### Unix - System wide from releases

```sh
sudo install -Dm755 pap* /usr/bin/pap
```

##### Unix - Local from releases

> **Warning**
> You may see an error that pap wasn't found, if you see this you may not have `~/.local/bin/` in your PATH.
> See [common issues](#local-installation-not-found) on how to add it.

```sh
install -Dm755 pap* ~/.local/bin/pap
```

#### Unix - From Source

> **Warning**
> `pap update` downloads and installs a binary, it does not compile it from source.
> If you need to compile pap from source, don't use `pap update`.

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
make
```

Finally, move it into your binary directory:

##### Unix - System wide from source

```sh
sudo make install PREFIX=/usr
```

##### Unix - Local from source

> **Warning**
> You may see an error that pap wasn't found, if you see this you may not have `~/.local/bin/` in your PATH.
> See [common issues](#local-installation-not-found) on how to add it.

```sh
make install
```

### Windows

pap **does** work on windows, but windows has a ~~bad~~ different way to CLI apps.

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
# Use `make` if you have it, otherwise:

mkdir -vp build
go build -o build
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

## Updating pap

> **Note**
> This will only work if you are running version 0.11.0 or higher. If not, just reinstall using the install guide.

If you used a release and followed the install guide, you should be able to simply run:

```sh
sudo pap update
```

or if you did a local install:

```sh
pap update
```

## Uninstalling pap

Simply delete the binary file you installed. pap does not create any files that you do not explicitly
tell it to.

So, if you did a system wide install do:

```sh
sudo rm /usr/bin/pap
```

or if you did a local install:

```sh
rm ~/.local/bin/pap
```

## Contributing

Anyone is welcome to contribute, and if someone can port pap to various package managers, it would be greatly appreciated.

If you want more info about how to contribute, take a look at [CONTRIBUTING.md](CONTRIBUTING.md).

If you would like to add a plugin to the repository, take a look at [PLUGINS.md](PLUGINS.md).

If you like pap, feel free to [star it on github](https://github.com/talwat/pap), or [vote for it on the AUR](https://aur.archlinux.org/packages/pap).

## Dependencies

- [schollz/progressbar](https://github.com/schollz/progressbar)
- [urfave/cli](https://github.com/urfave/cli)

## Packaging

pap is currently on one singular repository: The AUR.

If you would like to package & submit pap to a repository, [open an issue](https://github.com/talwat/pap/issues).

[![Packaging status](https://repology.org/badge/vertical-allrepos/pap.svg)](https://repology.org/project/pap/versions)
