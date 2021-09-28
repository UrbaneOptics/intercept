# MTA Bustracker Lambda
This project was originially created to track minute-by-minute locations for busses on a particular MTA bus line by leveraging MTA's [GTFS-RT SIRI API](https://api.prod.obanyc.com/wiki/Developers/GTFSRt)

While this current project would require a new lambda for each new bus line you'd like to track, it would also be possible to create a list of lines you'd like feeds for that you can iterate through in a single lambda.

# Setup
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- Create an S3 bucket for writing API responses to
- Upload the `mta-bustracker-lambda` to AWS Lambda


## Deploying
1. Run `sam build --profile <AWS_PROFILE>` to build the lambda locally
2. Run `sam deploy --guided --profile <AWS_PROFILE` to begin deploying the lambda
3. Once the lambda is deployed, navigate to the Cloudwatch / Eventbridge event that will trigger it every minute(should be found in [Amazon EventBridge -> Events -> Rules](https://console.aws.amazon.com/events/home?region=us-east-1#/rules)) and click 'edit'
4. Scroll down to the 'Select Targets' section and expand the 'Configure Input' section
5. Select 'Constant (JSON Text) and enter the following info:
```json
{ "data": "http://api.prod.obanyc.com/api/siri/vehicle-monitoring.json?key=<MTA_API_KEY>&LineRef=<MTA_LINE_REFERENCE>",
  "bucket_path": "<YOUR_BUCKET>/<FOLDER>"
}
```
