# MongoDB Replica Set with Kubernetes

## Clone
```bash
$ git clone git@github.com:devops-config/k8s-mongodb-replica-set.git
```

## Authenticate by using `key file` or `x.509 certificate`

### Keyfile Security

#### Prepare
Copy configurations to the root directory.
```bash
$ cp ./k8s-mongodb-replica-set/keyfile/* ./k8s-mongodb-replica-set/
```

#### Generate a key
```bash
$ openssl rand -base64 756

<< 'EXAMPLE-OUTPUT'
ON+rfCAVesB0+NaW/DaVO/OOzFXH4l6kKV/IQ7k2zwf1EQkWt/7xWJKJEVhbUhO/
rWlH/7n9JDK0Mh5AziuejWj70QDq5QCmb/T5/+ZUlhT5g5r9xIeKy2x3gFcfEbw0
TZ/wfb0U/o7+s/wgcOUJnuzVhxx65BREDfqYPnJB7rQTe+9o/xXI4yXddu/oTQtD
eIHJKUJc0KhxXymcf02zsqbw38n7z54DCL+J9i0okVE2MWbDEwYfIPHNk+U8fSR8
W3Asie3dWArAoC06y3jKRGcCqEZZk5RHmg1zCEpsok7sp92x+8Xa91RTpJXtzuv7
E08o5yZ6/aJriJnZRaeYky7QNb2z+u9xRMII1+T/ha/c91VRNCgLWNNAsYxp8JzC
vAjLSHWQccaBhZrkJfwyumAzHMQoQN4qmrLw36UdeN8no9ip+dsJZLRkAivdxvUI
4lUuVEOZsHyT9K1bSfOaGf+xkKLazLrdCyqzArAGzd7nHy0Ry9+8P86WMV6FbHbo
LzEXtWqa4Rav5Xh9nIUBcvoP34uVThAPeNyyqSbrujknnWeH8umtaDi/nc/zr9dX
zm12j6haJXeLootet8yS5bS89Utl6Li4Darj/mDIsSx3rQ7YSZEIC0LrDF6o44d4
8PtD/TsU+e1CQbZyvjtiNqoB6wSOj8wE8N9kSc8BepETkRlT2CmIvKfKhqSmSseQ
D4HaIthynSzL1TSGTzMZXd2CsMQBTyPv5k5n983JYm7VNipxNqj9Y1rmGIrPcRlI
kD+kfdzBykQcZrJLFmEzXu7ob0/tFMTTBrgJHpNs3k6o7bk6lyLBzvc9GK+cTqCM
bzOdqgshMZHTHJhhsj6JdujLObwZx8hURgmQCRlEG/FhV19nbwhgpoPE2+uy68J/
OHyKul4r6i9gnxQsxEVI3ktNGvbszvVxUc+n3hQlNLA8quEEbt6LMAUmm0Caz3UV
udpuXzsBozqP2CS+FtLjKeEW/W3q7m7mROG7nN8ah1dSZLqI
EXAMPLE-OUTPUT
```

Copy the above command's output then paste it into `data.keyFile` in the `4-mongo-security.yml` file.

#### Start replica set
```bash
$ kubectl apply -f ./k8s-mongodb-replica-set/
```

#### Connect to mongodb inside pod
```bash
$ kubectl exec -it mongo-0 -n mongo-ns -- bash

$ mongosh \
    --username ${MONGO_INITDB_ROOT_USERNAME} \
    --password ${MONGO_INITDB_ROOT_PASSWORD} \
    --authenticationDatabase 'admin'
```

### X.509 Certificate

#### Prepare
Copy configurations to the root directory.
```bash
$ cp ./k8s-mongodb-replica-set/x.509/* ./k8s-mongodb-replica-set/
```

#### Generate root CA
```bash
$ openssl req -sha256 -new -x509 -days 365 \
    -subj "/CN=mongo/O=WW" \
    -out root-ca.crt \
    -keyout root-ca.key
```

Note: Remember the password is required by the above command.

Copy the contents of `root-ca.crt` and `root-ca.key` then paste them into `data.root-ca.crt` and `data.root-ca.key` in the `4-mongo-security.yml` file respectively.

#### Create `.PEM` file for admin
```bash
$ openssl req -sha256 -new -nodes \
    -subj "/CN=admin" \
    -out admin.csr \
    -keyout admin.key

$ openssl x509 -req -sha256 -days 365 \
    -in  admin.csr \
    -CA root-ca.crt \
    -CAkey root-ca.key \
    -CAcreateserial \
    -out admin.crt

$ cat admin.crt admin.key > admin.pem
```
In the second command, openssl asks you for the password you just entered when creating the root CA.

Copy the content of `admin.pem` then paste it into `data.admin.pem` in the `4-mongo-security.yml` file.

#### Create `.PEM` file for specific user
```bash
$ openssl req -sha256 -new -nodes \
    -subj "/CN=ww_user" \
    -out ww_user.csr \
    -keyout ww_user.key

$ openssl x509 -req -sha256 -days 365 \
    -in  ww_user.csr \
    -CA root-ca.crt \
    -CAkey root-ca.key \
    -CAcreateserial \
    -out ww_user.crt

$ cat ww_user.crt ww_user.key > ww_user.pem
```

Note: If you change the `CN`'s value, you must also change `data.MONGO_USERNAME` in the file `2-mongo-env.yml` to the same.

Use this `.pem` file and `root-ca.crt` to connection to mongodb with specific roles set up by the admin.

#### Start replica set
```bash
$ kubectl apply -f ./k8s-mongodb-replica-set/
``` 

#### Connect to mongodb inside pod
```bash
$ kubectl exec -it mongo-0 -n mongo-ns -- bash

$ mongosh \
    --tls \
    --tlsCertificateKeyFile admin.pem \
    --tlsCAFile root-ca.crt \
    --authenticationDatabase '$external' \
    --authenticationMechanism MONGODB-X509 \
    mongo-0.mongo
```

## References:
- [https://www.spectrocloud.com/blog/kustomize-your-way-to-mongodb-replicaset/](https://www.spectrocloud.com/blog/kustomize-your-way-to-mongodb-replicaset/)
- [https://docs.mongodb.com/manual/tutorial/configure-x509-client-authentication/#std-label-x509-client-authentication](https://docs.mongodb.com/manual/tutorial/configure-x509-client-authentication/#std-label-x509-client-authentication)
- [https://goteleport.com/blog/securing-mongodb/](https://goteleport.com/blog/securing-mongodb/)
- [https://docs.mongodb.com/manual/tutorial/configure-x509-client-authentication/](https://docs.mongodb.com/manual/tutorial/configure-x509-client-authentication/)
