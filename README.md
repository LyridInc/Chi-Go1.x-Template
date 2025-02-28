# Lyrid Golang 1.x Chi Template

## Run locally with:
```
go get
go run ./main.go
```

Open http://localhost:3000

## Edit the names (optional):
Open .lyrid-definition and change the App and Module name, because this will override another applications with the same name in the platform.

## Then submit to Lyrid Platform:

```
lc code submit
```
Wait until the cloud platform to finish with the build and the default deployment.

## Start hacking:

Edit the route url, settings, and views at /entry folder with your custom APIs.

Add more middlewares or your business logic in there.

[Deploy to Lyrid!!](http://localhost:3000/login?one-click-deploy=true&origin=github&repository-url=https://github.com/LyridInc/Chi-Go1.x-Template.git&env=testing&project-type=GoLang&repo-name=Go-Chi)
