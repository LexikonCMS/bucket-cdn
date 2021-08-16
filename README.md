# bucketCDN
Lightweight reverse proxy to use a Google Cloud Storage bucket as a CDN backend. Intends to be an low-cost alternative to Cloud CDN while getting a project off the ground. As load later on increases, bucketCDN can be replaced with a [Cloud CDN](https://cloud.google.com/cdn/docs/setting-up-cdn-with-bucket) on the fly.

## Build image
> docker build -t lexikon/bucketcdn:latest .

## Start container
> docker run -e BUCKET_NAME=yourbucket -e HOST=cdn.example.com -p 8080:8080 lexikon/bucketcdn

## Configuration
| Env Variable  | Description                              |
|---------------|------------------------------------------|
| `BUCKET_NAME` | Name of the Google Cloud Storage bucket. |
| `HOST`        | Hostname on which the proxy listens.     |