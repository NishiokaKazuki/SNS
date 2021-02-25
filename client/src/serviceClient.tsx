import React from 'react';
import { ServiceClient } from './proto/services_pb_service';
import pb from './proto/messages_pb';

// export type ServiceClientAttached = {
//   client: ServiceClient;
// };

const client = new ServiceClient(`localhost:80`);

// const serviceClient = <P extends {}>(WrappedComponent: React.ComponentType<P & ServiceClientAttached>) =>
//   class MessageServiceAttached extends React.Component<P> {
//   render() {
//     return <WrappedComponent {...this.props} client={client} />;
//   }
// };

export interface iSignin {
  id: string,
  pw: string
}

class serviceClient {

}

export default new serviceClient()