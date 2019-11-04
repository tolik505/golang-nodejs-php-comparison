var http = require("http");

http.createServer(function(request,response) {
    response.writeHeader(200);

    response.write("You requested " + request.url);

    response.end();
}).listen(8080);
