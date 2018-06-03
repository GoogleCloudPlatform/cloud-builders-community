# firebase

This build step invokes `firebase` commands that can be used in [Google Cloud Container Builder](cloud.google.com/container-builder/).

Arguments passed to this builder will be passed to `firebase` directly,
allowing callers to run [any firebase
command](https://docs.docker.com/compose/reference/overview/).

## Usage

**Get the firebase token**

This command will generate a new CI token that will be encrypted by the KMS to be used within the CLI

```
firebase login:ci
```

**Enable the KMS API**

Click "setup" or "enable API" on https://console.cloud.google.com/security/kms 

**Create the secret on GCP**

This step will encrypt the token via KMS. Remember to replace `GENERATED_TOKEN` in the text

```bash
# create a keyring for cloudbuilder-related keys
gcloud kms keyrings create cloudbuilder --location global

# create a key for the firebase token
gcloud kms keys create firebase-token --location global --keyring cloudbuilder --purpose encryption

# create the encrypted token
echo -n $TOKEN | gcloud kms encrypt \
  --plaintext-file=- \
  --ciphertext-file=- \
  --location=global \
  --keyring=cloudbuilder \
  --key=firebase-token | base64
```

**Use the encrypted key**

The encrypted key (output from previous command) can now simply be used within the cloudbuilder configuration file like so:

> Note that you need to specify `[PROJECT_ID]` directly instead of using `$PROJECT_ID` within secrets

```yaml
secrets:
- kmsKeyName: 'projects/[PROJECT_ID]/locations/global/keyRings/cloudbuilder/cryptoKeys/firebase-token'
  secretEnv:
    FIREBASE_TOKEN: '<YOUR_ENCRYPTED_TOKEN>'
```

**Add permission to the cloudbuilder**

- Open GCP IAM menu
- Find email ending with `@cloudbuild.gserviceaccount.com`
- Add `Cloud KMS CryptoKey Decrypter` role to this account

## Examples

See examples in the `examples` subdirectory.
