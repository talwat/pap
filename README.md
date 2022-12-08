# pap

A simplistic **pa**permc hel**p**er.

## Examples

### Download the latest papermc jarfile

`pap download`

### Sign the EULA

`pap sign`

### Generate a script to run the jarfile

`pap run --jar server.jar`

## Why though?

pap was created to simplify the annoying parts of making a minecraft server, so you can do everything from an easy CLI.

And one of the most annoying parts, is downloading the latest papermc jarfile, which is the root purpose of pap.

pap looks at the papermc api, and gets the latest version automatically, without you needing to navigate to the papermc website and copy a link and paste it into a `wget` command.

## Install

### Source

#### Dependencies

* [Go](https://go.dev/) 1.17 or later
* [Git](https://git-scm.com/)

#### Unix-like

Just clone and compile pap

```sh
git clone https://github.com/talwat/pap
cd pap
go build .
```

and then move it into your binary directory,

##### System wide

```sh
sudo mv pap /usr/local/bin
```

##### Local

```sh
mv ~/.local/bin
```

#### Windows

While pap will probably work on windows (yet to be tested),
you will need to figure out how to compile and install pap yourself if you choose to do it by source.

pap is written in golang, so if you know how to compile a go program on windows you should be fine.
