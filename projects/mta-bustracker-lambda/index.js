'use strict';

const http = require('http');
var AWS = require('aws-sdk');

// Get reference to S3 client
var s3 = new AWS.S3();

/**
 * 
 * @param [Object] event
 *  => [string] data
 *    example: "http://api.prod.obanyc.com/api/siri/vehicle-monitoring.json?key=<MTA_API_KEY&LineRef=MTA%20NYCT_M5"
 *  => [string] bucket_path
 *    example: mta-test-bus-bucket/m5-interval-check
 */
exports.handler = (event, _, callback) => {
    const req = http.request(event.data, (res) => {
        let body = '';
        // NOTE: res.statusCode may be a helpful initial check
        res.setEncoding('utf8');
        res.on('data', (chunk) => body += chunk);
        res.on('end', () => {
            if (res.headers['content-type'] === 'application/json') {
                body = JSON.parse(body);
            } else {
              throw Error(`Invalid content-type returned. Received '${res.headers['content-type']}'. Expected 'application/json'`);
            }
            var s = new Date().getTime() + ".json"
            s3.putObject({
                Bucket: event.bucket_path,
                Key: s,
                Body: body
            }).promise()
              .then(() => callback( null, 'Object uploaded to S3' ) )
              .catch(e => {
                console.error( 'ERROR', e );
                callback( e );
              });
            callback(null, body);
        });
    });
    req.end();
    req.on('error', callback);
};
