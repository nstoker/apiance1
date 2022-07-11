# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## v0.0.4 2022-07-11

Update dependancies

## v0.0.3.1 2022-04-03

Further updates to dependancies.

## v0.0.3 2022-04-03

Updated further dependancies with `go list -u -m all | go get -u`. Added a variety of `.env` files

## 0.0.2 2022-04-02

Updated to go 1.18. Updated many dependancies.

Upgrade github.com/containerd/containerd to version 1.5.9 for [GHSA-mvff-h3cj-wj9c](https://github.com/advisories/GHSA-mvff-h3cj-wj9c).

Updated some packages.

Added alert if local `.env` missing (non-fatal).

## 0.0.1 2019-12-02

All basic tests working. Several are tagged with 'not implemented', but don't appear to be getting called.

## 0.0.0 2019-11-29

.1 release fixes some stupidity at the keyboard end of the project.

Part way through implementing the database connections, but getting issues with `go get` and the path `api/utils`.

Initial project build.
