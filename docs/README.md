# Docs

This directory contains the openapi specification for the API. The specification is written in the OpenAPI 3.0 format.

## Directory Structure

- `api.yaml`: The main OpenAPI specification entry point.
- `cfg.yaml`: Configuration file for the OpenAPI specification.
- `generated.go`: Generated Go code from the OpenAPI specification.
- `tool.go`: A tool to generate Go code from the OpenAPI specification.

- `parameters/`: Contains the parameter definitions.
- `path/`: Contains the path definitions.
- `requestBody/`: Contains the request body definitions.
- `response/`: Contains the response definitions.
- `schema/`: Contains the schema definitions.

## Generating Go Code

To generate Go code from the OpenAPI specification, run the following command:

```bash
go generate docs/tool.go
```

or generate globally with

```bash
go generate ./...
```

### Implementation

generate code would create a `ServerInterface` in `generated.go` file. This interface is used to implement the API server.

Use `internal/node/component/aggregator/compoent.go` to implement the API server.

## Oapi-codegen extensions

Refer to [OpenAPI extensions](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#openapi-extensions)

Here are some usage.

- `x-go-type` to define custom type.
- `x-go-type-import` to import custom type, usually used with `x-go-type` to define custom type.
- `x-go-type-skip-optional-pointer` remove pointer of fileds in structs.
- `x-go-name` override the generated name of a field or a type.
- `x-go-type-name` override the generated name of a type

## Handle http endpoint

Without store `openapi.json` file, `generate.go` will generate `GetSwagger` function to get the swagger object.

```go
    apiServer.GET("/openapi.json", func(c echo.Context) error {
        swagger, err := docs.GetSwagger()
        swagger.Servers = append(swagger.Servers, &openapi3.Server{
            URL: config.Discovery.Server.Endpoint,
        })

        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        return c.JSON(http.StatusOK, swagger)
    })
```
