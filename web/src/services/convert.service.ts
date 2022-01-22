import { defaultWebSocket } from "./WebSocketService";

export const convertLength = (params: { data: string }): Promise<String> => {
  return defaultWebSocket.request({method: 'Convert.Length', params});
}

export const convertToHex = (params: { data: string }): Promise<String> => {
  return defaultWebSocket.request({method: 'Convert.ToHex', params});
}

export const convertToJson = (params: { data: string }): Promise<String> => {
  return defaultWebSocket.request({method: 'Convert.ToJson', params});
}

export const convertBase64ToText = (params: { data: string }): Promise<String> => {
  return defaultWebSocket.request({method: 'Convert.Base64ToText', params});
}

export const convertBase64ToJson = (params: { data: string }): Promise<String> => {
  return defaultWebSocket.request({method: 'Convert.Base64ToJson', params});
}


export const convertToBinary = (params: { data: string }): Promise<String> => {
  return defaultWebSocket.request({method: 'Convert.ToBinary', params});
}

export const convert: { [index: string]: (params: any) => Promise<any> } = {
  length: convertLength,
  hex: convertToHex,
  json: convertToJson,
  base64ToText: convertBase64ToText,
  base64ToJson: convertBase64ToJson,
  binary: convertToBinary
}