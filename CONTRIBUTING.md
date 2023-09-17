---
title: Contributing
---

Thanks for your interest in helping developping this app.

Here the guidelines and requirements to do it the best way :rocket:

# Project overview

This project is a `Go galaxy connector`, its role is to be the _connector_ with the ESN Galaxy in order to be able to connector and get the roles from it.

# Guidelines

If you want to create a new feature, please create an issue and/or send [me](https://github.com/mickahell) a message.

If you want to contribute in the framework, check the issue and/or send [me](https://github.com/mickahell) a message.

## How to start

The project if based on `Go=1.19` (currently the latest on date 2022/11). And its running using `Docker`, for testing developpment, it can also be run locally.

- **Build with Docker**

```bash
make docker-build
```

- **Run with Docker**

```bash
make docker-run
```

- **Run locally**

```bash
make start-app
```

You can now call your api using :
- CAS client : `http://127.0.0.1:8181`

## Configuration files

To run, the app need a configuration file, thoses files can be found in the [`test/` folder](test) and named [`conf_local.yaml`](test/conf_local.yaml) and [`conf_docker.yaml`](test/conf_docker.yaml). They are only for testing, **never for production !**

## Units tests

The goal is to have every part of the API tested `>=80%`, currently some modules doesn't have this minimun testing requirement and so have to be boost up.

To run the unit tests, locally you can run :

```bash
make run-test
```
