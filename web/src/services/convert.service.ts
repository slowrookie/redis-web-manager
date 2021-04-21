import { defaultHeaders } from "./servcie";

export const convertLength = (params: { data: string }): Promise<String> => {
    return fetch("/convert/length", { method: 'post', headers: defaultHeaders, body: JSON.stringify(params) })
        .then(response => response.text());;
}

export const convertToHex = (params: { data: string }): Promise<String> => {
    return fetch("/convert/toHex", { method: 'post', headers: defaultHeaders, body: JSON.stringify(params) })
        .then(response => response.text());;
}

export const convertToJson = (params: { data: string }): Promise<String> => {
    return fetch("/convert/toJson", { method: 'post', headers: defaultHeaders, body: JSON.stringify(params) })
        .then(response => response.text());;
}

export const convertBase64ToText = (params: { data: string }): Promise<String> => {
    return fetch("/convert/base64ToText", { method: 'post', headers: defaultHeaders, body: JSON.stringify(params) })
        .then(response => response.text());;
}

export const convertBase64ToJson = (params: { data: string }): Promise<String> => {
    return fetch("/convert/base64ToJson", { method: 'post', headers: defaultHeaders, body: JSON.stringify(params) })
        .then(response => response.text());;
}

export const convertToBinary = (params: { data: string }): Promise<String> => {
    return fetch("/convert/toBinary", { method: 'post', headers: defaultHeaders, body: JSON.stringify(params) })
        .then(response => response.text());;
}

export const convert: { [index: string]: (params: any) => Promise<any> } = {
    length: convertLength,
    hex: convertToHex,
    json: convertToJson,
    base64ToText: convertBase64ToText,
    base64ToJson: convertBase64ToJson,
    binary: convertToBinary
}