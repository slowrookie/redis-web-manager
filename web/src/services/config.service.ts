import axios from "axios";

export interface Config {
  theme: string,
  language: string,
  port: number
}

export interface About {
  version: string
	commit: string
	date: string
	builtBy: string
	environment: string
}

export const getConfig = (): Promise<Config> => {
  return axios.get('/config/')
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

export const about = (): Promise<About> => {
  return axios.get('/about')
    .then(ret => ret.data);
}

export const checkPort = (port: number): Promise<any> => {
  return axios.get('/config/port/check', {params: {port}})
    .then(ret => {
      return true
    });
}