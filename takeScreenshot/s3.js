'use strict';

import fs from 'fs';
import util from 'util';
import AWS from 'aws-sdk';
import dotenv from 'dotenv';

dotenv.config();

const ID = process.env.S3_ID;
const SECRET = process.env.S3_SECRET;
const REGION = 'ap-northeast-2';
const BUCKET_NAME = process.env.BUCKET_NAME;

const s3 = new AWS.S3({
  accessKeyId: ID,
  secretAccessKey: SECRET,
  region: REGION
});

const uploadFile = async(file_name) => {
  // Read content from the file
  const file_content = fs.readFileSync(util.format(file_name));

  // Setting up S3 upload parameters
  const params = {
      Bucket: BUCKET_NAME,
      Key: util.format(file_name), // File name you want to save as in S3
      Body: file_content
  };

  // Uploading files to the bucket
  s3.upload(params, function(err, data) {
      if (err) {
          throw err;
      }
      console.log(`File uploaded successfully. ${data.Location}`);
  });
};

export default uploadFile;