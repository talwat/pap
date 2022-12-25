# Contributing

Some of these rules will be caught automatically by golangci-lint, but others will not.

## Variable & Function names

Please use descriptive names, but not names that are extremely long.

Single letter variables are only okay if the scope is small, or if it's obvious like `i` as the index.

## Functions

Try to seperate different parts of a function into their own seperate functions if they get too long or are utilized somewhere else.

Or, if you have several pieces of code that do a similar thing, put them into their own file.

For example, functions which get information from the PaperMC api are in `paper.go`.
