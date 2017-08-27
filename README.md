# duango
A minecraft server implement in golang.


## Project Structure
```
/
  protocol/ - Minecraft related packets
  auth/
  world/
  entity/

  server.go - Minecraft server
  main.go - main
  rpc.go - expose rpc interface to manage server
  plugin.go - plugin subsystem

```

## Reference

* http://wiki.vg/Protocol
* http://wiki.vg/Protocol_version_numbers
* https://github.com/skirmish/netherrack
* https://github.com/cuberite/cuberite
