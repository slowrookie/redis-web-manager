import { List, useTheme, Stack, ActionButton, GroupedList, Icon, Nav, INavLinkGroup, INavLink } from '@fluentui/react';
import * as _ from 'lodash';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import Tree, { TreeNode } from 'rc-tree';
import 'rc-tree/assets/index.css';

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
    [keyTree, setKeyTree] = useState<Array<any>>([]),
    [navLinkGroups, setNavLinkGroups] = useState<INavLinkGroup[]>([]);
  ;

  const [error, setError] = useState<Error>(),
    [selectedKey, setSelectedKey] = useState<string>();

  useEffect(() => {
    const keyArray = keys.map(k => k.split(connection.namespaceSeparator));
    let keyTreeArray: Array<any> = [];
    keyArray.forEach((k, kIndex) => {
      let currentNode: any = {};
      for (let index = k.length - 1; index >= 0; index--) {
        if (index === k.length - 1) {
          currentNode[k[index]] = { key: keys[kIndex] };
        } else {
          let parent: any = {}
          parent[k[index]] = { ...currentNode }
          currentNode = parent;
        }
      }
      keyTreeArray.push(currentNode);
    })

    const keyTreeObject = keyTreeArray.reduce((p, c) => _.defaultsDeep(p, c), {});

    const convertToRcTree = (obj: any, path: string): Array<any> => {
      return _.map(obj, (value: any, key: string) => {
        if (_.isString(value)) {
          return { key: `${path}${value}`, title: value };
        } else {
          if (!!value['key'] && Object.keys(value).length === 1) {
            return { key: `${path}${value['key']}`, title: value['key'] };
          }
          const children = convertToRcTree(value, `${path}${key}-`);
          return { key: `${path}${key}`, title: `${key} (${children.length})`, children };
        }
      })
    }

    const convertToNavGroup = (obj: any, path: string): Array<any> => {
      return _.map(obj, (value: any, key: string) => {
        if (_.isString(value)) {
          return { key: `${path}${value}`, name: value };
        } else {
          if (!!value['key'] && Object.keys(value).length === 1) {
            return { key: `${path}${value['key']}`, name: value['key'] };
          }
          const links = convertToNavGroup(value, `${path}${key}-`);
          return { key: `${path}${key}`, name: `${key} (${links.length})`, links };
        }
      })
    }

    setKeyTree(convertToRcTree(keyTreeObject, ""));

    console.log(convertToNavGroup(keyTreeObject, ""));

    setNavLinkGroups([{ links: convertToNavGroup(keyTreeObject, "") }]);
  }, [keys])

  return (<>
    {/* error */}
    <ErrorMessageBar error={error} />
    {/* {tree} */}
    {/* {
      <Tree defaultExpandAll={false} treeData={keyTree} showLine showIcon={false} switcherIcon={(obj) => {
        if (obj.isLeaf) {
          return <Icon iconName={"permissions"} />;
        }
        return "";
        // return obj.expanded ? <Icon iconName={"chevrondown"} /> : <Icon iconName={"chevronright"} />
      }}
      // onSelect={this.onSelect}
      />
    } */}

    <Nav
      styles={{
        link: { 
          height: 24
        },
        chevronButton: {
          height: 24,
          lineHeight: 24
        },
        chevronIcon: {
          height: 24,
          lineHeight: 24
        }
      }}
      // onLinkClick={() => { }}
      // selectedKey={undefined}
      groups={navLinkGroups}
    />

  </>)
}