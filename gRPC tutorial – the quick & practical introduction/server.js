const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')

const PROTO_PATH = 'chat.proto'
const Server_URI = '0.0.0.0:50051'

const userInChat = []

const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)

//
const joinChat = call => {
  console.log(`User ${call.request.user} has joined.`)

  call.on('cancelled', () => {
    userInChat = userInChat.filter(user => user !== call)
  })

  userInChat.push(call)
}

const sendMessage = (call, callback) => {
  const { message } = call.request

  if(!message) {
    return callback(new Error('You must provide a non-empty message.'))
  }

  const messageToSend = {
    ...call.request,
    timestamp: Math.floor(new Date().getTime() / 1000),
  }

  userInChat.forEach(user => user.write(messageToSend))

  callback(null, {})
}

const server = new grpc.Server()
server.addService(protoDescriptor.ChatService.service, {
  joinChat,
  sendMessage,
})
server.bind(Server_URI, grpc.ServerCredentials.createInsecure())

server.start()
console.log('server is running!')
