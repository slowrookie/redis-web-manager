import { ContextualMenu, ContextualMenuItemType, Icon } from '@fluentui/react';
import React, { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ConfigEvent, ConfigEventAction } from '../../events/ConfigEvent';
import { supportedLanguages } from '../../locales/resources';
import { Config } from '../../services/config.service';
import { AboutDialog } from './AboutDialog';

export interface IAppSettings {
  config: Config
}

export const AppSettings = (props: IAppSettings) => {
  const { t } = useTranslation(),
    [aboutDialogHidden, setAboutDialogHidden] = useState(true),
    [showContextualMenu, setShowContextualMenu] = useState(false),
    linkRef = useRef(null);

  const handleChangeLanguage = (language: string) => {
    ConfigEvent.next({ action: ConfigEventAction.Language, params: language })
  }

  const handleChangeTheme = (theme: string) => {
    ConfigEvent.next({ action: ConfigEventAction.Theme, params: theme })
  }
  
  return (<>
    {/* settings */}
    <div ref={linkRef} onClick={(e: React.MouseEvent) => {
      e.stopPropagation();
      e.preventDefault();
      setShowContextualMenu(true);
    }}>
      <Icon iconName="ChevronDown" className="icon" />

      <ContextualMenu
        items={[
          { key: 'divider_1', itemType: ContextualMenuItemType.Divider },
          {
            key: 'colorSolid', text: t('Theme'), iconProps: { iconName: 'ColorSolid', style: { lineHeight: '14px' } }, subMenuProps: {
              items: [
                { key: 'lightTheme', text: t('Theme-Light'), title: t('Theme-Light'), onClick: () => handleChangeTheme('light') },
                { key: 'darkTheme', text: t('Theme-Dark'), title: t('Theme-Dark'), onClick: () => handleChangeTheme('dark') },
                { key: 'WordTheme', text: t('Theme-Word'), title: t('Theme-Word'), onClick: () => handleChangeTheme('word') },
                { key: 'TeamsTheme', text: t('Theme-Teams'), title: t('Theme-Teams'), onClick: () => handleChangeTheme('teams') },
              ],
            },
          },
          {
            key: 'localeLanguage', text: t('Language'), iconProps: { iconName: 'LocaleLanguage', style: { lineHeight: '14px' } }, subMenuProps: {
              items: Object.keys(supportedLanguages).map(v => {
                return { key: v, text: supportedLanguages[v], onClick: () => handleChangeLanguage(v) }
              }),
            },
          },
          { key: 'divider_about', itemType: ContextualMenuItemType.Divider },
          {
            key: 'about', id: 'settingAbout', text: t('About'), iconProps: { iconName: 'Info', style: { lineHeight: '14px' } }, onClick: () => {
              setAboutDialogHidden(false);
            }
          },
        ]}
        hidden={!showContextualMenu}
        target={linkRef}
        onItemClick={() => { setShowContextualMenu(false) }}
        onDismiss={() => { setShowContextualMenu(false) }}
      />

      {/* about dialog  */}
      <AboutDialog hidden={aboutDialogHidden} onDismiss={() => setAboutDialogHidden(true)} />

    </div>
  </>)

}