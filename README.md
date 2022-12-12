# Deploy Go Applications on Clever Cloud

This repository is a code example written in Go. You can clone it and deploy it on Clever Cloud following these steps.

See an app demo [here](https://gogogo.cleverapps.io).

_NB: alternatively, this can be a reference for deploying your own code on a remote server._
You'll need a [Clever Cloud](https://clever-cloud.com) account to deploy on the platform.

## Deploy this repository on Clever Cloud

This few steps will guide you into deploying this app within minutes. 

### Create a new app from this repository

1. Fork  or clone this repository
2. On Clever Cloud, go to **create > an app** choose your forked repo from GitHub or **new app** if you deploy using Git.
3. Choose **GO** and follow the next steps. You don't need any addon.

### Inject the environment variables

Environment variables are injected dynamically on Clever Cloud. To deploy a Go app, the runtime will need a few instructions to know how to deploy you app.

At the **Environment variables** step, add the variable name `CC_GO_BUILD_TOOL` with value `gobuild`.

### ... and just... deploy!

Wait until the deployment is completed if you are deploying from GitHub, or copy paste the command from the Console to push the code if you are suning Git.  

## A few notes on Go deployment on Clever Cloud

This repository has already been configured for deployment. Here's how, so you can fix issues on your own code.

### Port and network listening

Notice this part of the code in `main.go`:

```go
}

func main() {
 http.HandleFunc("/", indexPage)
 http.ListenAndServe("0.0.0.0:8080", nil)

}
 ```

This is how you configure your app to listen to the wild network `0.0.0.0`, not only `localhost` or `127.0.0.1`.

### Paths and commands

#### Go build

For projects that don't support `go modules`, `go build` will be used. To use it on Clever Cloud, you use `gobuild`.  We will then deploy your project in a classic `GOPATH` hierarchy using go build. You can use other Go commands using the `CC_GO_BUILD_TOOL` variable environment.

#### Main file

By default, Clever Cloud will look for a `main.go` file containing the code to run. If you want to name it differentely, you'll need to set the `CC_GO_PKG` environment variable with the main file of your application as value.

You can find more instructions on Go application configuration in [our Go documentation](https://www.clever-cloud.com/doc/deploy/application/golang/go/#configure-your-go-application).
