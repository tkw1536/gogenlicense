module github.com/tkw1536/gogenlicense

go 1.20

require (
	github.com/golang/glog v1.1.1
	github.com/google/go-licenses v1.2.0
	golang.org/x/mod v0.12.0
)

// github.com/google/licenseclassifier is extremely fiddly
// and newer versions have a bug which prevents using it in module mode
// as such it is stuck on the old version for now

require (
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/licenseclassifier v0.0.0-20210722185704-3043a050f148 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20190725054713-01f96b0aa0cd // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/src-d/gcfg v1.4.0 // indirect
	github.com/xanzy/ssh-agent v0.2.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/src-d/go-billy.v4 v4.3.2 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)
