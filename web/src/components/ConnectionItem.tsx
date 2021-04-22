import { AnimationClassNames, Depths, INavLink, INavLinkGroup, Nav, Stack } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, openConnection } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Loading } from './common/Loading';
import { Console } from './Console';
import { Database, IDatabase } from './Database';
import { Info } from './info/Info';
import { parseInfo } from './utils';

export interface ConnectionItemProps {
  connection: Connection
}

export const ConnectionItem = (props: ConnectionItemProps) => {

  const { connection } = props;

  const [info, setInfo] = useState(),
    [databases, setDatabase] = useState<Array<IDatabase>>([{ db: 0, dbsize: 0 }]),
    [error, setError] = useState<string>(),
    [loading, setLoading] = useState(false),
    [selectedKey, setSelectedKey] = useState<string | undefined>('serverInfo'),
    { t } = useTranslation();

  useEffect(() => {
    setError('');
    setLoading(true);
    openConnection(connection.id).then(ret => {
      console.log(ret);
      var info: any = parseInfo(ret.info);
      setInfo(info);
      setDatabase([...Array(ret.database[1])].map((_, i) => {
        var reg = /[1-9][0-9]*/
        var keys = (info.Keyspace && info.Keyspace[`db${i}`] && info.Keyspace[`db${i}`].match(reg)[0]) || 0;
        return { db: i, dbsize: keys }
      }));
    })
      .catch(err => setError(err))
      .finally(() => { setLoading(false) });
  }, [connection.id, setLoading]);

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
        { name: t('Configuration'), key: 'configuration', url: '' },
      ]
    }
  ];


  return (
    <Stack horizontal style={{ height: '100%', position: 'relative' }}>

      <Loading loading={loading}></Loading>

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

      {/* error */}
      <Stack.Item grow={1} style={{ display: error ? 'block' : 'none' }}>
        <ErrorMessageBar error={error} />
      </Stack.Item>

    </Stack>
  )
}