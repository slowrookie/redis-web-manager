import { DefaultButton, Stack, TooltipHost } from '@fluentui/react';
import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Connection, executeCommand } from '../services/connection.service';
import { ErrorMessageBar } from './common/ErrorMessageBar';
import { FormatTextField } from './common/FormatTextField';
import { Loading } from './common/Loading';
import { KeyHeader } from './KeyHeader';

const buttonStyles = {
  root: {
    minWidth: 'auto',
    padding: '0 6px'
  }
}

export interface IStringKeyProps {
  connection: Connection
  db: number
  component: string,
  keyName: string,
  onKeyNameChanged: (oldKeyName: string, keyName: string) => void
  onDeletedKey: (keyName: string) => void
}

export interface IStringKey {
  keyName: string
  TTL: number
  length: number
  value: string
  initialValue: string
}

const defaultStringKey: IStringKey = {
  keyName: '',
  TTL: -1,
  value: '',
  initialValue: '',
  length: 0
}

// 2M
const MAX_LENGTH = 2 * 1024 * 1000;

export const StringKey = (props: IStringKeyProps) => {

  const { keyName, connection, db, onKeyNameChanged } = props;

  const { t } = useTranslation();

  const [keyProps, setKeyProps] = useState({ ...defaultStringKey, keyName: keyName }),
    [error, setError] = useState<Error>(),
    [loading, setLoading] = useState(false),
    [bigKey, setBigKey] = useState(false);

  useEffect(() => {
    setKeyProps({ ...defaultStringKey, keyName })
  }, [keyName])


  // load value
  const loadValue = useCallback(() => {
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['GET', keyProps.keyName],
      ]
    }).then((ret) => {
      setKeyProps(v => {
        v.value = ret[1];
        v.initialValue = ret[1];
        return {...v};
      })
    })
      .catch(err => setError(err))
      .finally(() => setLoading(false));
  }, [connection.id, db, keyProps.keyName])

  const load = useCallback(() => {
    setLoading(true);
    setError(undefined);
    executeCommand<Array<any>>({
      id: connection.id, commands: [
        ['SELECT', db],
        ['TTL', keyProps.keyName],
        ['STRLEN', keyProps.keyName],
      ]
    }).then((ret) => {
      if (!ret || !ret.length) return;
      console.log(ret);
      setKeyProps({
        keyName: keyProps.keyName,
        TTL: ret[1],
        length: ret[2],
        value: '',
        initialValue: '',
      });
      if (Number(ret[2]) > MAX_LENGTH) {
        setBigKey(true)
      } else {
        loadValue();
      }
    })
      .catch(err => setError(err))
      .finally(() => setLoading(false));

  }, [connection.id, db, keyProps.keyName, loadValue]);

  useEffect(() => {
    load();
  }, [load]);

  const handleValue = () => {
    setError(undefined);
    executeCommand({
      id: connection.id, commands: [
        ['SELECT', db],
        ['SET', keyProps.keyName, keyProps.value],
      ]
    }).then((ret) => {
      setKeyProps({ ...keyProps, initialValue: keyProps.value });
      load();
    })
      .catch(err => setError(err));
  }

  const textFieldActions = () => {
    return (
      <TooltipHost content={t('Update')}>
        <DefaultButton styles={buttonStyles} disabled={keyProps.initialValue === keyProps.value} iconProps={{ iconName: 'save', style: { height: 'auto' } }}
          onClick={handleValue}></DefaultButton>
      </TooltipHost>
    )
  }

  return (<>
    <Loading loading={loading}></Loading>
    <Stack style={{ height: '100%' }} tokens={{ padding: 5, childrenGap: 10 }}>
      {/* error */}
      <ErrorMessageBar error={error}></ErrorMessageBar>

      <KeyHeader {...props} keyName={keyProps.keyName} TTL={keyProps.TTL}
        onKeyNameChanged={(oldValue, newValue) => {
          setKeyProps({ ...keyProps, keyName: newValue });
          onKeyNameChanged(oldValue, newValue);
        }}
        onTTLChanged={(v) => { setKeyProps({ ...keyProps, TTL: v }) }}
        onRefresh={() => load()} />

      <Stack.Item grow={1}>
        <FormatTextField label={t('Value')} disabled={bigKey} multiline value={keyProps.value} length={keyProps.length}
          onChange={(e, v) => setKeyProps({ ...keyProps, value: v || '' })}
          actions={textFieldActions()}
        ></FormatTextField>
      </Stack.Item>

    </Stack >
  </>)
}