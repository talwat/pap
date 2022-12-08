# pap

A simplistic **pa**permc hel**p**er.

## Examples

### Download the latest papermc jarfile

`pap download`

### Sign the EULA

`pap sign`

### Generate a script to run the jarfile

`pap script --jar server.jar`

## Why though?

pap was created to simplify the annoying parts of making a minecraft server, so you can do everything from an easy CLI.

And one of the most annoying parts, is downloading the latest papermc jarfile, which is the root purpose of pap.

pap looks at the papermc api, and gets the latest version automatically, without you needing to navigate to the papermc website and copy a link and paste it into a `wget` command.

## Install

### Windows notice

pap **does** work on windows, but the install steps listed are for unix-like systems.

If you want to download from releases, download the fitting windows exe and [put it into path](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows#:~:text=Go%20to%20%22My%20computer%20%2D%3E,exe%20's%20directory%20into%20path.).

If you want to compile pap from source, you can run these commands:

```sh
git clone https://github.com/talwat/pap
cd pap
go build .
```

to clone and compile pap, and then [put it into path](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows#:~:text=Go%20to%20%22My%20computer%20%2D%3E,exe%20's%20directory%20into%20path.).

### macOS notice

If you are using macOS, please use the instructions for local installations.

### Local installation notice

You will need to add `~/.local/bin` to PATH like so if you use bash:

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
```

And like this if you use zsh:

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
```

### From releases

You can go to the [latest release](https://github.com/talwat/pap/releases/latest)
and download the fitting binary for your system from there.

pap is available on esentially all architectures and operating systems, so you will rarely need to compile it from source.

Simply mark the downloaded binary as executable and move it.

#### System wide from releases

```sh
sudo mv pap* /usr/local/bin/pap
sudo chmod +x /usr/local/bin/pap
```

#### Local from releases

```sh
mv pap* ~/.local/bin/pap
chmod +x ~/.local/bin/pap
```

### From Source

#### Build Dependencies

* [Go](https://go.dev/) 1.18 or later
* [Git](https://git-scm.com/)

#### Steps

Just clone and compile pap

```sh
git clone https://github.com/talwat/pap
cd pap
go build .
```

and then move it into your binary directory,

##### System wide from source

```sh
sudo mv pap /usr/local/bin/pap
sudo chmod +x /usr/local/bin/pap
```

##### Local from source

```sh
mv ~/.local/bin/pap
chmod +x ~/.local/bin/pap
```

## Contributing

Anyone is welcome to contribute, and if someone can port pap to various package managers, it would be greatly appreciated.

## Dependencies

* [schollz/progressbar](https://github.com/schollz/progressbar)
* [urfave/cli](https://github.com/urfave/cli)
