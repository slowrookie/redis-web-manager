import { Label, PrimaryButton, Stack, Text, TextField, useTheme } from '@fluentui/react';
import Editor from '@monaco-editor/react';
import React, { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { eidtLua, executionScript, Lua } from '../../services/lua.service';
import dayjs from 'dayjs';

export interface ILuaProps {
  lua: Lua,
  onChanged: (lua: Lua) => void
}

export const Scripting = (props: ILuaProps) => {
  const { lua, onChanged } = props;

  const editorRef = useRef<any>(null),
    theme = useTheme(),
    { t } = useTranslation(),
    [_lua, _setLua] = useState<Lua>(lua),
    [ret, setRet] = useState<any>(),
    [errorMsg, setErrorMsg] = useState<string>();

  useEffect(() => {
    console.log(lua);
    _setLua({ ...lua });
  }, [lua])

  const handleExecution = () => {
    setRet(null);
    setErrorMsg('');
    executionScript(_lua)
      .then(v => {
        setRet(v.result);
        _setLua(v)
      })
      .catch(err => {
        setErrorMsg(err);
        setRet(err);
      })
      .finally(() => {

      })
  }

  const handleSave = () => {
    eidtLua(_lua)
      .then(v => {
        onChanged(_lua);
      })
      .catch(err => {
        setErrorMsg(err)
        setRet(err)
      })
  }

  return (
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 10 }}>
      <TextField required prefix={t('Name')}
        value={_lua.name || ''}
        onChange={(e, v: string | undefined) => {
          _setLua({ ..._lua, name: v || '' })
        }} />

      <TextField prefix='KEYS'
        value={_lua.keys?.join(',')}
        onChange={(e, v: string | undefined) => {
          _setLua({ ..._lua, keys: v ? v.split(',') : [] })
        }}
        description={`${'multiple keys are separated by commas `,`'}`} />

      <TextField prefix='ARGV'
        value={_lua.args?.join(',')}
        onChange={(e, v: string | undefined) => {
          _setLua({ ..._lua, args: v ? v.split(',') : [] })
        }}
        description={`${'multiple args are separated by commas `,`'}`} />

      {/* eidtor */}
      <Stack.Item grow={1}>
        <Editor
          height={'100%'}
          theme={theme.name?.includes('dark') ? 'vs-dark' : 'light'}
          defaultLanguage="lua"
          value={_lua.script}
          defaultValue="-- lua script"
          onChange={(v) => {
            _setLua({ ..._lua, script: v || '' })
          }}
          onMount={(editor, monaco) => {
            editorRef.current = editor;
          }}
        />
      </Stack.Item>

      {/* buttons */}
      <Stack horizontal verticalAlign='center' tokens={{ childrenGap: 10 }} styles={{ root: { padding: '0 10px' } }}>
        {_lua && _lua.lastExecutionAt && (<>
          <Label>上次执行时间:</Label>
          <Text variant='small'>{dayjs(_lua?.lastExecutionAt).format("YYYY-MM-DD HH:mm:ss SSS")}</Text>
        </>)}
        {_lua && _lua.elapsed && (<>
          <Label>耗时:</Label>
          <Text variant='small'>{_lua?.elapsed}</Text>
        </>)}
        <Stack.Item grow={1}><span></span></Stack.Item>
        <PrimaryButton disabled={!_lua?.name} text={`${t('Save')}`} onClick={() => { handleSave() }}></PrimaryButton>
        <PrimaryButton disabled={!_lua?.name} text={t('Execution')} onClick={() => { handleExecution() }}></PrimaryButton>
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