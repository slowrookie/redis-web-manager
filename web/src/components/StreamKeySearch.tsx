import { ActionButton, Callout, IconButton, Slider, Stack, Text, TextField, TooltipHost, useTheme } from '@fluentui/react';
import React, { KeyboardEvent, MouseEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import dayjs from 'dayjs';

export interface IStreamKeySearch {
  id: string
  onSearch: (start: string | number, end: string | number, count: number) => void
  length: number
  firstTimestamp: number | null
  lastTimestamp: number | null
  defaultCount: number | null
}

export interface IStreamKeySearchCondition {
  start: string | number
  end: string | number
  count: number
  firstTimestamp: number | null
  lastTimestamp: number | null
  length: number
}

export const StreamKeySearch = (props: IStreamKeySearch) => {
  const { id, onSearch, length, firstTimestamp, lastTimestamp, defaultCount } = props;

  const theme = useTheme(),
    { t } = useTranslation();

  const textFieldStyles = {
    fieldGroup: { borderColor: theme.palette.neutralQuaternaryAlt }
  };

  const [searchVisible, setSearchVisible] = useState(false),
    [condition, setCondition] = useState<IStreamKeySearchCondition>({ start: '-', end: '+', count: defaultCount || 0, length, firstTimestamp, lastTimestamp });

  useEffect(() => {
    setCondition(condition => {
      return { ...condition, length: props.length, firstTimestamp: props.firstTimestamp, lastTimestamp: props.lastTimestamp };
    })
  }, [props])

  const handleSearch = (ev: MouseEvent<any> | KeyboardEvent) => {
    ev.preventDefault();
    onSearch(condition.start, condition.end, condition.count);
    setSearchVisible(false);
  }

  const handleRangeChange = (_: number, range?: Array<number>) => {
    if (!range) return;
    setCondition({ ...condition, start: range[0] === condition.firstTimestamp ? '-' : range[0], end: range[1] === condition.lastTimestamp ? '+' : range[1] })
  }

  return (<div>
    <TooltipHost content={t('Search')}>
      <IconButton id={id} iconProps={{ iconName: 'searchdata', style: { height: 'auto' } }} onClick={() => setSearchVisible(!searchVisible)} />
    </TooltipHost>

    {searchVisible && <Callout target={`#${id}`} onDismiss={() => setSearchVisible(false)} setInitialFocus={true}>
      <Stack tokens={{ childrenGap: 10, padding: 10 }} style={{ minWidth: 300 }}>

        <Slider ranged
          min={condition.firstTimestamp || 0}
          max={condition.lastTimestamp || 0}
          defaultLowerValue={condition.firstTimestamp || 0}
          lowerValue={condition.start === '-' ? (condition.firstTimestamp || 0) : Number(condition.start)}
          value={condition.end === '+' ? (condition.lastTimestamp || 0) : Number(condition.end)}
          onChange={handleRangeChange}
          showValue={false}
          step={100} />

        <TextField
          styles={textFieldStyles}
          label="START"
          value={`${condition.start}`}
          onChange={(e, v) => setCondition({ ...condition, start: Number(v) })}
          onRenderLabel={(props, defaultRender) => {
            const start = condition.start === '-' ? condition.firstTimestamp : condition.start;
            return (<Stack horizontal horizontalAlign="space-between" verticalAlign="center">
              {defaultRender && defaultRender(props)}
              <Text variant="xSmall">{dayjs(start).format("YYYY-MM-DD HH:mm:ss SSS")}</Text>
            </Stack>)
          }}
          onKeyDown={(ev: KeyboardEvent) => {
            if (ev.key === 'Enter') {
              handleSearch(ev);
            }
          }}
        />

        <TextField
          styles={textFieldStyles}
          label="END"
          value={`${condition.end}`}
          onChange={(e, v) => setCondition({ ...condition, end: Number(v) })}
          onRenderLabel={(props, defaultRender) => {
            const end = condition.end === '+' ? condition.lastTimestamp : condition.end
            return (<Stack horizontal horizontalAlign="space-between" verticalAlign="center">
              {defaultRender && defaultRender(props)}
              <Text variant="xSmall">{dayjs(end).format("YYYY-MM-DD HH:mm:ss SSS")}</Text>
            </Stack>)
          }}
          onKeyDown={(ev: KeyboardEvent) => {
            if (ev.key === 'Enter') {
              handleSearch(ev);
            }
          }}
        />

        <TextField styles={textFieldStyles} type="number" label="COUNT" suffix={`/ ${condition.length}`} value={`${condition.count}`}
          onChange={(e, v) => {
            if (Number(v) > condition.length) {
              v = `${condition.length}`
            }
            setCondition({ ...condition, count: Number(v) });
          }}
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