import { ITextFieldProps, ITextFieldStyles, Label, Stack, Text, TextField, useTheme } from '@fluentui/react';
import React, { FormEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

export interface IFormatTextFieldProps extends ITextFieldProps {
  value: string
  label: string
  onChange: (e: FormEvent, v?: string) => void
  length?: number
  actions?: any
}

export const FormatTextField = (props: IFormatTextFieldProps) => {
  const { length, label, onChange, actions } = props;
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

  const [_length, _setLength] = useState<number>();

  useEffect(() => {
    _setLength(length || 0);
  }, [length])

  const handleValueChange = (e: FormEvent, v?: string) => {
    onChange(e, v);
  }

  return (
    <Stack style={{ height: '100%', padding: '5px 0' }} tokens={{ childrenGap: 3 }}>

      <Stack horizontal horizontalAlign="space-between" verticalAlign="center" tokens={{ childrenGap: 5 }}>
        <Label>{label}</Label>
        {_length && <Text style={{ color: theme.palette.neutralQuaternaryAlt }}>{`${t('Size')} ${_length} bytes`}</Text>}
        <Stack.Item grow={1}><span /></Stack.Item>
        {actions && actions}
      </Stack>

      <Stack.Item grow={1}>
        <TextField styles={fieldStyles} {...props} onChange={handleValueChange} label={undefined} />
      </Stack.Item>

    </Stack>
  )
}