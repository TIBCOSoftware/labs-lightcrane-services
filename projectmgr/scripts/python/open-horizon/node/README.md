2. source SDO_DEVICE_STEPS

steven@steven-XS35:~$ source SDO_DEVICE_STEPS
Cleaning up any previous run...
(done)
 
Configuring environment to simulate the manufacture of an SDO device...
Your environment is configured (note which rendezvous server is used):
SDO_SAMPLE_MFG_KEEP_SVCS=false
SDO_RV_URL=http://sdo-sbx.trustedservices.intel.com:80
 
The "simulate-mfg.sh" script has been downloaded:
-rwxrwxr-x 1 steven steven 24971 Jun 24 15:39 ./simulate-mfg.sh
 
Now run this (to simulate manufacturing an SDO device):
  sudo -E ./simulate-mfg.sh | tee MFG_LOG && ls -l /var/sdo/voucher.json
After that, save the UUID and voucher, then shutdown and ship the device.

3. run simulate-mfg.sh : 

steven@steven-XS35:~$ sudo -E ./simulate-mfg.sh | tee MFG_LOG && ls -l /var/sdo/voucher.json
Found java 11
Found jq
Getting owner-boot-device script ...
Getting sdo_to.service ...
Verifying that rendezvous server sdo-sbx.trustedservices.intel.com is resolvable...
Pulling and tagging the SDO SCT services...
1.10: Pulling from openhorizon/manufacturer
Digest: sha256:224f5d035c259c2eb2a87655b0cb031e0fd5ce312cd9e74c97c3438929fdad3f
Status: Image is up to date for openhorizon/manufacturer:1.10
docker.io/openhorizon/manufacturer:1.10
1.10: Pulling from openhorizon/manufacturer-mariadb
Digest: sha256:8e2a79025972217d69c0aa767030899730b2297e511dea15aad67d48032a1cad
Status: Image is up to date for openhorizon/manufacturer-mariadb:1.10
docker.io/openhorizon/manufacturer-mariadb:1.10
Starting the SDO SCT services (could take about 75 seconds)...
manufacturer-mariadb is up-to-date
manufacturer is up-to-date
Updating the RV hostname in the mt_server_settings table...
Removing all rows from the rt_ownership_voucher table to enable redo...
Adding owner public key keys/owner-key.pub to the SCT services...
Running device initialization using java client...
24-06-2021 20:26:52.598 INFO  [main] org.sdo.pri.device.DeviceApp.logStarting - Starting DeviceApp on steven-XS35 with PID 14283 (/var/sdo/sdo_device_binaries_1.10_linux_x64/device/device.jar started by root in /var/sdo/sdo_device_binaries_1.10_linux_x64/device)
24-06-2021 20:26:52.622 INFO  [main] org.sdo.pri.device.DeviceApp.logStartupProfileInfo - No active profile set, falling back to default profiles: default
24-06-2021 20:26:57.179 INFO  [main] org.sdo.pri.device.DeviceApp.logStarted - Started DeviceApp in 8.986 seconds (JVM running for 12.858)
24-06-2021 20:26:57.524 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.cipher-block-mode = CTR
24-06-2021 20:26:57.530 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.di.uri = http://localhost:8039
24-06-2021 20:26:57.535 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.pm.credmfg.d = SDO Java Device
24-06-2021 20:26:57.540 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.output-dir = creds
24-06-2021 20:26:57.561 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.output-dir = [NativePRNG, Windows-PRNG, SHA1PRNG]
24-06-2021 20:26:57.568 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.serial = ab1d64b44e8b6a4f9a00b13912b9314ba2aece0141184eaf0df982985bacd
24-06-2021 20:26:57.573 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.stop-after-di = true
24-06-2021 20:26:57.580 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.cert = file:/var/sdo/sdo_device_binaries_1.10_linux_x64/device/creds/device.crt
24-06-2021 20:26:57.627 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.credentials is not set
24-06-2021 20:26:57.632 INFO  [main] org.sdo.pri.device.DeviceApp.logProperty - property org.sdo.device.key = file:/var/sdo/sdo_device_binaries_1.10_linux_x64/device/creds/device.key
24-06-2021 20:26:57.692 INFO  [main] org.sdo.pri.device.DeviceApp.lambda$applicationRunner$1 - device initialization begins
24-06-2021 20:26:58.245 INFO  [main] org.sdo.pri.device.DeviceApp.secureRandom - using SecureRandom NativePRNG
24-06-2021 20:26:58.256 WARN  [main] org.sdo.pri.device.DeviceApp.sslContext - UNSAFE: no-op TrustManager installed
24-06-2021 20:26:59.580 INFO  [main] o.sdo.pri.DeviceInitializationClient.call - DI URI is: http://localhost:8039
24-06-2021 20:26:59.790 INFO  [main] o.sdo.pri.DeviceInitializationClient.call - HttpRequest: http://localhost:8039/mp/113/msg/10 POST
{Content-Type=[application/json]}
{"m":"MTMAYWIxZDY0YjQ0ZThiNmE0ZjlhMDBiMTM5MTJiOTMxNGJhMmFlY2UwMTQxMTg0ZWFmMGRmOTgyOTg1YmFjZABTRE8gSmF2YSBEZXZpY2UALS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlHNk1HSUNBUUF3QURCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhBMElBQkhXam9WQjlJSVkrTWxycApTZWVkbkN1VGJDNnRoZi91bXJLdk9PcWlaMTJNZVFoa2dIaUVwOG8vZ00wNEc3dGpzNDhxSVZkT3lWMlhQL1RHCkVTOFovck9nQURBS0JnZ3Foa2pPUFFRREFnTklBREJGQWlCaFg3TitUTkFGRll5NmtVdFJzNmF6NWQ1QWswdEcKU3VQV2krUXNrMTNOWFFJaEFPQWZDRDRXM21valJKSVpuT2ZFTTBqL0lFZ3JqLzhPZ2RtT3R3Q1FER3JJCi0tLS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo="}
24-06-2021 20:27:12.680 INFO  [main] o.sdo.pri.DeviceInitializationClient.call - HttpResponse: (POST http://localhost:8039/mp/113/msg/10) 200
{authorization=[/v8AYQBiADEAZAA2ADQAYgA0ADQAZQA4AGIANgBhADQAZgA5AGEAMAAwAGIAMQAzADkAMQAyAGIAOQAzADEANABiAGEAMgBhAGUAYwBlADAAMQA0ADEAMQA4ADQAZQBhAGYAMABkAGYAOQA4ADIAOQA4ADUAYgBhAGMAZA], content-length=[365], content-type=[application/json], date=[Fri, 25 Jun 2021 00:27:12 GMT]}
{"oh":{"pv":113,"pe":1,"r":[1,[4,{"dn":"sdo-sbx.trustedservices.intel.com","po":80,"pow":80,"pr":"http"}]],"g":"hxe+KiHoSE+JU42pyibwjQ==","d":"SDO Java Device","pk":[13,1,[91,"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6+Yq304FMQeer35r6h/gt/s0lruFYoy6rJBiy2YvUCaCSwqpX5c62u/4xNbVzDj2Nw0TVod6cGlNo+WknJkZcA=="]],"hdc":[32,8,"gZ7XHK/T+HENMlO7kpZGlGYmPpdz+/uW6cSyV4UgVEI="]}}
24-06-2021 20:27:15.292 INFO  [main] o.sdo.pri.DeviceInitializationClient.call - HttpRequest: http://localhost:8039/mp/113/msg/12 POST
{Authorization=[/v8AYQBiADEAZAA2ADQAYgA0ADQAZQA4AGIANgBhADQAZgA5AGEAMAAwAGIAMQAzADkAMQAyAGIAOQAzADEANABiAGEAMgBhAGUAYwBlADAAMQA0ADEAMQA4ADQAZQBhAGYAMABkAGYAOQA4ADIAOQA4ADUAYgBhAGMAZA], Content-Type=[application/json]}
{"hmac":[32,108,"U0JuSUJRpKLzuIYArVFizl1rbvfF4vZ7VY9iiQlA/TI="]}
24-06-2021 20:27:27.519 INFO  [main] o.sdo.pri.DeviceInitializationClient.call - HttpResponse: (POST http://localhost:8039/mp/113/msg/12) 200
{content-length=[0], content-type=[application/json], date=[Fri, 25 Jun 2021 00:27:27 GMT]}

24-06-2021 20:27:28.063 INFO  [main] o.s.p.d.DeviceCredentialsStorage.store - credentials saved to /var/sdo/sdo_device_binaries_1.10_linux_x64/device/creds/8717be2a-21e8-484f-8953-8da9ca26f08d.oc
24-06-2021 20:27:28.072 INFO  [main] org.sdo.pri.device.DeviceApp.lambda$applicationRunner$1 - device initialization ends
Device UUID: 8717be2a-21e8-484f-8953-8da9ca26f08d
Extending the voucher to the owner...
Device serial and owner in the DB: ab1d64b44e8b6a4f9a00b13912b9314ba2aece0141184eaf0df982985bacd	all
Switching the device into owner mode with credential file 8717be2a-21e8-484f-8953-8da9ca26f08d.oc ...
Shutting down SDO SCT services...
Stopping manufacturer         ... done
Stopping manufacturer-mariadb ... done
Removing manufacturer         ... done
Removing manufacturer-mariadb ... done
Removing network manufacturer_network
-------------------------------------------------
Device UUID: 8717be2a-21e8-484f-8953-8da9ca26f08d
-------------------------------------------------
The extended ownership voucher is in file: /var/sdo/voucher.json
Device manufacturing initialization complete.
-rw-r--r-- 1 root root 1936 Jun 24 20:27 /var/sdo/voucher.json
steven@steven-XS35:~$ 
