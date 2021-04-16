
export interface Connection {
    id: string,
    name: string,
    host: string,
    port: number,
    auth: string,
    keysPattern: string,
    namespaceSeparator: string,
    timeoutConnect: number,
    timeoutExecute: number,
    dbScanLimit: number,
    dataScanLimit: number,
}

export interface IConnectionService {
    connections: Promise<Array<Connection>>
}

const useConfigService = (): IConnectionService => {
    return {
        connections: fetch("").then(response => response.json() as Promise<Array<Connection>>)
    }
}

export default useConfigService;