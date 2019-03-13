# d

[![Build Status](https://travis-ci.org/ejunjsh/d.svg?branch=master)](https://travis-ci.org/ejunjsh/d)

my docker container practice


# usage

    NAME:
       d - is a simple container runtime implementation.
             The purpose of this project is to learn how docker works and how to write a docker by ourselves
             Enjoy it, just for fun.
    
    USAGE:
       d [global options] command [command options] [arguments...]
    
    VERSION:
       0.0.0
    
    COMMANDS:
         init     Init container process run user's process in container. Do not call it outside
         run      Create a container with namespace and cgroups limit ie: mydocker run -ti [image] [command]
         ps       list all the containers
         exec     exec a command into container
         rm       remove unused containers
         network  container network commands
         install  install an image
         help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version
       
# reference

https://github.com/xianlubird/mydocker