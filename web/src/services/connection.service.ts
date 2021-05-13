import { defaultHeaders, serverErrorHandle } from "./servcie";
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
  return fetch("/connections").then(serverErrorHandle);
}

export const testConnection = (connection: Connection): Promise<any> => {
  return fetch("/connection/test", { method: 'post', headers: defaultHeaders, body: JSON.stringify(connection) });
}

export const saveConnection = (connection: Connection): Promise<Connection> => {
  return fetch("/connection", { method: 'post', headers: defaultHeaders, body: JSON.stringify(connection) })
    .then(serverErrorHandle);
}

export const deleteConnection = (id: string): Promise<any> => {
  return fetch(`/connection/${id}`, { method: 'delete' });
}

export const openConnection = (id: string): Promise<{ database: Array<any>, info: string }> => {
  return fetch(`/connection/${id}/open`, { method: 'post' })
    .then(serverErrorHandle);
}

export const disconnectionConnection = (id: string): Promise<any> => {
  return fetch(`/connection/${id}/disconnection`, { method: 'post' });
}

export const copyConnection = (connection: Connection): Promise<Connection> => {
  return fetch(`/connection/copy`, { method: 'post', headers: defaultHeaders, body: JSON.stringify(connection) })
    .then(serverErrorHandle);
}

export const executeCommand = <T>(command: Command): Promise<T> => {
  return fetch(`/connection/command`, { method: 'post', headers: defaultHeaders, body: JSON.stringify(command) })
    .then(serverErrorHandle);
}
