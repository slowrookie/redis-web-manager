import { INavLink, INavLinkGroup, Nav, Stack } from '@fluentui/react';
import { Depths } from '@fluentui/theme';
import React, { useEffect, useState } from 'react';
import { Connection, executeCommand } from '../../services/connection.service';
import { Network } from './Network';

export interface IDatabaseConfigProps {
  connection: Connection
}

export const DatabaseConfiguration = (props: IDatabaseConfigProps) => {

  const [configs, setConfigs] = useState<{ [index: string]: any }>(),
    [selectedKey, setSelectedKey] = useState<string | undefined>('Network');

  useEffect(() => {
    // executeCommand<Array<any>>({ id: props.connection.id, commands: [['CONFIG', 'GET', '*']] })
    //   .then(ret => {
    //     if (!ret || !ret.length) return;
    //     console.log(ret)
    //     const length: number = ret[0].length / 2;
    //     const kv: { [index: string]: any } = {};
    //     [...Array(length)].forEach((_, i) => {
    //       const index = 2 * i;
    //       kv[ret[0][index]] = ret[0][index + 1];
    //     });
    //     console.log(kv);
    //     setConfigs({ ...kv })
    //   })
  }, []);

  const configGroups: INavLinkGroup[] = [
    {
      links: [
        { name: 'Network', key: 'Network', isExpanded: true, url: '' },
        { name: 'General', key: 'General', url: '' },
        { name: 'Snapshotting', key: 'Snapshotting', url: '' },
        { name: 'Replication', key: 'REPLICATION', url: '' },
        { name: 'Security', key: 'SECURITY', url: '' },
        { name: 'Clients', key: 'CLIENTS', url: '' },
        { name: 'Memory management', key: 'MEMORY MANAGEMENT', url: '' },
        { name: 'Lazy freeing', key: 'LAZY FREEING', url: '' },
        { name: 'Append only mode', key: 'APPEND ONLY MODE', url: '' },
        { name: 'Lua scripting', key: 'LUA SCRIPTING', url: '' },
        { name: 'Redis cluster', key: 'REDIS CLUSTER', url: '' },
        { name: 'Cluster docker/nat', key: 'CLUSTER DOCKER/NAT', url: '' },
        { name: 'Slowlog', key: 'SLOWLOG', url: '' },
        { name: 'Lntency monitor', key: 'LNTENCY MONITOR', url: '' },
        { name: 'Event notification', key: 'EVENT NOTIFICATION', url: '' },
        { name: 'Defragmentation', key: 'DEFRAGMENTATION', url: '' },
        { name: 'Advanced config', key: 'ADVANCED CONFIG', url: '' },
      ]
    }
  ];

  const getConfigComponent = () => {
    switch (selectedKey) {
      case 'Network':
        return <Network {...configs} />
    
      default:
        break;
    }
  }

  return (<Stack horizontal tokens={{ padding: 5 }}>
    <Nav
      selectedKey={selectedKey}
      styles={{
        groupContent: { marginBottom: 0 },
        root: { width: 180, boxShadow: Depths.depth8 },
        link: { '&:after': { border: 0 }, lineHeight: '32px', height: 32, fontSize: 12 }
      }}
      groups={configGroups}
      onLinkClick={(ev?: React.MouseEvent<HTMLElement>, item?: INavLink) => {
        if (item?.links && item?.links.length) return;
        setSelectedKey(item?.key);
      }}
    />

    

  </Stack>)
}