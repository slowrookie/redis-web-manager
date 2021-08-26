import { DetailsList, DetailsListLayoutMode, SelectionMode, Toggle } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';

export interface ISlowLogProps {
    connection: Connection
}

export const SlowLog = (props: ISlowLogProps) => {
    const { t } = useTranslation();

    const [autoRefresh, setAutoRefresh] = useState(false),
        [error, setError] = useState<Error | undefined>(),
        [items, setItems] = useState<Array<{ time: string, duration: string, command: string, addr: string }>>([]);

    const slowlogList = useCallback((id) => {
        setError(undefined);
        executeCommand<Array<Array<any>>>({ id: id, commands: [['SLOWLOG', 'GET']] })
            .then((ret) => {
                if (!ret || !ret.length) return;
                var logs = ret[0].map(v => {
                    return {
                        time: v[1],
                        duration: v[2],
                        command: v[3].join(' '),
                        addr: v[4],
                    }
                })
                setItems(logs)
            })
            .catch(err => setError(err));
    }, [])

    useEffect(() => {
        slowlogList(props.connection.id);
        const timer = setInterval(() => {
            slowlogList(props.connection.id);
        }, 2000);
        if (!autoRefresh) {
            clearInterval(timer);
        };
        return () => {
            clearInterval(timer)
        };
    }, [props.connection, autoRefresh, slowlogList])

    return (<>
        <ErrorMessageBar error={error} />
        <Toggle label={t('Auto refresh')} checked={autoRefresh} onChange={() => { setAutoRefresh(!autoRefresh) }} />
        <DetailsList
            compact={true}
            items={items}
            columns={[
                { key: 'addr', name: t('Client address'), fieldName: 'addr', minWidth: 100, isResizable: true },
                { key: 'command', name: t('Command'), fieldName: 'command', minWidth: 150, isResizable: true },
                { key: 'time', name: t('Dealt with'), fieldName: 'time', minWidth: 80, maxWidth: 100 },
                { key: 'duration', name: t('Duration (μs)'), fieldName: 'duration', minWidth: 60, maxWidth: 60 },
            ]}
            selectionMode={SelectionMode.none}
            layoutMode={DetailsListLayoutMode.justified}
        />
    </>)
}