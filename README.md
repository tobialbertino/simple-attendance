# Attendance App

on this page:
- [dependencies](#dependencies)
- [Framework](#framework)
- [Architecture](#architecture-folder)
- [modules / internal description](#modules--internal-description)

# dependencies

using gofiber v2  
```
go get github.com/gofiber/fiber/v2
```

## Framework

- Framework : GoFiber
- Configuration : GoDotEnv

## Architecture

per-modules / internal:  
Delivery -> UseCase -> Repository

## modules / internal description:

- user : user management
- auth : auth management
- attendance : attendance management
