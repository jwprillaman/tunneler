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

Tunneler requires one of set or unset modes to be set.


#### Flags

##### set  -s  
sets the config to be the contents on resolv file

##### unset -u
return config to its original contents before set

##### config -c
set the config file to set resolv from. Defaults to the first file found in base directory with .conf suffix.

##### resolv -r
set the resolv file to be set or unset. Defaults to /etc/resolv.conf





