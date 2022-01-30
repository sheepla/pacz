# pacz

**pacz** is an Arch Linux fuzzy searcher with fzf-like UI.

*This repository is still under development!*

## Installation

Clone this repository then run the following:

```
go install
```

## Usage

Specify a query as an argument.

```
pacz QUERY
```

To install the selected packages in bulk, run the following:

```bash
pacz QUERY | sudo pacman -S --needed -
```


