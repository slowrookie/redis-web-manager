import { defaultHeaders } from "./servcie";

export interface Config {
    theme: string,
    language: string
}

export interface IConfigService {
    getConfig: () => Promise<Config>
    setConfig: (config: Config) => Promise<Config>
}

const configService = (): IConfigService => {
    return {
        getConfig: () => fetch("/config").then(response => response.json() as Promise<Config>),
        setConfig: (config: Config) => fetch("/config", { method: 'post', headers: defaultHeaders, body: JSON.stringify(config) })
            .then(response => response.json() as Promise<Config>),
    }
}

export default configService;