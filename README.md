# grasshopper

Grasshopper is an implementation of the [Nulecule Specification](http://www.projectatomic.io/docs/nulecule/). It can be used to bootstrap container applications and to install and run them. Grasshopper is designed to be run in a container context.

## Getting Started

Grasshopper is packaged as a container. End-users typically do not install the software from source. Instead use the Grasshopper container as the `FROM` line in a Dockerfile and package your application on top. For example:

```
FROM goern/grasshopper:0.0.4

MAINTAINER Your Name <you@example.com>

ADD /Nulecule /Dockerfile README.md /grasshopping/
ADD /artifacts /grasshopping/artifacts
```


## Developers

[![Build Status](https://travis-ci.org/goern/grasshopper.svg?branch=master)](https://travis-ci.org/goern/grasshopper)
[![Coverage Status](https://coveralls.io/repos/goern/grasshopper/badge.svg?branch=master&service=github)](https://coveralls.io/github/goern/grasshopper?branch=master)

First of all, clone the github repository: `git clone https://github.com/goern/grasshopper`.

### Build

To build a Grasshopper base image just run `make image`, this will compile the
Grasshopper binary and generate a Docker container image containing it.


## Providers

Providers represent various deployment targets.

## Dependencies

Please have a look at Godeps/Godeps.json to figure out the dependencies.

##Communication channels

* IRC: #nulecule (On Freenode)
* Mailing List: N/A

# Copyright

Copyright (C) 2015 Red Hat Inc.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

The GNU Lesser General Public License is provided within the file lgpl-3.0.txt.
