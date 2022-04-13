import { CheckboxVisibility, DefaultButton, Depths, DetailsList, DetailsListLayoutMode, PrimaryButton, Selection, SelectionMode, Spinner, SpinnerSize, Stack, TooltipHost, useTheme } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection } from '../../services/connection.service';
import { deleteLua, loadLuas, Lua } from '../../services/lua.service';
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

  const [selection] = useState(new Selection({
    onSelectionChanged: () => {
      setSelectedLua(selection.getSelection()[0] as Lua);
    }
  }))

  const handleAdd = () => {
    selection.setAllSelected(false);
    setSelectedLua({ connectionID: connection.id } as Lua)
  }

  const handleDelete = () => {
    selectedLua && deleteLua(selectedLua)
      .then(() => {
        handleLoadLuas();
        selection.setAllSelected(false);
      })
      .catch((err: Error) => { setError(err); })
      .finally()
  }

  return (
    <Stack horizontal style={{ height: '100%' }}>
      <Stack styles={{ root: { width: 220, boxShadow: Depths.depth4, padding: 5 } }} tokens={{ childrenGap: 5 }}>
        {/* error */}
        <ErrorMessageBar error={error}></ErrorMessageBar>
        {/* header */}
        <Stack tokens={{ childrenGap: 5 }} >
          <TooltipHost styles={{ root: { width: '100%' } }} key="circleAddition" content={t('Add')}>
            <PrimaryButton styles={{ root: { width: '100%' } }} iconProps={{ iconName: "circleAddition", style: { height: 'auto' } }}
              onClick={handleAdd} />
          </TooltipHost>

          <TooltipHost styles={{ root: { width: '100%' } }} key="remove" content={t('Delete')}>
            <DefaultButton disabled={!selectedLua?.id} styles={{ root: { width: '100%' } }} iconProps={{ iconName: "delete", style: { height: 'auto', color: theme.palette.redDark } }}
              onClick={handleDelete} />
          </TooltipHost>
        </Stack>
        {/* {list} */}
        <Stack.Item grow={1} style={{ overflow: 'auto' }}>
          {loading ? <Spinner size={SpinnerSize.small} /> :
            <DetailsList compact
              layoutMode={DetailsListLayoutMode.justified}
              selection={selection}
              selectionMode={SelectionMode.single}
              checkboxVisibility={CheckboxVisibility.hidden}
              enterModalSelectionOnTouch={true}
              selectionPreservedOnEmptyClick={true}
              isHeaderVisible={false}
              setKey="luas"
              columns={[
                {
                  key: 'id', name: 'id', minWidth: 190, onRender: (item) => {
                    return <TooltipHost content={item.name}>
                      {item.name}
                    </TooltipHost>
                  }
                }]}
              items={luas} />
          }
        </Stack.Item>
      </Stack>

      <Stack.Item grow={1}>
        {selectedLua && <Scripting lua={selectedLua} onChanged={() => {
          handleLoadLuas();
        }} />}
        {/* {scripting()} */}
      </Stack.Item>
    </Stack>
  )
}