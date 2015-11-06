# Insman
AWS instance manager.

# Configuring Credentials
From the official documentation:

> Before using the SDK, ensure that you've configured credentials. The best way to configure credentials on a development machine is to use the ~/.aws/credentials file, which might look like:

    [default]
    aws_access_key_id = AKID1234567890
    aws_secret_access_key = MY-SECRET-KEY
 
> Alternatively, you can set the following environment variables:

    AWS_ACCESS_KEY_ID=AKID1234567890
    AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY

Set your AWS region by setting the environment variable `AWS_REGION`:

    export AWS_REGION="us-west-2"

In order to send commands to your instances you need an SSH key, set the `AWS_KEY_FILE` with the absolute path of your key file:

    export AWS_KEY_FILE="/home/user/.ssh/myawskeyfile.pem"
