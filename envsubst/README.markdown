# envsubst cloud builder

## envsubst cloud builder
This builder can be used to pre-process files for environment variables using `envsubst`.

### Building this builder
To build this builder, run the following command in this directory.
```sh
$ gcloud builds submit --config=cloudbuild.yaml
```

## Using this builder

Assuming you have the file `planetary-message.txt` you wish to pre-process in your build:
```
This is a text message from planet ${PLANET}.
```

Use the following step to do it:
```yaml
- id: preprocess-resources
  name: gcr.io/${PROJECT_ID}/envsubst
  env: ["PLANET=Mars"]
  args: ["message.txt"]
```

This builder can also accept wildcards! Lets say you have another file called `info.txt`:
```
The planet ${PLANET} is the next one on the solar system!
```

You can pass a wildcard, like so:
```yaml
- id: preprocess-resources
  name: gcr.io/${PROJECT_ID}/envsubst
  env: ["PLANET=Mars"]
  args: ["*.txt"]
```

This builder can also read environment variables from a `.env` file. Let's assume you have another file called `test.env`:

```dotenv
PLANET=Mars
```

You can instruct the builder to use this file to get environment variables from, like so:
```yaml
- id: preprocess-resources
  name: gcr.io/${PROJECT_ID}/envsubst
  args: ["-e", "test.env", "*.txt"]
```
