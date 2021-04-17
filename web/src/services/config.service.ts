import { defaultHeaders } from "./servcie";

export interface Config {
    theme: string,
    language: string
}

export const getConfig = ():Promise<Config> => {
    return fetch("/config").then(response => response.json() as Promise<Config>);   
}

export const setConfig = (config: Config): Promise<Config> => {
    return fetch("/config", { method: 'post', headers: defaultHeaders, body: JSON.stringify(config) })
    .then(response => response.json() as Promise<Config>);
}