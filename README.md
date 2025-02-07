# Healthcare

## Usage
### Stateless components
Stateless components are components that can report their health on demand, for example: 
storage component that can ping database anytime.
```go
collector := func() (ComponentHealth, error) {
    return ComponentHealth{
        Status: StatusHealthy,
    }, nil
}

healthcare := New()
healthcare.AddStateless("sample", collector)

health := healthcare.Collect()
```

The response will be:

```json
{
    "sample": {
        "status": "HEALTHY"
    }
}
```

...or you can add a message to describe the reason of the status:

```go
collector := func() (healthcare.ComponentHealth, error) {
    return healthcare.ComponentHealth{
        Status: healthcare.StatusUnhealthy,
        Message: "connection closed"
    }, nil
}

healthcare := healthcare.New()
err := healthcare.AddStateless("sample", collector)
if err != nil {
    return fmt.Errorf(err)
}

health := healthcare.Collect()
```

The response will be:

```json
{
    "sample": {
        "status": "UNHEALTHY",
        "message": "connection closed"
    }
}
```
### Stateful components
Stateful components are components that can't report their status on demand, for example: external API clients,
which cannot call the external service on demand. In that way, you can interpret their last communication attempt
as you wish and report the status of the service, healthcare will store it for you.

```go
healthcare := healthcare.New()
reciever, err := healthcare.AddStateful("sample")
if err != nil {
    return fmt.Errorf(err)
}

reciever.SetHealthy("")

health := healthcare.Collect()
```

The response will be:

```json
{
    "sample": {
        "status": "HEALTHY"
    }
}
```

### Combining different components
You can combine different types of components in one healthcare instance.

```go
collector := func() (healthcare.ComponentHealth, error) {
    return healthcare.ComponentHealth{
        Status: healthcare.StatusUnhealthy,
        Message: "connection closed"
    }, nil
}

healthcare := healthcare.New()

reciever, err := healthcare.AddStateful("sampleStateful")
if err != nil {
    return fmt.Errorf(err)
}
err = healthcare.AddStateless("sampleStateless", collector)
reciever.SetHealthy("")

healthcare.Collect()
```

Response: 
```json
{
    "sampleStateless": {
        "status": "UNHEALTHY",
        "message": "connection closed"
    },
    "sampleStateful": {
        "status": "HEALTHY"
    }
}
```

### Get specific component status
```go
collector := func() (healthcare.ComponentHealth, error) {
    return healthcare.ComponentHealth{
        Status: healthcare.StatusUnhealthy,
        Message: "connection closed"
    }, nil
}

healthcare := healthcare.New()

err = healthcare.AddStateless("sampleStateless", collector)

healthcare.CollectSpecific("SampleStateless")
```
