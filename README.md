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
4. `docker-compose up` (use `docker-compose up -d` to run as a daemon)

# Ports

These are the ports each front-end is hosted on by default.

- Bibliogram (Instagram): 10407
- Invidious (Youtube): 3000
- Nitter (Twitter): 8082
- Scribe (Medium): 8081
- Teddit (Reddit): 8080

To visit, go to localhost:PORT on the machine running the docker containers, or IP:PORT if connecting from another machine on the same network, where IP is the local IP of the host machine.

