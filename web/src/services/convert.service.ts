import axios from "axios";

export const convertLength = (params: { data: string }): Promise<String> => {
  return axios.post('/convert/length', params)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const convertToHex = (params: { data: string }): Promise<String> => {
  return axios.post('/convert/toHex', params)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const convertToJson = (params: { data: string }): Promise<String> => {
  return axios.post('/convert/toJson', params)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const convertBase64ToText = (params: { data: string }): Promise<String> => {
  return axios.post('/convert/base64ToText', params)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const convertBase64ToJson = (params: { data: string }): Promise<String> => {
  return axios.post('/convert/base64ToJson', params)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const convertToBinary = (params: { data: string }): Promise<String> => {
  return axios.post('/convert/toBinary', params)
    .then(ret => ret.data)
    .catch(err => {
      throw new Error(err?.response?.data);
    });
}

export const convert: { [index: string]: (params: any) => Promise<any> } = {
  length: convertLength,
  hex: convertToHex,
  json: convertToJson,
  base64ToText: convertBase64ToText,
  base64ToJson: convertBase64ToJson,
  binary: convertToBinary
}