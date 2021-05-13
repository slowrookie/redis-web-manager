import { defaultHeaders, serverErrorHandle } from "./servcie";

export interface Config {
    theme: string,
    language: string
}

export const getConfig = (): Promise<Config> => {
    return fetch("/config").then(serverErrorHandle);
}

export const setConfig = (config: Config): Promise<Config> => {
    return fetch("/config", { method: 'post', headers: defaultHeaders, body: JSON.stringify(config) })
        .then(serverErrorHandle);
}