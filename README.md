# ngsw-rehash

This is a simple program to update the hashes inside of a ngsw.json file. This file is used by Angular webapps to configure a Service Worker and to update the files from the server correctly.

Features:
- Dead-Simple CLI interface, just pass the path to the ngsw.json file of the dist folder
- No dependency needed (single statically linked binary)
- Small footprint (ready be used inside docker images)


## Why update this file?
Image you want to change some of the final dist files of your angular webapp on runtime. For example to the the correct API endpoint URL on the startup of the frontend docker container. This will alter the hash of the edited file and will break the PWA caching of that file. To fix this, you can simply run ngsw-rehash after altering the file and the hash will be correct :rocket:.


## Example
```bash
sed "s,BACKEND_URL,https://fancy-api-url.example.com/,g' dist/your-app/index.html
ngsw-rehash dist/your-app/ngsw.json
```


## Alternatives?
There is a command inside the Angular CLI that does basically the same as this tool. If you have the whole Angular CLI available while you want to recreate the hashes, maybe you should better use this one:

```bash
node_modules/.bin/ngsw-config dist src/ngsw-config.json
```

But for this command you have to install the Angular CLI, which blows up for example your Docker image. As normally this is not needed, this simple tool can help keep the images small and still be able to recreate the hashes for changed files.
