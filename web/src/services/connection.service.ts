import axios from "axios";
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
}

export interface Command {
  id: string,
  commands: Array<any>
}

export const getConnections = (): Promise<any> => {
  return axios.get('/connections').then(ret => ret.data);
}

export const testConnection = (connection: Connection): Promise<any> => {
  return axios.post('/connection/test', connection)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const saveConnection = (connection: Connection): Promise<Connection> => {
  return axios.post('/connection', connection)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const deleteConnection = (id: string): Promise<any> => {
  return axios.delete(`/connection/${id}`)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const openConnection = (id: string): Promise<{ database: Array<any>, info: string }> => {
  return axios.post(`/connection/${id}/open`)
  .then(ret => ret.data)
  .catch(err => {
    throw new Error(err?.response?.data);
  });
}

export const disconnectionConnection = (id: string): Promise<any> => {
  return axios.post(`/connection/${id}/disconnection`)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const copyConnection = (connection: Connection): Promise<Connection> => {
  return axios.post(`/connection/copy`, connection)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const executeCommand = <T>(command: Command): Promise<T> => {
  return axios.post('/connection/command', command)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data || 'The service cannot be accessed, please check the network and service');
    });
}
