import { ActionButton, Callout, Checkbox, IconButton, Stack, TextField, TooltipHost, useTheme } from '@fluentui/react';
import React, { KeyboardEvent, MouseEvent, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { KeyTypes } from './utils';


export interface IDatabaseFilterProps {
  id: string
  defaultValue: string
  defaultCount: number
  onFilter: (value: string, count: number, types: Array<string> | string) => void
}

export const DatabaseFilter = (props: IDatabaseFilterProps) => {

  const { id, defaultValue, defaultCount, onFilter } = props;

  const theme = useTheme(),
    { t } = useTranslation();

  const textFieldStyles = {
    suffix: { padding: 0 },
    fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt }
  };

  const [filterVisible, setFilterVisible] = useState(false),
    [value, setValue] = useState<string | undefined>(defaultValue),
    [count, setCount] = useState<number>(defaultCount),
    [types, setTypes] = useState<Array<{ type: string, checked: boolean }>>(Object.keys(KeyTypes).map(type => {
      return { type, checked: true }
    }));

  const handleFilter = (ev: KeyboardEvent | MouseEvent<any>) => {
    ev.preventDefault();
    value && onFilter(value, count, types.filter(t => t.checked).map(t => t.type));
    setFilterVisible(false);
  }

  return (<div>
    <TextField value={value}
      styles={textFieldStyles}
      onChange={(ev, v) => setValue(v)}
      onKeyDown={(ev: KeyboardEvent) => {
        if (ev.key === 'Enter') {
          handleFilter(ev);
        }
      }}
      onRenderSuffix={(textFieldProps, defaultRender) => {
        return (
          <TooltipHost content={t('Filter')}>
            <IconButton id={id} styles={{ root: { height: 24 } }} iconProps={{ iconName: 'filter', style: { height: 'auto' } }} onClick={() => setFilterVisible(!filterVisible)} />
          </TooltipHost>
        )
      }} />

    {filterVisible && <Callout target={`#${id}`} onDismiss={() => setFilterVisible(false)} setInitialFocus={true}>
      <Stack tokens={{ childrenGap: 10, padding: 10 }} style={{ width: 150 }}>

        <TextField type="number" label="COUNT" value={`${count}`} onChange={(e, v) => setCount(Number(v))} styles={textFieldStyles}
          onKeyDown={(ev: KeyboardEvent) => {
            if (ev.key === 'Enter') {
              handleFilter(ev);
            }
          }}
        />

        {types.map((t, i) => {
          return <Checkbox disabled key={i} label={t.type} checked={t.checked} onChange={(ev, checked) => {
            if (undefined !== checked) {
              types[i].checked = checked;
              setTypes([...types]);
            }
          }} />
        })}

        <Stack horizontal horizontalAlign="space-between" >
          <TooltipHost content={t('Cancel')}>
            <ActionButton text={t('Cancel')} onClick={() => setFilterVisible(false)} />
          </TooltipHost>
          <TooltipHost content={t('OK')}>
            <ActionButton text={t('OK')} onClick={handleFilter} />
          </TooltipHost>
        </Stack>
      </Stack>
    </Callout>}
  </div>)
}