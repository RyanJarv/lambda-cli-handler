# lambda-cli-handler

This is an example of how to call lambda functions locally, it uses a simple wrapper to redirect to the local
cli or lambda handler depending on the precense of the AWS_LAMBDA_FUNCTION_NAME environment variable.

## Cli Example

```
% cat events.json | go run ./main.go
{"body":"hello world","code":200}
{"body":"","code":200}
```

## Input parsing

This uses the same payload parsing locally as it does in lambda, so Unmarshalling into the Input struct is handled for us.

## Lambda

I didn't test this exact code in lambda, but I've used this for other things. It should work as expected.
