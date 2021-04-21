import {
  CheckboxVisibility, DetailsListLayoutMode,
  DetailsRow,
  DialogType,
  IconButton,
  IDetailsRowStyles,
  SearchBox, Selection,
  SelectionMode, ShimmeredDetailsList, Stack, TooltipHost, useTheme
} from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ConfirmButton } from './common/ConfirmButton';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Pagination } from './common/Pagination';
import { KeyHeader } from './KeyHeader';
import { ListKeyPanel } from './panel/ListKeyPanel';

export interface IListKeyProps {
  connection: Connection
  db: number
  keyName: string,
  component: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface IListKeyItem {
  row: number,
  value: string
}

export interface IListKey {
  keyName: string,
  TTL: number,
  length: number,
  values: Array<IListKeyItem>,
  start: number,
  end: number,
  memoryUsage: number
}

const defaultListKey: IListKey = {
  keyName: '',
  TTL: -1,
  length: 0,
  values: [],
  start: 0,
  end: 20,
  memoryUsage: 0
}

export const ListKey = (props: IListKeyProps) => {

  const { connection, db, keyName, component, onKeyNameChanged } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyProps, setKeyProps] = useState<IListKey>({ ...defaultListKey, keyName }),
    [selectedValue, setSelectedValue] = useState<IListKeyItem | undefined>(),
    [error, setError] = useState<string>(),
    [showEditPanel, setShowEditPanel] = useState(false),
    [search, setSearch] = useState<{ start: number, end: number }>({ start: 0, end: 20 }),
    [items, setItems] = useState<Array<IListKeyItem>>([]),
    [filterValue, setFilterValue] = useState<string | undefined>(undefined);

  const load = useCallback(() => {
    if (!search) return;
    setError('');
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['TTL', keyProps.keyName],
        ['LLEN', keyProps.keyName],
        ['LRANGE', keyProps.keyName, search.start, search.end],
        // ['MEMORY', 'USAGE', keyName] version > 4.0
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      const vs = ret[3].map((v: string, i: number) => {
        return { row: search.start + i, value: v }
      });
      console.log(ret);
      setKeyProps(kprops => {
        return { ...kprops, TTL: ret[1], length: ret[2], values: vs }
      })
      setItems(vs);
    })
      .catch(err => setError(err))
  }, [connection.id, db, keyProps.keyName, search]);

  useEffect(() => {
    load();
  }, [load])

  // filter
  useEffect(() => {
    setItems(filterValue ? keyProps.values.filter(v => v.value.indexOf(filterValue) > -1) : keyProps.values);
  }, [filterValue, keyProps.values])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedValue(selection.getSelection()[0] as IListKeyItem);
    }
  });

  const handlePaginationChange = (start: number, end: number) => {
    setSelectedValue(undefined);
    setSearch({ start, end });
  }

  const handleDeleteConfirm = () => {
    if (!selectedValue) return;
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['LREM', keyProps.keyName, 1, selectedValue.value],
      ]
    })
      .then((ret) => {
        setKeyProps({ ...keyProps, values: keyProps.values.filter(v => v.row !== selectedValue.row), length: keyProps.length ? keyProps.length - 1 : keyProps.length });
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
        onRefresh={() => setSearch({ ...search })} />

      <Stack horizontal horizontalAlign="space-between" verticalAlign="center" tokens={{ childrenGap: 5 }}>
        <Stack.Item grow={1}><span />
          <SearchBox styles={{ root: { borderColor: theme.palette.neutralQuaternaryAlt }, iconContainer: { lineHeight: '32px' } }}
            placeholder={t('Current page search...')}
            iconProps={{ iconName: 'search' }}
            onEscape={(ev) => { setFilterValue(undefined); }}
            onClear={(ev) => { setFilterValue(undefined); }}
            onChange={(ev, nv) => { setFilterValue(nv) }}
          />
        </Stack.Item>
        <TooltipHost content={t('Add')}>
          <IconButton iconProps={{ iconName: 'circleAddition', style: { height: 'auto' } }}
            onClick={() => {
              setShowEditPanel(true);
              setSelectedValue(undefined);
            }} />
        </TooltipHost>
        <TooltipHost content={t('Edit')}>
          <IconButton iconProps={{ iconName: 'edit', style: { height: 'auto' } }}
            disabled={!selectedValue}
            onClick={() => {
              setShowEditPanel(true);
            }} />
        </TooltipHost>
        <ConfirmButton label={t('Delete')} disabled={!selectedValue} dialogContentProps={{
          type: DialogType.close, title: t('Delete this item?'),
          subText: `${selectedValue?.row}: ${selectedValue?.value}`
        }} onConfirm={handleDeleteConfirm} />
      </Stack>

      <Stack.Item grow={1} style={{ height: '100%', overflow: 'auto' }}>
        <ShimmeredDetailsList
          selection={selection}
          compact
          isHeaderVisible={false}
          setKey="items"
          items={items}
          columns={[
            { key: 'row', name: 'row', fieldName: 'row', minWidth: 35, maxWidth: 35 },
            { key: 'value', name: 'value', fieldName: 'value', minWidth: 100 }
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
      <ListKeyPanel {...props} isOpen={showEditPanel} keyType={component} keyValue={selectedValue ? selectedValue.value : ''}
        index={selectedValue?.row}
        disabledKeyName={true}
        onDismiss={handleEditPanelDismiss}
        onSave={handleSave} />


      <Pagination totalElements={keyProps.length} onChange={handlePaginationChange} />

    </Stack >
  )
}