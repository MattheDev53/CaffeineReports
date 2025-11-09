# Caffeine Reports

Simple Go CLI for generating PDF Reports with Typst

## Prerequisites

This program assumes that the Typst executable is accessible from your system's PATH. Ensure that you have [Typst](https://typst.app/) installed on any machine that your project will be executed. You can install it by following [the instructions in the Typst repository](https://github.com/typst/typst?tab=readme-ov-file#installation).

## Building

### For Linux

As simple as installing Nix and running

```bash
nix build
```

### For Windows

I just did this & it worked.

```nushell
sh -c 'GOOS=windows GOARCH=amd64 go build'
```

### For any other platform

Good luck, make a Pull Request if you find out.

## How to Use

```
Usage of caffeine-reports:
caffeine-reports help : Prints this help message
caffeine-reports generate [template] [data] [output] : generates a report from the provided data and template
caffeine-reports version : Prints Version Information
caffeine-reports credits : Prints Credits
```
