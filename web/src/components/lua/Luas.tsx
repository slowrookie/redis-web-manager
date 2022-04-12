import { CheckboxVisibility, Depths, DetailsList, DetailsListLayoutMode, PrimaryButton, Selection, SelectionMode, Spinner, SpinnerSize, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection } from '../../services/connection.service';
import { loadLuas, Lua } from '../../services/lua.service';
import { ErrorMessageBar } from '../common/ErrorMessageBar';
import { Scripting } from './Scripting';


export interface ILuaProps {
  connection: Connection
}

export const Luas = (props: ILuaProps) => {
  const { connection } = props;

  const theme = useTheme(),
    [loading, setLoading] = useState(false),
    { t } = useTranslation(),
    [luas, setLuas] = useState<Lua[]>([]),
    [error, setError] = useState<Error>(),
    [selectedLua, setSelectedLua] = useState<Lua>();

  const handleLoadLuas = useCallback(() => {
    setError(undefined);
    setLoading(true);
    loadLuas(connection.id)
      .then((v: Lua[]) => setLuas(v))
      .catch((err: Error) => { setError(err); })
      .finally(() => { setLoading(false) });
  }, [connection.id])

  useEffect(() => {
    handleLoadLuas();
  }, [handleLoadLuas])

  const selection = new Selection({
    onSelectionChanged: () => {
      setSelectedLua(selection.getSelection()[0] as Lua);
    }
  })

  const handleAdd = () => {
    setSelectedLua({ connectionID: connection.id } as Lua)
  }

  return (
    <Stack horizontal style={{ height: '100%' }}>
      <Stack styles={{ root: { width: 200, boxShadow: Depths.depth4, padding: 5 } }}>
        {/* error */}
        <ErrorMessageBar error={error}></ErrorMessageBar>
        {/* header */}
        <Stack horizontal tokens={{ childrenGap: 5 }} >
          <TooltipHost styles={{ root: { width: '100%' } }} key="circleAddition" content={t('Add')}>
            <PrimaryButton styles={{ root: { width: '100%' } }} iconProps={{ iconName: "circleAddition", style: { height: 'auto' } }}
              onClick={handleAdd} />
          </TooltipHost>
        </Stack>
        {/* {list} */}
        {loading ? <Spinner size={SpinnerSize.small} /> :
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
                key: 'id', name: 'id', minWidth: 190, onRender: (item) => {
                  return <TooltipHost content={item.script}>
                    {item.script}
                  </TooltipHost>
                }
              }]}
            items={[...luas]} />
        }
      </Stack>

      <Stack.Item grow={1}>
        {selectedLua && <Scripting lua={selectedLua} onChanged={() => {
          loadLuas(connection.id);
        }} />}
      </Stack.Item>
    </Stack>
  )
}