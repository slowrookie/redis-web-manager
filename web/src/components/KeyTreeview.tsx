import { List, useTheme, Stack, ActionButton } from '@fluentui/react';
import * as _ from 'lodash';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';

export interface IKeyGroupedListProps {
  connection: Connection
  db: number
  keys: Array<string>,
  onSelectedKey: (type: string, keyName: string) => void,
  onDeletedKey: (keyName: string) => void,
}

export const KeyTreeview = (props: IKeyGroupedListProps) => {
  const theme = useTheme(),
    { t } = useTranslation(),
    { connection, db, keys, onSelectedKey, onDeletedKey } = props,
    [keyTree, setKeyTree] = useState<any>({});
  ;

  const [error, setError] = useState<Error>(),
    [selectedKey, setSelectedKey] = useState<string>();

  useEffect(() => {
    const keyArray = keys.map(k => k.split(connection.namespaceSeparator));
    let keyTreeArray: Array<any> = [];
    keyArray.forEach((k, kIndex) => {
      let keyNode: any = {};
      for (let index = k.length - 1; index >= 0; index--) {
        if (index === k.length - 1) {
          keyNode[k[index]] = { key: keys[kIndex] };
        } else {
          let parent: any = {}
          parent[k[index]] = { ...keyNode }
          keyNode = parent;
        }
      }
      keyTreeArray.push(keyNode);
    })
    setKeyTree(keyTreeArray.reduce((p, c) => _.defaultsDeep(p, c), {}));
  }, [keys])

  const treeView = (treeObj: any, deep: number) => {
    console.log(treeObj);

    return <List items={Object.keys(treeObj)} onRenderCell={(item, i) => {
      if (!item) return (<></>);
      return (<>
        {item === 'key' && <Stack horizontal verticalAlign="center" style={{ background: item === selectedKey ? theme.palette.neutralLight : '' }}>
          <Stack.Item styles={{ root: { width: 10 * deep } }}><span /></Stack.Item>
          <Stack.Item grow={1}>
            <ActionButton
              iconProps={{ iconName: 'permissions', style: { height: 'auto', color: theme.palette.yellow } }}
              style={{ width: '100%', height: 28 }}
              onClick={() => { }}
              text={treeObj.key} />
          </Stack.Item>
        </Stack>}

        {/* group */}
        {item !== 'key' && <Stack horizontal verticalAlign="center" style={{ background: item === selectedKey ? theme.palette.neutralLight : '' }}>
          <Stack.Item styles={{ root: { width: 10 * deep } }}><span /></Stack.Item>
          <Stack.Item grow={1}>
            <ActionButton
              iconProps={{ iconName: 'permissions', style: { height: 'auto', color: theme.palette.yellow } }}
              style={{ width: '100%', height: 28 }}
              onClick={() => { }}
              text={item} />
          </Stack.Item>
        </Stack>}

        {treeObj[item] && _.isObject(treeObj[item]) && treeView(treeObj[item], deep + 1)}
      </>);
    }} />
  }

  return (<>
    {/* error */}
    <ErrorMessageBar error={error} />
    {/* {tree} */}
    {treeView(keyTree, 0)}
  </>)
}