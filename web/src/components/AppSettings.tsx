import { ContextualMenuItemType, Dialog, PrimaryButton, Stack, Text, TooltipHost } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { supportedLanguages } from '../locales/resources';
import { About, about } from '../services/config.service';

export interface IAppSettings {
  onChangeLanguage?: (language: string) => void
  onChangeTheme?: (theme: string) => void
}

const headerButtonStyles = {
  root: {
    minWidth: 'auto',
    borderRadius: 0,
    padding: '0 8px',
    height: 42
  }
}

export const AppSettings = (props: IAppSettings) => {

  const { t } = useTranslation(),
    [aboutData, setAboutData] = useState<About>(),
    [aboutDialogHidden, setAboutDialogHidden] = useState(true);

  useEffect(() => {
    about().then((ret) => {
      if (!ret) return;
      setAboutData(ret)
    })
      .finally(() => {

      })
  }, []);


  return (<>
    {/* settings */}
    <TooltipHost content={t('More')}>
      <PrimaryButton iconProps={{ iconName: 'settings', style: { height: 'auto' } }} styles={headerButtonStyles} menuProps={{
        items: [
          { key: 'divider_1', itemType: ContextualMenuItemType.Divider },
          {
            key: 'colorSolid', text: t('Theme'), iconProps: { iconName: 'ColorSolid', style: { lineHeight: '14px' } }, subMenuProps: {
              items: [
                { key: 'darkTheme', text: t('Theme-Dark'), title: t('Theme-Dark'), onClick: () => props.onChangeTheme && props.onChangeTheme('dark') },
                { key: 'lightTheme', text: t('Theme-Light'), title: t('Theme-Light'), onClick: () => props.onChangeTheme && props.onChangeTheme('light') },
              ],
            },
          },
          {
            key: 'localeLanguage', text: t('Lanaguage'), iconProps: { iconName: 'LocaleLanguage', style: { lineHeight: '14px' } }, subMenuProps: {
              items: Object.keys(supportedLanguages).map(v => {
                return { key: v, text: supportedLanguages[v], onClick: () => props.onChangeLanguage && props.onChangeLanguage(v) }
              }),
            },
          },
          { key: 'divider_1', itemType: ContextualMenuItemType.Divider },
          {
            key: 'about', id: 'settingAbout', text: t('About'), iconProps: { iconName: 'Info', style: { lineHeight: '14px' } }, onClick: () => {
              setAboutDialogHidden(false);
            }
          },
        ]
      }} onRenderMenuIcon={() => <></>} />
    </TooltipHost>

    {/* about dialog  */}
    <Dialog hidden={aboutDialogHidden} onDismiss={() => { setAboutDialogHidden(!aboutDialogHidden) }} dialogContentProps={{ title: "Redis Web Manager" }} modalProps={{ isBlocking: true }}>
      <Stack>
        <Stack horizontal horizontalAlign="space-between">
          <Text>Version: </Text>
          <Text>{aboutData?.Version}</Text>
        </Stack>
        <Stack horizontal horizontalAlign="space-between">
          <Text>Commit: </Text>
          <Text>{aboutData?.Commit}</Text>
        </Stack>
        <Stack horizontal horizontalAlign="space-between">
          <Text>Date: </Text>
          <Text>{aboutData?.Date}</Text>
        </Stack>
        <Stack horizontal horizontalAlign="space-between">
          <Text>BuiltBy: </Text>
          <Text>{aboutData?.BuiltBy}</Text>
        </Stack>
        <Stack horizontal horizontalAlign="space-between">
          <Text>Environment: </Text>
          <Text>{aboutData?.Environment}</Text>
        </Stack>
      </Stack>
    </Dialog>
  </>)

}