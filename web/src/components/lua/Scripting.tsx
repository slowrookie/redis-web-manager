import { Label, PrimaryButton, Stack, Text, TextField, useTheme } from '@fluentui/react';
import Editor, {loader} from '@monaco-editor/react';
import * as monaco from 'monaco-editor';
import dayjs from 'dayjs';
import React, { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { debounceTime, fromEvent, map } from 'rxjs';
import { eidtLua, executionScript, Lua } from '../../services/lua.service';

loader.config({monaco})

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

  // re-render eiditor after window resize
  const [display, setDisplay] = useState<string>('block');
  useEffect(() => {
    const sub = fromEvent(window, "resize").pipe(
      map(v => {
        setDisplay('none');
        return v;
      }),
      debounceTime(250)
    ).subscribe(v => {
      setDisplay('block');
    })
    return () => sub && sub.unsubscribe();
  })

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
        <div style={{ height: '100%', display }}>
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
        </div>
      </Stack.Item>

      {/* buttons */}
      <Stack horizontal verticalAlign='center' tokens={{ childrenGap: 10 }} styles={{ root: { padding: '0 10px' } }}>
        {_lua && _lua.lastExecutionAt && (<>
          <Label>{`${t('Last Execution At')}:`}</Label>
          <Text variant='small'>{dayjs(_lua?.lastExecutionAt).format("YYYY-MM-DD HH:mm:ss SSS")}</Text>
        </>)}
        {_lua && _lua.elapsed && (<>
          <Label>{`${t('Elapsed')}:`}:</Label>
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