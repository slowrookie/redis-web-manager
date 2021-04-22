import { DetailsList, DetailsListLayoutMode, SelectionMode } from '@fluentui/react';
import React from 'react';

export interface INetworkProps {
    'protected-mod'?: 'no' | 'yes'
    unixsocket?: string
    port?: number
    'tcp-keepalive'?: number
    unixsocketperm?: string
    bind?: string
}

export const Network = (props: INetworkProps) => {
    return (
        <DetailsList
            compact={true}
            items={items}
            columns={[
                { key: 'addr', name: t('Client address'), fieldName: 'addr', minWidth: 100 },
                { key: 'age', name: t('Connected Duration (sec)'), fieldName: 'age', minWidth: 180 },
            ]}
            isHeaderVisible={false}
            selectionMode={SelectionMode.none}
            layoutMode={DetailsListLayoutMode.justified}
        />
    )
}