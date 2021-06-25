import { ActionButton, CheckboxVisibility, DetailsListLayoutMode, DetailsRow, IDetailsRowStyles, ScrollablePane, ScrollbarVisibility, SearchBox, Selection, SelectionMode, ShimmeredDetailsList, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../../services/connection.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';

export interface IDatabaseConfigProps {
  connection: Connection
}

interface Config {
  key: string
  value: any
}

export const DatabaseConfiguration = (props: IDatabaseConfigProps) => {

  const theme = useTheme(), { t } = useTranslation();

  const [configs, setConfigs] = useState<Array<Config>>([]),
    // [selectedKey, setSelectedKey] = useState<Config>(),
    [filter, setFilter] = useState<string>("*"),
    [loading, setLoading] = useState<boolean>(false),
    [error, setError] = useState<Error>();

  const load = useCallback(() => {
    setLoading(true);
    executeCommand<Array<any>>({ id: props.connection.id, commands: [['CONFIG', 'GET', filter]] })
      .then(ret => {
        if (!ret || !ret.length) return;
        const length: number = ret[0].length / 2;
        const confs: Array<Config> = [];
        [...Array(length)].forEach((_, i) => {
          const index = 2 * i;
          const key = ret[0][index];
          const value = ret[0][index + 1];
          confs.push({ key, value })
        });
        setConfigs([...confs]);
      })
      .catch(err => setError(err))
      .finally(() => setLoading(false));
  }, [props.connection.id, filter]);

  useEffect(() => {
    load();
  }, [load]);

  const selection = new Selection({
    onSelectionChanged: () => {
      // TODO
      // setSelectedKey(selection.getSelection()[0] as Config);
    }
  });

  return (
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 5 }}>

      <ErrorMessageBar error={error}></ErrorMessageBar>

      <Stack horizontal horizontalAlign="space-between" verticalAlign="center" tokens={{ childrenGap: 5 }}>
        <Stack.Item grow={1}>
          <SearchBox value={filter} styles={{ root: { borderColor: theme.palette.neutralQuaternaryAlt }, iconContainer: { lineHeight: '32px' } }}
            placeholder={`${t('Search')}....`}
            iconProps={{ iconName: 'search' }}
            onChange={(_, nv) => { setFilter(nv || '*') }}
          />
        </Stack.Item>
        {/* <TooltipHost content={t('Rewrite')}>
          <ActionButton iconProps={{ iconName: 'Save', style: { height: 'auto' } }}
            onClick={() => {
            }} />
        </TooltipHost> */}
        <TooltipHost content={t('Refresh')}>
          <ActionButton iconProps={{ iconName: 'Refresh', style: { height: 'auto' } }} onClick={load} />
        </TooltipHost>
      </Stack>
      <Stack.Item grow={1} style={{ position: 'relative' }}>
        <ScrollablePane scrollbarVisibility={ScrollbarVisibility.auto}>
          <ShimmeredDetailsList
            selection={selection}
            compact
            isHeaderVisible={false}
            setKey="items"
            items={configs}
            columns={[
              { key: 'key', name: 'key', fieldName: 'key', minWidth: 120, maxWidth: 250, isResizable: true },
              { key: 'value', name: 'value', fieldName: 'value', minWidth: 100, isResizable: true }
            ]}
            selectionPreservedOnEmptyClick={false}
            checkboxVisibility={CheckboxVisibility.hidden}
            selectionMode={SelectionMode.single}
            layoutMode={DetailsListLayoutMode.justified}
            enableShimmer={loading}
            onItemInvoked={(item) => { }}
            onRenderRow={props => {
              const detailsRowStyles: Partial<IDetailsRowStyles> = {
                root: {
                  backgroundColor: (props && props.itemIndex % 2 === 0) ? theme.palette.neutralLighterAlt : ''
                }
              };
              return props ? <DetailsRow {...props} styles={detailsRowStyles} /> : (<></>)
            }}
          />
        </ScrollablePane>
      </Stack.Item>
    </Stack>
  )
}