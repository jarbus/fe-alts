# fe-alts
Opt out of Twitter, Medium, Invidious, Reddit, and Instagram front-ends with a single Docker Compose. Designed by [nfusionz](github.com/nfusionz).

# Overview

This repository is designed to make it dead-simple to host alternative front ends to some of the biggest websites on the internet. With a single `docker-compose up -d`, it's now possible to launch the following five services at once:

 - [teddit (Reddit)](https://codeberg.org/teddit/teddit)
 - [Bibliogram (Instagram)](https://sr.ht/~cadence/bibliogram/)
 - [Scribe (Medium)](https://sr.ht/~edwardloveall/Scribe/)
 - [Invidious (Invidious)](https://github.com/iv-org/invidious)
 - [Nitter (Nitter)](https://github.com/zedeus/nitter)

# Installation

1. Follow installation steps for [Docker-Compose](https://docs.docker.com/compose/install/), if not already installed.
2. `git clone https://github.com/jarbus/fe-alts.git`
3. `cd fe-alts`
4. `docker-compose up` (use `docker-compose up -d` to run as a daemon). You may need to run this multiple times if there are errors building the containers.

# Ports

These are the ports each front-end is hosted on by default.

- Bibliogram (Instagram): 10407
- Invidious (Youtube): 3000
- Nitter (Twitter): 8082
- Scribe (Medium): 8081
- Teddit (Reddit): 8080

To visit, go to localhost:PORT on the machine running the docker containers, or IP:PORT if connecting from another machine on the same network, where IP is the local IP of the host machine.

# Example Reverse Proxy Configuration

If your host machine is running a reverse proxy, such as Nginx, and the router forwards all traffic on port 443 to the host, then you only need to add one server block to use a custom domain for each service.

For example, `/etc/nginx/conf.d/teddit.conf`[:

```
server {
    listen 443 ssl;
    server_name teddit.yourdomain.com;

    location / {
        proxy_pass http://127.0.0.1:8080;
    }
}
```
Note that you will need to use something like [Certbot](https://certbot.eff.org/) to generate certificates for each sub-domain.
