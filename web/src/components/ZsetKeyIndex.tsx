import {
  CheckboxVisibility, DetailsListLayoutMode,
  DetailsRow,
  DialogType,
  IconButton,
  IDetailsRowStyles,
  Selection,
  SelectionMode, ShimmeredDetailsList, Stack, TooltipHost, useTheme
} from '@fluentui/react';
import React, { useCallback, useEffect, useImperativeHandle, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ConfirmButton } from './common/ConfirmButton';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { Pagination } from './common/Pagination';
import { ZsetKeyPanel } from './panel/ZsetKeyPanel';
import { IZsetKeyItem } from './ZsetKey';

export interface IZsetKeyIndexProps {
  connection: Connection
  db: number
  keyName: string
  component: string
  childRef: any
  onChangeSearchType: (type: string) => void
}

const defaultZsetKeyIndexSearchProps = {
  keyName: '',
  length: 0,
  start: 0,
  end: 20,
}

export const ZsetKeyIndex = (props: IZsetKeyIndexProps) => {

  const { connection, db, keyName, component, onChangeSearchType, childRef } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyProps, setKeyProps] = useState({ ...defaultZsetKeyIndexSearchProps, keyName }),
    [selectedValue, setSelectedValue] = useState<IZsetKeyItem>(),
    [error, setError] = useState<string>(),
    [showEditPanel, setShowEditPanel] = useState(false),
    [search, setSearch] = useState({ start: 0, end: 20, asc: true }),
    [items, setItems] = useState<Array<IZsetKeyItem>>([]);

  useEffect(() => {
    setKeyProps(pops => {
      return { ...pops, keyName }
    });
  }, [keyName])

  const load = useCallback(() => {
    if (!search.end) return;
    setError('');
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['ZCARD', keyProps.keyName],
        [search.asc ? 'ZRANGE' : 'ZREVRANGE', keyProps.keyName, search.start, search.end, 'WITHSCORES']
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      console.log(ret);
      const length = ret[2].length / 2;
      const vs = [...Array(length)].map((_, i) => {
        const index = 2 * i;
        return { index: i, row: i + 1, value: ret[2][index], score: ret[2][index + 1] };
      })

      setKeyProps(kProps => {
        return { ...kProps, length: ret[1] };
      });
      setItems(vs);
    })
      .catch(err => setError(err))
  }, [connection.id, db, keyProps.keyName, search]);

  useEffect(() => {
    load();
  }, [load])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedValue(selection.getSelection()[0] as IZsetKeyItem);
    }
  });

  const handlePaginationChange = (start: number, end: number) => {
    setSelectedValue(undefined);
    setSearch({ ...search, start, end });
  }

  const handleDeleteConfirm = () => {
    if (!selectedValue) return;
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['ZREM', keyProps.keyName, selectedValue.value],
      ]
    })
      .then((ret) => {
        setKeyProps({ ...keyProps, length: keyProps.length ? keyProps.length - 1 : keyProps.length });
        setItems(items.filter(v => v.row !== selectedValue.row))
        setSelectedValue(undefined);
      })
  }

  const handleEditPanelDismiss = () => {
    setShowEditPanel(false);
    setSelectedValue(undefined);
  }

  const handleSave = (newValue: string) => {
    setSearch({ ...search })
  }

  useImperativeHandle(childRef, () => {
    return {
      refresh() {
        setSearch({ ...search })
      }
    }
  });

  return (<>
    {/* error */}
    <ErrorMessageBar error={error}></ErrorMessageBar>

    <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 5 }}>

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

      <TooltipHost content={t('Cursor search')}>
        <IconButton iconProps={{ iconName: 'GroupList', style: { height: 'auto' } }} onClick={() => { onChangeSearchType('CURSOR') }} />
      </TooltipHost>
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
          { key: 'score', name: 'score', fieldName: 'score', minWidth: 150, maxWidth: 150 },
          { key: 'value', name: 'value', fieldName: 'value', minWidth: 100 }
        ]}
        selectionPreservedOnEmptyClick={false}
        checkboxVisibility={CheckboxVisibility.hidden}
        selectionMode={SelectionMode.single}
        layoutMode={DetailsListLayoutMode.justified}
        enableShimmer={!items.length}
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
    <ZsetKeyPanel {...props} isOpen={showEditPanel} keyType={component}
      keyValue={selectedValue ? selectedValue.value : ''}
      keyScore={selectedValue ? selectedValue.score : ''}
      index={selectedValue ? selectedValue.row - 1 : undefined}
      disabledKeyName={true}
      onDismiss={handleEditPanelDismiss}
      onSave={handleSave} />

    <Pagination totalElements={keyProps.length} onChange={handlePaginationChange} />

  </>
  )
}