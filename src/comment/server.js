var PROTO_PATH = __dirname + '/comment.proto';

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }
);
var commentProto = grpc.loadPackageDefinition(packageDefinition).comment;

/**
 * Implements the SayHello RPC method.
 */
function sayHello(call, callback) {
  callback(null, { code: JSON.stringify(call) });
}

/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
function main() {
  var server = new grpc.Server();
  server.addService(commentProto.CommentService.service, {
    AddComment: sayHello
  });
  server.bind('0.0.0.0:9090', grpc.ServerCredentials.createInsecure());
  server.start();
}

main();
