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
  return (<>
   
    </>
  )
}