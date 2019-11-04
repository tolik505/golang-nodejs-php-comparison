var http = require("http")
var client = require("redis").createClient(6379, "172.17.0.1");

client.select(8);

var server = http.createServer(function(req,rep) {
        rep.writeHead(200);

        client.get("count", function(err,reply) {
                if (!reply) reply = 0;
                rep.end(reply.toString());
                reply++;
                client.set("count", reply);
        });
}).listen(8080);