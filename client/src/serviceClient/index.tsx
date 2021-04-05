import { ServiceClient } from '../proto/services_pb_service';
import pb from '../proto/messages_pb';
import { grpc } from '@improbable-eng/grpc-web';
import { List } from '@material-ui/core';
import { rejects } from 'assert';

// export type ServiceClientAttached = {
//   client: ServiceClient;
// };

const client = new ServiceClient('http://localhost:80', {});

// const serviceClient = <P extends {}>(WrappedComponent: React.ComponentType<P & ServiceClientAttached>) =>
//   class MessageServiceAttached extends React.Component<P> {
//   render() {
//     return <WrappedComponent {...this.props} client={client} />;
//   }
// };

export interface iSignin {
  handle: string,
  pw: string
}

class serviceClient {
  public signinRequest = (args: iSignin) => {
    return new Promise(resolve =>{
      const req = new pb.SignInRequest()
      req.setHandle(args.handle)
      req.setPassword(args.pw)
      client.signIn(req,  (err: any, res: any) => {
        if (err || res === null) {
            throw err
        }
        resolve(res)
      })
    })

  }

  public userRequest = (token: string) => {
    return new Promise(resolve =>{
      const req = new pb.UserRequest()
      const meta = new grpc.Metadata();
      meta.set('authorization', 'Bearer '+token);
      const deadline = new Date();
      deadline.setSeconds(deadline.getSeconds() + 10);
      meta.set('deadline', deadline.getTime().toString());
      console.log(meta)

      client.user(req, meta, (err: any, res: any) => {
        if (err || res === null) {
            rejects(err)
        }
        resolve(res)
      })
    })

  }
}

export default new serviceClient()