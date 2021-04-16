import en_US from './en_US.json';
import zh_hans from './zh-Hans.json';
import zh_hant from './zh-Hant.json';

export const resources = {
    en_US: {
        translation: en_US
    },
    zh_Hans: {
        translation: zh_hans
    },
    zh_hant: {
        translation: zh_hant
    }
}

export const supportedLanguages: {[index: string]: string} = {
    en_US: "English",
    zh_Hans: "简体中文",
    zh_hant: "繁體中文",
}