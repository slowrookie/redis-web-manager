import {
  ActionButton,
  CheckboxVisibility, DetailsListLayoutMode,
  DetailsRow,
  DialogType,
  IconButton,
  IDetailsRowStyles,
  Selection,
  SelectionMode, ShimmeredDetailsList, Stack, Text, TooltipHost, useTheme
} from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ConfirmButton } from './common/ConfirmButton';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { KeyHeader } from './KeyHeader';
import { SetKeyPanel } from './panel/SetKeyPanel';
import { SetKeySearch } from './SetKeySearch';

const defaultMatchPattern = "*";

export interface ISetKeyProps {
  connection: Connection
  db: number
  keyName: string,
  component: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface ISetKeyItem {
  row: number,
  value: string
}

export interface ISetKey {
  keyName: string,
  TTL: number,
  length: number,
  memoryUsage: number
  currentCursor: number,
  scanned: number
}

const defaultSetKey: ISetKey = {
  keyName: '',
  TTL: -1,
  length: 0,
  memoryUsage: 0,
  currentCursor: 0,
  scanned: 0
}

export const SetKey = (props: ISetKeyProps) => {

  const { connection, db, keyName, component, onKeyNameChanged } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyProps, setKeyProps] = useState({ ...defaultSetKey, keyName }),
    [selectedValue, setSelectedValue] = useState<ISetKeyItem | undefined>(),
    [error, setError] = useState<Error>(),
    [showEditPanel, setShowEditPanel] = useState(false),
    [search, setSearch] = useState({ cursor: 0, pattern: defaultMatchPattern, count: connection.dataScanLimit }),
    [items, setItems] = useState<Array<ISetKeyItem>>([]);

  const load = useCallback(() => {
    if (!search.pattern || !search.count) return;
    setError(undefined);
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['TTL', keyProps.keyName],
        ['SCARD', keyProps.keyName],
        ['SSCAN', keyProps.keyName, search.cursor, 'MATCH', search.pattern, 'COUNT', search.count],
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      const vs = ret[3][1].map((v: string, i: number) => {
        return { row: i + 1, value: v }
      });
      setItems(items => {
        items = search.cursor === 0 ? vs : items.concat(vs);
        return items.map((v: ISetKeyItem, i: number) => {
          v.row = i
          return v;
        });
      });
      setKeyProps(kp => {
        let scanned = vs.length > search.count ? vs.length : search.count * 1;
        if (search.cursor !== 0) {
          scanned += kp.scanned * 1;
        }
        scanned = scanned > ret[2] ? ret[2] : scanned;
        return { ...kp, keyName: keyProps.keyName, TTL: ret[1], length: ret[2], scanned, currentCursor: ret[3][0] };
      })
    })
      .catch(err => setError(err));
  }, [connection.id, db, search, keyProps.keyName]);

  useEffect(() => {
    load();
  }, [load])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedValue(selection.getSelection()[0] as ISetKeyItem);
    }
  });

  const handleDeleteConfirm = () => {
    if (!selectedValue) return;
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['SREM', keyProps.keyName, selectedValue.value],
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

  const handleSave = () => {
    setSearch({ ...search, cursor: 0 });
  }

  const handleSearch = (pattern: string, count: number) => {
    setSearch({ ...search, cursor: 0, pattern, count });
  }

  return (<>
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 5 }}>
      {/* error */}
      <ErrorMessageBar error={error}></ErrorMessageBar>

      <KeyHeader {...props} keyName={keyProps.keyName} TTL={keyProps.TTL}
        onKeyNameChanged={(oldValue, newValue) => {
          setKeyProps({ ...keyProps, keyName: newValue });
          onKeyNameChanged(oldValue, newValue);
        }}
        onTTLChanged={(v) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => { load() }} />

      <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 5 }}>

        <SetKeySearch id={`set-key-search-${connection.id}-${db}`}
          defaultMatchPattern={defaultMatchPattern}
          defaultCount={connection.dataScanLimit}
          length={keyProps.length}
          onSearch={handleSearch}
        />

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

      <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 10 }} >
        <TooltipHost content="Loaded / Scanned / Total">
          <Text variant="tiny" style={{ fontWeight: 'bold' }} nowrap block>{`${items.length} / ${keyProps.scanned} /${keyProps.length} `}</Text>
        </TooltipHost>
        <TooltipHost content={t('Load more')}>
          <ActionButton
            style={{ height: 32 }}
            iconProps={{ iconName: 'syncicon', style: { height: 'auto' } }}
            text={t('Load more')}
            disabled={keyProps.scanned === keyProps.length}
            onClick={() => { setSearch({ ...search, cursor: keyProps.currentCursor }) }} />
        </TooltipHost>
      </Stack>
    </Stack >

    {/* edit key */}
    <SetKeyPanel {...props} isOpen={showEditPanel} keyType={component} keyValue={selectedValue ? selectedValue.value : ''}
      index={selectedValue ? selectedValue.row - 1 : undefined}
      disabledKeyName={true}
      onDismiss={handleEditPanelDismiss}
      onSave={handleSave} />
  </>)
}