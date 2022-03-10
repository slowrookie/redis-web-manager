import { defaultService } from "./api.service";

export const convertLength = (params: { data: string }): Promise<String> => {
  return defaultService.request({method: 'ConvertLength', params});
}

export const convertToHex = (params: { data: string }): Promise<String> => {
  return defaultService.request({method: 'ConvertToHex', params});
}

export const convertToJson = (params: { data: string }): Promise<String> => {
  return defaultService.request({method: 'ConvertToJson', params});
}

export const convertBase64ToText = (params: { data: string }): Promise<String> => {
  return defaultService.request({method: 'ConvertBase64ToText', params});
}

export const convertBase64ToJson = (params: { data: string }): Promise<String> => {
  return defaultService.request({method: 'ConvertBase64ToJson', params});
}


export const convertToBinary = (params: { data: string }): Promise<String> => {
  return defaultService.request({method: 'ConvertToBinary', params});
}

export const convert: { [index: string]: (params: any) => Promise<any> } = {
  length: convertLength,
  hex: convertToHex,
  json: convertToJson,
  base64ToText: convertBase64ToText,
  base64ToJson: convertBase64ToJson,
  binary: convertToBinary
}