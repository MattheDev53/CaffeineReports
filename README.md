# CaffeineReports

Simple Go CLI for generating PDF Reports with Typst

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
Usage of CaffeineReports:
CaffeineReports help : Prints this help message
CaffeineReports generate [template] [data] [output] : generates a report from the provided data and template
CaffeineReports version : Prints Version Information
CaffeineReports credits : Prints Credits
```
