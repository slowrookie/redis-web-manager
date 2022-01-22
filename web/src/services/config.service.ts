import { defaultWebSocket } from "./WebSocketService";

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
  return defaultWebSocket.request({method: 'Config'});
}

export const setConfig = (config: Config): Promise<Config> => {
  return defaultWebSocket.request({method: 'Config.Set', params: config});
}

export const about = (): Promise<About> => {
  return defaultWebSocket.request({method: 'About'});
}

export const checkPort = (port: number): Promise<any> => {
  return defaultWebSocket.request({method: 'Config.CheckPort', params: port});
}