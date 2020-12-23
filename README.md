# gcpsecret

gcpsecret downloads a specified secret from Google Cloud Secret Manager.

I use this to install SSL certificates on various devices like Raspberry Pis.  I
use certbot to create a wildcard certificate on a single machine, and store that
certificate in Google Cloud Secret Manager.  Then I pull the key and certificate
down to the various devices.

This means the devices don't need to be publicly accessible, nor do they need to
have credentials that allow modifying DNS entries.  They only need a service
account credential which has read access to the specific secrets in Secret
Manager.

On the individual devices, I have a cronjob similar to:

```
0 2 * * * GOOGLE_APPLICATION_CREDENTIALS=CREDENTIAL_FILE gcpsecret -project PROJECT -secret CERT_SECRET > /tmp/cert && mv /tmp/cert /etc/ssl/certs/CERT_FILE
0 2 * * * GOOGLE_APPLICATION_CREDENTIALS=CREDENTIAL_FILE gcpsecret -project PROJECT -secret KEY_SECRET > /tmp/key && mv /tmp/key /etc/ssl/certs/KEY_FILE
```
