# lambda-cli-handler

This is an example of how to call lambda functions locally, it uses a simple wrapper to redirect to the local
cli or lambda handler depending on the precense of the AWS_LAMBDA_FUNCTION_NAME environment variable.

## Cli Example

```
% go build -o main main.go
% ./main                  
test
got payload test
test
hello
got payload hello
hello
```

## Lambda

I didn't test this exact code in lambda, but I've used this for other things. It should work as expected.
