import { Stack } from '@fluentui/react';
import React, { useCallback, useEffect, useRef, useState } from 'react';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { KeyHeader } from './KeyHeader';
import { ZsetKeyCursor } from './ZsetKeyCursor';
import { ZsetKeyIndex } from './ZsetKeyIndex';

export interface IZsetKeyProps {
  connection: Connection
  db: number
  component: string,
  keyName: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface IZsetKeyItem {
  row: number,
  value: string,
  score: string
}

const DefaultZsetKeyProps = {
  keyName: '',
  TTL: -1,
}

export const ZsetKey = (props: IZsetKeyProps) => {

  const { connection, db, keyName, onKeyNameChanged } = props;

  const childRef = useRef<any>();

  const [keyProps, setKeyProps] = useState({ ...DefaultZsetKeyProps, keyName }),
    [error, setError] = useState<Error>(),
    [searchType, setSearchType] = useState('CURSOR');

  const load = useCallback(() => {
    setError(undefined);
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['TTL', keyProps.keyName],
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      console.log(ret);
      setKeyProps(kprop => { return { ...kprop, TTL: ret[1] } })
    })
      .catch(err => setError(err))
  }, [connection.id, db, keyProps.keyName]);

  useEffect(() => {
    load();
  }, [load])

  return (
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 10 }}>
      {/* error */}
      <ErrorMessageBar error={error}></ErrorMessageBar>

      <KeyHeader {...props} keyName={keyProps.keyName} TTL={keyProps.TTL}
        onKeyNameChanged={(oldValue, newValue) => {
          setKeyProps({ ...keyProps, keyName: newValue });
          onKeyNameChanged(oldValue, newValue);
        }}
        onTTLChanged={(v) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => { childRef && childRef.current && childRef.current.refresh(); load(); }} />

      {searchType === 'INDEX' && <ZsetKeyIndex  {...props} onChangeSearchType={(v: string) => setSearchType(v)} childRef={childRef} />}
      {searchType === 'CURSOR' && <ZsetKeyCursor {...props} onChangeSearchType={(v: string) => setSearchType(v)} childRef={childRef} />}

    </Stack >
  )
}