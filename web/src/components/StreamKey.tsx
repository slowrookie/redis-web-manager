import {
  CheckboxVisibility, DetailsListLayoutMode,
  DetailsRow,
  DialogType,
  IconButton,
  IDetailsRowStyles,
  Selection,
  SelectionMode, ShimmeredDetailsList, Stack, TooltipHost, useTheme
} from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { parseArrayToObject } from './utils';
import { ConfirmButton } from './common/ConfirmButton';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { KeyHeader } from './KeyHeader';
import { StreamKeyPanel } from './panel/StreamKeyPanel';
import { StreamKeySearch } from './StreamKeySearch';


export interface IStreamKeyProps {
  connection: Connection
  db: number
  keyName: string,
  component: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface IStreamKeyItem {
  row: number,
  id: string,
  value: string
}

export interface IStreamKey {
  keyName: string
  TTL: number
  length: number
  values: Array<IStreamKeyItem>
  memoryUsage: number
  firstTimestamp: number
  lastTimestamp: number
}

const defaultStreamKey: IStreamKey = {
  keyName: '',
  TTL: -1,
  length: 0,
  values: [],
  memoryUsage: 0,
  firstTimestamp: 0,
  lastTimestamp: 0
}

export interface IStreamKeySearch {
  start: string | number
  end: string | number
  count: number,
  asc: boolean
}

export const StreamKey = (props: IStreamKeyProps) => {

  const { connection, db, keyName, component, onKeyNameChanged } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyProps, setKeyProps] = useState({ ...defaultStreamKey, keyName }),
    [selectedValue, setSelectedValue] = useState<IStreamKeyItem | undefined>(),
    [error, setError] = useState<Error>(),
    [showEditPanel, setShowEditPanel] = useState(false),
    [search, setSearch] = useState<IStreamKeySearch>({ start: '-', end: '+', count: connection.dataScanLimit, asc: true }),
    [items, setItems] = useState<Array<IStreamKeyItem>>([]);

  const load = useCallback(() => {
    if (!search.start || !search.end || !search.count) return;
    setError(undefined);
    const commands = [
      ['SELECT', db],
      ['TTL', keyProps.keyName],
      ['XINFO', 'STREAM', keyProps.keyName]
    ];
    search.asc && commands.push(['XRANGE', keyProps.keyName, search.start, search.end, 'COUNT', search.count])
    !search.asc && commands.push(['XREVRANGE', keyProps.keyName, search.end, search.start, 'COUNT', search.count])
    executeCommand<Array<any>>({ id: connection.id, commands }).then((ret) => {
      if (!ret || !ret.length) return;
      const vs = ret[3].map((v: any, i: number) => {
        const length = v[1].length / 2;
        const rv: any = {};
        [...Array(length)].forEach((_, i) => {
          const index = 2 * i;
          rv[v[1][index]] = v[1][index + 1];
        });
        return { row: i + 1, id: v[0], value: JSON.stringify(rv) }
      });
      const streamInfo = parseArrayToObject(ret[2]);
      setKeyProps(kp => {
        return {
          ...kp, keyName: keyProps.keyName, TTL: ret[1], length: streamInfo.length, values: vs,
          firstTimestamp: Number(Object.keys(streamInfo['first-entry'])[0].split("-")[0]),
          lastTimestamp: Number(Object.keys(streamInfo['last-entry'])[0].split("-")[0])
        };
      })
      console.log(ret, streamInfo);
    })
      .catch(err => setError(err));
  }, [connection.id, db, search, keyProps.keyName]);

  useEffect(() => {
    load();
  }, [load])

