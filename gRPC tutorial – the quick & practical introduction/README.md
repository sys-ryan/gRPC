## gRPC tutorial – the quick & practical introduction   
### “gRPC hello world” – a simplified chat application   
### using Node.js and JavaScript  

#### purpose : Learning the concept of gRPC through the tutorial    

#### reference : https://tsh.io/blog/grpc-tutorial/   
#### by Rafał Ostrowski  






```
npm init --yes   
npm install grpc @grpc/proto-loader --save  
```




`joinChat` – a server-side streaming method (notice the stream keyword in its definition). It means that once client sends a single request to the server, it will receive a stream of responses over time.  

`sendMessage` – a simple unary method that sends a single request and receives a single response (should be quite familiar to anyone who has ever worked with REST APIs)
