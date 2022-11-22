# Initial Setup

> GOPATH=$HOME
> git clone https://github.com/eliranwong/gobible
> cd gobible
> go mod init github.com/eliranwong/gobible
> go install
> export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))

# References

https://go.dev/talks/2014/organizeio.slide#1

https://go.dev/doc/code
