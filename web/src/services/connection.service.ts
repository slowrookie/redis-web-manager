import { defaultWebSocket } from "./WebSocketService";
export interface Connection {
  id: string,
  name: string,
  host: string,
  port: number,
  username: string,
  auth: string,
  keysPattern: string,
  namespaceSeparator: string,
  timeoutConnect: number,
  timeoutExecute: number,
  dbScanLimit: number,
  dataScanLimit: number,
  tls: {
    enable: boolean
    cert?: string
    key?: string
    ca?: string
  }
}

export interface Command {
  id: string,
  commands: Array<any>
}

export const getConnections = (): Promise<any> => {
  return defaultWebSocket.request({method: 'Connections'});
}

export const testConnection = (connection: Connection): Promise<any> => {
  return defaultWebSocket.request({method: 'Connection.Test', params: connection})
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const saveConnection = (connection: Connection): Promise<Connection> => {
  return defaultWebSocket.request({method: 'Connection.Edit', params: connection})
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const deleteConnection = (id: string): Promise<any> => {
  return defaultWebSocket.request({method: 'Connection.Delete', params: id})
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const openConnection = (id: string): Promise<{ database: Array<any>, info: string }> => {
  return defaultWebSocket.request({method: 'Connection.Open', params: id})
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const disconnectionConnection = (id: string): Promise<any> => {
  return defaultWebSocket.request({method: 'Connection.Disconnection', params: id})
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const copyConnection = (connection: Connection): Promise<Connection> => {
  return defaultWebSocket.request({method: 'Connection.Copy', params: connection})
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const executeCommand = <T>(command: Command): Promise<T> => {
  return defaultWebSocket.request({method: 'Connection.Command', params: command})
    .catch(err => {
      throw new Error(err?.response?.data || 'The service cannot be accessed, please check the network and service');
    });
}