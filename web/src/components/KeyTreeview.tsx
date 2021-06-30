import { useTheme, GroupedList, IGroup, DetailsRow, IToggleStyles, SelectionMode, IColumn } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Connection } from '../services/connection.service';

export interface IKeyGroupedListProps {
  connection: Connection
  db: number
  keys: Array<string>,
  onSelectedKey: (type: string, keyName: string) => void,
  onDeletedKey: (keyName: string) => void,
}

interface IKeyTreeviewNode {
  key: string,
  children?: Array<IKeyTreeviewNode>,
  level: number
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

  useEffect(() => {
    const tree: IKeyTreeviewNode[] = [];
    const keyArray = keys.map(k => k.split(connection.namespaceSeparator));
    console.log(keyArray);

    keyArray.forEach(k => {
      if (k.length === 1) {
        tree.push({ key: k[0], level: 0 })
        return;
      }
      // ???
      k.forEach((ksp, index) => {
        if (tree.find(v => v.key === ksp)) {

        } else {

        }
      })
    })
    console.log(tree);
  }, [keys])

  return (<>
    {/* error */}
    <ErrorMessageBar error={error} />
    {/* {tree} */}

  </>)
}