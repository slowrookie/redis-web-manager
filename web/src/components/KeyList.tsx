import { CheckboxVisibility, DetailsList, DetailsListLayoutMode, Selection, SelectionMode, TooltipHost } from '@fluentui/react';
import React, { useState } from 'react';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { ComponentType } from './utils';

export interface IKeyListProps {
  connection: Connection
  db: number
  keys: Array<string>,
  onSelectedKey: (type: string, keyName: string) => void,
}

export const KeyList = (props: IKeyListProps) => {
  const { connection, db, keys, onSelectedKey } = props;
  const [error, setError] = useState<Error>();

  const selection = new Selection({
    onSelectionChanged: () => {
      if (selection.getSelection().length) {
        const v = selection.getSelection()[0] as string;
        setError(undefined);
        executeCommand<Array<any>>({ id: connection.id, commands: [['SELECT', db], ['TYPE', v]] })
          .then((ret) => {
            if (!ret || !ret.length) return;
            onSelectedKey(ComponentType[ret[1].toUpperCase()], v);
          })
          .catch(err => setError(err))
          .finally(() => { });
      }
    }
  })

  return (<>
    {/* error */}
    <ErrorMessageBar error={error}></ErrorMessageBar>
    {/* {list} */}
    <DetailsList compact
      layoutMode={DetailsListLayoutMode.justified}
      selection={selection}
      selectionMode={SelectionMode.single}
      checkboxVisibility={CheckboxVisibility.hidden}
      enterModalSelectionOnTouch={true}
      selectionPreservedOnEmptyClick={true}
      isHeaderVisible={false}
      columns={[
        {
          key: 'Key', name: 'Key', minWidth: 190, onRender: (item) => {
            return <TooltipHost content={item}>
              {item}
            </TooltipHost>
          }
        }]}
      items={[...keys]} />
  </>)
}