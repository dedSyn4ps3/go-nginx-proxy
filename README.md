# Go-Nginx Proxy

<br>

![Banner](https://github.com/dedSyn4ps3/go-nginx-proxy/raw/main/backend/main/assets/img/landing-page.png "Banner") 

<br>

## Summary

<br>

> This repo is meant to hold the code files for the Medium article [Power to the Proxy](https://medium.com/@erutherford_nullreturn).

The overall structure of the project is as follows:

```
.
├── api
├── backend
│   └── main
│       └── assets
│           ├── css
│           ├── fonts
│           ├── img
│           └── js
├── devices
└── nginx
    └── sites-enabled

```

As far as resource directories go, the `assets` folder is pretty straightforward. It's parent directory, `backend/main` , houses the basic website template used to simulate a public-facing site used in a real deployment.

<br>

At the root of the project is a `docker-compose.yml` file that is used to deploy two different Docker containers:

- One for the website's contact form API
- One for a separate 'devices' API to simulate an additional back end service a user may wish to utilize without opening another public facing port

<br>

Finally, the `nginx` directory contains an example configuration file that shows how to go about declaring various endpoints for the Nginx service to proxy requests to. This is where we declare the route (and port number) of our different containers to allow web requests to be proxied properly.

<br>

## Credits
The skeleton used for the example website, including the CSS stylesheets, were created by [TemplateMo](https://templatemo.com/). Be sure to check out there many different HTML templates if you're ever in need of inspiration or template code!

<br>

## Licensing

- Copyright 2022 dedSyn4ps3 (github.com/dedsyn4ps3)
- Licensed under MIT
