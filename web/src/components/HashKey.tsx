import { ActionButton, CheckboxVisibility, DetailsListLayoutMode, DetailsRow, DialogType, IconButton, IDetailsRowStyles, MessageBar, MessageBarType, Selection, SelectionMode, ShimmeredDetailsList, Stack, Text, TooltipHost, useTheme } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ConfirmButton } from './common/ConfirmButton';
import { HashKeySearch } from './HashKeySearch';
import { KeyHeader } from './KeyHeader';
import { HashKeyPanel } from './panel/HashKeyPanel';

const defaultMatchPattern = "*";

const defaultStringKeyProps = {
  keyName: '',
  TTL: -1,
  length: 0,
  memoryUsage: 0,
  currentCursor: 0,
  scanned: 0
}

export interface IHashKeyProps {
  connection: Connection
  db: number
  keyName: string,
  component: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface IHashKeyItem {
  row?: number,
  field: string,
  value: string
}

export const HashKey = (props: IHashKeyProps) => {

  const { connection, db, keyName, component, onKeyNameChanged } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyProps, setKeyProps] = useState({ ...defaultStringKeyProps, keyName }),
    [selectedValue, setSelectedValue] = useState<IHashKeyItem | null>(),
    [error, setError] = useState<string>(),
    [showEditPanel, setShowEditPanel] = useState(false),
    [search, setSearch] = useState({ cursor: 0, pattern: defaultMatchPattern, count: connection.dataScanLimit }),
    [items, setItems] = useState<Array<IHashKeyItem>>([]);

  const load = useCallback(() => {
    if (!search.pattern || !search.count) return;
    setError('');
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['TTL', keyProps.keyName],
        ['HLEN', keyProps.keyName],
        ['HSCAN', keyProps.keyName, search.cursor, 'MATCH', search.pattern, 'COUNT', search.count],
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      const length = ret[3][1].length / 2;
      const vs = [...Array(length)].map((_, i) => {
        const index = 2 * i;
        return { field: ret[3][1][index], value: ret[3][1][index + 1] };
      })
      setItems(items => {
        items = search.cursor === 0 ? vs : items.concat(vs);
        return items.map((v, i) => {
          v.row = i
          return v;
        });
      });
      setKeyProps(kp => {
        var scanned = vs.length > search.count ? vs.length : search.count * 1;
        if (search.cursor !== 0) {
          scanned += kp.scanned * 1;
        }
        scanned = scanned > ret[2] ? ret[2] : scanned;
        return { ...kp, keyName: keyProps.keyName, TTL: ret[1], length: ret[2], scanned, currentCursor: ret[3][0] };
      })
      console.log(ret);
    })
      .catch((err: Error) => { setError(err.message); });
  }, [connection.id, db, search, keyProps.keyName]);

  useEffect(() => {
    load();
  }, [load])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedValue(selection.getSelection()[0] as IHashKeyItem);
    }
  });

  const handleDeleteConfirm = () => {
    if (!!!selectedValue) return;

    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['HDEL', keyProps.keyName, selectedValue.field],
      ]
    })
      .then((ret) => {
        setItems(items.filter(v => v.row !== selectedValue.row).map((v, i) => {
          v.row = i;
          return v;
        }));
        setKeyProps({ ...keyProps, length: keyProps.length ? keyProps.length - 1 : keyProps.length, scanned: keyProps.scanned - 1 });
        setSelectedValue(null);
      })
  }

  const handleEditPanelDismiss = () => {
    setShowEditPanel(false);
    setSelectedValue(null);
  }

  const handleSearch = (pattern: string, count: number) => {
    setSearch({ ...search, cursor: 0, pattern, count });
  }

  const handleSave = (field: string, value: string, index?: number) => {
    setSearch({ ...search, cursor: 0 });
  }

  return (
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 10 }}>
      {/* error */}
      {error && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.blocked} isMultiline={false} onDismiss={() => { setError(""); }} truncated={true}>
        {error}
      </MessageBar>}

      <KeyHeader {...props} keyName={keyProps.keyName} TTL={keyProps.TTL}
        onKeyNameChanged={(oldKeyName: string, keyName: string) => {
          setKeyProps({ ...keyProps, keyName });
          onKeyNameChanged(oldKeyName, keyName);
        }}
        onTTLChanged={(v: number) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => { load() }} />

      <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 5 }}>

        <HashKeySearch id={`hash-key-search-${connection.id}-${db}`}
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
            { key: 'row', name: 'Row', fieldName: 'row', minWidth: 35, maxWidth: 35 },
            { key: 'field', name: 'Field', fieldName: 'field', minWidth: 150, maxWidth: 150, isResizable: true },
            { key: 'value', name: 'Value', fieldName: 'value', minWidth: 100 }
          ]}
          selectionPreservedOnEmptyClick={false}
          checkboxVisibility={CheckboxVisibility.hidden}
          selectionMode={SelectionMode.single}
          layoutMode={DetailsListLayoutMode.justified}
          enableShimmer={!items.length}
          onItemInvoked={(item) => {
            setSelectedValue({ ...item });
            setShowEditPanel(true);
          }}
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

      {/* edit key */}
      {selectedValue && <HashKeyPanel {...props} isOpen={showEditPanel} keyType={component}
        KeyFieldName={selectedValue?.field}
        KeyFieldValue={selectedValue?.value}
        index={selectedValue?.row}
        disabledKeyName={true}
        onDismiss={handleEditPanelDismiss}
        onSave={handleSave} />}

    </Stack >
  )
}