  // filter and loaded pagination
  useEffect(() => {
    setItems(keyProps.values);
  }, [keyProps.values])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedValue(selection.getSelection()[0] as IStreamKeyItem);
    }
  });

  const handleDeleteConfirm = () => {
    if (!selectedValue) return;
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['XDEL', keyProps.keyName, selectedValue.id],
      ]
    })
      .then((ret) => {
        setKeyProps({ ...keyProps, values: keyProps.values.filter(v => v.id !== selectedValue.id), length: keyProps.length ? keyProps.length - 1 : keyProps.length });
        setSelectedValue(undefined);
      })
  }

  const handleEditPanelDismiss = () => {
    setShowEditPanel(false);
    setSelectedValue(undefined);
  }

  const handleSave = () => {
    setSearch({ ...search });
  }

  const handleSearch = (start: string | number, end: string | number, count: number) => {
    setSearch({ ...search, start, end, count })
  }

  return (
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 5 }}>
      {/* error */}
      <ErrorMessageBar error={error}></ErrorMessageBar>

      <KeyHeader {...props} keyName={keyProps.keyName} TTL={keyProps.TTL}
        onKeyNameChanged={(oldValue, newValue) => {
          setKeyProps({ ...keyProps, keyName: newValue });
          onKeyNameChanged(oldValue, newValue);
        }}
        onTTLChanged={(v: number) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => { load() }} />

      <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 5 }}>

        <StreamKeySearch id={`stream-key-search-${connection.id}-${db}`}
          length={keyProps.length}
          firstTimestamp={keyProps.firstTimestamp}
          lastTimestamp={keyProps.lastTimestamp}
          defaultCount={connection.dataScanLimit}
          onSearch={handleSearch}
        />

        <TooltipHost content={search.asc ? t('Ascending') : t('Descending')}>
          <IconButton toggle checked={!search.asc} iconProps={{ iconName: search.asc ? 'Ascending' : 'Descending', style: { height: 'auto' } }}
            onClick={() => {
              setSearch({ ...search, asc: !search.asc });
            }} />
        </TooltipHost>

        <TooltipHost content={t('Add')}>
          <IconButton iconProps={{ iconName: 'circleAddition', style: { height: 'auto' } }}
            onClick={() => {
              setShowEditPanel(true);
              setSelectedValue(undefined);
            }} />
        </TooltipHost>
        <TooltipHost content={t('View')}>
          <IconButton iconProps={{ iconName: 'view', style: { height: 'auto' } }}
            disabled={!selectedValue}
            onClick={() => {
              setShowEditPanel(true);
            }} />
        </TooltipHost>
        <ConfirmButton label={t('Delete')} disabled={!selectedValue} dialogContentProps={{
          type: DialogType.close, title: 'Delete this item?',
          subText: `${selectedValue?.row}: ${selectedValue?.value}`
        }} onConfirm={handleDeleteConfirm} />
      </Stack>

      <Stack.Item grow={1} style={{ height: '100%', overflow: 'auto' }}>
        <ShimmeredDetailsList
          selection={selection}
          compact
          isHeaderVisible={false}
          setKey="streamItems"
          items={items}
          columns={[
            { key: 'id', name: 'ID', fieldName: 'id', minWidth: 80, maxWidth: 150, isResizable: true },
            { key: 'value', name: 'Value', fieldName: 'value', minWidth: 80, isResizable: true }
          ]}
          selectionPreservedOnEmptyClick={false}
          checkboxVisibility={CheckboxVisibility.hidden}
          selectionMode={SelectionMode.single}
          layoutMode={DetailsListLayoutMode.justified}
          enableShimmer={!keyProps.values.length}
          onItemInvoked={(item) => { setSelectedValue(item); setShowEditPanel(true) }}
          onRenderRow={props => {
            const detailsRowStyles: Partial<IDetailsRowStyles> = {
              root: {
                backgroundColor: (props && props.itemIndex % 2 === 0) ? theme.palette.neutralLighterAlt : ''
              }
            };
            return props ? <DetailsRow {...props} styles={detailsRowStyles} /> : (<></>)
          }}
        />
      </Stack.Item>

      {/* edit key */}
      <StreamKeyPanel {...props} isOpen={showEditPanel} onlyView={!!selectedValue} keyType={component} keyId={selectedValue ? selectedValue.id : ''}
        keyValue={selectedValue ? selectedValue.value : ''}
        index={selectedValue ? selectedValue.row - 1 : undefined}
        disabledKeyName={true}
        onDismiss={handleEditPanelDismiss}
        onSave={handleSave} />

    </Stack >
  )
}