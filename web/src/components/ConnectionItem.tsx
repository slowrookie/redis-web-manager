import { AnimationClassNames, Depths, INavLink, INavLinkGroup, Nav, Stack } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, openConnection } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Loading } from './common/Loading';
import { DatabaseConfiguration } from './configuration/DatabaseConfiguration';
import { Console } from './Console';
import { Database, IDatabase } from './Database';
import { Info } from './info/Info';
import { parseInfo } from './utils';

export interface IConnectionItemProps {
  connection: Connection,
  info: any,
  databases: Array<IDatabase>
}

export const ConnectionItem = (props: IConnectionItemProps) => {
  const { connection, info, databases } = props;

  const [selectedKey, setSelectedKey] = useState<string | undefined>('serverInfo'),
    { t } = useTranslation();

  const navLinkGroups: INavLinkGroup[] = [
    {
      links: [
        {
          name: t('Overview'), key: 'overview', isExpanded: true, links: [
            { name: t('Server Info'), key: 'serverInfo', url: '' },
            { name: t('Database'), key: 'database', url: '' },
            { name: t('CLI'), key: 'cli', url: '' },
          ],
          url: ''
        },
        { name: t('Configuration'), key: 'databaseConfig', url: '' },
      ]
    }
  ];

  return (
    <Stack horizontal style={{ height: '100%', position: 'relative' }}>

      <Nav selectedKey={selectedKey} styles={{ root: { width: 150, height: '100%', boxShadow: Depths.depth8 } }}
        groups={navLinkGroups}
        onLinkClick={(ev?: React.MouseEvent<HTMLElement>, item?: INavLink) => {
          if (item?.links && item?.links.length) return;
          setSelectedKey(item?.key);
        }}
      />

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'serverInfo' ? 'block' : 'none' }}>
        {info && <Info {...props} info={info} />}
      </Stack.Item>

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'database' ? 'block' : 'none' }}>
        {databases.length && <Database {...props} databases={databases} />}
      </Stack.Item>

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'cli' ? 'block' : 'none' }}>
        <Console {...props} />
      </Stack.Item>

      <Stack.Item grow={1}
        className={selectedKey ? AnimationClassNames.fadeIn100 : AnimationClassNames.fadeOut100}
        style={{ display: selectedKey === 'databaseConfig' ? 'block' : 'none' }}>
        <DatabaseConfiguration {...props} />
      </Stack.Item>

    </Stack>
  )
}