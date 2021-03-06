import { DetailsList, DetailsListLayoutMode, SelectionMode, Toggle } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';
import { parseClient } from '../utils';

export interface IClientProps {
  connection: Connection
}

export const Client = (props: IClientProps) => {
  const { connection } = props,
    { t } = useTranslation();

  const [autoRefresh, setAutoRefresh] = useState(false),
    [error, setError] = useState<Error | undefined>(),
    [items, setItems] = useState<Array<any>>([]);

  const clientList = useCallback((id) => {
    setError(undefined);
    executeCommand<Array<string>>({ id: id, commands: [['CLIENT', 'LIST']] }).then((ret) => {
      if (!ret || !ret.length) return;
      setItems(parseClient(ret[0]))
    })
      .catch(err => setError(err));
  }, [])

  useEffect(() => {
    clientList(connection.id);
    const timer = setInterval(() => {
      clientList(connection.id);
    }, 2000);
    if (!autoRefresh) {
      clearInterval(timer);
    };
    return () => clearInterval(timer);
  }, [connection, autoRefresh, clientList])

  return (<>
    <ErrorMessageBar error={error} />
    <Toggle label={t('Auto refresh')} checked={autoRefresh} onChange={() => { setAutoRefresh(!autoRefresh) }} />
    <DetailsList
      compact={true}
      items={items}
      columns={[
        { key: 'addr', name: t('Client address'), fieldName: 'addr', minWidth: 100 },
        { key: 'age', name: t('Connected Duration (sec)'), fieldName: 'age', minWidth: 180 },
        { key: 'idle', name: t('Idle Duration (sec)'), fieldName: 'idle', minWidth: 160 },
        { key: 'flags', name: t('Flags'), fieldName: 'flags', minWidth: 40 },
        { key: 'db', name: t('Current Database'), fieldName: 'db', minWidth: 90 },
      ]}
      selectionMode={SelectionMode.none}
      layoutMode={DetailsListLayoutMode.justified}
    />
  </>)
}