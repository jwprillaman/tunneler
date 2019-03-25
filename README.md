#### Tunneler
##### Creator : James Prillaman

automated tool for sercuring vpn access

### Quickstart

Set resolv to current config
```
$sudo tunneler -s
```

Sunset the resolv
```
$sudo tunneler -u
```

Tunneler has two modes used to set the configuration. If no configuration flag (-c) is set for a config location it will look in the current directory for a file this suffix .conf. The resolve defaults to /etc/resonv if none is set with the flag (-r).


#### Set
sets the config to be the contents on resolv file

flag -s


#### Unset
return config to its original contents before set

flag -u



