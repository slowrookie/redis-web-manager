import { IColumn, IGroup, useTheme } from '@fluentui/react';
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

interface IKeyTreeviewNode {
  path?: string,
  key?: string,
  children?: Array<IKeyTreeviewNode>,
}

const columns: Array<IColumn> = [{ key: 'name', name: "name", fieldName: "name", minWidth: 80 }];

export const KeyTreeview = (props: IKeyGroupedListProps) => {
  const theme = useTheme(),
    { t } = useTranslation(),
    { connection, db, keys, onSelectedKey, onDeletedKey } = props,
    // [items, setItems] = useState([])
    [keyTree, setKeyTree] = useState<IGroup[]>();
  ;

  const [error, setError] = useState<Error>(),
    [selectedKey, setSelectedKey] = useState<string>();


  const merge = (existsNode: IKeyTreeviewNode, node: IKeyTreeviewNode) => {

  }

  useEffect(() => {
    const tree: IKeyTreeviewNode[] = [];
    console.log(keys);
    const keyArray = keys.map(k => k.split(connection.namespaceSeparator));
    console.log(keyArray);

    let keyTree: Array<any> = [];
    keyArray.forEach((k, kIndex) => {
      let keyNode: any = {};
      for (let index = k.length - 1; index > 0; index--) {
          if (index === k.length - 1) {
            keyNode[k[index]] = {key: keys[kIndex]};
          } else {
            keyNode[k[index]] = keyNode
          }
      }
      keyTree.push(keyNode);
    })

    console.log(keyTree);
    

    // keyArray.forEach((k, kIndex) => {
    //   let node: IKeyTreeviewNode = {};
    //   for (let index = 0; index < k.length; index++) {
    //     if (index == 0) {
    //       node = { key: keys[kIndex] };
    //     } else {
    //       node = { path: k.filter((_, i) => i < (k.length - index)).join(connection.namespaceSeparator), children: [node] };
    //     }
    //   }
    //   tree.push(node);
    // })
    // console.log(tree);

  }, [keys])

  return (<>
    {/* error */}
    <ErrorMessageBar error={error} />
    {/* {tree} */}

  </>)
}