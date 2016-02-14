# AWS EC2 Lister

This is a utility program designed to transform json output from `aws ec2 describe-instances`
and show very brief listing information about the ec2/ecs instances.

## How to build a binary for my platform?

Checkout the syntax required for `go build`.

An example build for Darwin would look like this:

   ```bash
   GOOS=darwin; GOARCH=386; go build -o darwin-ec2-lister
   ```

## How do I deploy/administrate it?

Once the binary is built, I upload the binary up to an S3 bucket.

    aws s3 cp darwin-ec2-lister s3://my-executables/

>>Note: Both arguments should match the installer file.

Next, I give users the `install-example.sh` script, which they copy to their
machine, rename it if they want to, give it execute permission, and just run it.

>>Note: It assumes these users are authorized to access the associate AWS resources.

To make life easier...I setup an alias for this in my `.bash_profile` file:

    alias aws-instance-list='. ~/bin/aws-instance-lister.sh'

On any updates to the lister...all the user has to run is and they get the latest:

    /path/to/lister-script.sh --latest

>>Note: You could implement a nicer versioning scheme/strategy instead.

## How to contribute?

While this wasn't really designed for a general public audience, if there's a way it can be achieved, then please feel free to submit a pull request for review!

## License

MIT
