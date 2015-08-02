glot-cli
==========


## Overview
glot-cli is a command line utility for interacting with the glot API.
Supported actions include running code through the [run api](https://github.com/prasmussen/glot-run/tree/master/api_docs) and
publishing snippets to the [snippets api](https://github.com/prasmussen/glot-snippets/tree/master/api_docs).

## Installation
- Save the 'glot' binary to a location in your PATH (i.e. `/usr/local/bin/`)

### Downloads
- [glot-freebsd-386 v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnVFRpM184dmNxakU)
- [glot-freebsd-x64 v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnbkwyQXQ4bm9LU28)
- [glot-linux-386 v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnLUdDTHFkWHdPVU0)
- [glot-linux-arm v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnN1RzbmNqVVN6Nzg)
- [glot-linux-rpi v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6Embndl92SXlYYXBUc2s)
- [glot-linux-x64 v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbncmpnRUxKTEt0SGc)
- [glot-osx-386 v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnLTVFTFYwcXo0cVE)
- [glot-osx-x64 v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6Embnbm9JdENQSTVNdUU)
- [glot-windows-386.exe v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnbWRNSEtEbi0ya0k)
- [glot-windows-x64.exe v1.1.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnaVQ1RFNOekR3amM)

## Usage
    glot new <language>  (Create new snippet)
    glot run <path>  (Run code)
    glot run <path> --version <version>  (Run code code with a specific language version)
    glot publish --title <title>  (Publish snippet)
    glot languages  (List available languages available to run)
    glot versions <language>  (List available versions for a language)
    glot help  (Print help)
    glot --version  (Print application version)


## Examples
###### Run code
    $ echo 'print("foo " * 5)' > foo.py
    $ glot run foo.py
    foo foo foo foo foo

###### Publish snippet
    $ glot publish --title "Print foo"
    Publishing...
    Id: e4t1cn7jgt
    Title: Print foo
    Language: python
    Public: False
    Created: 2015-06-27T10:02:15Z
    Modified: 2015-06-27T10:02:15Z
    Owner: 427ca0e3-b254-4e97-9326-88c814758af5
    Files hash: 70c69c1dfc09a411cf990d0d5f812d6faaf09cc2
    Web Url: https://glot.io/snippets/e4t1cn7jgt
    Api Url: https://snippets.glot.io/snippets/e4t1cn7jgt
