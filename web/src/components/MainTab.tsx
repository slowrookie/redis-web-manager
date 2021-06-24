import { ContextualMenuItemType, Depths, Icon, Pivot, PivotItem, PrimaryButton, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { supportedLanguages } from '../locales/resources';
import { Connection } from '../services/connection.service';
import { ConnectionItem, IConnectionItemProps } from './ConnectionItem';
import { ConnectionList } from './ConnectionList';
import { IDatabase } from './Database';
export interface IMainTabProps {
  onChangeLanguage?: (language: string) => void
  onChangeTheme?: (theme: string) => void
}

interface IMainTab {
  connectionItems: Array<IConnectionItemProps>,
  selectedKey: string | undefined,
  showConnectionList: boolean,
}

const headerButtonStyles = {
  root: {
    minWidth: 'auto',
    borderRadius: 0,
    padding: '0 8px',
    height: 42
  }
}

export const MainTab = (props: IMainTabProps) => {
  const
    [mainTab, setMainTab] = useState<IMainTab>({ connectionItems: [], selectedKey: '', showConnectionList: true }),
    { t } = useTranslation(),
    theme = useTheme();

  useEffect(() => {
    if (!mainTab.connectionItems.length) {
      setMainTab(m => {
        return { ...m, selectedKey: '', showConnectionList: true }
      })
    }
  }, [mainTab.connectionItems])

  const handleConnectionClick = (v: Connection, info: any, databases: Array<IDatabase>) => {
    const exists = mainTab.connectionItems.filter(item => item.connection.id === v.id);
    const connections = mainTab.connectionItems;
    !exists.length && connections.push({ connection: v, info, databases });
    setMainTab({ ...mainTab, connectionItems: [...connections], selectedKey: v.id, showConnectionList: false });
  }

  return (<>
    <Stack horizontal style={{ backgroundColor: theme.palette.themeSecondary, height: 42, boxShadow: Depths.depth8 }}>
      {/* home */}
      <TooltipHost content={'连接管理'}>
        <PrimaryButton iconProps={{ iconName: 'home' }} styles={headerButtonStyles} onClick={() => setMainTab({ ...mainTab, selectedKey: '', showConnectionList: true })} />
      </TooltipHost>
      {/* tabs */}
      <Stack.Item grow={1}>
        <Pivot overflowBehavior='menu' headersOnly={true} styles={{ link: { height: 42, lineHeight: '42px', background: 'transparent' } }}
          linkFormat='tabs'
          selectedKey={mainTab.selectedKey}
          getTabId={itemKey => `connection-tab-${itemKey}`}
          onLinkClick={item => { 
            setMainTab({ ...mainTab, selectedKey: item?.props.itemKey, showConnectionList: false }) 
          }} >
          {mainTab.connectionItems.map((v, index) => {
            return <PivotItem headerText={v.connection.name} key={v.connection.id} itemKey={v.connection.id}
              onRenderItemLink={(link, defaultRenderer: any) => (
                <span>
                  {defaultRenderer(link)}
                  <Icon style={{ marginLeft: 5 }} iconName="Cancel" className="icon" onClick={(e: React.MouseEvent) => {
                    e.stopPropagation();
                    e.preventDefault();
                    const connections = mainTab.connectionItems.filter((_, i) => index !== i);
                    mainTab.connectionItems = [...connections];
                    setMainTab({ ...mainTab, connectionItems: [...connections], selectedKey: connections.length > 0 ? connections[0].connection.id : undefined, showConnectionList: !connections.length && false });
                  }}
                  />
                </span>
              )}
            />
          })}
        </Pivot>
      </Stack.Item>
      
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
          ]
        }} onRenderMenuIcon={() => <></>} />
      </TooltipHost>
    </Stack>

    <div style={{ display: mainTab.showConnectionList ? '' : 'none', height: 'calc(100vh - 42px)', overflowY: 'auto', overflowX: 'hidden' }} >
      <ConnectionList onConnectionClick={handleConnectionClick} />
    </div>

    {
      mainTab.connectionItems.map((v, i) => {
        return (<div key={v.connection.id} style={{ display: mainTab.selectedKey === v.connection.id ? '' : 'none', height: 'calc(100vh - 42px)' }} >
          <ConnectionItem {...v} />
        </div>)
      })
    }
  </>);
}
