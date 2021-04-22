import { ActionButton, Callout, IconButton, Stack, TextField, TooltipHost, useTheme } from '@fluentui/react';
import React, { KeyboardEvent, MouseEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

export interface IZsetKeyCursorProps {
  id: string
  onSearch: (pattern: string, count: number) => void
  length: number
  defaultMatchPattern: string
  defaultCount: number
}

export const ZsetKeySearch = (props: IZsetKeyCursorProps) => {
  const { id, onSearch, length, defaultMatchPattern, defaultCount } = props;

  const theme = useTheme(),
    { t } = useTranslation();

  const textFieldStyles = {
    fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt }
  };

  const [searchVisible, setSearchVisible] = useState(false),
    [condition, setCondition] = useState({ pattern: defaultMatchPattern, count: defaultCount, length });

  useEffect(() => {
    setCondition(condition => {
      return { ...condition, length: props.length };
    })
  }, [props])

  const handleSearch = (ev: MouseEvent<any> | KeyboardEvent) => {
    ev.preventDefault();
    onSearch(condition.pattern, condition.count);
    setSearchVisible(false);
  }

  return (<div>
    <TooltipHost content={t('Search')}>
      <IconButton id={id} iconProps={{ iconName: 'searchdata', style: { height: 'auto' } }} onClick={() => setSearchVisible(true)} />
    </TooltipHost>

    {searchVisible && <Callout target={`#${id}`} onDismiss={() => setSearchVisible(false)} setInitialFocus={true}>
      <Stack tokens={{ childrenGap: 10, padding: 10 }} style={{ minWidth: 300 }}>

        <TextField styles={textFieldStyles} label="MATCH" defaultValue={defaultMatchPattern} value={condition.pattern}
          onChange={(e, v) => v && setCondition({ ...condition, pattern: v })}
          onKeyDown={(ev: KeyboardEvent) => {
            if (ev.key === 'Enter') {
              handleSearch(ev);
            }
          }} />

        <TextField styles={textFieldStyles} label="COUNT" type="number" suffix={`/ ${length}`} value={`${condition.count}`}
          onChange={(e, v) => { v && setCondition({ ...condition, count: Number(v) }); }}
          onKeyDown={(ev: KeyboardEvent) => {
            if (ev.key === 'Enter') {
              handleSearch(ev);
            }
          }}
        />

        <Stack horizontal horizontalAlign="space-between" >
          <TooltipHost content={t('Cancel')}>
            <ActionButton text={t('Cancel')} onClick={() => setSearchVisible(false)} />
          </TooltipHost>
          <TooltipHost content={t('OK')}>
            <ActionButton text={t('OK')} onClick={handleSearch} />
          </TooltipHost>
        </Stack>
      </Stack>
    </Callout>}
  </div>)
}