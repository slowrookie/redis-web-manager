import { DefaultButton, Depths, IButtonStyles, Icon, Pivot, PivotItem, PrimaryButton, Stack, TooltipHost } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { Connection } from '../services/connection.service';
import { AppSettings } from './AppSettings';
import { ConnectionItem, IConnectionItemProps } from './ConnectionItem';
import { ConnectionList } from './ConnectionList';
export interface IMainTabProps {
  onChangeLanguage?: (language: string) => void
  onChangeTheme?: (theme: string) => void
}

interface IMainTab {
  connectionItems: Array<IConnectionItemProps>,
  selectedKey: string | undefined,
  showConnectionList: boolean,
}

const headerButtonStyles: IButtonStyles = {
  root: {
    minWidth: 'auto',
    borderRadius: 0,
    padding: '0 8px',
    height: 42,
    border: 0
  }
}

export const MainTab = (props: IMainTabProps) => {
  const
    [mainTab, setMainTab] = useState<IMainTab>({ connectionItems: [], selectedKey: '', showConnectionList: true });

  useEffect(() => {
    if (!mainTab.connectionItems.length) {
      setMainTab(m => {
        return { ...m, selectedKey: '', showConnectionList: true }
      })
    }
  }, [mainTab.connectionItems])

  const handleConnectionClick = (v: Connection) => {
    const exists = mainTab.connectionItems.filter(item => item.connection.id === v.id);
    const connections = mainTab.connectionItems;
    !exists.length && connections.push({ connection: v });
    setMainTab({ ...mainTab, connectionItems: [...connections], selectedKey: v.id, showConnectionList: false });
  }

  return (<>
    <Stack horizontal style={{  height: 42, boxShadow: Depths.depth8 }}>
      {/* home */}
      <TooltipHost content={'连接管理'}>
        <DefaultButton 
          iconProps={{ iconName: 'ProductList' }} 
          styles={headerButtonStyles} 
          onClick={() => setMainTab({ ...mainTab, selectedKey: '', showConnectionList: true })} />
      </TooltipHost>
      {/* tabs */}
      <Stack.Item grow={1}>
        <Pivot overflowBehavior='menu' headersOnly styles={{ link: { height: 42, lineHeight: '42px' } }}
          linkFormat='links'
          selectedKey={mainTab.selectedKey}
          onLinkClick={item => {
            setMainTab({ ...mainTab, selectedKey: item?.props.itemKey, showConnectionList: false })
          }} >
          {mainTab.connectionItems.map((v, index) => {
            return <PivotItem alwaysRender={true} headerText={v.connection.name} key={v.connection.id} itemKey={v.connection.id}
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
      <AppSettings {...props} />
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
