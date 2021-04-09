ipinfo Terraform Provider
==================

Makes the ipinfo.io IP address API available to Terraform as a datasource.

_Why?_ I used this to create my own dynamic DNS solution with the Namechaep Terraform provider as a learning experiment.

Prerequisites
------------------

You need to sign up for an account and receive [an access token on ipinfo.io](https://ipinfo.io/account/token).

Use that token to either fill in the provider block or set the environment variable `IPINFO_TOKEN`.

## Requirements

* Terraform v0.12.X or greater
* Go 1.16.X or greater

Usage
---------------------

Then inside a Terraform file within your project (Ex. `providers.tf`):

Using the provider
----------------------

Make sure your API details are correct in the provider block.

```hcl
terraform {
  required_providers {
    ipinfo = {
      source  = "robgmills/ipinfo"
      version = "1.0.0"
    }
  }
}

provider "ipinfo" {
  api_token = "YoUrToKeNhErE" # Also set by the env variable `IPINFO_TOKEN`
}

# Look up and output your public IP info
data "ipinfo" "myinfo" {}

output "myinfohostname" {
  value = data.ipinfo.myinfo.hostname
}

output "myinfoip" {
  value = data.ipinfo.myinfo.ip
}

# Look up and output the IP info for Google's DNS
data "ipinfo" "googledns" {
    ip = "8.8.8.8"
}

output "googlednshostname" {
    value = data.ipinfo.googledns.hostname
}
```

Setup terraform and view the plan output.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.15+ is recommended). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules), added in Go 1.11.

To compile the provider, run `make testacc`. This will build and test the provider.

To contribute changes, please open a PR by forking the repository, adding the fork to your local copy of the git repository, create a branch, commit your changes, and open a PR:

```bash
$ git remote add fork git@github.com/youruser/terraform-provider-namechep
$ git checkout -b your-new-feature
$ git add .
$ git commit -m "Add a new feature"
$ git push -u fork your-new-feature
...
```