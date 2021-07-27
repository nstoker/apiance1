# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

[CVE-2020-26160](https://github.com/advisories/GHSA-w73w-5m7g-f7qc). jwt-go allows attackers to bypass intended access restrictions in situations with []string{} for m["aud"] (which is allowed by the specification). Because the type assertion fails, "" is the value of aud. This is a security problem if the JWT token is presented to a service that lacks its own audience check. There is no patch available and users of jwt-go are advised to migrate to [golang-jwt](https://github.com/golang-jwt/jwt) at version 3.2.1.

## 0.0.1 2019-12-02

All basic tests working. Several are tagged with 'not implemented', but don't appear to be getting called.

## 0.0.0 2019-11-29

.1 release fixes some stupidity at the keyboard end of the project.

Part way through implementing the database connections, but getting issues with `go get` and the path `api/utils`.

Initial project build.
