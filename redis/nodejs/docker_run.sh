docker run -it --rm -p 8888:8080 -v "$PWD/src":/usr/src/myapp -w /usr/src/myapp node:latest node index.js