# Go Splash

Go-Splash is tool which provides fast and simple interactions with Splash.

Splash is a javascript rendering service which runs on docker and uses an HTTP API. 

## Getting Started

You will need to have both Docker and Splash installed and running. 

[How to Install Docker](https://docs.docker.com/install/)

[How to Install Splash](https://splash.readthedocs.io/en/stable/install.html)

### Installation

Go-Splash interacts with an API, and as such there are many options. The default options chosen are sensible options, however to fully realise the power of Go-Splash you may want to [build from source](https://github.com/sp1ash/go-splash#build-from-source) 

If you have Go 1.9 or later installed and all paths configured, you can install Go-Splash with go-get:

`go get -u github.com/sp1at/go-splash

You can also download the latest binary [here](https://github.com/sp1at/releases)


### Configuration

Go-Splash looks for the config file `~/.config/go-splash/config.yaml`


#### Config Options
These options are useful if you inted to run Splash on a droplet. 
``` config.yaml
SplashHost: "localhost"
SplashPort: "8050"
```

| **Note:** If the config file is not found, the default values will be used.

### Basic Usage

`go-splash https://google.com`

This will create new directories in the current working directory.

`./out/google.com`

Inside this directory will be four files:

* google.com.har: A full [Har Report](https://en.wikipedia.org/wiki/HAR_(file_format))
* google.com.html: Full HTML output
* google.com.jpg: A full page screenshot (Default width: 1200px, Default time to wait: 3s)
* google.com.json: All previous files saved in one convenient json file. All images are encoded using Base64.

### Built-in Help
To see the flags and arguements:

`go-splash -help`

```
Usage:
  go-splash [flags] url

Flags:
	--har        Outputs Har file
	--html       Outputs HTML file
	--jpg, jpeg  Outputs JPG file
	--json       Outputs JSON file
	--png        Outputs PNG file

```

### Building From Source

This tool is written in Go you need to install the Go language/compiler/etc. Full details of installation and set up can be found on the [Golang website](https://golang.org/doc/install).

Compiling
go-splash uses external dependencies, and so you will need to pull them in first:

`go get && go build`

Once the binary is built, you can install it into your go path using:

`go install`

### Contributions

This tool is my first attempt at tool-making as well as in GoLang. Any contributions or feedback will be highly appricated. 

### Thanks

Thanks for trying go-splash!

 I would love to know if you find the tool useful or if there is anything you would like it to do. You can find me at:


**email:** `sp1at <at> sp1at <dot> com`

**twitter:** `@sp_l_at`

---
| The [Golang Logo and Mascot](https://golang.org/doc/faq#gopher) were created by Renée French and are licensed under [CC3](https://creativecommons.org/licenses/by/3.0/legalcode) and were adapted for this project. This project is in no way endorced or affiliated with the licensor.