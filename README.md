# stormcli

Is a tool for manage [Stormkit.io](https://stormkit.io) via CLI.

## install

To install, just install go on you machine. Then clone the repository.

```
# ssh
git clone git@github.com:giuliobosco/stormcli.git
# or http
git clone https://github.com/giuliobosco/stormcli.git
```

And just install it or build it

```
# install
go install .
# or build
go build .
```

## commands

### apps

list 

```
$ stormcli apps
github/stormkit-dev/sample-project
github/giuliobosco/web-project
$ stormcli apps -numbers -details # stormcli apps -nd
0 | github/stormkit-dev/sample-project
        Status: true
        AutoDeploy: commit
        DefaultEnv: production
        Endpoint:
        DisplayName:
        CreatedAt: 2020-03-17 10:54:20 +0100 CET
        DeployedAt: 2020-06-07 06:38:46 +0200 CEST

1 | github/giuliobosco/todoWEB
        Status: true
        AutoDeploy: commit
        DefaultEnv: production
        Endpoint:
        DisplayName:
        CreatedAt: 2020-03-17 10:54:20 +0100 CET
        DeployedAt: 2020-06-07 06:38:46 +0200 CEST

```


## configuration

File `~/stormcli.yaml`:

```
app:
  bearer_token: "token" #bearer authentication token
```
