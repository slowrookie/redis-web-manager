import axios from "axios";

export interface Config {
  theme: string,
  language: string
}

export const getConfig = (): Promise<Config> => {
  return axios.get('/config')
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const setConfig = (config: Config): Promise<Config> => {
  return axios.post('/config', config)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}