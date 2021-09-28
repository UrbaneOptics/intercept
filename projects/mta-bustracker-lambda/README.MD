# MTA Bustracker Lambda
This project was originially created to track minute-by-minute locations for a particular MTA bus line by leveraging MTA's [GTFS-RT SIRI API](https://api.prod.obanyc.com/wiki/Developers/GTFSRt)

While this current project would require a new lambda for each new bus line you'd like to track, it would also be possible to create a list of lines you'd like feeds for that you can iterate through in a single lambda.

# Setup
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- Create an S3 bucket for writing API responses to
- Upload the `mta-bustracker-lambda` to AWS Lambda

# Architecture
