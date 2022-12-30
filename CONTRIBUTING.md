# Contributing

Some of these rules will be caught automatically by golangci-lint, but others will not.

## Variable & Function names

Please use descriptive names, but not names that are extremely long.

Single letter variables are only okay if it's extremely obvious like `i` as the index, otherwise please use
longer names, abbreviations like `pkg` are okay.

## Functions

Try to seperate different parts of a function into their own seperate functions if they get too long or are utilized somewhere else.

Or, if you have several pieces of code that do a similar thing, put them into their own file.

For example, functions which get information from the PaperMC api are in `paper/api.go`.

## Adding support for other jarfile types

If you wish to do this, make sure you include a `GetURL` function, and beyond that it's up to you.

Please try and include a few unit tests aswell.

Additionally, you must get the jarfile and all information directly from **official sources**.

## Logging

Log major steps in an action. You can use the functions in the `log` package to do this.

It's a bad idea to have long periods of time with no logs, because this gives off the impression that nothing is happening,
and it's better to be transparent about what the program is doing when.

Logs should not end with a `.` and must always begin with `pap:` which is automatically done by the `log` package.

The only exeption is for `...`.

## Styling

Just use [gofumpt](https://github.com/mvdan/gofumpt).

## Commits

Please use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for all commits.
