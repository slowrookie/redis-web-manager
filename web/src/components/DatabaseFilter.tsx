import { Callout, Checkbox, IconButton, Stack, TextField, TooltipHost, useTheme, ChoiceGroup, DirectionalHint } from '@fluentui/react';
import _ from 'lodash';
import React, { KeyboardEvent, MouseEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { KeyTypes } from './utils';


export interface IDatabaseFilterProps {
  id: string
  defaultValue: string
  defaultCount: number
  namespaceSeparator: string
  keys: Array<string>
  onFilter: (value: string, count: number, types: Array<string> | string) => void
}

export const DatabaseFilter = (props: IDatabaseFilterProps) => {

  const { id, defaultValue, defaultCount, namespaceSeparator, keys, onFilter } = props;

  const theme = useTheme(),
    { t } = useTranslation();

  const textFieldStyles = {
    suffix: { padding: 0 },
    fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt }
  };

  const [filterVisible, setFilterVisible] = useState(false),
    [filterGroupVisible, setFilterGroupVisible] = useState(false),
    [value, setValue] = useState<string | undefined>(defaultValue),
    [count, setCount] = useState<number>(defaultCount),
    [types, setTypes] = useState<Array<{ type: string, checked: boolean }>>(Object.keys(KeyTypes).map(type => {
      return { type, checked: true }
    })),
    [groups, setGroups] = useState<Array<string>>()
    ;

  useEffect(() => {
    const groupKyes = keys
      .filter(k => k.indexOf(namespaceSeparator) > 0)
      .map(g => {
        const keySplit = g.split(namespaceSeparator);
        keySplit.splice(keySplit.length - 1, 1, "*")
        return keySplit.join(namespaceSeparator);
      });

    setGroups(["*", ..._.uniq(groupKyes)]);
  }, [keys, namespaceSeparator])

  const handleFilter = (ev: KeyboardEvent | MouseEvent<any>) => {
    ev.preventDefault();
    value && onFilter(value, count, types.filter(t => t.checked).map(t => t.type));
    setFilterVisible(false);
  }

  return (<div>
    <TextField
      value={value}
      styles={textFieldStyles}
      onChange={(ev, v) => setValue(v)}
      onKeyDown={(ev: KeyboardEvent) => {
        if (ev.key === 'Enter') {
          handleFilter(ev);
        }
      }}
      onRenderSuffix={(textFieldProps, defaultRender) => {
        return (<>
           <TooltipHost content={t('Groups')}>
            <IconButton id={`${id}-group`} styles={{ root: { height: 24 } }} iconProps={{ iconName: 'grouplist', style: { height: 'auto' } }} onClick={() => setFilterGroupVisible(!filterVisible)} />
          </TooltipHost>
          <TooltipHost content={t('Filter')}>
            <IconButton id={`${id}-filter`} styles={{ root: { height: 24 } }} iconProps={{ iconName: 'filter', style: { height: 'auto' } }} onClick={() => setFilterVisible(!filterVisible)} />
          </TooltipHost>
          <TooltipHost content={t('Search')}>
            <IconButton styles={{ root: { height: 24 } }}
              iconProps={{ iconName: 'skypeCircleCheck', style: { height: 'auto' } }} onClick={handleFilter} />
          </TooltipHost>
        </>)
      }} />

    {filterGroupVisible && <Callout target={`#${id}-group`} onDismiss={() => setFilterGroupVisible(false)} directionalHint={DirectionalHint.bottomLeftEdge}>
      <Stack tokens={{ childrenGap: 10, padding: 10 }}>
        <ChoiceGroup
          // defaultSelectedKey="*"
          selectedKey={value}
          options={groups?.map((v) => { return { key: v, text: v } })}
          onChange={(_, o) => {
            setValue(o?.key);
          }}
          label={t('Groups')} />
      </Stack>
    </Callout>}

    {filterVisible && <Callout target={`#${id}-filter`} onDismiss={() => setFilterVisible(false)} setInitialFocus={true}>
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

        {/* <Stack horizontal horizontalAlign="space-between" >
          <TooltipHost content={t('Cancel')}>
            <ActionButton text={t('Cancel')} onClick={() => setFilterVisible(false)} />
          </TooltipHost>
          <TooltipHost content={t('OK')}>
            <ActionButton text={t('OK')} onClick={handleFilter} />
          </TooltipHost>
        </Stack> */}

      </Stack>
    </Callout>}
  </div>)
}