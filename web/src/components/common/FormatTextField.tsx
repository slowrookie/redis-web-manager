import { Dropdown, IDropdownOption, ITextFieldProps, Label, Stack, Text, useTheme } from '@fluentui/react';
import Editor, {loader} from '@monaco-editor/react';
import * as monaco from 'monaco-editor';
import { FormEvent, useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';

loader.config({monaco})
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

  const editorRef = useRef<any>(null),
    [_length, _setLength] = useState<number>(),
    [contentType, setContentType] = useState('text');

  useEffect(() => {
    _setLength(length || 0);
  }, [length])

  return (
    <Stack style={{ height: '100%', padding: '5px 0' }} tokens={{ childrenGap: 3 }}>

      <Stack horizontal horizontalAlign="space-between" verticalAlign="center" tokens={{ childrenGap: 5 }}>
        <Label>{label}</Label>
        {_length && <Text style={{ color: theme.palette.neutralQuaternaryAlt }}>{`${t('Size')} ${_length} bytes`}</Text>}
        <Stack.Item grow={1}><span /></Stack.Item>
        <Dropdown selectedKey={contentType} styles={{ root: { width: 70 }, title: { borderColor: theme.palette.neutralQuaternaryAlt } }}
          options={[
            { key: 'text', text: 'TEXT' },
            { key: 'json', text: 'JSON' },
          ]}
          onChange={(event: React.FormEvent<HTMLDivElement>, option?: IDropdownOption) => {
            option && setContentType(`${option.key}`)
          }} 
          />
        {actions && actions}
      </Stack>

      <Stack.Item grow={1} style={{ border: '1px solid rgb(225, 223, 221)' }}>
        <Editor
          height={'100%'}
          theme={theme.name?.includes('dark') ? 'vs-dark' : 'light'}
          defaultLanguage="text"
          language={contentType}
          value={props.value}
          onChange={(v, e: any) => {
            onChange && onChange(e, v);
          }}
          onMount={(editor, monaco) => {
            editor.updateOptions({
              minimap: { enabled: false },
              lineNumbers: 'off',
              glyphMargin: false,
              folding: false,
              lineDecorationsWidth: 0,
              scrollbar: {
                verticalScrollbarSize: 5,
		            horizontalScrollbarSize: 5,
              },
            });
            editorRef.current = editor;
          }}
        />
      </Stack.Item>

    </Stack>
  )
}