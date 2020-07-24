# tfto
cli tool to transform tf state file into static inventory file

## mvp
- create ansible inventory from tfstate file
- start with only one cloud provider
- first cloud provier: hetzner

## v2.0 :D
- support multiple cloud providers
- provider are recognized automatically
- e.g. build cluster from one DO instance, one hetzner instance, etc ...
- limit + handle stdin read

