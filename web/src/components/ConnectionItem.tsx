import { Depths, DirectionalHint, IButtonStyles, IconButton, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Loading } from './common/Loading';
import { DatabaseConfiguration } from './configuration/DatabaseConfiguration';
import { Console } from './Console';
import { Database, IDatabase } from './Database';
import { Info } from './info/Info';
import { parseInfo } from './utils';

export interface IConnectionItemProps {
  connection: Connection,
}

const buttonStyles: IButtonStyles = {
  root: { width: 42, height: 42 }
}

export const ConnectionItem = (props: IConnectionItemProps) => {

  const { connection } = props;

  const [selectedKey, setSelectedKey] = useState<string | undefined>('serverInfo'),
    [info, setInfo] = useState<any>(),
    [databases, setDatabases] = useState<Array<IDatabase>>([]),
    [error, setError] = useState<Error>(),
    [loading, setLoading] = useState<boolean>(false),
    { t } = useTranslation(),
    theme = useTheme();

  useEffect(() => {
    setLoading(true);
    executeCommand({ id: connection.id, commands: [['CONFIG', 'GET', 'DATABASES'], ['INFO']] })
      .then((ret: any) => {
        let info: any = parseInfo(ret[1]);
        let databases = ([...Array(ret[0].length ? Number(ret[0][1]) : 1)].map((_, i) => {
          var reg = /[1-9][0-9]*/
          var keys = (info.Keyspace && info.Keyspace[`db${i}`] && info.Keyspace[`db${i}`].match(reg)[0]) || 0;
          return { db: i, dbsize: keys };
        }));
        setInfo(info);
        setDatabases(databases);
        console.log(ret);
      })
      .catch(err => setError(err))
      .finally(() => { setLoading(false) });
  }, [connection])

  const selectedStyle = (key: string) => {
    return selectedKey === key ? { borderRight: `2px solid ${theme.palette.themePrimary}` } : { borderRight: '0px' };
  }

  return (<>
    <ErrorMessageBar error={error}></ErrorMessageBar>
    <Loading loading={loading} />
    <Stack horizontal style={{ height: '100%', position: 'relative' }}>

      <Stack style={{ height: '100%', boxShadow: Depths.depth8 }}>
        {/* server info */}
        <TooltipHost content={t('Server Info')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('serverInfo')} iconProps={{ iconName: 'ServerProcesses' }} onClick={() => {
            setSelectedKey('serverInfo');
          }} />
        </TooltipHost>

        <TooltipHost content={t('Database')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('database')} iconProps={{ iconName: 'Database' }} onClick={() => {
            setSelectedKey('database');
          }} />
        </TooltipHost>

        <TooltipHost content={t('CLI')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('cli')} iconProps={{ iconName: 'Code' }} onClick={() => {
            setSelectedKey('cli');
          }} />
        </TooltipHost>

        <Stack.Item grow={1}><span /></Stack.Item>

        <TooltipHost content={t('Configuration')} directionalHint={DirectionalHint.rightCenter}>
          <IconButton styles={buttonStyles} style={selectedStyle('databaseConfig')} iconProps={{ iconName: 'DataManagementSettings' }} onClick={() => {
            setSelectedKey('databaseConfig');
          }} />
        </TooltipHost>

      </Stack>

      <Stack.Item grow={1} style={{ display: selectedKey === 'serverInfo' ? 'block' : 'none' }}>
        {info && <Info {...props} info={info} />}
      </Stack.Item>

      <Stack.Item grow={1} style={{ display: selectedKey === 'database' ? 'block' : 'none' }}>
        {databases.length && <Database {...props} databases={databases} />}
      </Stack.Item>

      <Stack.Item grow={1} style={{ display: selectedKey === 'cli' ? 'block' : 'none' }}>
        <Console {...props} />
      </Stack.Item>

      <Stack.Item grow={1} style={{ display: selectedKey === 'databaseConfig' ? 'block' : 'none' }}>
        <DatabaseConfiguration {...props} />
      </Stack.Item>

    </Stack>
  </>)
}