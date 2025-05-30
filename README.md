<p align="center">
  <a href="https://www.abs-lang.org/">
    <img alt="abs language logo" src="https://github.com/abs-lang/abs/blob/master/bin/ABS.png?raw=true" width="310">
  </a>
</p>

<p align="center">
  <a href="https://github.com/abs-lang/abs">
    <img alt="GitHub Workflow Status (branch)" src="https://img.shields.io/github/actions/workflow/status/abs-lang/abs/tests.yml?branch=master">
  </a>
  <a href="https://github.com/abs-lang/abs">
    <img alt="License" src="https://img.shields.io/github/license/abs-lang/abs.svg">
  </a>
  <a href="https://github.com/abs-lang/abs">
    <img alt="Version" src="https://img.shields.io/github/release-pre/abs-lang/abs.svg">
  </a>
  <img alt="undefined" src="https://img.shields.io/github/release-date/abs-lang/abs.svg?style=flat">
  <img alt="undefined" src="https://img.shields.io/github/downloads/abs-lang/abs/total.svg?style=flat">
  <br />
  <img alt="undefined" src="https://img.shields.io/badge/platform-linux | macosx | windows-red.svg">
  <img alt="undefined"  src="https://img.shields.io/github/last-commit/abs-lang/abs.svg?style=flat">
  <a href='https://coveralls.io/github/abs-lang/abs'><img src='https://coveralls.io/repos/github/abs-lang/abs/badge.svg' alt='Coverage Status' /></a>
  <br />
  <img alt="undefined" src="https://img.shields.io/github/contributors/abs-lang/abs.svg?style=flat">
  <img alt="undefined" src="https://img.shields.io/github/issues/abs-lang/abs.svg?style=flat">
  <img alt="undefined" src="https://img.shields.io/github/issues-pr-closed/abs-lang/abs.svg?style=flat">
  <img alt="undefined" src="https://img.shields.io/github/stars/abs-lang/abs.svg?style=social">

</p>

# The ABS programming language

ABS is a programming language that works best when you're scripting on
your terminal. It tries to combine the elegance of languages
such as Python, or Ruby with the convenience of Bash.

``` bash
tz = `cat /etc/timezone`
continent, city = tz.split("/")

echo("Best city in the world?")

selection = stdin()

if selection == city {
  echo("You might be biased...")
}
```

See it in action:

![sample](https://raw.githubusercontent.com/abs-lang/abs/refs/heads/master/docs/vhs/images/readme-sample.gif)

Let's try to fetch our IP address and print the sum of its
parts if it's higher than 100. Here's how you do it
in Bash:

``` bash
# Simple program that fetches your IP and sums it up
RES=$(curl -s 'https://api.ipify.org?format=json' || "ERR")

if [ "$RES" = "ERR" ]; then
    echo "An error occurred"
    exit 1
fi

IP=$(echo $RES | jq -r ".ip")
IFS=. read first second third fourth <<EOF
${IP##*-}
EOF

total=$((first + second + third + fourth))
if [ $total -gt 100 ]; then
    echo "The sum of [$IP] is a large number, $total."
fi
```

And here's the same code in ABS:

``` bash
# Simple program that fetches your IP and sums it up
res = `curl -s 'https://api.ipify.org?format=json'`

if !res.ok {
  echo("An error occurred: %s", res)
  exit(1)
}

ip = res.json().ip
total = ip.split(".").map(int).sum()
if total > 100 {
    echo("The sum of [$ip] is a large number, $total.")
}
```

Wondering how to run this code? Grab the latest
[release](https://github.com/abs-lang/abs/releases) and run:

```
$ abs script.abs
```

## Installation

The easiest way is to:

``` bash
bash <(curl https://www.abs-lang.org/installer.sh)
```

then you can:

``` bash
$ abs path/to/scripts.abs
```

## Documentation

Visit [abs-lang.org](https://www.abs-lang.org) or check the [examples directory](https://github.com/abs-lang/abs/tree/master/examples),
which contains a few short scripts.

## Contributing

Want to hack on ABS locally? The recommended development
environment is inside a Docker container — simply:

* `make build` (builds the container)
* `make run` (sends you inside the development container)
* `make test` (runs the abs tests)

After you make any change, run `make test` and check
if any errors pop up. If everything looks fine that means
you're ready to [open a pull request](https://github.com/abs-lang/abs/pulls)!

## Status

ABS is fresh and under active development; exciting
things happen on a weekly basis!

Have a look at the roadmap [here](https://github.com/abs-lang/abs/milestones):
to know of what version we're currently working on take a look at [abs-lang.org/misc/technical-details](https://www.abs-lang.org/misc/technical-details).
