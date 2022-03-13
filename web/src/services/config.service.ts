import { defaultService } from "./api.service";

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
  return defaultService.request({method: 'Config'});
}

export const setConfig = (config: Config): Promise<Config> => {
  return defaultService.request({method: 'SetConfig', params: config});
}

export const about = (): Promise<About> => {
  return defaultService.request({method: 'AboutInfo'});
}