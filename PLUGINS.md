# Plugins

## Table of contents

- [Plugins](#plugins)
  - [Table of contents](#table-of-contents)
  - [A plugin is out of date](#a-plugin-is-out-of-date)
  - [A plugin isn't working](#a-plugin-isnt-working)
  - [Creating a plugin](#creating-a-plugin)
    - [Fields](#fields)
      - [`name`](#name)
      - [`version`](#version)
      - [`description`](#description)
      - [`license`](#license)
      - [`authors`](#authors)
      - [`site` _(optional)_](#site-optional)
      - [`dependencies` _(optional)_](#dependencies-optional)
      - [`optionalDependencies` _(optional)_](#optionaldependencies-optional)
      - [`downloads`](#downloads)
        - [Jenkins](#jenkins)
        - [URL](#url)
        - [Full example](#full-example)
      - [`install`](#install)
      - [`uninstall`](#uninstall)
      - [`note`](#note)
    - [Testing your plugin](#testing-your-plugin)

## A plugin is out of date

If you notice a plugin is out of date, usually you can update it by simply changing the `version` property to whatever the current version is.

## A plugin isn't working

If a plugin is broken, please create an [issue](https://github.com/talwat/pap/issues).

## Creating a plugin

Creating a plugin is very easy, and most of the information asked can be found in `plugin.yml`.

Every plugin has a json file in the `plugins` directory. This file tells pap how to install it, uninstall it, and so on.

You can see [`plugins/example.jsonc`](plugins/example.jsonc) for a commented example on how to create a plugin and each field needed.

Or, continue here for more detailed explanations.

### Fields

Each plugin has some fields that give metadata & important information, this section lists them.

#### `name`

The name of the plugin. This should be all lowercase, without spaces.

The name of the plugin should also be identical to the json file itself.

#### `version`

The version of the plugin. Use `latest` if you either:

- Use jenkins to automatically build and distribute your plugin.
- Have a static link to the latest version of your plugin that doesn't change.

If you don't meet ethier of those requirements, you can pick a version that matches your versioning scheme, and then access it in a URL by using `{version}`.

For example:

```json
{
  "version": "0.1"
}
```

#### `description`

A short description of your plugin, it should end with a period (`.`)

#### `license`

The license your plugin uses.

If you actually have a license, use it's [SPDX Identifier](https://spdx.org/licenses/).

If the plugin is proprietary, use `proprietary`.

If you have no idea what license it uses, just use `unknown`.

#### `authors`

A list of authors that created the plugin.

This can be usernames, real names, etc...

#### `site` _(optional)_

The site of the plugin, this is optional.

If you don't have a website, this can also be your repository.

You should include the protocol (usually `https://`)

For example:

```json
{
  "site": "https://www.example.com"
}
```

#### `dependencies` _(optional)_

The dependencies of your plugin. This is only optional if you don't have any.

If your plugin has a dependency that isn't in pap yet, you can either:

1. Implement that plugin yourself
2. Or open an [issue](https://github.com/talwat/pap/issues) if it's a very common dependency

This is a list of strings.

For example:

```json
{
  "dependencies": ["exampledependency"]
}
```

#### `optionalDependencies` _(optional)_

An optional dependency of your plugin. This enhances functionality or adds new features.

For example:

```json
{
  "optionalDependencies": ["vault"]
}
```

#### `downloads`

This is **the** most important part.

This is a list of downloads, so you can have multiple.

Downloads have two types: `jenkins` and `urls`.

As mentioned in the [`version`](#version) field, you can use jenkins of a fixed url.

This is defined in the `type` attribute, so it can be `jenkins` or `url`.

The `filename` attribute defines what name to save the downloaded file as, so it's predictable.

##### Jenkins

If you are using jenkins, you can define your job in the `job` property.

Additionally, you can select an artifact with the `artifact` property, which is a regex.

Please only use basic regex's, because more complex ones hurt compatibility.

For example:

```json
{
  "type": "jenkins",
  "job": "https://ci.athion.net/job/Example",
  "artifact": "jarfile-bukkit-v.*",
  "filename": "plugin.jar"
}
```

##### URL

If you are using the `url` method, just define the `url` property as your URL.

You can use `{version}` in the URL which will be substituted with whatever the `version` property is set to.

For example:

```json
{
  "type": "url",
  "url": "https://www.example.com/plugin-{version}.jar",
  "filename": "plugin.jar"
}
```

##### Full example

```json
{
  "downloads": [
    {
      "type": "jenkins",
      "job": "https://ci.athion.net/job/Example",
      "artifact": "jarfile-bukkit-.*",
      "filename": "plugin.jar"
    },
    {
      "type": "url",
      "url": "https://www.example.com/plugin-{version}.jar",
      "filename": "plugin.jar"
    }
  ]
}
```

#### `install`

If your plugin just downloads a jarfile like most, you can get away with setting the `type` attribute to `simple` which downloads the jarfile and exits.

For example:

```json
{
  "install": {
    "type": "simple"
  }
}
```

If you need to unzip a file or run some commands, you can use `complex` for the `type` attribute.

This allows you to define commands to run on windows and unix like operating systems.

On unix, `sh` is used for the shell. On windows, it's `powershell`.

For example (non functional):

```json
{
  "install": {
    "type": "complex",
    "commands": {
      "windows": ["move my_plugin/*.jar .", "rmdir my_plugin"],
      "unix": ["mv my_plugin/*.jar .", "rm -rf my_plugin"]
    }
  }
}
```

#### `uninstall`

How to uninstall your plugin.

You do this by defining some files/directories to delete.

Each file has a `path` which is relative to the `plugins` directory, and a `type` which can be `main`, `config`, or `data`. It can also be `other`.

For example:

```json
{
  "uninstall": {
    "files": [
      {
        "type": "main",
        "path": "myPlugin.jar"
      },
      {
        "type": "config",
        "path": "myPlugin"
      }
    ]
  }
}
```

#### `note`

The `note` attribute will be displayed at the end of the command, and is useful for displaying key information.

For example, if your plugin needs a specific property to be turned off/on, mention it here.

It is an array, and each item will be displayed on a seperate line.

Example:

```json
{
  "note": [
    "you need to disable pvp for this plugin to work correctly",
    "or else bad things will happen"
  ]
}
```

### Testing your plugin

Now that your plugin file is complete, it's time to test.

First, make a `test` directory so you don't accidentally put downloaded files in the wrong place.

This directory is on the `.gitignore` so you don't have to worry about any binary files being commited.

```sh
mkdir test
cd test
```

Then, you can run the plugin install command but with a path instead:

```sh
pap plugin install ../plugins/myplugin.json
```

And then create a PR.
