import { Depths, Icon, Pivot, PivotItem, Stack } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { Connection } from '../services/connection.service';
import { AppSettings } from './settings/AppSettings';
import { ConnectionItem, IConnectionItemProps } from './ConnectionItem';
import { ConnectionList } from './ConnectionList';
import { Config } from '../services/config.service';
export interface IMainTabProps {
  config: Config
}

interface IMainTab {
  connectionItems: Array<IConnectionItemProps>,
  selectedKey: string | undefined,
  showConnectionList: boolean,
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
    <Pivot overflowBehavior='menu' styles={{
      root: {boxShadow: Depths.depth8},
      link: { height: 42, lineHeight: '42px', minWidth: 42 }
    }}
      linkFormat='tabs'
      selectedKey={mainTab.selectedKey}
      onLinkClick={item => {
        if (item?.props.itemKey === "Home") {
          setMainTab({ ...mainTab, selectedKey: 'Home', showConnectionList: true });
          return;
        }
        setMainTab({ ...mainTab, selectedKey: item?.props.itemKey, showConnectionList: false })
      }} >

      {/* Home */}
      <PivotItem itemIcon="ProductList" key="Home" itemKey="Home"
        onRenderItemLink={(link, defaultRenderer: any) => (
          <Stack horizontal tokens={{ childrenGap: 10 }}>
            {defaultRenderer(link)}
            {/* Settings */}
            <AppSettings {...props} />
          </Stack>
        )}>
        <div style={{ height: 'calc(100vh - 42px)', overflowY: 'auto', overflowX: 'hidden' }} >
          <ConnectionList onConnectionClick={handleConnectionClick} />
        </div>
      </PivotItem>

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
        >
          <div style={{ height: 'calc(100vh - 42px)', overflow: 'hidden' }} >
            <ConnectionItem {...v} />
          </div>
        </PivotItem>
      })}
    </Pivot>
  </>);
}
