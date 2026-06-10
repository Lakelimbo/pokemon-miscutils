# Pokémon Miscellaneous Utils

Just some handy utils for Pokémon statistical data (usually things that are annoying to calculate every time, or too expensive to do on the client-side).

## Running

```sh
miscutils -h
# NAME:
#    Pokémon Miscellaneous Utils - A bunch of random things for Pokémon data
# USAGE:
#    Pokémon Miscellaneous Utils [global options] [command [command options]]
# COMMANDS:
#    effectiveness, types, effect  calculates type effectiveness
#    help, h                       Shows a list of commands or help for one command
# GLOBAL OPTIONS:
#    --help, -h  show help
```

(You can also place `--help` on each command to show more information)

## Commands
### Type effectiveness
- Command: `effectiveness`
- Aliases: `effect`, `types`

Will calculate the Type Effectiveness (in a defensive manner) for one or two types.

```sh
miscutils effectiveness
# NAME:
#    Pokémon Miscellaneous Utils effectiveness - calculates type effectiveness
# USAGE:
#    Pokémon Miscellaneous Utils effectiveness [options] types [types ...]
# OPTIONS:
#    --ability string, -a string  possible ability that affects the type effectiveness
#    --help, -h                   show help
```

Example:
```sh
miscutils type Dragon
miscutils effect Dragon Dark
```

You can also optionally provide an [ability that affects damage output](https://bulbapedia.bulbagarden.net/wiki/Category:Abilities_that_alter_damage_taken) with the `--ability` flag (alias `-a`):

```sh
miscutils type Bug Dark --ability "Wonder Guard"

# You can also use a trimmed string if you prefer
miscutils type Normal Electric -a dryskin
```

## Build
Requirements:
- Go v1.26 or higher
- golangci-lint

To build:
```sh
make build

# or if you don't wnat to use Makefile:
CGO_ENABLED=0 go build -ldflags="-s -w"
```

For a smaller binary size, I recommend disabling the `s` and `w` flags, as well as CGO, even though this app is not using it. But there should be no breaking change if those are not set.

Before commiting any changes, run the linter with golangci-lint:
```sh
make lint
```
