import { createTheme } from "@fluentui/style-utilities";
import { Theme } from "@fluentui/theme";
import { DefaultTheme, DarkTheme, WordTheme, TeamsTheme } from '@fluentui/theme-samples';

DarkTheme.name = 'dark';

export const themes: { [index: string]: Theme } = {
  default: DefaultTheme,
  light: DefaultTheme,
  dark: DarkTheme,
  word: WordTheme,
  teams: TeamsTheme,
  customDark: createTheme({
    palette: {
      themePrimary: '#007acc',
      themeLighterAlt: '#f3f9fd',
      themeLighter: '#cfe7f7',
      themeLight: '#a8d3f0',
      themeTertiary: '#5aabe0',
      themeSecondary: '#1988d2',
      themeDarkAlt: '#006eb8',
      themeDark: '#005d9b',
      themeDarker: '#004572',
      neutralLighterAlt: '#282828',
      neutralLighter: '#313131',
      neutralLight: '#3f3f3f',
      neutralQuaternaryAlt: '#484848',
      neutralQuaternary: '#4f4f4f',
      neutralTertiaryAlt: '#6d6d6d',
      neutralTertiary: '#3d3d3d',
      neutralSecondary: '#7a7a7a',
      neutralPrimaryAlt: '#b4b4b4',
      neutralPrimary: '#cccccc',
      neutralDark: '#d8d8d8',
      black: '#e2e2e2',
      white: '#1e1e1e',
    }
  }),
  CustomLight: createTheme({
    palette: {
      themePrimary: '#0078d4',
      themeLighterAlt: '#eff6fc',
      themeLighter: '#deecf9',
      themeLight: '#c7e0f4',
      themeTertiary: '#71afe5',
      themeSecondary: '#2b88d8',
      themeDarkAlt: '#106ebe',
      themeDark: '#005a9e',
      themeDarker: '#004578',
      neutralLighterAlt: '#faf9f8',
      neutralLighter: '#f3f2f1',
      neutralLight: '#edebe9',
      neutralQuaternaryAlt: '#e1dfdd',
      neutralQuaternary: '#d0d0d0',
      neutralTertiaryAlt: '#c8c6c4',
      neutralTertiary: '#a19f9d',
      neutralSecondary: '#605e5c',
      neutralPrimaryAlt: '#3b3a39',
      neutralPrimary: '#323130',
      neutralDark: '#201f1e',
      black: '#000000',
      white: '#ffffff',
    }
  })
}

export const getCustomTheme = (theme: string) => {
  theme = theme || 'light'
  const th = themes[theme]
  th.fonts.medium.fontSize = 12;
  return th
}