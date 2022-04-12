import { PrimaryButton, Stack, TextField } from '@fluentui/react';
import Editor from '@monaco-editor/react';
import React, { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { eidtLua, executionScript, Lua } from '../../services/lua.service';


export interface ILuaProps {
  lua: Lua,
  onChanged: (lua: Lua) => void
}

export const Scripting = (props: ILuaProps) => {
  const { lua, onChanged } = props;

  const editorRef = useRef<any>(null),
    { t } = useTranslation(),
    [ret, setRet] = useState<any>(),
    [errorMsg, setErrorMsg] = useState<string>();

  const handleExecution = (save?: boolean) => {
    if (!editorRef.current) return;
    setRet(null);
    setErrorMsg('');
    lua.Script = editorRef.current.getValue();
    (save ? eidtLua(lua).then(() => executionScript(lua)) : executionScript(lua))
      .then(v => {
        setRet(v);
        save && onChanged(lua);
      })
      .catch(err => {
        setErrorMsg(err);
        setRet(err);
      })
      .finally(() => {

      })
  }

  return (
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 10 }}>
      <TextField prefix='KEYS'
        value={lua.Keys?.join(',')}
        onChange={(e, v: string | undefined) => {
          lua.Keys = v ? v.split(',') : [];
        }}
        description={`${t('multiple keys are separated by commas `,`')}`} />
      <TextField prefix='ARGV'
        value={lua.Args?.join(',')}
        onChange={(e, v: string | undefined) => {
          lua.Args = v ? v.split(',') : [];
        }}
        description={`${t('multiple args are separated by commas `,`')}`} />
      {/* eidtor */}
      <Stack.Item grow={1}>
        <Editor
          height={'100%'}
          defaultLanguage="lua"
          value={lua.Script}
          defaultValue="-- lua script"
          onMount={(editor, monaco) => {
            editorRef.current = editor;
          }}
        />
      </Stack.Item>

      {/* buttons */}
      <Stack horizontal horizontalAlign='end' tokens={{ childrenGap: 10 }}>
        <PrimaryButton text={t('Execution')} onClick={() => { handleExecution() }}></PrimaryButton>
        <PrimaryButton text={`${t('Execution')} & ${t('Save')}`} onClick={() => { handleExecution(true) }}></PrimaryButton>
      </Stack>

      {/* result */}
      <TextField multiline rows={5}
        autoAdjustHeight
        readOnly
        errorMessage={errorMsg}
        value={ret} />
    </Stack>
  )
}