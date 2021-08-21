import { DetailsList, DetailsListLayoutMode, Label, Pivot, PivotItem, SelectionMode, Stack, Text, Toggle, useTheme } from '@fluentui/react';
import dayjs from 'dayjs';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts';
import { Connection, executeCommand } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';
import { parseInfo } from '../utils';
import { Client } from './Client';
import { PubSub } from './PubSub';
import { SlowLog } from './SlowLog';

export interface IInfoProps {
  connection: Connection
  info: any
}

export const Info = (props: IInfoProps) => {

  const theme = useTheme(),
    { t } = useTranslation();

  const [info, setInfo] = useState<any>(props.info || {}),
    [error, setError] = useState<Error | undefined>(),
    [autoRefresh, setAutoRefresh] = useState(false),
    [memoryData, setMemoryData] = useState<Array<{ time: string, used_memory: number, used_memory_lua: number, used_memory_rss: number, used_memory_peak: number }>>([]);

  const getInfo = useCallback((id) => {
    setError(undefined);
    executeCommand<Array<string>>({ id: id, commands: [["INFO"]] }).then((ret) => {
      if (!ret || !ret.length) return;
      const retInfo = parseInfo(ret[0]);
      setInfo(retInfo);
      setMemoryData(md => {
        md.push({
          time: dayjs().format("HH:mm:ss"),
          used_memory: Math.round(retInfo.Memory.used_memory / 1024 / 1024 * 100) / 100,
          used_memory_lua: Math.round(retInfo.Memory.used_memory_lua / 1024 / 1024 * 100) / 100,
          used_memory_rss: Math.round(retInfo.Memory.used_memory_rss / 1024 / 1024 * 100) / 100,
          used_memory_peak: Math.round(retInfo.Memory.used_memory_peak / 1024 / 1024 * 100) / 100
        })
        return [...md]
      });
    })
      .catch(err => setError(err));
  }, [setMemoryData])

  useEffect(() => {
    const timer = setInterval(() => {
      getInfo(props.connection.id);
    }, 2000);
    if (!autoRefresh) {
      clearInterval(timer);
    };
    return () => clearInterval(timer);
  }, [props.connection, getInfo, autoRefresh])

  return (
    <Stack tokens={{ padding: 10, childrenGap: 10 }} style={{ minWidth: 450, height: "100%", overflow: 'auto' }}>
      <Stack horizontal horizontalAlign="space-between">
        <Stack horizontalAlign="center">
          <Label>{t('Redis version')}</Label>
          <Text>{info?.Server.redis_version}</Text>
        </Stack>
        <Stack horizontalAlign="center">
          <Label>{t('Used memory')}</Label>
          <Text>{info?.Memory.used_memory_human}</Text>
        </Stack>
        <Stack horizontalAlign="center">
          <Label>{t('Clients')}</Label>
          <Text>{info?.Clients.connected_clients}</Text>
        </Stack>
        <Stack horizontalAlign="center">
          <Label>{t('Commands processed')}</Label>
          <Text>{info?.Stats.total_commands_processed}</Text>
        </Stack>
        <Stack horizontalAlign="center">
          <Label>{t('Uptime in days')}</Label>
          <Text>{info?.Server.uptime_in_days}</Text>
        </Stack>
        <Stack horizontal horizontalAlign="center" verticalAlign="center">
          <Toggle label={t('Auto refresh')} onChange={() => { setAutoRefresh(!autoRefresh) }} />
        </Stack>
      </Stack>

      <ErrorMessageBar error={error} />

      <Stack.Item grow={1}>
        <Pivot linkFormat="tabs" style={{ height: '100%' }} styles={{ link: { height: 32, lineHeight: '32px' }, itemContainer: { height: 'calc(100% - 32px)' } }}>

          <PivotItem alwaysRender style={{ height: '100%' }} itemKey="memory" headerText={t('Memory usage')}>
            <ResponsiveContainer width="100%" height="100%">
              <LineChart data={memoryData} margin={{ top: 20, right: 10 }}>
                <CartesianGrid strokeDasharray="4 4" />
                <XAxis dataKey="time" />
                <YAxis label={{ value: "MB", angle: -90, position: 'insideLeft' }} />
                <Tooltip />
                <Legend />
                <Line type="basis" dataKey="used_memory" stroke={theme.palette.themePrimary} activeDot={true} dot={false} />
                <Line type="monotone" dataKey="used_memory_lua" stroke={theme.palette.orange} activeDot={true} dot={false} />
                <Line type="monotone" dataKey="used_memory_rss" stroke={theme.palette.green} activeDot={true} dot={false} />
                <Line type="monotone" dataKey="used_memory_peak" stroke={theme.palette.purple} activeDot={true} dot={false} />
              </LineChart>
            </ResponsiveContainer>
          </PivotItem>

          <PivotItem alwaysRender key="info" itemKey="info" headerText={t('Service Info')}>
            <Pivot style={{ paddingTop: 10 }} styles={{ link: { height: 24, lineHeight: '24px' } }}>
              {info && Object.keys(info).map(v => {
                var items = Object.keys(info[v]).map(i => {
                  return { key: i, name: i, value: info[v][i] }
                }) || [];
                return (<PivotItem alwaysRender key={v} headerText={v}>
                  <DetailsList
                    isHeaderVisible={false}
                    compact={true}
                    items={items}
                    columns={[
                      { key: 'name', name: t('Property'), fieldName: 'name', minWidth: 200, maxWidth: 200 },
                      { key: 'value', name: t('Value'), fieldName: 'value', minWidth: 100 }
                    ]}
                    selectionMode={SelectionMode.none}
                    layoutMode={DetailsListLayoutMode.justified}
                  />
                </PivotItem>)
              })}
            </Pivot>
          </PivotItem>
          <PivotItem alwaysRender itemKey="slowlog" headerText={t('Slow log')}>
            <SlowLog connection={props.connection}></SlowLog>
          </PivotItem>
          <PivotItem alwaysRender itemKey="clients" headerText={t('Clients')}>
            <Client connection={props.connection} />
          </PivotItem>
          <PivotItem alwaysRender itemKey="pub/sub" headerText={t('Push/subscribe channel')}>
            <PubSub connection={props.connection}></PubSub>
          </PivotItem>
        </Pivot>
      </Stack.Item>
    </Stack>
  )
}