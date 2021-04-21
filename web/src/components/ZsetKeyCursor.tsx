import {
  ActionButton,
  CheckboxVisibility, DetailsListLayoutMode,
  DetailsRow,
  DialogType,
  IconButton,
  Selection,
  SelectionMode, ShimmeredDetailsList, Stack, Text, TooltipHost, useTheme
} from '@fluentui/react';
import React, { useCallback, useEffect, useImperativeHandle, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { backendAPI } from '../api';
import { ConfirmButton } from './common/ConfirmButton';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { ZsetKeyPanel } from './panel/ZsetKeyPanel';
import { ZsetKeySearch } from './ZsetKeyCursorSearch';

const defaultMatchPattern = "*";

const defaultStringKeyProps = {
  keyName: '',
  length: 0,
  currentCursor: 0,
  scanned: 0,
}

export const ZsetKeyCursor = (props) => {

  const { connection, db, keyName, component, onChangeSearchType, childRef } = props;
  const { t } = useTranslation(), theme = useTheme();

  const [keyProps, setKeyProps] = useState({ ...defaultStringKeyProps, keyName }),
    [selectedValue, setSelectedValue] = useState(),
    [error, setError] = useState(),
    [showEditPanel, setShowEditPanel] = useState(false),
    [search, setSearch] = useState({ cursor: 0, pattern: defaultMatchPattern, count: connection.dataScanLimit }),
    [items, setItems] = useState([]);

  const load = useCallback(() => {
    if (!search.pattern || !search.count) return;
    setError('');
    backendAPI.Connection.Command({
      id: connection.id, commands: [
        ['SELECT', db],
        ['ZCARD', keyProps.keyName],
        ['ZSCAN', keyProps.keyName, search.cursor, 'MATCH', search.pattern, 'COUNT', search.count],
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      const length = ret[2][1].length / 2;
      const vs = [...Array(length)].map((_, i) => {
        const index = 2 * i;
        return { row: i + 1, value: ret[2][1][index], score: ret[2][1][index + 1] };
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
        scanned = scanned > ret[1] ? ret[1] : scanned;
        return { ...kp, keyName: keyProps.keyName, length: ret[1], scanned, currentCursor: ret[2][0] };
      })
      console.log(ret);
    })
      .catch(err => setError(err));
  }, [connection.id, db, search, keyProps.keyName]);

  useEffect(() => {
    load();
  }, [load])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedValue(selection.getSelection()[0]);
    }
  });

  const handleDeleteConfirm = () => {
    backendAPI.Connection.Command({
      id: connection.id, commands: [
        ['SELECT', db],
        ['ZREM', keyProps.keyName, selectedValue.value],
      ]
    })
      .then((ret) => {
        setKeyProps({ ...keyProps, values: keyProps.values.filter(v => v.row !== selectedValue.row), length: keyProps.length ? keyProps.length - 1 : keyProps.length });
        setSelectedValue();
      })
  }

  const handleEditPanelDismiss = () => {
    setShowEditPanel(false);
    setSelectedValue();
  }

  const handleSave = (newValue) => {
    setSearch({ ...search, cursor: 0 });
  }

  useImperativeHandle(childRef, () => {
    return {
      refresh() {
        load();
      }
    }
  });

  const handleSearch = (pattern, count) => {
    setSearch({ ...search, cursor: 0, pattern, count });
  }

  return (<>
    {/* error */}
    <ErrorMessageBar error={error}></ErrorMessageBar>

    <Stack horizontal horizontalAlign="end" verticalAlign="center" tokens={{ childrenGap: 5 }}>

      <ZsetKeySearch id={`zset-key-cursor-search-${connection.id}-${db}`}
        defaultMatchPattern={defaultMatchPattern}
        defaultCount={connection.dataScanLimit}
        length={keyProps.length}
        onSearch={handleSearch}
      />

      <TooltipHost content={t('Add')}>
        <IconButton iconProps={{ iconName: 'circleAddition', style: { height: 'auto' } }}
          onClick={() => {
            setShowEditPanel(true);
            setSelectedValue();
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

      <TooltipHost content={t('Index search')}>
        <IconButton iconProps={{ iconName: 'QueryList', style: { height: 'auto' } }} onClick={() => { onChangeSearchType('INDEX') }} />
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
          return <DetailsRow {...props} styles={{
            root: {
              backgroundColor: (props && props.itemIndex % 2 === 0) ? theme.palette.neutralLighterAlt : ''
            }
          }} />
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
    <ZsetKeyPanel {...props} isOpen={showEditPanel} keyType={component} keyValue={selectedValue?.value} keyScore={selectedValue?.score}
      index={selectedValue?.row - 1}
      disabledKeyName={true}
      onDismiss={handleEditPanelDismiss}
      onSave={handleSave} />


  </>)
}