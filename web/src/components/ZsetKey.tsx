import {

  Stack
} from '@fluentui/react';
import React, { useCallback, useEffect, useRef, useState } from 'react';
import { backendAPI } from '../api';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { KeyHeader } from './KeyHeader';
import { ZsetKeyCursor } from './ZsetKeyCursor';
import { ZsetKeyIndex } from './ZsetKeyIndex';

const DefaultZsetKeyProps = {
  keyName: '',
  TTL: -1,
}

export const ZsetKey = (props) => {

  const { connection, db, keyName, onKeyNameChanged } = props;

  const childRef = useRef();

  const [keyProps, setKeyProps] = useState({ ...DefaultZsetKeyProps, keyName }),
    [error, setError] = useState(),
    [searchType, setSearchType] = useState('CURSOR');

  const load = useCallback(() => {
    setError('');
    backendAPI.Connection.Command({
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
        onTTLChange={(v) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => { childRef.current.refresh(); load(); }} />

      {searchType === 'INDEX' && <ZsetKeyIndex  {...props} onChangeSearchType={v => setSearchType(v)} childRef={childRef} />}
      {searchType === 'CURSOR' && <ZsetKeyCursor {...props} onChangeSearchType={v => setSearchType(v)} childRef={childRef} />}

    </Stack >
  )
}