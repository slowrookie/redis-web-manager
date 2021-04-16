import { ContextualMenu, ContextualMenuItemType, Depths, DirectionalHint, Pivot, PivotItem, PrimaryButton, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { supportedLanguages } from '../locales/resources';
import { Connection } from '../services/connection.service';

export interface IMainTabProps {
  onChangeLanguage?: (language: string) => void
  onChangeTheme?: (theme: string) => void
}

interface IMainTab {
  connections: Array<Connection>,
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

const headerMenuItemProps = {
  styles: {
    item: {
      height: 28, lineHeight: '28px'
    },
    root: {
      height: 28, lineHeight: '28px'
    }
  }
}

export const MainTab = (props: IMainTabProps) => {
  const
    [mainTab, setMainTab] = useState<IMainTab>({ connections: [], selectedKey: '', showConnectionList: true }),
    [contextmenu, setContextmenu] = useState<any>({ hidden: true, target: EventTarget || null, index: -1 }),
    { t } = useTranslation(),
    theme = useTheme();

  useEffect(() => {
    if (!mainTab.connections.length) {
      setMainTab(m => {
        return { ...m, selectedKey: '', showConnectionList: true }
      })
    }
    const connectionsTabHeadersEls = mainTab.connections.map((v: any, i) => {
      const el = document.getElementById(`connection-tab-${v.id}`);
      el?.addEventListener("contextmenu", (e: MouseEvent) => {
        e.stopPropagation();
        e.preventDefault();
        setContextmenu({ hidden: false, target: e?.currentTarget, index: i });
      });
      return el;
    })

    return () => {
      connectionsTabHeadersEls.forEach(el => el?.removeEventListener("contextmenu", () => { }));
    }
  }, [mainTab.connections])

  const handleConnectionClick = (v: Connection) => {
    const exists = mainTab.connections.filter(item => item.id === v.id);
    const connections = mainTab.connections;
    !exists.length && connections.push(v);
    setMainTab({ ...mainTab, connections: [...connections], selectedKey: v.id, showConnectionList: false });
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
          onLinkClick={item => { setMainTab({ ...mainTab, selectedKey: item?.props.itemKey, showConnectionList: false }) }} >
          {mainTab.connections.map((v, i) => {
            return <PivotItem headerText={v.name} key={v.id} itemKey={v.id} />
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

    <ContextualMenu
      directionalHint={DirectionalHint.bottomRightEdge}
      isBeakVisible={true}
      items={[
        {
          key: 'close', text: t('Close'), itemProps: headerMenuItemProps, onClick: (e) => {
            const cons = mainTab.connections.filter((_, index) => index !== contextmenu.index);
            mainTab.connections = cons;
            if (cons.length) {
              mainTab.selectedKey = cons[0].id;
            } else {
              mainTab.showConnectionList = true;
            }
            setMainTab({ ...mainTab })
          }
        },
        {
          key: 'closeOtherTabs', text: t('Close other tabs'), itemProps: headerMenuItemProps, onClick: (e) => {
            setMainTab({ ...mainTab, connections: mainTab.connections.filter((_, index) => index === contextmenu.index) })
          }
        },
        {
          key: 'closeAll', text: t('Close all'), itemProps: headerMenuItemProps, onClick: (e) => {
            setMainTab({ ...mainTab, connections: [], showConnectionList: true })
          }
        },
      ]}
      hidden={contextmenu.hidden}
      target={contextmenu.target}
      onDismiss={() => setContextmenu({ ...contextmenu, hidden: true })}
    />

    <div style={{ display: mainTab.showConnectionList ? '' : 'none', height: 'calc(100vh - 40px)', overflowY: 'auto', overflowX: 'hidden', position: 'relative' }} >
      {/* <ConnectionList onConnectionClick={handleConnectionClick} /> */}
    </div>

    {
      mainTab.connections.length > 0 && mainTab.connections.map((v, i) => {
        return (<div key={v.id} style={{ display: mainTab.selectedKey === v.id ? '' : 'none', height: 'calc(100vh - 40px)' }} >
          {/* <ConnectionItem connection={v} /> */}
        </div>)
      })
    }
  </>);
}
