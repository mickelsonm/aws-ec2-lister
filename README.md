# AWS EC2 Lister

This is a utility program designed to transform json output from `aws ec2 describe-instances` 
and show very brief information about the ec2 name (from tag), instance ID, private IP, and public IP.

## How to build a binary for my platform?

Checkout the syntax required for `go build`. 

An example build for Darwin would look like this:

   ```bash
   GOOS=darwin; GOARCH=386; go build -o lister-darwin
   ```

## How to contribute?

While this wasn't really designed for a general public audience, if there's a way it can be achieved, then please
feel free to submit a pull request for review!

## License

MIT
