import { defaultService } from "./api.service";


export interface Connection {
  id: string,
  name: string,
  host: string,
  port: number,
  addrs?: Array<string>,
  username: string,
  auth: string,
  keysPattern: string,
  namespaceSeparator: string,
  timeoutConnect: number,
  timeoutExecute: number,
  dbScanLimit: number,
  dataScanLimit: number,
  tls?: {
    enable: boolean
    cert?: string
    key?: string
    ca?: string
  },
  ssh?: {
    enable: boolean
    user?: string
    password?: string
    privateKey?: string
    host?: string
    port?: number
  },
  isCluster?: boolean
  isSentinel?: boolean
  sentinelPassword?: string
  masterName?: string
  routeByLatency?: boolean
  routeRandomly?: boolean
}
export interface Command {
  id: string,
  commands: Array<any>
}

export const getConnections = (): Promise<any> => {
  return defaultService.request({ method: 'Connections' })
}

export const testConnection = (connection: Connection): Promise<any> => {
  return defaultService.request({ method: 'TestConnection', params: connection });
}

export const saveConnection = (connection: Connection): Promise<Connection> => {
  return defaultService.request({ method: connection.id ? 'EditConnection' : 'NewConnection', params: connection });
}

export const deleteConnection = (id: string): Promise<any> => {
  return defaultService.request({ method: 'DeleteConnection', params: id });
}

export const openConnection = (id: string): Promise<{ database: Array<any>, info: string }> => {
  return defaultService.request({ method: 'OpenConnection', params: id });
}

export const disconnectionConnection = (id: string): Promise<any> => {
  return defaultService.request({ method: 'DisConnection', params: id });
}

export const copyConnection = (connection: Connection): Promise<Connection> => {
  return defaultService.request({ method: 'NewConnection', params: connection });
}

export const executeCommand = <T>(command: Command): Promise<T> => {
  return defaultService.request({ method: 'CommandConnection', params: command });
}

export const suggestions = <T>(command: string): Promise<T> => {
  return defaultService.request({ method: 'Suggestions', params: command });
}