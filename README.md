# shortening
A basic url shortener built for speed. This application uses redis as a cache, then cockroachDB (postgres) as a backend if redis is stale. Port/Adapter architecture makes it so that the cache can be swapped out with other modules (eg. memcached for redis) if necessary for a specific application of the project. url shortening done with sha256, but more advanced hashing techniques can be swapped in through the port/adapter model. authentication is supported via openid connect for Google and Github respectively through an implementation of the authorization code flow. 

overview of server interfaces: https://www.figma.com/file/k2PleKToxuaKmg30cngTyz/system-architecture?node-id=0%3A1


to run, clone the repo, and run 
```bash
go run pkg/main.go
```
specific services need environment varibles set

eg. redis will need REDIS_ADDR and REDIS_PASSWORD

(readme WIP)

