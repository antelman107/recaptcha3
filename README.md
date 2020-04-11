# reCAPTCHA v3 verifier
Call `Verify` to perform verify request ( https://developers.google.com/recaptcha/docs/verify#api_request ).
## Basic usage 
```go
verifier := NewVerifier(&http.Client{})

resp, err := verifier.Verify(context.Background(), "secret", "token", "")
if err != nil {
    return err
} 
if !resp.Success {
    return errors.New("you are bot")
}

```

## Specify request timeout 
```go
verifier := NewVerifier(&http.Client{})

ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
defer cancelFunc()

resp, err := verifier.Verify(ctx, "secret", "token", "")
if err != nil {
    return err
} 
```