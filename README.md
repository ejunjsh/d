# d

[![Build Status](https://travis-ci.org/ejunjsh/d.svg?branch=master)](https://travis-ci.org/ejunjsh/d)

my docker container practice


## precondition

you need to install `docker`  and `iptables` commandline tools

enable ip forwarding 

    sysctl net.ipv4.conf.all.forwarding=1

## install

    go get github.com/ejunjsh/d/cmds/d
   
## usage

    NAME:
       d - a simple container runtime implementation.
             The purpose of this project is to learn how docker works and how to write a docker by ourselves
             Enjoy it, just for fun.
    
    USAGE:
       d [global options] command [command options] [arguments...]
    
    VERSION:
       0.0.0
    
    COMMANDS:
         init     Init container process run user's process in container. Do not call it outside
         run      Create a container with namespace and cgroups limit ie: d run -ti [image] [command]
         ps       list all the containers
         exec     exec a command into container
         rm       remove unused containers
         network  container network commands
         install  install an image into d
         help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version


## example

notice: this example is run at `Ubuntu 18.04.2 LTS` and the `d` is tested in this environment only. 

I'm not sure if `d` is compatible all the linux distribution.

### create nginx container

    # docker create nginx
    Unable to find image 'nginx:latest' locally
    latest: Pulling from library/nginx
    f7e2b70d04ae: Pull complete 
    08dd01e3f3ac: Pull complete 
    d9ef3a1eb792: Pull complete 
    Digest: sha256:98efe605f61725fd817ea69521b0eeb32bef007af0e3d0aeb6258c6e6fe7fc1a
    Status: Downloaded newer image for nginx:latest
    a46332a8ebfae2c83365affbdfe2cd285894c85563d6c898665816f3b7609b47
    
notice the last line of output, it is a container id
    
### generate tar archive 

    # docker export a46332a8ebfae2c83365affbdfe2cd285894c85563d6c898665816f3b7609b47 > nginx.tar
    
use the container id to generate the tar

### install tar archive into d
     
    # d install nginx.tar
    
### run a container with d

    # d run -i nginx bash
    missing mydocker_pid env skip nsenter
    {"level":"info","msg":"createTty true","time":"2019-03-13T17:39:00+08:00"}
    {"level":"info","msg":"dirs=/var/lib/d/c/1672401789/writeLayer:/var/lib/d/i/nginx","time":"2019-03-13T17:39:00+08:00"}
    missing mydocker_pid env skip nsenter
    {"level":"info","msg":"command all is bash","time":"2019-03-13T17:39:00+08:00"}
    {"level":"info","msg":"init come on","time":"2019-03-13T17:39:00+08:00"}
    {"level":"info","msg":"Current location is /var/lib/d/c/1672401789/mnt","time":"2019-03-13T17:39:00+08:00"}
    {"level":"info","msg":"Find path /bin/bash","time":"2019-03-13T17:39:00+08:00"}
    root@1672401789:/#
    
you will see it returns a bash and the host name is the container name `1672401789`
    
## reference

https://github.com/xianlubird/mydocker