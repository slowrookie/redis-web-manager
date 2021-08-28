import { Dropdown, IDropdownOption, ITextFieldProps, ITextFieldStyles, Label, Stack, Text, TextField, useTheme } from '@fluentui/react';
import React, { FormEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { convert, convertLength } from '../../services/convert.service';

export interface IFormatTextFieldProps extends ITextFieldProps {
  value: string
  label: string
  onChange: (e: FormEvent, v?: string) => void
  length?: number
  actions?: any
}

export const FormatTextField = (props: IFormatTextFieldProps) => {
  const { length, value, label, onChange, actions } = props;
  const theme = useTheme();
  const { t } = useTranslation();

  const fieldStyles: Partial<ITextFieldStyles> = {
    root: { height: '100%' },
    wrapper: { height: '100%' },
    fieldGroup: { height: '100%', borderColor: theme.palette.neutralQuaternaryAlt },
    subComponentStyles: {
      label: { root: { padding: '1px 0' } }
    }
  }

  const [viewValue, setViewValue] = useState<string>(),
    [viewSelectedKey, setViewSelectedKey] = useState<string | number>('plainText'),
    [_length, _setLength] = useState<number>(0);

  useEffect(() => {
    _setLength(length || 0);
  }, [length])

  const viewOptions: Array<IDropdownOption> = [
    { key: 'plainText', text: 'Plain Text' },
    { key: 'hex', text: 'HEX' },
    { key: 'json', text: 'JSON' },
    { key: 'base64ToText', text: 'BASE64 TO TEXT' },
    { key: 'base64ToJson', text: 'BASE64 TO JSON' },
    { key: 'binary', text: 'BINARY' },
  ]

  const handleChangeView = (e: FormEvent<HTMLDivElement>, option?: IDropdownOption) => {
    if (!option) return;

    convert[option.key] && convert[option.key](value).then((ret) => {
      setViewValue(ret);
    });
    setViewSelectedKey(option.key);
  }

  const handleValueChange = (e: FormEvent, v?: string) => {
    onChange(e, v);
    getLength(v || '');
  }

  const getLength = (v: string) => {
    v && convertLength({ data: v }).then((ret) => {
      _setLength(Number(ret));
    });
  }

  return (
    <Stack style={{ height: '100%', padding: '5px 0' }} tokens={{childrenGap: 3}}>

      <Stack horizontal horizontalAlign="space-between" verticalAlign="center" tokens={{ childrenGap: 5 }}>
        <Label>{label}</Label>
        {_length >= 0 && <Text style={{ color: theme.palette.neutralQuaternaryAlt }}>{`${t('Size')} ${_length} bytes`}</Text>}
        <Stack.Item grow={1}><span /></Stack.Item>
        <Text>{t('View')}</Text>
        <Dropdown
          styles={{
            title: {
              borderColor: theme.palette.neutralQuaternaryAlt
            }
          }}
          style={{ minWidth: 140 }}
          selectedKey={viewSelectedKey}
          options={viewOptions} onChange={handleChangeView} />
        {actions && actions}
      </Stack>

      <Stack.Item grow={1}>
        {viewValue ? <TextField readOnly styles={fieldStyles} {...props} value={viewValue} label={undefined}></TextField> :
          <TextField styles={fieldStyles} {...props} onChange={handleValueChange} label={undefined} ></TextField>}
      </Stack.Item>

    </Stack>
  )
}