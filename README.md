# mysql-test

This is a CLI for testing MySQL connectivity. It connects to the target host:port/schema with the supplied username/password and performs a ping. Exit status 0 means success. Any errors will be printed in stdout.

# Build
Run `make` at the root of this repo. It compiles a binary under `build/linux` targeting `linux/amd64`.

# Usage example

Requires `MYSQL_PASSWORD` environment variable to be defined.

```sh
mysql-test -u myuser -d mydbname -h myhostname -p 3306
```
