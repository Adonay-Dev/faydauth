# Faydauth

[![Go Reference](https://pkg.go.dev/badge/github.com/Adonay-Dev/faydauth.svg)](https://pkg.go.dev/github.com/Adonay-Dev/faydauth)  
[![Go Report Card](https://goreportcard.com/badge/github.com/Adonay-Dev/faydauth)](https://goreportcard.com/report/github.com/Adonay-Dev/faydauth)  
[![Build Status](https://github.com/Adonay-Dev/faydauth/actions/workflows/ci.yml/badge.svg)](https://github.com/Adonay-Dev/faydauth/actions/workflows/ci.yml)

`faydauth` is a **Go library** to integrate **Fayda Digital ID** authentication, SSO, and token management (JWT + refresh tokens). Supports **Redis** or **in-memory caching** for storing tokens.  

---

## Features

- Authenticate users via **Fayda API** (VeriFayda / eSignet)  
- Generate **JWT access tokens**  
- Generate **refresh tokens** with TTL  
- Validate and **revoke refresh tokens**  
- Support for **Redis** or **in-memory cache**  
- Easy integration into your Go services  

---

## Installation

```bash
go get github.com/Adonay-Dev/faydauth
