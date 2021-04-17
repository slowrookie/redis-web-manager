
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

export const getConnections = (): Promise<Array<Connection>> => {
    return fetch("/connections").then(response => response.json() as Promise<Array<Connection>>);
}