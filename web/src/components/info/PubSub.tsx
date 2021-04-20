import { DetailsList, DetailsListLayoutMode, MessageBar, MessageBarType, SelectionMode, Toggle } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../../services/connection.service';

export interface IPubSub {
    connection: Connection
}

export const PubSub = (props: IPubSub) => {
    const { t } = useTranslation();

    const [autoRefresh, setAutoRefresh] = useState(true),
        [error, setError] = useState<string>(),
        [items, setItems] = useState<Array<{ name: string }>>([]);

    const pubsubList = useCallback((id) => {
        setError('');
        executeCommand<Array<Array<string>>>({ id: id, commands: [['PUBSUB', 'CHANNELS']] })
            .then((ret) => {
                if (!ret || !ret.length) return;
                var channels = ret[0].map(v => {
                    return { name: v }
                })
                setItems(channels)
            })
            .catch(err => setError(err));
    }, [])

    useEffect(() => {
        pubsubList(props.connection.id);
        const timer = setInterval(() => {
            pubsubList(props.connection.id);
        }, 2000);
        if (!autoRefresh) {
            clearInterval(timer);
        };
        return () => {
            clearInterval(timer)
        };
    }, [props.connection, autoRefresh, pubsubList])

    return (<>
        {error && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.blocked} isMultiline={true} onDismiss={() => { setError(""); }} truncated={true}>
            {error}
        </MessageBar>}
        <Toggle label={t('Auto refresh')} checked={autoRefresh} onChange={() => { setAutoRefresh(!autoRefresh) }} />
        <DetailsList
            compact={true}
            items={items}
            columns={[
                { key: 'name', name: t('Channel'), fieldName: 'name', minWidth: 100 },
            ]}
            selectionMode={SelectionMode.none}
            layoutMode={DetailsListLayoutMode.justified}
        />
    </>)
}